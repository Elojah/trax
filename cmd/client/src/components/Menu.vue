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
	<div id="app-sidebar-3"
		class="w-[280px] bg-surface-900 h-screen hidden lg:block flex-shrink-0 absolute lg:static left-0 top-0 z-10 border-r border-surface-800 select-none">
		<div class="flex flex-col h-full">
			<div class="p-4 flex items-center gap-4">
				<svg xmlns="http://www.w3.org/2000/svg" width="43" height="43" viewBox="0 0 43 43" fill="none"
					class="w-10 h-10">
					<path fill-rule="evenodd" clip-rule="evenodd"
						d="M21.5 42.0498C33.098 42.0498 42.5 32.6477 42.5 21.0498C42.5 9.45183 33.098 0.0498047 21.5 0.0498047C9.902 0.0498047 0.5 9.45183 0.5 21.0498C0.5 32.6477 9.902 42.0498 21.5 42.0498ZM28.0513 9.83248C28.3702 8.69975 27.2709 8.02994 26.267 8.74516L12.2528 18.7288C11.164 19.5045 11.3353 21.0498 12.51 21.0498H16.2003V21.0212H23.3926L17.5323 23.089L14.9487 32.2671C14.6299 33.3999 15.729 34.0697 16.733 33.3544L30.7472 23.3708C31.836 22.5951 31.6646 21.0498 30.49 21.0498H24.8937L28.0513 9.83248Z"
						class="fill-surface-0" />
				</svg>
				<span class="text-lg font-semibold leading-tight text-surface-0">ZenTrail</span>
			</div>
			<div class="overflow-y-auto flex-1 p-2 flex flex-col gap-4">
				<ul class="list-none m-0 flex flex-col gap-1">
					<li>
						<div v-styleclass="{
							selector: '@next',
							enterFromClass: 'hidden',
							enterActiveClass: 'animate-slidedown',
							leaveToClass: 'hidden',
							leaveActiveClass: 'animate-slideup'
						}" class="flex items-center cursor-pointer p-3 gap-4 rounded-lg text-surface-0 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150">
							<span class="font-semibold text-base leading-tight">Home</span>
							<i class="pi pi-angle-down !text-base !leading-none text-surface-400 ml-auto" />
						</div>
						<ul class="list-none m-0 overflow-hidden flex flex-col gap-1 mt-1">
							<li>
								<a
									class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
									<i
										class="pi pi-home !text-base !leading-none text-surface-400 hover:text-surface-0" />
									<span class="font-medium text-base leading-tight">Dashboard</span>
								</a>
							</li>
							<li>
								<a
									class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
									<i
										class="pi pi-users !text-base !leading-none text-surface-400 hover:text-surface-0" />
									<span class="font-medium text-base leading-tight">Team</span>
								</a>
							</li>
						</ul>
					</li>
				</ul>

				<hr class="border-t border-surface-800 my-0" />

				<ul class="list-none m-0 flex flex-col gap-1">
					<li>
						<div v-styleclass="{
							selector: '@next',
							enterFromClass: 'hidden',
							enterActiveClass: 'animate-slidedown',
							leaveToClass: 'hidden',
							leaveActiveClass: 'animate-slideup'
						}" class="flex items-center cursor-pointer p-3 gap-4 rounded-lg text-surface-0 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150">
							<span class="font-semibold text-base leading-tight">Your favorites</span>
							<i class="pi pi-angle-down !text-base !leading-none text-surface-400 ml-auto" />
						</div>
						<ul class="list-none m-0 overflow-hidden flex flex-col gap-1 mt-1">
							<li>
								<a
									class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
									<span
										class="inline-flex items-center justify-center p-1 bg-violet-500/30 rounded-md border border-violet-500/30">
										<i class="pi pi-sparkles !text-xs !leading-none text-violet-200" />
									</span>
									<span class="font-medium text-base leading-tight">Usage</span>
								</a>
							</li>
							<li>
								<a
									class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
									<span
										class="inline-flex items-center justify-center p-1 bg-amber-500/30 rounded-md border border-amber-500/30">
										<i class="pi pi-wrench !text-xs !leading-none text-amber-200" />
									</span>
									<span class="font-medium text-base leading-tight">Equipments</span>
								</a>
							</li>
							<li>
								<a
									class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
									<span
										class="inline-flex items-center justify-center p-1 bg-lime-500/30 rounded-md border border-lime-500/30">
										<i class="pi pi-palette !text-xs !leading-none text-lime-200" />
									</span>
									<span class="font-medium text-base leading-tight">Consumption</span>
								</a>
							</li>
						</ul>
					</li>
				</ul>
			</div>
			<div class="p-2 mt-auto border-surface-800 has-[ul.hidden]:border-t">
				<ul class="list-none p-2 m-0 hidden origin-bottom animate-duration-150 border-t border-surface-800">
					<li>
						<a
							class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
							<i class="pi pi-user !text-base !leading-none" />
							<span class="font-medium text-base leading-tight">Profile</span>
						</a>
					</li>
					<li>
						<a
							class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
							<i class="pi pi-cog !text-base !leading-none" />
							<span class="font-medium text-base leading-tight">Settings</span>
						</a>
					</li>
					<li>
						<a
							class="flex items-center cursor-pointer p-3 gap-2 rounded-lg text-surface-400 hover:bg-surface-800 border border-transparent hover:border-surface-700 transition-colors duration-150 hover:text-surface-0">
							<i class="pi pi-sign-out !text-base !leading-none" />
							<span class="font-medium text-base leading-tight">Sign Out</span>
						</a>
					</li>
				</ul>
				<a v-styleclass="{
					selector: '@prev',
					enterFromClass: 'hidden',
					enterActiveClass: 'animate-scalein',
					leaveToClass: 'hidden',
					leaveActiveClass: 'animate-fadeout'
				}" class="flex items-center cursor-pointer p-2 gap-2 text-surface-400 hover:text-surface-0">
					<img src="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
						class="w-8 h-8 rounded-full" />
					<span class="font-medium text-base leading-tight">Amy Elsner</span>
					<i class="pi pi-angle-up !text-base !leading-none text-surface-400 ml-auto" />
				</a>
			</div>
		</div>
	</div>
</template>

<style scoped></style>
