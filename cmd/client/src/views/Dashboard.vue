<script setup lang="ts">
import Dashboard from '@/components/Dashboard.vue';
</script>

<template>
  <v-main class="bg">
    <v-container>
      <Dashboard />
    </v-container>
  </v-main>
</template>

<style scoped>
.bg {
  background: url('@/assets/img/main-background.svg') !important;
}
</style>
