import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Entity } from '@internal/user/entity'
import {
  CreateEntityReq,
  DeleteEntityReq,
  ListEntityReq,
  UpdateEntityReq
} from '@internal/user/dto/entity'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { parse, ulid } from '@/utils/ulid'

export const useEntityStore = defineStore('entity', () => {
  const entities = ref<Map<string, Entity>>(new Map())

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

  const updateEntity = async (req: UpdateEntityReq) => {
    try {
      const resp = await api.updateEntity(req, { meta: { token: token.value } })

      entities.value?.set(ulid(resp.response.iD), resp.response)
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const deleteEntity = async (req: DeleteEntityReq) => {
    try {
      const resp = await api.deleteEntity(req, { meta: { token: token.value } })

      entities.value?.delete(ulid(resp.response.Entity.iD))
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const listEntity = async function (req: ListEntityReq): Promise<string[]> {
    try {
      const resp = await api.listEntity(req, { meta: { token: token.value } })

      resp.response.entities?.forEach((entity: Entity) => {
        entities.value?.set(ulid(entity.iD), entity)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.entities.map((entity: Entity) => ulid(entity.iD))
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
    return []
  }

  const populateEntity = async function (ids: string[]) {
    const newIDs = ids.reduce((acc: Uint8Array[], id: string) => {
      if (entities.value?.has(id)) {
        return acc
      } else {
        return [...acc, parse(id)]
      }
    }, [])

    if (newIDs.length === 0) {
      return
    }

    await listEntity(
      ListEntityReq.create({
        iDs: newIDs,
        paginate: {
          start: BigInt(0),
          end: BigInt(newIDs.length),
          order: true,
          sort: 'created_at'
        }
      })
    )
  }

  return {
    entities,
    total,
    createEntity,
    listEntity,
    updateEntity,
    deleteEntity,
    populateEntity
  }
})
