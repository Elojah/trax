import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Group } from '@internal/user/group'
import {
  CreateGroupReq,
  DeleteGroupReq,
  ListGroupReq,
  ListGroupResp,
  UpdateGroupReq
} from '@internal/user/dto/group'
import { useAuthStore } from './auth'
import { computed, ref } from 'vue'
import { parse, ulid } from '@/utils/ulid'

export const useGroupStore = defineStore('group', () => {
  const groups = ref<Map<string, Group>>(new Map())
  const total = ref<bigint>(BigInt(0))

  const selected = ref<Group[]>([])

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const create = async function (name: string, avatarURL: string, description: string) {
    try {
      const req = CreateGroupReq.create({
        name: name,
        avatarURL: avatarURL,
        description: description
      })

      return await api.createGroup(req, { meta: { token: token.value } })
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const update = async (req: UpdateGroupReq) => {
    try {
      const resp = await api.updateGroup(req, { meta: { token: token.value } })

      groups.value?.set(ulid(resp.response.iD), resp.response)
    } catch (err: any) {
      logger.error(err)
      throw err;
    }
  }

  const delete_ = async (req: DeleteGroupReq) => {
    try {
      const resp: { response: Group } = await api.deleteGroup(req, { meta: { token: token.value } })

      groups.value?.delete(ulid(resp.response.iD))
    } catch (err: any) {
      logger.error(err)
      throw err;
    }
  }

  const list = async function (req: ListGroupReq): Promise<string[]> {
    try {
      const resp: { response: ListGroupResp } = await api.listGroup(req, { meta: { token: token.value } })

      resp.response.groups?.forEach((group: Group) => {
        groups.value?.set(ulid(group.iD), group)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.groups.map((group: Group) => ulid(group.iD))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const populate = async function (ids: string[]) {
    const newIDs = ids.reduce((acc: Uint8Array[], id: string) => {
      if (groups.value?.has(id)) {
        return acc
      } else {
        return [...acc, parse(id)]
      }
    }, [])

    if (newIDs.length === 0) {
      return
    }

    await list(
      ListGroupReq.create({
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
    groups,
    total,
    selected,
    create,
    list,
    update,
    delete_,
    populate
  }
})
