<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useRouter, useRoute } from 'vue-router'
import { toRefs, ref, computed } from 'vue';
import Avatar from 'primevue/avatar';

const router = useRouter()
const route = useRoute()

const authStore = useAuthStore();
const {
	user: user,
} = toRefs(authStore);
const redirect = function (name: string) {
	router.push({ name: name });
}

const selectedItem = computed(() => route.name?.toString() || 'home');

interface MenuItem {
	label: string;
	icon: string;
	command?: () => void;
}

const selectItem = (item: MenuItem) => {
	if (item.command) {
		item.command();
	}
};

const menuItems = ref([
	{
		label: 'home',
		icon: 'pi pi-fw pi-home',
		command: () => redirect('home')
	},
	{
		label: 'map',
		icon: 'pi pi-map',
		command: () => redirect('map')
	},
	{
		label: 'group',
		icon: 'pi pi-building',
		command: () => redirect('group')
	},
	{
		label: 'settings',
		icon: 'pi pi-fw pi-cog',
	},
]);

</script>

<template>
	<div v-if="user" id="app-sidebar-9"
		class="h-screen bg-surface-0 dark:bg-surface-950 hidden lg:block flex-shrink-0 absolute lg:static left-0 top-0 z-10 border-r border-surface w-24 lg:w-24 w-72 select-none">
		<div class="flex flex-col h-full">
			<div class="flex items-center justify-center p-4 flex-shrink-0 bg-surface-0 dark:bg-surface-900">
				<svg xmlns="http://www.w3.org/2000/svg" width="33" height="32" viewBox="0 0 33 32" fill="none">
					<path fill-rule="evenodd" clip-rule="evenodd"
						d="M7.34219 2.87829C6.19766 3.67858 5.1627 4.62478 4.26426 5.68992C7.9357 5.34906 12.6001 5.90564 18.0155 8.61335C23.7984 11.5047 28.455 11.6025 31.6958 10.9773C31.4017 10.087 31.0315 9.23135 30.593 8.41791C26.8832 8.80919 22.1272 8.29127 16.5845 5.51998C13.0648 3.76014 9.96221 3.03521 7.34219 2.87829ZM28.4259 5.33332C25.4962 2.06 21.2387 0 16.5 0C15.1084 0 13.7581 0.177686 12.4709 0.511584C14.2143 0.987269 16.0663 1.68319 18.0155 2.65781C22.0736 4.68682 25.5771 5.34013 28.4259 5.33332ZM32.3887 14.1025C28.4735 14.8756 23.067 14.7168 16.5845 11.4755C10.524 8.44527 5.70035 8.48343 2.44712 9.20639C2.2792 9.24367 2.11523 9.28287 1.95522 9.32367C1.5293 10.25 1.18931 11.2241 0.945362 12.2356C1.20591 12.166 1.47514 12.0998 1.75293 12.0381C5.69966 11.161 11.2761 11.1991 18.0155 14.5689C24.0761 17.5991 28.8997 17.561 32.1529 16.838C32.2644 16.8133 32.3742 16.7877 32.4822 16.7613C32.4941 16.509 32.5 16.2552 32.5 16C32.5 15.358 32.4622 14.7248 32.3887 14.1025ZM31.9598 20.1378C28.0826 20.8157 22.8336 20.5555 16.5845 17.431C10.524 14.4008 5.70035 14.439 2.44712 15.1619C1.725 15.3223 1.07539 15.5178 0.502344 15.7241C0.500782 15.8158 0.5 15.9078 0.5 16C0.5 24.8366 7.66344 32 16.5 32C23.9057 32 30.1362 26.9687 31.9598 20.1378Z"
						class="fill-surface-900 dark:fill-surface-0" />
				</svg>
			</div>
			<div class="flex-1 py-4 flex flex-col">
				<div v-for="(item) in menuItems" :key="item.label"
					class="flex items-center cursor-pointer group lg:flex hidden" @click="selectItem(item)">
					<div class="w-1 h-full" :class="selectedItem === item.label ? 'bg-primary' : 'bg-transparent'" />
					<div class="flex-1 flex flex-col items-center py-3 gap-2 transition-colors duration-150">
						<i
							:class="[item.icon, selectedItem === item.label ? 'text-primary' : 'text-surface-600 dark:text-surface-300 group-hover:text-primary', '!text-2xl !leading-none']" />
						<span class="text-sm font-medium text-center"
							:class="selectedItem === item.label ? 'text-primary' : 'text-surface-600 dark:text-surface-300 group-hover:text-primary'">{{
								item.label.charAt(0).toUpperCase() + item.label.slice(1) }}</span>
					</div>
					<div class="w-1 h-full bg-transparent" />
				</div>

				<div class="lg:hidden flex flex-col gap-2 mt-4">
					<a v-for="(item) in menuItems" :key="item.label"
						class="w-full flex items-center cursor-pointer p-3 gap-3"
						:class="[selectedItem === item.label ? 'text-primary  border-l-4 border-primary' : 'text-surface-700 dark:text-surface-300 hover:text-primary border-l-4 border-transparent']"
						@click="selectItem(item)">
						<i
							:class="[item.icon, '!text-lg !leading-none', selectedItem === item.label ? 'text-primary' : 'text-surface-600 dark:text-surface-300']" />
						<span class="font-medium text-base leading-tight">{{ item.label.charAt(0).toUpperCase() +
							item.label.slice(1) }}</span>
					</a>
				</div>
			</div>
			<div class="py-4 border-t border-surface flex justify-center items-center">
				<Avatar :image="user?.avatarURL" class="!w-10 !h-10 cursor-pointer" shape="circle"
					@click="selectItem({ label: 'user', icon: 'pi pi-fw pi-user', command: () => redirect('user') })" />
			</div>
		</div>
	</div>
</template>

<style scoped></style>
