<script setup lang="ts">
// Vue and Store imports
import { computed, ref, onMounted, toRefs } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useErrorsStore } from '@/stores/errors';
import { useUserStore } from '@/stores/user';
import { useGroupStore } from '@/stores/group';

// Internal utilities and types
import { formatDate } from '@/utils/date';
import { back } from '@/utils/router';
import type { U } from '@internal/user/user';

// PrimeVue UI Components
import Button from 'primevue/button';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';
import Card from 'primevue/card';
import Skeleton from 'primevue/skeleton';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';

// Table Components
import UserRoleTable from '@/components/user/RoleTable.vue';
import { useAuthStore } from '@/stores/auth';
import { hasClaim } from '@/utils/claims';
import { Command, Resource } from '@internal/user/role';
import { parse } from '@/utils/ulid';

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();
const groupStore = useGroupStore();
const errorsStore = useErrorsStore();
const authStore = useAuthStore();

const { users } = toRefs(userStore);
const { claims } = toRefs(authStore);
const { success, message } = toRefs(errorsStore);

const loading = ref(true);
const userId = route.params.userId as string
const groupId = route.params.groupId as string
const activeTab = ref("0");

const user = computed(() => {
	return users.value.get(userId) || null;
});

const readUserPermission = computed(() => {
	return hasClaim(claims.value, parse(groupId), Resource.R_user, Command.C_read);
});

const readRolePermission = computed(() => {
	return hasClaim(claims.value, parse(groupId), Resource.R_role, Command.C_read);
});

const loadUser = async () => {
	if (!userId || !groupId) {
		router.push({ name: 'group-details', params: { id: groupId } });
		return;
	}

	loading.value = true;

	try {
		// Load group data
		await groupStore.populate([groupId]);
	} catch (e) {
		errorsStore.showGRPC(e);
		router.push({ name: 'group-details', params: { id: groupId } });
	}

	try {
		await userStore.populate([userId], [groupId]);
	} catch (e) {
		errorsStore.showGRPC(e);
		router.push({ name: 'group-details', params: { id: groupId } });
	}

	loading.value = false;
};

const getUserInitials = (user: U): string => {
	if (user.firstName && user.lastName) {
		return (user.firstName.charAt(0) + user.lastName.charAt(0)).toUpperCase();
	}
	return user.email.charAt(0).toUpperCase();
};

const getUserDisplayName = (user: U): string => {
	if (user.firstName && user.lastName) {
		return `${user.firstName} ${user.lastName}`;
	}
	return user.email;
};

// Initialize data on component mount
onMounted(async () => {
	await loadUser();
});
</script>

