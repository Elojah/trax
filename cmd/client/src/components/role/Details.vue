<script setup lang="ts">
// Vue and Store imports
import { computed, ref, onMounted, toRefs } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useRoleStore } from '@/stores/role';
import { useGroupStore } from '@/stores/group';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { DeleteRoleReq, UpdateRoleReq } from '@internal/user/dto/role';
import { String$ } from '@pkg/pbtypes/string';

// PrimeVue UI Components
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';
import Card from 'primevue/card';
import Skeleton from 'primevue/skeleton';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

// Table Components
import PermissionTable from '@/components/permission/Table.vue';
import RoleUserTable from '@/components/role/UserTable.vue';

// Form Validation
import { zodResolver } from '@primevue/forms/resolvers/zod';
import z from 'zod';
import { Form, FormField } from '@primevue/forms';
import type { FormSubmitEvent } from '@primevue/forms';
import { useConfirm } from 'primevue/useconfirm';

const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();
const roleStore = useRoleStore();
const groupStore = useGroupStore();
const errorsStore = useErrorsStore();
const confirm = useConfirm();

const { roles } = roleStore;
const { groups } = groupStore;
const { success, message } = toRefs(errorsStore);

const loading = ref(true);
const roleId = computed(() => route.params.roleId as string);
const groupId = computed(() => route.params.groupId as string);
const activeTab = ref("0");

// Dialog states
const dialogManageRole = ref(false);
const dialogEditPermissions = ref(false);

// Form validation
const resolver = zodResolver(
	z.object({
		name: z.string().min(1, { message: 'Name is required.' }),
	})
);

const initialValues = ref({
	name: '',
});

// Reference to permission table for editing
const permissionTableRef = ref();

const role = computed(() => {
	if (!roleId.value) return null;
	return roles.get(roleId.value) || null;
});

const group = computed(() => {
	if (!groupId.value) return null;
	return groups.get(groupId.value) || null;
});


const loadRole = async () => {
	if (!roleId.value || !groupId.value) {
		router.push({ name: 'group-details', params: { id: groupId.value } });
		return;
	}

	loading.value = true;

	try {
		// Load group data
		await groupStore.populate([groupId.value]);
	} catch (e) {
		errorsStore.showGRPC(e);
		router.push({ name: 'group-details', params: { id: groupId.value } });
	}

	try {
		// Load role data
		await roleStore.populate([roleId.value], [groupId.value]);
	} catch (e) {
		errorsStore.showGRPC(e);
		router.push({ name: 'group-details', params: { id: groupId.value } });
		return;
	}

	loading.value = false;
};

const goBack = () => {
	router.go(-1);
};

// Action handlers
const openManageRole = () => {
	if (!role.value?.role) return;

	initialValues.value = {
		name: role.value.role.name || '',
	};
	dialogManageRole.value = true;
};

const openEditPermissions = () => {
	if (!role.value?.role) return;
	dialogEditPermissions.value = true;
};

