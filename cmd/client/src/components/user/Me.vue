<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { computed, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);

const authStore = useAuthStore();
const {
	user: user,
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
		await authStore.updateUser(undefined, user.value?.lastName);
	}

	lastNameEdit.value = !lastNameEdit.value
};

const updateFirstName = async function () {
	if (firstNameEdit.value) {
		await authStore.updateUser(user.value?.firstName, undefined);
	}

	firstNameEdit.value = !firstNameEdit.value
};

const initials = computed(() => (user?.value?.firstName.at(0)?.toUpperCase() ?? '').concat(user?.value?.lastName.at(0)?.toUpperCase() ?? ''))

</script>

<template>
	<v-card v-if="user" class="px-6 py-6 rounded-lg" variant="elevated">
		<v-card-title class="justify-center mb-2 text-center">
			<v-form v-if="user" ref="form" v-model="valid" lazy-validation>
				<v-avatar class="mb-4" size="96" :color="!user.avatarURL ? 'primary' : ''">
					<img v-if="user.avatarURL" :src="user.avatarURL" alt="Avatar">
					<span v-else-if="!user.avatarURL" class=" mx-auto text-center text-h5">
						{{ initials }}
					</span>
				</v-avatar>
				<v-text-field class="justify-center text-h6" density="compact" v-model="user.lastName"
					:rules="nameRules" :variant="!lastNameEdit ? 'plain' : 'underlined'" :readonly="!lastNameEdit">
					<template v-slot:prepend-inner>
						<v-icon color="primary" size="large" @click="updateLastName"
							:icon="!lastNameEdit ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
					</template>
				</v-text-field>
				<v-text-field class="justify-center text-h6" density="compact" v-model="user.firstName"
					:rules="nameRules" :variant="!firstNameEdit ? 'plain' : 'underlined'" :readonly="!firstNameEdit">
					<template v-slot:prepend-inner>
						<v-icon color="primary" size="large" @click="updateFirstName"
							:icon="!firstNameEdit ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
					</template>
				</v-text-field>
			</v-form>
		</v-card-title>
		<v-card-actions class="justify-center mt-6">
			<v-btn class="px-4" size="large" variant="tonal" append-icon="mdi-export" @click="signout">Signout
				<template v-slot:append>
					<v-icon color="error"></v-icon>
				</template>
			</v-btn>

		</v-card-actions>
	</v-card>
</template>

<style scoped></style>