<template>
	<div class="min-h-screen flex flex-col relative flex-auto">
		<!-- Header -->
		<div
			class="flex justify-between items-center px-8 py-4 bg-surface-0 dark:bg-surface-950 border-b border-surface-200 dark:border-surface-700">
			<div class="flex items-center gap-4">
				<Button icon="pi pi-arrow-left" severity="secondary" text rounded class="w-10 h-10" @click="back"
					v-tooltip.bottom="'Back'" />
				<div class="flex items-center gap-3">
					<div
						class="flex items-center justify-center w-10 h-10 rounded-lg bg-blue-100 dark:bg-blue-400/30 border border-blue-200 dark:border-blue-400/20">
						<i class="pi pi-user text-lg text-blue-600 dark:text-blue-300"></i>
					</div>
					<div class="flex flex-col">
						<h2 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">
							{{ user ? getUserDisplayName(user) : 'Loading...' }}
						</h2>
						<span class="text-sm text-surface-500 dark:text-surface-400">User management</span>
					</div>
				</div>
			</div>
		</div>

		<!-- Content -->
		<div class="flex flex-col flex-auto">
			<div class="mx-auto w-full">
				<!-- Success Message -->
				<Message v-if="success && message" severity="success" class="mb-4">{{ message }}</Message>

				<!-- Loading State -->
				<Card v-if="loading" class="mb-6">
					<template #content>
						<div class="flex items-center gap-6 p-6">
							<Skeleton shape="circle" size="5rem" />
							<div class="flex-1">
								<Skeleton width="12rem" height="1.5rem" class="mb-2" />
								<Skeleton width="20rem" height="1rem" class="mb-4" />
								<Skeleton width="100%" height="4rem" />
							</div>
						</div>
					</template>
				</Card>

				<!-- User Not Found -->
				<Card v-else-if="!user" class="mb-6">
					<template #content>
						<div class="text-center p-8">
							<i class="pi pi-exclamation-circle text-4xl text-surface-400 mb-4"></i>
							<h3 class="text-xl font-semibold text-surface-900 dark:text-surface-0 mb-2">User Not Found
							</h3>
							<p class="text-surface-500 dark:text-surface-400 mb-4">
								The user you're looking for doesn't exist or you don't have permission to view it.
							</p>
							<Button label="Back" icon="pi pi-arrow-left" @click="back" />
						</div>
					</template>
				</Card>

				<!-- User Content with Tabs -->
				<Tabs v-else v-model:value="activeTab" class="w-full">
					<TabList>
						<Tab :disabled="!readUserPermission" value="0">
							<div class="flex items-center gap-2">
								<i class="pi pi-info-circle"></i>
								<span>Details</span>
							</div>
						</Tab>
						<Tab :disabled="!readRolePermission" value="1">
							<div class="flex items-center gap-2">
								<i class="pi pi-shield"></i>
								<span>Roles</span>
							</div>
						</Tab>
					</TabList>
					<TabPanels>
						<!-- Details Tab -->
						<TabPanel value="0">
							<Card v-if="readUserPermission" class="shadow-lg border-0">
								<template #content>
									<div class="p-8">
										<!-- Hero Section with Edit Button -->
										<div class="flex flex-col md:flex-row items-start gap-8 mb-8">
											<!-- Avatar -->
											<div class="flex-shrink-0 relative">
												<Avatar v-if="user.avatarURL" :image="user.avatarURL" size="xlarge"
													shape="circle"
													class="border-4 border-surface-200 dark:border-surface-700 shadow-lg" />
												<Avatar v-else :label="getUserInitials(user)" size="xlarge"
													shape="circle"
													class="bg-gradient-to-br from-blue-400 to-blue-600 text-white border-4 border-surface-200 dark:border-surface-700 shadow-lg text-4xl font-bold" />
											</div>

											<!-- User Info -->
											<div class="flex-1 min-w-0">
												<div
													class="flex flex-col md:flex-row md:items-start justify-between gap-4 mb-6">
													<div>
														<h1
															class="text-3xl font-bold text-surface-900 dark:text-surface-0 mb-2">
															{{ getUserDisplayName(user) }}
														</h1>
														<p class="text-lg text-surface-600 dark:text-surface-300 mb-4">
															{{ user.email }}
														</p>
														<div
															class="flex flex-wrap gap-4 text-sm text-surface-500 dark:text-surface-400">
															<div class="flex items-center gap-2">
																<i class="pi pi-calendar"></i>
																<span>Joined {{ formatDate(user.createdAt) }}</span>
															</div>
														</div>
													</div>
												</div>

												<!-- Details Grid -->
												<div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
													<div class="bg-surface-50 dark:bg-surface-800 rounded-lg p-4">
														<h3
															class="font-semibold text-surface-900 dark:text-surface-0 mb-2 flex items-center gap-2">
															<i class="pi pi-id-card text-surface-500"></i>
															Personal Information
														</h3>
														<div class="space-y-2 text-sm">
															<div class="flex justify-between">
																<span
																	class="text-surface-500 dark:text-surface-400">First
																	Name:</span>
																<span class="text-surface-900 dark:text-surface-0">{{
																	user.firstName
																	|| 'Not provided' }}</span>
															</div>
															<div class="flex justify-between">
																<span
																	class="text-surface-500 dark:text-surface-400">Last
																	Name:</span>
																<span class="text-surface-900 dark:text-surface-0">{{
																	user.lastName
																	|| 'Not provided' }}</span>
															</div>
															<div class="flex justify-between">
																<span
																	class="text-surface-500 dark:text-surface-400">Email:</span>
																<span class="text-surface-900 dark:text-surface-0">{{
																	user.email
																}}</span>
															</div>
														</div>
													</div>

													<div class="bg-surface-50 dark:bg-surface-800 rounded-lg p-4">
														<h3
															class="font-semibold text-surface-900 dark:text-surface-0 mb-2 flex items-center gap-2">
															<i class="pi pi-clock text-surface-500"></i>
															Account Information
														</h3>
														<div class="space-y-2 text-sm">
															<div class="flex justify-between">
																<span
																	class="text-surface-500 dark:text-surface-400">Created:</span>
																<span class="text-surface-900 dark:text-surface-0">{{
																	formatDate(user.createdAt) }}</span>
															</div>
															<div class="flex justify-between">
																<span
																	class="text-surface-500 dark:text-surface-400">Updated:</span>
																<span class="text-surface-900 dark:text-surface-0">{{
																	formatDate(user.updatedAt) }}</span>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
									</div>
								</template>
							</Card>
						</TabPanel>

						<!-- Roles Tab -->
						<TabPanel value="1">
							<UserRoleTable v-if="readRolePermission" :user-id="userId" :group-id="groupId" />
						</TabPanel>
					</TabPanels>
				</Tabs>
			</div>
		</div>
	</div>
</template>

<style scoped></style>
