<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import router from '@/router';

const authStore = useAuthStore();
const profile = authStore.profile;

const signout = () => {
	authStore.signout();
	router.push({ name: 'sign' });
};
</script>

<template>
	<v-card v-if="profile" class="px-6 py-6 rounded-xl" variant="elevated">
		<v-card-title class="justify-center mb-2 text-center">
			<v-avatar size="96">
				<img :src="profile?.avatarURL" alt="Avatar">
			</v-avatar>
			<v-list-item :title="profile.firstName + ' ' + profile.lastName"></v-list-item>
		</v-card-title>
		<v-card-actions class="justify-center">
			<v-btn size="large" variant="tonal" append-icon="mdi-account-circle" @click="signout">Signout
				<template v-slot:append>
					<v-icon color="error"></v-icon>
				</template>
			</v-btn>

		</v-card-actions>
	</v-card>
</template>

<style scoped></style>
