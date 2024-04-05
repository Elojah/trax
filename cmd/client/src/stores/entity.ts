import { defineStore } from 'pinia';
import { config, logger } from "@/config"
import { APIClient } from '@api/api.client';
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { Entity } from '@internal/user/entity';
import { ListEntityReq } from '@internal/user/dto/entity';
import { useAuthStore } from './auth';
import { computed, ref } from 'vue';

export const useEntityStore = defineStore('entity', () => {
  const entity = ref({} as Entity)
  const entities = ref([] as Entity[] | null)

  const api = new APIClient(new GrpcWebFetchTransport({
    baseUrl: config.api_url,
  }))
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const createEntity = async function () {
    try {
      const req = Entity.create({
        name: entity.value?.name
      });

      const resp = await api.createEntity(req, { meta: { token: token.value } })
      entity.value = resp.response;
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err);
      }
    }
  };

  const listEntity = async function () {
    try {
      const req = ListEntityReq.create({
        paginate: {
          start: 1,
          end: 10,
          sort: "created_at",
          order: false,
        }
      });

      const resp = await api.listEntity(req, { meta: { token: token.value } })
      console.log(resp)
      entities.value = resp.response.entities;
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err);
      }
    }
  };

  return {
    entity,
    entities,
    createEntity,
    listEntity,
  }
});
