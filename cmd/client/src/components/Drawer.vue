<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router'

const router = useRouter()

const authStore = useAuthStore()

const redirect = function (name: string) {
	router.push({ name: name });
}

</script>

<template>
	<v-navigation-drawer floating>
		<v-list nav>
			<v-list-item prepend-icon="mdi-home" title="Home" value="home" @click="redirect('home')">
				<template v-slot:prepend>
					<v-icon color="success"></v-icon>
				</template>
			</v-list-item>
		</v-list>
		<v-divider></v-divider>
		<template v-slot:append v-if="authStore.profile">
			<v-divider></v-divider>
			<v-list-item prepend-icon="mdi-account" title="Profile" value="profile" @click="redirect('profile')">
				<template v-slot:prepend>
					<v-icon color="success"></v-icon>
				</template>
			</v-list-item>
		</template>
	</v-navigation-drawer>
</template>

<style scoped></style>
