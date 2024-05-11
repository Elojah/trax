import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Permission } from '@internal/user/permission'
import { CreatePermissionReq, ListPermissionReq } from '@internal/user/dto/permission'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'

export const usePermissionStore = defineStore('permission', () => {
  const permissions = ref<Map<string, Permission>>(new Map())
  const total = ref<bigint>(BigInt(0))

  const selected = ref<Permission[]>([])

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const create = async function (name: string) {
    try {
      const req = CreatePermissionReq.create({
        name: name
      })

      return await api.createPermission(req, { meta: { token: token.value } })
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  // Return permissions ids and entity ids
  const list = async function (req: ListPermissionReq): Promise<[string[], string[]]> {
    try {
      const resp = await api.listPermission(req, { meta: { token: token.value } })

      resp.response.permissions?.forEach((permission: Permission) => {
        permissions.value?.set(ulid(permission.iD), permission)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.permissions.reduce(
        (acc: string[][], permission: Permission) => {
          acc[0].push(ulid(permission.iD))
          acc[1].push(ulid(permission.entityID))
          return acc
        },
        [[], []]
      )
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
    return [[], []]
  }

  return {
    permissions,
    total,
    selected,
    create,
    list
  }
})
