<script setup lang="ts">
import Menu from '@/components/Menu.vue';
import Errors from '@/components/Errors.vue';

import { onMounted, onUnmounted, ref } from 'vue'
import { useAuthStore } from './stores/auth';

const authStore = useAuthStore();
const timer = ref();

onMounted(async () => {
  timer.value = setInterval(() => {
    authStore.refreshToken();
  }, 10 * 60 * 1000); // TODO: configurable
});

onUnmounted(() => clearInterval(timer.value))
</script>

<template>
  <div class="layout">
    <Menu />
    <div class="layout-content">
      <!-- <router-view /> -->
      <!-- <Errors /> -->
    </div>
  </div>
</template>

<style scoped></style>
