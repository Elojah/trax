<script setup lang="ts">
import Drawer from '@/components/Drawer.vue';
import Errors from '@/components/Errors.vue';
import Menu from 'primevue/menu';

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

const items = ref([
    { label: 'New', icon: 'pi pi-plus' },
    { label: 'Search', icon: 'pi pi-search' }
]);
</script>

<template>
      <div class="card flex justify-center">
        <Menu :model="items" />
    </div>
</template>

<style scoped></style>
