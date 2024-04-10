import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Entity } from '@internal/user/entity'
import { Paginate } from '@pkg/paginate/paginate'
import { ListEntityReq } from '@internal/user/dto/entity'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'

export const useEntityStore = defineStore('entity', () => {
  const entity = ref(null as Entity | null)
  const entities = ref([] as Entity[] | null)

  const newEntity = ref({} as Entity)

  const total = ref(BigInt(0) as bigint)

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const createEntity = async function () {
    try {
      const req = Entity.create({
        name: newEntity.value?.name,
        avatarURL: newEntity.value?.avatarURL,
        description: newEntity.value?.description
      })

      const resp = await api.createEntity(req, { meta: { token: token.value } })
      newEntity.value = resp.response
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

      entities.value = resp.response.entities
      total.value = resp.response.total
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
    newEntity,
    total,
    createEntity,
    listEntity
  }
})
