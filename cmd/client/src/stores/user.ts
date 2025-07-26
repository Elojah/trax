import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { ListUserReq, ListUserResp } from '@internal/user/dto/user'
import { useAuthStore } from '@/stores/auth'
import { computed, ref } from 'vue'
import { parse, ulid, zero } from '@/utils/ulid'
import { CreateRoleUserReq, DeleteRoleUserReq, RoleUserResp } from '@internal/user/dto/role'
import type { U } from '@internal/user/user'
import type { Role } from '@internal/user/role'

export const useUserStore = defineStore('user', () => {
  const users = ref<Map<string, U>>(new Map())
  const total = ref<bigint>(BigInt(0))

  // users by role
  const usersByRole = ref<Map<string, Map<string, boolean>>>(new Map())

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const invite = async function (name: string, roles: Role[]) { throw new Error('not implemented') }

  // Return users ids and group ids
  const list = async function (req: ListUserReq): Promise<string[]> {
    try {
      const resp: { response: ListUserResp } = await api.listUser(req, { meta: { token: token.value } })

      resp.response.users?.forEach((user: U) => {
        users.value?.set(ulid(user?.iD), user)
      })

      const userIDs: string[] = resp.response.users.map((user: U) => ulid(user?.iD))

      if (req.roleID.length > 0) {
        usersByRole.value.set(ulid(req.roleID), userIDs.reduce((acc: Map<string, boolean>, userID: string) => {
          acc.set(userID, true);

          return acc;
        }, new Map<string, boolean>()))
      }

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return userIDs
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const addRole = async function (userID: Uint8Array, roleID: Uint8Array) {
    try {
      // zero case exception, dry update local only
      if (ulid(roleID) === ulid(zero)) {
        const users = usersByRole.value.get(ulid(roleID)) ?? new Map()
        users?.set(ulid(zero), true)
        usersByRole.value.set(ulid(roleID), users)

        return
      }

      const req = CreateRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp: { response: RoleUserResp } = await api.createRoleUser(req, { meta: { token: token.value } })

      const users = usersByRole.value.get(ulid(resp.response.role?.role?.iD)) ?? new Map()
      users?.set(ulid(resp.response.user?.iD), true)
      usersByRole.value.set(ulid(resp.response.role?.role?.iD), users)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const deleteRole = async function (userID: Uint8Array, roleID: Uint8Array) {

    try {
      // zero case exception, dry update local only
      if (ulid(roleID) === ulid(zero)) {
        const users = usersByRole.value.get(ulid(roleID)) ?? new Map()
        users?.delete(ulid(zero))
        usersByRole.value.set(ulid(roleID), users)

        return
      }

      const req = DeleteRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp: { response: RoleUserResp } = await api.deleteRoleUser(req, { meta: { token: token.value } })

      const users = usersByRole.value.get(ulid(resp.response.role?.role?.iD)) ?? new Map()
      users?.delete(ulid(resp.response.user?.iD))
      usersByRole.value.set(ulid(resp.response.role?.role?.iD), users)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }


  const populate = async function (ids: string[], groupIDs: string[]) {
    const newIDs = ids.reduce((acc: Uint8Array[], id: string) => {
      if (users.value?.has(id)) {
        return acc
      } else {
        return [...acc, parse(id)]
      }
    }, [])

    if (newIDs.length === 0) {
      return
    }

    await list(
      ListUserReq.create({
        iDs: newIDs,
        groupIDs: groupIDs.map(id => parse(id)),
        paginate: {
          start: BigInt(0),
          end: BigInt(newIDs.length),
          order: true,
          sort: 'created_at'
        }
      })
    )
  }



  return {
    users,
    total,
    usersByRole,
    invite,
    list,
    addRole,
    deleteRole,
    populate,
  }
})
