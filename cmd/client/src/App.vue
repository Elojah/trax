<script setup lang="ts">
import Drawer from '@/components/Drawer.vue';
</script>

<template>
  <v-app id="trax">
    <Drawer />
    <RouterView />
  </v-app>
</template>

<style scoped></style>
