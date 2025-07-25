<script setup lang="ts">
// Vue and Store imports
import { computed, ref, onMounted, toRefs } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { DeleteGroupReq, UpdateGroupReq } from '@internal/user/dto/group';
import { Group } from '@internal/user/group';

// PrimeVue UI Components
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';
import Card from 'primevue/card';
import Skeleton from 'primevue/skeleton';
import Textarea from 'primevue/textarea';
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
import RoleTable from './RoleTable.vue';
import UserTable from './UserTable.vue';

// Form Validation
import { zodResolver } from '@primevue/forms/resolvers/zod';
import z from 'zod';
import { Form, FormField } from '@primevue/forms';
import type { FormSubmitEvent } from '@primevue/forms';
import { useConfirm } from 'primevue/useconfirm';

const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();
const store = useGroupStore();
const errorsStore = useErrorsStore();
const confirm = useConfirm();

const { groups } = store;
const { success, message } = toRefs(errorsStore);

const loading = ref(true);
const groupId = computed(() => route.params.id as string);
const activeTab = ref("0");

// Dialog states
const dialogManageGroup = ref(false);

// Form validation
const resolver = zodResolver(
	z.object({
		name: z.string().min(1, { message: 'Name is required.' }),
		"avatar-url": z.string().url({ message: 'Avatar URL must be a valid URL.' }).optional(),
		description: z.string().max(256, { message: 'Description must be at most 256 characters long.' }).optional(),
	})
);

const initialValues = ref({
	name: '',
	"avatar-url": '',
	description: '',
});

const group = computed(() => {
	if (!groupId.value) return null;
	return groups.get(groupId.value) || null;
});

const loadGroup = async () => {
	if (!groupId.value) {
		router.push({ name: 'group' });
		return;
	}

	loading.value = true;

	try {
		await store.populate([groupId.value]);
	} catch (e) {
		errorsStore.showGRPC(e);
		router.push({ name: 'group' });
	}

	loading.value = false;
};

const goBack = () => {
	router.push({ name: 'group' });
};

// Action handlers
const openManageGroup = () => {
	if (group.value) {
		// Populate form with current group data
		initialValues.value = {
			name: group.value.group?.name || '',
			"avatar-url": group.value.group?.avatarURL || '',
			description: group.value.group?.description || '',
		};
		dialogManageGroup.value = true;
	}
};

const update = async (e: FormSubmitEvent) => {
	if (!e.valid || !group.value) {
		logger.error('Update group form is invalid', e);
		return;
	}

	try {
		await store.update(UpdateGroupReq.create({
			iD: group.value?.group?.iD,
			name: { value: e.states.name.value },
			avatarURL: { value: e.states["avatar-url"].value },
			description: { value: e.states.description.value },
		}));
	} catch (e) {
		errorsStore.showGRPC(e)
		return
	}

	message.value = `Group ${e.states.name.value} updated successfully`;
	success.value = true;
	dialogManageGroup.value = false;
	e.reset()
	await authStore.refreshToken();
	await loadGroup();
};

