<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { separator } from '@primeuix/themes/aura/breadcrumb';
import { useRouter } from 'vue-router'

const router = useRouter()

const authStore = useAuthStore()

const redirect = function (name: string) {
	router.push({ name: name });
}

const active = true;

const menu = [
	{
		separator: true
	},
	{
		label: 'Dashboard',
		icon: 'pi pi-fw pi-home',
		command: () => redirect('dashboard')
	},
		{
		separator: true
	},
	{
		label: 'Settings',
		icon: 'pi pi-fw pi-cog',
		items: [
			{
				label: 'Profile',
				icon: 'pi pi-fw pi-user',
				command: () => redirect('profile')
			},
			{
				label: 'Preferences',
				icon: 'pi pi-fw pi-sliders-h',
				command: () => redirect('preferences')
			}
		]
	},
	{
		label: 'Logout',
		icon: 'pi pi-fw pi-sign-out',
		command: () => authStore.signout()
	}
];

</script>

<template>
	<Menu :model="menu">
		<template #start>
			<img src="@/assets/logo.svg" alt="Logo" class="logo" />
		</template>
	</Menu>
</template>

<style scoped>
</style>
