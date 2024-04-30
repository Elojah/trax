import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { U } from '@internal/user/user'
import { Paginate } from '@pkg/paginate/paginate'
import { ListUserReq } from '@internal/user/dto/user'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'

export const useUserStore = defineStore('user', () => {
  const users = ref<Map<string, U>>(new Map())
  const total = ref<bigint>(BigInt(0))

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const inviteUser = async function (name: string) {
    // try {
    //   const req = CreateUserReq.create({
    //     name: name
    //   })
    //   return await api.createUser(req, { meta: { token: token.value } })
    // } catch (err: any) {
    //   switch (err.code) {
    //     default:
    //       logger.error(err)
    //   }
    // }
  }

  // Return users ids and entity ids
  const listUser = async function (
    ids: Uint8Array[] | null,
    search: string,
    p: Paginate
  ): Promise<[string[], string[]]> {
    try {
      const req = ListUserReq.create({
        search: search,
        paginate: p,
        iDs: ids
      })

      const resp = await api.listUser(req, { meta: { token: token.value } })

      resp.response.users?.forEach((user: U) => {
        users.value?.set(ulid(user.iD), user)
      })

      if (ids === null) {
        total.value = resp.response.total
      }

      return resp.response.users.reduce(
        (acc: string[][], user: U) => {
          acc[0].push(ulid(user.iD))
          // acc[1].push(ulid(user.entityID))
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
    users,
    total,
    inviteUser,
    listUser
  }
})
