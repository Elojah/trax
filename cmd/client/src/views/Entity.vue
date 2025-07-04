<script setup lang="ts">
import EntityTable from '@/components/entity/Table.vue';
import EntityDetails from '@/components/entity/Details.vue';
import RoleTable from '@/components/role/Table.vue';
import UserTable from '@/components/user/Table.vue';
import { ref } from 'vue';

// [Info, Roles, User]
const view = ref(0)

const tableEntities = ref();
</script>

<template>
</template>

<style scoped>
</style>
