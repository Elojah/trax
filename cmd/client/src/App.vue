<script setup lang="ts">
import Sidebar from '@/components/Sidebar.vue';
import Errors from '@/components/Errors.vue';
import ConfirmDialog from 'primevue/confirmdialog';
import { onMounted, onUnmounted, ref } from 'vue'
import { useAuthStore } from '@/stores/auth';

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
    <div class="min-h-screen bg-surface-0 dark:bg-surface-950">
      <Sidebar />
      <div class="lg:ml-24 h-screen overflow-y-auto">
        <RouterView />
      </div>
      <Errors />
      <ConfirmDialog />
    </div>
  </v-app>
</template>

<style scoped></style>
