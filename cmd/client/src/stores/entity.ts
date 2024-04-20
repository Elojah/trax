import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Entity } from '@internal/user/entity'
import { Paginate } from '@pkg/paginate/paginate'
import { CreateEntityReq, ListEntityReq } from '@internal/user/dto/entity'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { ulid } from '@/utils/ulid'

export const useEntityStore = defineStore('entity', () => {
  const entities = ref<Entity[] | null>([])
  const entitiesByID = ref<Map<string, Entity>>(new Map())

  const total = ref<bigint>(BigInt(0))

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

  const listEntityByID = async function (ids: Uint8Array[], search: string, p: Paginate) {
    try {
      const req = ListEntityReq.create({
        search: search,
        paginate: p,
        ids: ids
      })

      const resp = await api.listEntity(req, { meta: { token: token.value } })

      // TODO: manage errors

      resp.response.entities?.forEach((entity: Entity) => {
        entitiesByID.value?.set(ulid(entity.iD), entity)
      })

      return resp
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const populateEntityByID = async function (ids: Uint8Array[]) {
    const newIDs = ids.reduce((acc: Uint8Array[], id: Uint8Array) => {
      if (entitiesByID.value?.has(ulid(id))) {
        return acc
      } else {
        return [...acc, id]
      }
    }, [])

    if (newIDs.length === 0) {
      return
    }

    await listEntityByID(newIDs, '', {
      start: BigInt(0),
      end: BigInt(newIDs.length),
      order: true,
      sort: 'created_at'
    })
  }

  return {
    entities,
    entitiesByID,
    total,
    createEntity,
    listEntity,
    listEntityByID,
    populateEntityByID
  }
})
