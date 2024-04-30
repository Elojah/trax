import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Role } from '@internal/user/role'
import { CreateRoleReq, ListRoleReq } from '@internal/user/dto/role'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'

export const useRoleStore = defineStore('role', () => {
  const roles = ref<Map<string, Role>>(new Map())
  const total = ref<bigint>(BigInt(0))

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const createRole = async function (name: string) {
    try {
      const req = CreateRoleReq.create({
        name: name
      })

      return await api.createRole(req, { meta: { token: token.value } })
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  // Return roles ids and entity ids
  const listRole = async function (req: ListRoleReq): Promise<[string[], string[]]> {
    try {
      const resp = await api.listRole(req, { meta: { token: token.value } })

      resp.response.roles?.forEach((role: Role) => {
        roles.value?.set(ulid(role.iD), role)
      })

      if (req.iDs === null) {
        total.value = resp.response.total
      }

      return resp.response.roles.reduce(
        (acc: string[][], role: Role) => {
          acc[0].push(ulid(role.iD))
          acc[1].push(ulid(role.entityID))
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
    roles,
    total,
    createRole,
    listRole
  }
})
