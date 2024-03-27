<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { ref } from 'vue';

const authStore = useAuthStore();
const profile = authStore.profile;

const signout = () => {
	authStore.signout();
	router.push({ name: 'sign' });
};

const firstNameEdit = ref(false);
const lastNameEdit = ref(false);

</script>

<template>
	<v-card v-if="profile" class="px-6 py-6 rounded-xl" variant="elevated">
		<v-card-title class="justify-center mb-2 text-center">
			<v-avatar class="mb-4" size="96">
				<img :src="profile?.avatarURL" alt="Avatar">
			</v-avatar>
			<v-col cols="12" class="pa-0">
				<v-text-field v-model="profile.lastName" variant="solo" append-icon="mdi-pencil-circle"
					@click:append="firstNameEdit = !firstNameEdit" :readonly="lastNameEdit"></v-text-field>
			</v-col>
			<v-col cols="12" class="pa-0">
				<v-text-field v-model="profile.firstName" variant="solo" append-icon="mdi-pencil-circle"
					@click:append="lastNameEdit = !lastNameEdit" :readonly="firstNameEdit"></v-text-field>
			</v-col>
		</v-card-title>
		<v-card-actions class="justify-center mt-6">
			<v-btn size="large" variant="tonal" append-icon="mdi-account-circle" @click="signout">Signout
				<template v-slot:append>
					<v-icon color="error"></v-icon>
				</template>
			</v-btn>

		</v-card-actions>
	</v-card>
</template>

<style scoped></style>
