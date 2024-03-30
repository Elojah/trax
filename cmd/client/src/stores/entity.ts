import { defineStore } from 'pinia';
import { config, logger } from "@/config"
import { APIClient } from '@api/api.client';
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { Entity } from '@internal/user/entity';
import { useAuthStore } from './auth';
import { computed, ref } from 'vue';

export const useEntityStore = defineStore('entity', () => {
  const entity = ref(null as Entity | null)

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

  return {
    entity,
    createEntity
  }
});
