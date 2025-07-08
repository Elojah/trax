<script setup lang="ts">
import GroupTable from '@/components/group/Table.vue';
import GroupDetails from '@/components/group/Details.vue';
import RoleTable from '@/components/role/Table.vue';
import UserTable from '@/components/user/Table.vue';
import { ref } from 'vue';

// [Info, Roles, User]
const view = ref(0)

const tableGroups = ref();
</script>

<template>
</template>

<style scoped>
</style>
