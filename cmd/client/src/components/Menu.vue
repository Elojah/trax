<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router'

const router = useRouter()

const authStore = useAuthStore()

const redirect = function (name: string) {
	router.push({ name: name });
}

const active = true;

const menu = [
	{
		label: 'Dashboard',
		icon: 'pi pi-fw pi-home',
		command: () => redirect('dashboard')
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
    <aside class="layout-sidebar" :class="{ active: active }">
        <nav>
            <ol class="layout-menu">
                <AppMenuItem :menu="menu"></AppMenuItem>
            </ol>
        </nav>
    </aside>
</template>

<style scoped>
</style>
