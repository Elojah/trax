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
	<v-snackbar :timeout="2000" v-model="success" class="mt-6" color="primary">
		{{ message }}
		<template v-slot:actions>
			<v-btn variant="text" @click="success = false">Close</v-btn>
		</template>
	</v-snackbar>
	<v-snackbar :timeout="2000" v-model="error" class="mt-6" color="error">
		{{ message }}
		<template v-slot:actions>
			<v-btn variant="text" @click="error = false">Close</v-btn>
		</template>
	</v-snackbar>
</template>
