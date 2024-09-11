import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { ListUserReq } from '@internal/user/dto/user'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid, zero } from '@/utils/ulid'
import { CreateRoleUserReq, DeleteRoleUserReq } from '@internal/user/dto/role'
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

  // Return users ids and entity ids
  const list = async function (req: ListUserReq): Promise<string[]> {
    try {
      const resp = await api.listUser(req, { meta: { token: token.value } })

      resp.response.users?.forEach((user: U) => {
        users.value?.set(ulid(user?.iD), user)
      })

      const userIDs = resp.response.users.map((user: U) => ulid(user?.iD))

      if (req.roleID) {
        usersByRole.value.set(ulid(req.roleID), userIDs.reduce((acc: Map<string, boolean>, userID: Uint8Array) => {
          acc.set(ulid(userID), true);

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
      if (ulid(userID) === ulid(zero)) {
        const users = usersByRole.value.get(ulid(roleID)) ?? new Map()
        users?.set(ulid(zero), true)
        usersByRole.value.set(ulid(roleID), users)

        return
      }

      const req = CreateRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.createRoleUser(req, { meta: { token: token.value } })

      const users = usersByRole.value.get(ulid(resp.response.role?.iD)) ?? new Map()
      users?.set(ulid(resp.response.user?.iD), true)
      usersByRole.value.set(ulid(resp.response.role?.iD), users)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const deleteRole = async function (userID: Uint8Array, roleID: Uint8Array) {
    try {
      // zero case exception, dry update local only
      if (ulid(userID) === ulid(zero)) {
        const users = usersByRole.value.get(ulid(roleID)) ?? new Map()
        users?.delete(ulid(zero))
        usersByRole.value.set(ulid(roleID), users)

        return
      }

      const req = DeleteRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.deleteRoleUser(req, { meta: { token: token.value } })

      const users = usersByRole.value.get(ulid(resp.response.role?.iD)) ?? new Map()
      users?.delete(ulid(resp.response.user?.iD))
      usersByRole.value.set(ulid(resp.response.role?.iD), users)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }


  return {
    users,
    total,
    usersByRole,
    invite,
    list,
    addRole,
    deleteRole,
  }
})
