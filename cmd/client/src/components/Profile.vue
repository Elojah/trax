<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);

const authStore = useAuthStore();
const {
	profile: profile,
} = toRefs(authStore);

const signout = () => {
	authStore.signout();
	router.push({ name: 'sign' });
};

const valid = ref(null as boolean | null)

const firstNameEdit = ref(false);
const lastNameEdit = ref(false);

const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const updateLastName = async function () {
	if (lastNameEdit.value) {
		await authStore.updateProfile();
	}

	lastNameEdit.value = !lastNameEdit.value
};

const updateFirstName = async function () {
	if (firstNameEdit.value) {
		await authStore.updateProfile();
	}

	firstNameEdit.value = !firstNameEdit.value
};

</script>

<template>
	<v-card v-if="profile" class="px-6 py-6 rounded-xl" variant="elevated">
		<v-card-title class="justify-center mb-2 text-center">
			<v-form v-if="profile" ref="form" v-model="valid" lazy-validation>
				<v-avatar class="mb-4" size="96">
					<img :src="profile.avatarURL" alt="Avatar">
				</v-avatar>
				<v-col cols="12" class="pa-0">
					<v-text-field v-model="profile.lastName" :rules="nameRules"
						:variant="!lastNameEdit ? 'solo' : 'solo-filled'" :readonly="!lastNameEdit">
						<template v-slot:append-inner>
							<v-icon color="success" large @click="updateLastName"
								:icon="!lastNameEdit ? 'mdi-pencil-circle-outline' : 'mdi-check-circle-outline'"></v-icon>
						</template>
					</v-text-field>
				</v-col>
				<v-col cols="12" class="pa-0">
					<v-text-field v-model="profile.firstName" :rules="nameRules"
						:variant="!firstNameEdit ? 'solo' : 'solo-filled'" :readonly="!firstNameEdit">
						<template v-slot:append-inner>
							<v-icon color="success" large @click="updateFirstName"
								:icon="!firstNameEdit ? 'mdi-pencil-circle-outline' : 'mdi-check-circle-outline'"></v-icon>
						</template>
					</v-text-field>
				</v-col>
			</v-form>
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
