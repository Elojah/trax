<script setup lang="ts">
import Sidebar from '@/components/Sidebar.vue';
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
  <v-app id="trax">
    <div class="min-h-screen flex relative lg:static bg-surface-0 dark:bg-surface-950">
      <Sidebar />
      <RouterView />
      <Errors />
    </div>
  </v-app>
</template>

<style scoped></style>