const deleteGroup = () => {
	if (!group.value) return;

	confirm.require({
		message: `Are you sure you want to delete the group "${group.value.group?.name}"? This action cannot be undone.`,
		header: 'Delete Group',
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
			if (!group.value) return;

			const name = group.value.group?.name || '';

			try {
				await store.delete_(DeleteGroupReq.create({ iD: group.value?.group?.iD }));
			} catch (e) {
				errorsStore.showGRPC(e);
				return
			}

			message.value = `Group "${name}" deleted successfully`;
			success.value = true;
			await authStore.refreshToken();
			router.push({ name: 'group' });
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
	loadGroup();
});
</script>

<template>
	<div class="min-h-screen flex flex-col relative flex-auto">
		<!-- Header -->
		<div
			class="flex justify-between items-center px-8 py-4 bg-surface-0 dark:bg-surface-950 border-b border-surface-200 dark:border-surface-700">
			<div class="flex items-center gap-4">
				<Button icon="pi pi-arrow-left" severity="secondary" text rounded class="w-10 h-10" @click="goBack"
					v-tooltip.bottom="'Back to groups'" />
				<div class="flex items-center gap-3">
					<div
						class="flex items-center justify-center w-10 h-10 rounded-lg bg-blue-100 dark:bg-blue-400/30 border border-blue-200 dark:border-blue-400/20">
						<i class="pi pi-building text-lg text-blue-600 dark:text-blue-300"></i>
					</div>
					<div class="flex flex-col">
						<h2 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">
							{{ group?.group?.name || 'Loading...' }}
						</h2>
						<span class="text-sm text-surface-500 dark:text-surface-400">Group management</span>
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

				<!-- Group Not Found -->
				<Card v-else-if="!group" class="mb-6">
					<template #content>
						<div class="text-center p-8">
							<i class="pi pi-exclamation-circle text-4xl text-surface-400 mb-4"></i>
							<h3 class="text-xl font-semibold text-surface-900 dark:text-surface-0 mb-2">Group Not Found
							</h3>
							<p class="text-surface-500 dark:text-surface-400 mb-4">
								The group you're looking for doesn't exist or you don't have permission to view it.
							</p>
							<Button label="Back to Groups" icon="pi pi-arrow-left" @click="goBack" />
						</div>
					</template>
				</Card>

				<!-- Group Content with Tabs -->
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
								<i class="pi pi-shield"></i>
								<span>Roles</span>
							</div>
						</Tab>
						<Tab value="2">
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
												<Avatar v-if="group.group?.avatarURL" :image="group.group?.avatarURL"
													size="xlarge" shape="circle"
													class="border-4 border-surface-200 dark:border-surface-700 shadow-lg" />
												<Avatar v-else :label="group.group?.name.charAt(0).toUpperCase()"
													size="xlarge" shape="circle"
													class="bg-gradient-to-br from-primary-400 to-primary-600 text-white border-4 border-surface-200 dark:border-surface-700 shadow-lg text-4xl font-bold" />
											</div>

											<!-- Group Info -->
											<div class="flex-1 min-w-0">
												<div
													class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
													<div>
														<h1
															class="text-4xl font-bold text-surface-900 dark:text-surface-0 mb-2">
															{{ group.group?.name }}
														</h1>
														<div
															class="flex items-center gap-2 text-surface-500 dark:text-surface-400">
															<i class="pi pi-calendar text-sm"></i>
															<span class="text-sm">
																Created {{ formatDate(group.group?.createdAt) }}
															</span>
														</div>
													</div>
													<div class="flex items-center gap-3">
														<Button label="" icon="pi pi-cog" outlined severity="primary"
															class="font-medium" @click="openManageGroup"
															v-tooltip.bottom="'Edit group settings'" />
													</div>
												</div>

												<!-- Description -->
												<div class="mb-8">
													<div
														class="bg-surface-50 dark:bg-surface-800 rounded-xl p-6 border border-surface-200 dark:border-surface-700">
														<h3
															class="text-lg font-semibold text-surface-900 dark:text-surface-0 mb-3 flex items-center gap-2">
															<i class="pi pi-info-circle text-primary-500"></i>
															Description
														</h3>
														<p
															class="text-surface-700 dark:text-surface-200 leading-relaxed text-base">
															{{ group.group?.description || 'No description.' }}
														</p>
													</div>
												</div>

												<!-- Statistics Cards -->
												<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-6 mb-8">
													<div
														class="bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-900/20 dark:to-blue-800/20 rounded-xl p-6 border border-blue-200 dark:border-blue-700/50">
														<div class="flex items-center gap-3">
															<div
																class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
																<i class="pi pi-users text-white text-xl"></i>
															</div>
															<div>
																<p
																	class="text-2xl font-bold text-blue-700 dark:text-blue-300">
																	{{ group.userCount }}
																</p>
																<p
																	class="text-blue-600 dark:text-blue-400 text-sm font-medium">
																	Member(s)</p>
															</div>
														</div>
													</div>

													<div
														class="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-900/20 dark:to-purple-800/20 rounded-xl p-6 border border-purple-200 dark:border-purple-700/50">
														<div class="flex items-center gap-3">
															<div
																class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
																<i class="pi pi-shield text-white text-xl"></i>
															</div>
															<div>
																<p
																	class="text-2xl font-bold text-purple-700 dark:text-purple-300">
																	{{ group.roleCount }}
																</p>
																<p
																	class="text-purple-600 dark:text-purple-400 text-sm font-medium">
																	Role(s)</p>
															</div>
														</div>
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
																	{{ formatDate(group.group?.createdAt) }}
																</span>
															</div>
														</div>

														<div class="flex flex-col gap-2"
															v-if="group.group?.updatedAt && group.group?.updatedAt !== group.group?.createdAt">
															<h4
																class="text-sm font-semibold text-surface-900 dark:text-surface-0 uppercase tracking-wide">
																Last Updated
															</h4>
															<div class="flex items-center gap-2">
																<i
																	class="pi pi-clock text-surface-500 dark:text-surface-400"></i>
																<span class="text-surface-700 dark:text-surface-200">
																	{{ formatDate(group.group?.updatedAt) }}
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

						<!-- Roles Tab -->
						<TabPanel value="1">
							<RoleTable :groupId="groupId" />
						</TabPanel>

						<!-- Users Tab -->
						<TabPanel value="2">
							<UserTable :groupId="groupId" />
						</TabPanel>
					</TabPanels>
				</Tabs>
			</div>
		</div>

		<!-- Manage Group Dialog -->
		<Dialog v-model:visible="dialogManageGroup" append-to="body" modal
			:breakpoints="{ '960px': '75vw', '640px': '80vw' }" :style="{ width: '40rem' }" :draggable="false"
			:resizable="false" :show-header="false" class="shadow-sm rounded-2xl"
			:pt="{ content: '!p-6', footer: '!pb-6 !px-6' }">
			<div class="flex justify-between items-start gap-4">
				<div class="flex gap-4">
					<div
						class="flex items-center justify-center w-9 h-9 bg-orange-100 dark:bg-orange-400/30 text-orange-600 dark:text-orange-300 rounded-3xl border border-orange-200 dark:border-orange-400/20 shrink-0">
						<i class="pi pi-cog !text-xl !leading-none" />
					</div>
				</div>
				<Button icon="pi pi-times" text rounded severity="secondary" class="w-10 h-10 !p-2"
					@click="dialogManageGroup = false" />
			</div>
			<Form v-slot="$form" :resolver="resolver" :initialValues="initialValues" @submit="update"
				class="flex flex-col gap-6">
				<div class="flex items-start gap-4">
					<div class="flex-1 flex flex-col gap-2">
						<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Manage
							group
						</h1>
						<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">Edit group details
							and manage settings.</span>
					</div>
				</div>
				<div class="flex flex-col gap-6">
					<FormField v-slot="$field" name="name" class="flex flex-col gap-2">
						<label for="manage-name" class="text-color text-base">Name</label>
						<IconField icon-position="left" class="w-full">
							<InputIcon class="pi pi-user" />
							<InputText id="manage-name" name="name" placeholder="Enter group name" type="text"
								class="w-full" />
						</IconField>
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message
						}}
						</Message>
					</FormField>
					<FormField v-slot="$field" name="avatar-url" class="flex flex-col gap-2">
						<label for="manage-avatar-url" class="text-color text-base">Avatar URL</label>
						<IconField icon-position="left" class="w-full">
							<InputIcon class="pi pi-image" />
							<InputText id="manage-avatar-url" name="avatar-url" placeholder="Enter avatar URL"
								type="text" class="w-full" />
						</IconField>
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message
						}}
						</Message>
					</FormField>
					<FormField v-slot="$field" name="description" class="flex flex-col gap-2">
						<label for="manage-description" class="text-color text-base">Description</label>
						<Textarea id="manage-description" name="description" placeholder="Enter group description"
							rows="3" class="w-full" />
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message
						}}
						</Message>
					</FormField>
				</div>
				<div class="flex justify-between items-center">
					<Button label="Delete Group" icon="pi pi-trash" severity="danger" outlined class="font-medium"
						@click="deleteGroup" />
					<div class="flex gap-4">
						<Button label="Cancel" outlined @click="dialogManageGroup = false" />
						<Button label="Update" type="submit" />
					</div>
				</div>
			</Form>
		</Dialog>
	</div>
</template>

<style scoped></style>
