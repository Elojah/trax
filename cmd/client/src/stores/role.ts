import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Role } from '@internal/user/role'
import { Paginate } from '@pkg/paginate/paginate'
import { CreateRoleReq, ListRoleReq } from '@internal/user/dto/role'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'

export const useRoleStore = defineStore('role', () => {
  const role = ref<Role | null>(null)
  const roles = ref<Role[] | null>([])
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

  const listRole = async function (search: string, p: Paginate) {
    try {
      const req = ListRoleReq.create({
        search: search,
        paginate: p
      })

      const resp = await api.listRole(req, { meta: { token: token.value } })

      // TODO: manage errors

      roles.value = resp.response.roles
      total.value = resp.response.total

      return resp
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  return {
    role,
    roles,
    total,
    createRole,
    listRole
  }
})
