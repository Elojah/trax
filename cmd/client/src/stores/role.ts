import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { CreateRoleReq, CreateRoleUserReq, DeleteRoleReq, DeleteRoleUserReq, ListRoleReq, ListRoleResp, RolePermission, RoleUserResp, UpdateRoleReq } from '@internal/user/dto/role'
import { useAuthStore } from '@/stores/auth'
import { computed, ref } from 'vue'
import { parse, ulid, zero } from '@/utils/ulid'
import type { Permission } from '@internal/user/role'

export const useRoleStore = defineStore('role', () => {
  const roles = ref<Map<string, RolePermission>>(new Map())
  const total = ref<bigint>(BigInt(0))

  // roles by user
  // user_id -> role_id -> true
  const rolesByUser = ref<Map<string, Map<string, boolean>>>(new Map())

  const selected = ref<RolePermission[]>([])

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const create = async function (groupID: Uint8Array, name: string, permissions: Permission[]) {
    try {
      const req = CreateRoleReq.create({
        groupID: groupID,
        name: name,
        permissions: permissions
      })

      return await api.createRole(req, { meta: { token: token.value } })
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const update = async (req: UpdateRoleReq) => {
    try {
      const resp = await api.updateRole(req, { meta: { token: token.value } })

      roles.value?.set(ulid(resp.response.role?.iD), resp.response)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Return roles ids and group ids
  const list = async function (req: ListRoleReq): Promise<string[]> {
    try {
      const resp: { response: ListRoleResp } = await api.listRole(req, { meta: { token: token.value } })

      resp.response.roles?.forEach((role: RolePermission) => {
        roles.value?.set(ulid(role.role?.iD), role)
      })

      const roleIDs: string[] = resp.response.roles.map((role: RolePermission) => ulid(role.role?.iD))

      if (req.userID.length > 0) {
        rolesByUser.value.set(ulid(req.userID), roleIDs.reduce((acc: Map<string, boolean>, roleID: string) => {
          acc.set(roleID, true);

          return acc;
        }, new Map<string, boolean>()))
      }

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return roleIDs
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Add user to role
  const addUser = async function (roleID: Uint8Array, userID: Uint8Array, dry: boolean = false) {
    try {
      // zero case exception, dry update local only
      if (dry) {
        const roles = rolesByUser.value.get(ulid(userID)) ?? new Map()
        roles?.set(ulid(roleID), true)
        rolesByUser.value.set(ulid(userID), roles)

        return
      }

      const req = CreateRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp: { response: RoleUserResp } = await api.createRoleUser(req, { meta: { token: token.value } })

      const roles = rolesByUser.value.get(ulid(resp.response.user?.iD)) ?? new Map()
      roles?.set(ulid(resp.response.role?.role?.iD), true)
      rolesByUser.value.set(ulid(resp.response.user?.iD), roles)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }


  const deleteUser = async function (roleID: Uint8Array, userID: Uint8Array, dry: boolean = false) {
    try {
      // zero case exception, dry update local only
      if (dry) {
        const roles = rolesByUser.value.get(ulid(userID)) ?? new Map()
        roles?.delete(ulid(roleID))
        rolesByUser.value.set(ulid(userID), roles)

        return
      }

      const req = DeleteRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp: { response: RoleUserResp } = await api.deleteRoleUser(req, { meta: { token: token.value } })

      const roles = rolesByUser.value.get(ulid(resp.response.user?.iD)) ?? new Map()
      roles?.delete(ulid(resp.response.role?.role?.iD))
      rolesByUser.value.set(ulid(resp.response.user?.iD), roles)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }


  const delete_ = async (req: DeleteRoleReq) => {
    try {
      const resp: { response: RolePermission } = await api.deleteRole(req, { meta: { token: token.value } })

      roles.value?.delete(ulid(resp.response.role?.iD))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const populate = async function (ids: string[], groupIDs: string[]) {
    const newIDs = ids.reduce((acc: Uint8Array[], id: string) => {
      if (roles.value?.has(id)) {
        return acc
      } else {
        return [...acc, parse(id)]
      }
    }, [])

    if (newIDs.length === 0) {
      return
    }

    await list(
      ListRoleReq.create({
        iDs: newIDs,
        groupIDs: groupIDs.map(id => parse(id)),
        ownGroup: true,
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
    roles,
    rolesByUser,
    total,
    selected,
    create,
    update,
    list,
    delete_,
    addUser,
    deleteUser,
    populate,
  }
})
