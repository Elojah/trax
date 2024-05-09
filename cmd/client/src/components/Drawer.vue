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
	<v-navigation-drawer v-if="authStore.user" expand-on-hover rail>
		<v-list nav>
			<v-list-item prepend-icon="mdi-domain" title="Entities" value="entities" @click="redirect('entity')"
				active-class="active-bg">
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-list-item>
			<v-divider></v-divider>
			<v-list-item prepend-icon="mdi-chart-box-multiple" title="Dashboard" value="dashboard"
				@click="redirect('dashboard')" active-class="active-bg">
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-list-item>
			<v-divider></v-divider>
			<v-list-item v-if="authStore.user" prepend-icon="mdi-map" title="Map" value="map" @click="redirect('map')"
				active-class="active-bg">
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-list-item>
			<v-divider></v-divider>
			<v-list-item v-if="authStore.user" prepend-icon="mdi-library-shelves" title="Stocks" value="stocks"
				@click="redirect('stocks')" active-class="active-bg">
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-list-item>
		</v-list>
		<v-divider thickness="3"></v-divider>
		<template v-slot:append>
			<v-divider thickness="3"></v-divider>
			<v-list nav>
				<v-list-item prepend-icon="mdi-account-circle" title="User" value="user" @click="redirect('user')">
					<template v-slot:prepend>
						<v-icon color="primary"></v-icon>
					</template>
				</v-list-item>
			</v-list>
		</template>
	</v-navigation-drawer>
</template>

<style scoped>
.active-bg {
	background-color: #950c75;
}
</style>