const updatePermissions = async () => {
	if (!role.value?.role || !permissionTableRef.value) return;

	try {
		const selectedPermissions = permissionTableRef.value.selected();
		await roleStore.update(UpdateRoleReq.create({
			iD: role.value.role.iD,
			permissions: selectedPermissions,
		}));
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `Role permissions updated successfully`;
	success.value = true;
	dialogEditPermissions.value = false;
	await authStore.refreshToken();
	await loadRole();
};

const update = async (e: FormSubmitEvent) => {
	if (!e.valid || !role.value?.role) {
		logger.error('Update role form is invalid', e);
		return;
	}

	try {
		await roleStore.update(UpdateRoleReq.create({
			iD: role.value.role.iD,
			name: String$.create({ value: e.states.name.value }),
		}));
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `Role ${e.states.name.value} updated successfully`;
	success.value = true;
	dialogManageRole.value = false;
	e.reset();
	await authStore.refreshToken();
	await loadRole();
};

const deleteRoleAction = () => {
	if (!role.value?.role) return;

	confirm.require({
		message: `Are you sure you want to delete the role "${role.value.role.name}"? This action cannot be undone.`,
		header: 'Delete Role',
		icon: 'pi pi-exclamation-triangle',
		rejectProps: {
			label: 'Cancel',
			severity: 'secondary',
			outlined: true
		},
		acceptProps: {
			label: 'Delete',
			severity: 'danger'
		},
		accept: async () => {
			if (!role.value?.role) return;

			const name = role.value?.role?.name

			try {
				await roleStore.delete_(DeleteRoleReq.create({ iD: role.value.role.iD }));
			} catch (e) {
				errorsStore.showGRPC(e);
				return;
			}

			message.value = `Role "${name}" deleted successfully`;
			success.value = true;
			await authStore.refreshToken();
			// Navigate back to group after deletion
			router.push({ name: 'group-details', params: { id: groupId.value } });
		}
	});
};

const formatDate = (timestamp: bigint | undefined): string => {
	if (!timestamp) return 'Unknown';

	return new Date(Number(timestamp) * 1000).toLocaleDateString('en-GB', {
		day: 'numeric',
		month: 'long',
		year: 'numeric',
		hour: '2-digit',
		minute: '2-digit'
	});
};

// Initialize data on component mount
onMounted(() => {
	loadRole();
});
</script>

<template>
	<div class="min-h-screen flex flex-col relative flex-auto">
		<!-- Header -->
		<div
			class="flex justify-between items-center px-8 py-4 bg-surface-0 dark:bg-surface-950 border-b border-surface-200 dark:border-surface-700">
			<div class="flex items-center gap-4">
				<Button icon="pi pi-arrow-left" severity="secondary" text rounded class="w-10 h-10" @click="goBack"
					v-tooltip.bottom="'Back to group'" />
				<div class="flex items-center gap-3">
					<div
						class="flex items-center justify-center w-10 h-10 rounded-lg bg-purple-100 dark:bg-purple-400/30 border border-purple-200 dark:border-purple-400/20">
						<i class="pi pi-shield text-lg text-purple-600 dark:text-purple-300"></i>
					</div>
					<div class="flex flex-col">
						<h2 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">
							{{ role?.role?.name || 'Loading...' }}
						</h2>
						<span class="text-sm text-surface-500 dark:text-surface-400">Role management</span>
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

				<!-- Role Not Found -->
				<Card v-else-if="!role" class="mb-6">
					<template #content>
						<div class="text-center p-8">
							<i class="pi pi-exclamation-circle text-4xl text-surface-400 mb-4"></i>
							<h3 class="text-xl font-semibold text-surface-900 dark:text-surface-0 mb-2">Role Not Found
							</h3>
							<p class="text-surface-500 dark:text-surface-400 mb-4">
								The role you're looking for doesn't exist or you don't have permission to view it.
							</p>
							<Button label="Back to Group" icon="pi pi-arrow-left" @click="goBack" />
						</div>
					</template>
				</Card>

				<!-- Role Content with Tabs -->
				<Tabs v-else v-model:value="activeTab" class="w-full">
					<TabList>
						<Tab value="0">
							<div class="flex items-center gap-2">
								<i class="pi pi-info-circle"></i>
								<span>Details</span>
							</div>
						</Tab>
						<Tab value="1">
							<div class="flex items-center gap-2">
								<i class="pi pi-users"></i>
								<span>Users</span>
							</div>
						</Tab>
					</TabList>
					<TabPanels>
						<!-- Details Tab -->
						<TabPanel value="0">
							<Card class="shadow-lg border-0">
								<template #content>
									<div class="p-8">
										<!-- Hero Section with Edit Button -->
										<div class="flex flex-col md:flex-row items-start gap-8 mb-8">
											<!-- Avatar -->
											<div class="flex-shrink-0 relative">
												<Avatar :label="role.role?.name.charAt(0).toUpperCase()" size="xlarge"
													shape="circle"
													class="bg-gradient-to-br from-purple-400 to-purple-600 text-white border-4 border-surface-200 dark:border-surface-700 shadow-lg text-4xl font-bold" />
											</div>

											<!-- Role Info -->
											<div class="flex-1 min-w-0">
												<div
													class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
													<div>
														<h1
															class="text-4xl font-bold text-surface-900 dark:text-surface-0 mb-2">
															{{ role.role?.name }}
														</h1>
														<div
															class="flex items-center gap-2 text-surface-500 dark:text-surface-400">
															<i class="pi pi-calendar text-sm"></i>
															<span class="text-sm">
																Created {{ formatDate(role.role?.createdAt) }}
															</span>
														</div>
													</div>
													<div class="flex items-center gap-3">
														<Button label="" icon="pi pi-cog" outlined severity="primary"
															class="font-medium" @click="openManageRole"
															v-tooltip.bottom="'Edit role settings'" />
													</div>
												</div>

												<!-- Statistics Cards -->
												<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-6 mb-8">
													<div
														class="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-900/20 dark:to-purple-800/20 rounded-xl p-6 border border-purple-200 dark:border-purple-700/50">
														<div class="flex items-center gap-3">
															<div
																class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
																<i class="pi pi-key text-white text-xl"></i>
															</div>
															<div>
																<p
																	class="text-2xl font-bold text-purple-700 dark:text-purple-300">
																	{{ role.permissions?.length || 0 }}
																</p>
																<p
																	class="text-purple-600 dark:text-purple-400 text-sm font-medium">
																	Permission(s)</p>
															</div>
														</div>
													</div>

													<div
														class="bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-900/20 dark:to-blue-800/20 rounded-xl p-6 border border-blue-200 dark:border-blue-700/50">
														<div class="flex items-center gap-3">
															<div
																class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
																<i class="pi pi-building text-white text-xl"></i>
															</div>
															<div>
																<p
																	class="text-2xl font-bold text-blue-700 dark:text-blue-300">
																	{{ group?.group?.name || 'Unknown' }}
																</p>
																<p
																	class="text-blue-600 dark:text-blue-400 text-sm font-medium">
																	Group</p>
															</div>
														</div>
													</div>
												</div>

												<!-- Permissions Section -->
												<div class="mb-8">
													<div
														class="bg-surface-50 dark:bg-surface-800 rounded-xl p-6 border border-surface-200 dark:border-surface-700">
														<div class="flex justify-between items-center mb-4">
															<h3
																class="text-lg font-semibold text-surface-900 dark:text-surface-0 flex items-center gap-2">
																<i class="pi pi-key text-primary-500"></i>
																Permissions
															</h3>
															<Button label="" icon="pi pi-pencil" outlined
																severity="primary" class="font-medium"
																@click="openEditPermissions"
																v-tooltip.bottom="'Edit permissions'" />
														</div>
														<PermissionTable :disabled="true"
															:permissions="role.permissions" />
													</div>
												</div>

												<!-- Metadata -->
												<div
													class="bg-surface-50 dark:bg-surface-800 rounded-xl p-6 border border-surface-200 dark:border-surface-700">
													<h3
														class="text-lg font-semibold text-surface-900 dark:text-surface-0 mb-4 flex items-center gap-2">
														<i class="pi pi-info text-primary-500"></i>
														Information
													</h3>
													<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
														<div class="flex flex-col gap-2">
															<h4
																class="text-sm font-semibold text-surface-900 dark:text-surface-0 uppercase tracking-wide">
																Created
															</h4>
															<div class="flex items-center gap-2">
																<i
																	class="pi pi-calendar text-surface-500 dark:text-surface-400"></i>
																<span class="text-surface-700 dark:text-surface-200">
																	{{ formatDate(role.role?.createdAt) }}
																</span>
															</div>
														</div>

														<div class="flex flex-col gap-2"
															v-if="role.role?.updatedAt && role.role?.updatedAt !== role.role?.createdAt">
															<h4
																class="text-sm font-semibold text-surface-900 dark:text-surface-0 uppercase tracking-wide">
																Last Updated
															</h4>
															<div class="flex items-center gap-2">
																<i
																	class="pi pi-clock text-surface-500 dark:text-surface-400"></i>
																<span class="text-surface-700 dark:text-surface-200">
																	{{ formatDate(role.role?.updatedAt) }}
																</span>
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

						<!-- Users Tab -->
						<TabPanel value="1">
							<RoleUserTable :role-id="roleId" :group-id="groupId" />
						</TabPanel>
					</TabPanels>
				</Tabs>
			</div>
		</div>

		<!-- Manage Role Dialog -->
		<Dialog v-model:visible="dialogManageRole" append-to="body" modal
			:breakpoints="{ '960px': '75vw', '640px': '80vw' }" :style="{ width: '40rem' }" :draggable="false"
			:resizable="false" :show-header="false" class="shadow-sm rounded-2xl"
			:pt="{ content: '!p-6', footer: '!pb-6 !px-6' }">
			<div class="flex justify-between items-start gap-4">
				<div class="flex gap-4">
					<div
						class="flex items-center justify-center w-9 h-9 bg-purple-100 dark:bg-purple-400/30 text-purple-600 dark:text-purple-300 rounded-3xl border border-purple-200 dark:border-purple-400/20 shrink-0">
						<i class="pi pi-shield !text-xl !leading-none" />
					</div>
				</div>
				<Button icon="pi pi-times" text rounded severity="secondary" class="w-10 h-10 !p-2"
					@click="dialogManageRole = false" />
			</div>
			<Form v-slot="$form" :resolver="resolver" :initialValues="initialValues" @submit="update"
				class="flex flex-col gap-6">
				<div class="flex items-start gap-4">
					<div class="flex-1 flex flex-col gap-2">
						<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Manage
							role
						</h1>
						<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">Edit role details
							and permissions.</span>
					</div>
				</div>
				<div class="flex flex-col gap-6">
					<FormField v-slot="$field" name="name" class="flex flex-col gap-2">
						<label for="manage-role-name" class="text-color text-base">Name</label>
						<IconField icon-position="left" class="w-full">
							<InputIcon class="pi pi-shield" />
							<InputText id="manage-role-name" name="name" placeholder="Enter role name" type="text"
								class="w-full" />
						</IconField>
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message
						}}
						</Message>
					</FormField>
				</div>
				<div class="flex justify-between items-center">
					<Button label="Delete Role" icon="pi pi-trash" severity="danger" outlined class="font-medium"
						@click="deleteRoleAction" />
					<div class="flex gap-4">
						<Button label="Cancel" outlined @click="dialogManageRole = false" />
						<Button label="Update" type="submit" />
					</div>
				</div>
			</Form>
		</Dialog>

		<!-- Edit Permissions Dialog -->
		<Dialog v-model:visible="dialogEditPermissions" append-to="body" modal
			:breakpoints="{ '960px': '75vw', '640px': '90vw' }" :style="{ width: '50rem' }" :draggable="false"
			:resizable="false" :show-header="false" class="shadow-sm rounded-2xl"
			:pt="{ content: '!p-6', footer: '!pb-6 !px-6' }">
			<div class="flex justify-between items-start gap-4">
				<div class="flex gap-4">
					<div
						class="flex items-center justify-center w-9 h-9 bg-purple-100 dark:bg-purple-400/30 text-purple-600 dark:text-purple-300 rounded-3xl border border-purple-200 dark:border-purple-400/20 shrink-0">
						<i class="pi pi-key !text-xl !leading-none" />
					</div>
				</div>
				<Button icon="pi pi-times" text rounded severity="secondary" class="w-10 h-10 !p-2"
					@click="dialogEditPermissions = false" />
			</div>
			<div class="flex flex-col gap-6">
				<div class="flex items-start gap-4">
					<div class="flex-1 flex flex-col gap-2">
						<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Edit
							permissions
						</h1>
						<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">Modify role
							permissions
							for {{ role?.role?.name }}.</span>
					</div>
				</div>
				<div class="flex flex-col gap-6">
					<PermissionTable :disabled="false" :permissions="role?.permissions" ref="permissionTableRef" />
				</div>
				<div class="flex justify-end gap-4">
					<Button label="Cancel" outlined @click="dialogEditPermissions = false" />
					<Button label="Update Permissions" @click="updatePermissions" />
				</div>
			</div>
		</Dialog>
	</div>
</template>

<style scoped></style>
