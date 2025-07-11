<script setup lang="ts">
import Menu from '@/components/Menu.vue';
import Errors from '@/components/Errors.vue';
import Badge from 'primevue/badge';

import { onMounted, onUnmounted, ref } from 'vue'
import { useAuthStore } from './stores/auth';

const authStore = useAuthStore();
const timer = ref();

onMounted(async () => {
  timer.value = setInterval(() => {
    // authStore.refreshToken();
  }, 10 * 60 * 1000); // TODO: configurable
});

onUnmounted(() => clearInterval(timer.value))
</script>

<template>
  <v-app id="trax">
    <RouterView />
    <Errors />
  </v-app>
</template>

<style scoped></style>
