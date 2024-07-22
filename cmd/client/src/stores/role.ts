import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { CreateRoleReq, DeleteRoleReq, ListRoleReq, RolePermission, UpdateRoleReq } from '@internal/user/dto/role'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'
import type { Permission } from '@internal/user/role'

export const useRoleStore = defineStore('role', () => {
  const roles = ref<Map<string, RolePermission>>(new Map())
  const total = ref<bigint>(BigInt(0))

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
  const list = async function (req: ListRoleReq): Promise<[string[], string[]]> {
    try {
      const resp = await api.listRole(req, { meta: { token: token.value } })

      resp.response.roles?.forEach((role: RolePermission) => {
        roles.value?.set(ulid(role.role?.iD), role)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.roles.reduce(
        (acc: string[][], role: RolePermission) => {
          acc[0].push(ulid(role.role?.iD))
          acc[1].push(ulid(role.role?.entityID))
          return acc
        },
        [[], []]
      )
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
    total,
    selected,
    create,
    update,
    list,
    delete_,
  }
})
