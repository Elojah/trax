import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { ListUserReq, UserRoles } from '@internal/user/dto/user'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'
import { CreateRoleUserReq } from '@internal/user/dto/role'

export const useUserStore = defineStore('user', () => {
  const users = ref<Map<string, UserRoles>>(new Map())
  const total = ref<bigint>(BigInt(0))

  const selected = ref<UserRoles[]>([])

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const invite = async function (name: string) { throw new Error('not implemented') }

  // Return users ids and entity ids
  const list = async function (req: ListUserReq): Promise<string[]> {
    try {
      const resp = await api.listUser(req, { meta: { token: token.value } })

      resp.response.users?.forEach((user: UserRoles) => {
        users.value?.set(ulid(user.user?.iD), user)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.users.map((user: UserRoles) => ulid(user.user?.iD))
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
    return []
  }

  const addRole = async function (userID: Uint8Array, roleID: Uint8Array) {
    try {
      const req = CreateRoleUserReq.create({
        userID: userID,
        roleID: roleID
      })

      const resp = await api.createRoleUser(req, { meta: { token: token.value } })
      users.value.set(ulid(resp.response.user?.iD), resp.response)
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  return {
    users,
    total,
    selected,
    invite,
    list,
    addRole
  }
})
