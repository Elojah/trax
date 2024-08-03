import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { ListUserReq, UserRoles } from '@internal/user/dto/user'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'
import { CreateRoleUserReq, DeleteRoleUserReq } from '@internal/user/dto/role'
import type { U } from '@internal/user/user'
import type { Role } from '@internal/user/role'

export const useUserStore = defineStore('user', () => {
  const users = ref<Map<string, U>>(new Map())
  const total = ref<bigint>(BigInt(0))

  // roles by user
  const roles = ref<Map<string, Role[]>>(new Map())

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

      resp.response.users?.forEach((user: UserRoles) => {
        roles.value?.set(ulid(user.user?.iD), user.roles)
        users.value?.set(ulid(user.user?.iD), user.user!)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.users.map((user: UserRoles) => ulid(user.user?.iD))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Add role refresh cache in selected user only
  const addRole = async function (userID: Uint8Array, roleID: Uint8Array) {
    try {
      const req = CreateRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.createRoleUser(req, { meta: { token: token.value } })
      roles.value.set(ulid(resp.response.user?.iD), resp.response.roles)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const deleteRole = async function (userID: Uint8Array, roleID: Uint8Array) {
    try {
      const req = DeleteRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.deleteRoleUser(req, { meta: { token: token.value } })
      roles.value.set(ulid(resp.response.user?.iD), resp.response.roles)

      // if no roles left, remove user from cache
      if (roles.value.get(ulid(resp.response.user?.iD))?.length === 0) {
        roles.value.delete(ulid(resp.response.user?.iD))
        users.value.delete(ulid(resp.response.user?.iD))
      }
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Add role modifies local cache only, no API call
  const addRoleDry = async function (userID: Uint8Array, role: Role) {
    try {
      const rs = roles.value.get(ulid(userID))

      if (!rs) {
        roles.value.set(ulid(userID), [role])
      } else {
        rs.push(role)
        roles.value.set(ulid(userID), rs)
      }
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Delete role modifies local cache only, no API call
  const deleteRoleDry = async function (userID: Uint8Array, roleID: Uint8Array) {
    try {
      const rs = roles.value.get(ulid(userID))
      const id = ulid(roleID)

      if (!rs) {
        return
      } else {
        const newRoles = rs.filter((r) => ulid(r.iD) !== id)
        roles.value.set(ulid(userID), newRoles)
      }
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Reset roles for user, no API call
  const resetRoleDry = async function (userID: Uint8Array) {
    try {
      roles.value.delete(ulid(userID))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Delete role deletes role for all users, no API call
  const deleteRoleGlobal = async function (roleID: Uint8Array) {
    try {
      roles.value.forEach((rs, userID) => {
        const id = ulid(roleID)

        if (!rs) {
          return
        } else {
          const newRoles = rs.filter((r) => ulid(r.iD) !== id)
          roles.value.set(userID, newRoles)
        }
      })
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  return {
    users,
    total,
    roles,
    invite,
    list,
    addRole,
    deleteRole,
    addRoleDry,
    deleteRoleDry,
    resetRoleDry,
    deleteRoleGlobal,
  }
})
