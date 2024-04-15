import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Entity } from '@internal/user/entity'
import { Paginate } from '@pkg/paginate/paginate'
import { CreateEntityReq, ListEntityReq } from '@internal/user/dto/entity'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'

export const useEntityStore = defineStore('entity', () => {
  const entity = ref(null as Entity | null)

  const entities = ref([] as Entity[] | null)

  const total = ref(BigInt(0) as bigint)

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const createEntity = async function (name: string, avatarURL: string, description: string) {
    try {
      const req = CreateEntityReq.create({
        name: name,
        avatarURL: avatarURL,
        description: description
      })

      return await api.createEntity(req, { meta: { token: token.value } })
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const listEntity = async function (search: string, p: Paginate) {
    try {
      const req = ListEntityReq.create({
        search: search,
        paginate: p
      })

      const resp = await api.listEntity(req, { meta: { token: token.value } })

      // TODO: manage errors

      entities.value = resp.response.entities
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
    entity,
    entities,
    total,
    createEntity,
    listEntity
  }
})
