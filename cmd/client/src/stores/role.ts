import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { CreateRoleReq, CreateRoleUserReq, DeleteRoleReq, DeleteRoleUserReq, ListRoleReq, RolePermission, RoleUsers, UpdateRoleReq } from '@internal/user/dto/role'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'
import type { Permission } from '@internal/user/role'
import type { U } from '@internal/user/user'

export const useRoleStore = defineStore('role', () => {
  const roles = ref<Map<string, RolePermission>>(new Map())
  const total = ref<bigint>(BigInt(0))

  // users by role
  const users = ref<Map<string, U[]>>(new Map())

  const selected = ref<RolePermission[]>([])

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const create = async function (entityID: Uint8Array, name: string, permissions: Permission[]) {
    try {
      const req = CreateRoleReq.create({
        entityID: entityID,
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

      roles.value?.set(ulid(resp.response.iD), resp.response)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Return roles ids and entity ids
  const list = async function (req: ListRoleReq): Promise<string[]> {
    try {
      const resp = await api.listRole(req, { meta: { token: token.value } })

      resp.response.roles?.forEach((role: RoleUsers) => {
        roles.value?.set(ulid(role.role?.role?.iD), role?.role!)
        users.value?.set(ulid(role.role?.role?.iD), role.users)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.roles.map((role: RoleUsers) => ulid(role.role?.role?.iD))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Add user refresh cache in selected role only
  const addUser = async function (roleID: Uint8Array, userID: Uint8Array) {
    try {
      const req = CreateRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.createRoleUser(req, { meta: { token: token.value } })
      users.value.set(ulid(resp.response.role?.iD), resp.response.users)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }


  const deleteUser = async function (roleID: Uint8Array, userID: Uint8Array) {
    try {
      const req = DeleteRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.deleteRoleUser(req, { meta: { token: token.value } })
      users.value.set(ulid(resp.response.role?.iD), resp.response.users)

      // if no users left, remove user from cache
      if (users.value.get(ulid(resp.response.role?.iD))?.length === 0) {
        users.value.delete(ulid(resp.response.role?.iD))
      }
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }


  const delete_ = async (req: DeleteRoleReq) => {
    try {
      const resp = await api.deleteRole(req, { meta: { token: token.value } })

      roles.value?.delete(ulid(resp.response.role?.iD))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  return {
    roles,
    users,
    total,
    selected,
    create,
    update,
    list,
    addUser,
    deleteUser,
    delete_,
  }
})
