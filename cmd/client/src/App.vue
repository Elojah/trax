<script setup lang="ts">
import Drawer from '@/components/Drawer.vue';

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
  <v-app id="trax">
    <Drawer />
    <RouterView />
  </v-app>
</template>

<style scoped></style>
