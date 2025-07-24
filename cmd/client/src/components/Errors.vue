<script setup lang="ts">
import { useErrorsStore } from '@/stores/errors';
import { useRouter } from 'vue-router'
import { toRefs, watch } from 'vue';
import Toast from 'primevue/toast';
import { useToast } from 'primevue';

const errorsStore = useErrorsStore()
const {
	message,
	success,
	error,
} = toRefs(errorsStore)

const toast = useToast()

watch(success, () => {
	if (success.value) {
		toast.add({
			severity: 'success',
			summary: 'Success',
			detail: message.value,
			life: 3000,
		});

		success.value = false;
	}
}, { immediate: false });

watch(error, () => {
	if (error.value) {
		toast.add({
			severity: 'error',
			summary: 'Error',
			detail: message.value,
			life: 3000,
		});

		error.value = false;
	}
}, { immediate: false });

</script>

<template>
	<Toast />
</template>
