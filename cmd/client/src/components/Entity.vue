<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);

const authStore = useAuthStore();
const {
	profile: profile,
} = toRefs(authStore);

const valid = ref(null as boolean | null)
</script>

<template>
</template>

<style scoped></style>
