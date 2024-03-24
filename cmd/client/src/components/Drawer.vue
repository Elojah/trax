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
	<v-navigation-drawer rail v-if="authStore.profile">
		<v-list nav>
			<v-list-item prepend-icon="mdi-chart-box-multiple" title="Dashboard" value="dashboard"
				@click="redirect('dashboard')">
				<template v-slot:prepend>
					<v-icon color="success"></v-icon>
				</template>
			</v-list-item>
			<v-divider></v-divider>
			<v-list-item v-if="authStore.profile" prepend-icon="mdi-map" title="Map" value="map"
				@click="redirect('map')">
				<template v-slot:prepend>
					<v-icon color="success"></v-icon>
				</template>
			</v-list-item>
			<v-divider></v-divider>
			<v-list-item v-if="authStore.profile" prepend-icon="mdi-library-shelves" title="Stocks" value="stocks"
				@click="redirect('stocks')">
				<template v-slot:prepend>
					<v-icon color="success"></v-icon>
				</template>
			</v-list-item>
		</v-list>
		<v-divider thickness="3"></v-divider>
		<template v-slot:append>
			<v-divider thickness="3"></v-divider>
			<v-list nav>
				<v-list-item prepend-icon="mdi-account-circle" title="Profile" value="profile"
					@click="redirect('profile')">
					<template v-slot:prepend>
						<v-icon color="success"></v-icon>
					</template>
				</v-list-item>
			</v-list>
		</template>
	</v-navigation-drawer>
</template>

<style scoped></style>
