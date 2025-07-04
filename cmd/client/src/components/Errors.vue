<script setup lang="ts">
import { useErrorsStore } from '@/stores/errors';
import { useRouter } from 'vue-router'
import { toRefs } from 'vue';

const router = useRouter()

const errorsStore = useErrorsStore()
const {
	message,
	success,
	error,
} = toRefs(errorsStore)

</script>

<template>
</template>
