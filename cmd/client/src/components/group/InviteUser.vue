<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useRoleStore } from '@/stores/role';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { ListRoleReq } from '@internal/user/dto/role';
import type { RolePermission } from '@internal/user/dto/role';
import { Command } from '@internal/user/role';

// PrimeVue UI Components
import DataTable, {
	type DataTableFilterEvent,
	type DataTablePageEvent,
	type DataTableProps,
	type DataTableSortEvent
} from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';
import Checkbox from 'primevue/checkbox';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

// Form Validation
import { zodResolver } from '@primevue/forms/resolvers/zod';
import z from 'zod';
import { Form, FormField } from '@primevue/forms';
import type { FormSubmitEvent } from '@primevue/forms';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const roleStore = useRoleStore();
const { roles, total } = toRefs(roleStore);

const groupId = computed(() => route.params.groupId as string);

const loading = ref(false);
const search = ref('');
const viewIDs = ref<string[]>([]);
const selectedRoles = ref<string[]>([]);

const group = computed(() => {
	return groups.value.get(groupId.value) || null;
});

const views = computed(() => {
	return viewIDs.value.map((roleId: string) => roles.value?.get(roleId));
});

const properties = ref<DataTableProps>({});

// Form validation
const resolver = zodResolver(
	z.object({
		email: z.string().email({ message: 'Please enter a valid email address.' }),
	})
);

const initialValues = ref({
	email: '',
});

const list = async (p: DataTableProps = {
	first: 0,
	rows: 10,
	sortField: 'created_at',
	sortOrder: -1,
}) => {
	if (!group.value) {
		viewIDs.value = [];
		return;
	}

	loading.value = true;

	try {
		const page = Math.floor((p.first ?? 0) / (p.rows ?? 1)) + 1;
		const sortBy = p.sortField ? [{
			key: p.sortField,
			order: p.sortOrder === 1 ? 'asc' : 'desc'
		}] : [{ key: 'created_at', order: 'desc' }];

		const newIDs = await roleStore.list(ListRoleReq.create({
			groupIDs: [group.value.group?.iD],
			search: search.value,
			paginate: {
				start: BigInt(((page - 1) * (p.rows ?? 10)) + 1),
				end: BigInt(page * (p.rows ?? 10)),
				sort: sortBy?.at(0)?.key ?? '',
				order: sortBy?.at(0)?.order === 'asc' ? true : false,
			}
		}));

		viewIDs.value = newIDs;

		if (p) {
			properties.value = p;
		}

	} catch (e) {
		errorsStore.showGRPC(e);
	}

	loading.value = false;
};

// Handle DataTable lazy loading events
const onPage = (event: DataTablePageEvent) => {
	const props: DataTableProps = {
		first: event.first,
		rows: event.rows,
		sortField: event.sortField,
		sortOrder: event.sortOrder ?? 0,
		filters: event.filters,
	};
	list(props);
};

const onSort = (event: DataTableSortEvent) => {
	const props: DataTableProps = {
		first: event.first,
		rows: event.rows,
		sortField: event.sortField,
		sortOrder: event.sortOrder ?? 0,
		filters: event.filters,
	};
	list(props);
};

const onFilter = (event: DataTableFilterEvent) => {
	const props: DataTableProps = {
		first: event.first,
		rows: event.rows,
		sortField: event.sortField,
		sortOrder: event.sortOrder ?? 0,
		filters: event.filters,
	};
	list(props);
};

// Handle search input changes
const onSearch = () => {
	list({
		first: 0,
		rows: properties.value.rows ?? 10,
		sortField: properties.value.sortField,
		sortOrder: properties.value.sortOrder ?? -1,
		filters: properties.value.filters,
	});
};

const toggleRoleSelection = (roleId: string) => {
	const index = selectedRoles.value.indexOf(roleId);
	if (index > -1) {
		selectedRoles.value.splice(index, 1);
	} else {
		selectedRoles.value.push(roleId);
	}
};

const isRoleSelected = (roleId: string): boolean => {
	return selectedRoles.value.includes(roleId);
};

const goBack = () => {
	router.push({
		name: 'group-details',
		params: { id: groupId.value }
	});
};

const inviteUser = async (e: FormSubmitEvent) => {
	if (!e.valid || !group.value) {
		logger.error('Invite user form is invalid', e);
		return;
	}

	try {
		// TODO: Implement user invitation logic with selected roles
		// await userStore.invite(e.states.email.value, group.value.iD, selectedRoles.value);
		console.log('Inviting user:', e.states.email.value, 'to group:', group.value.group?.name);
		if (selectedRoles.value.length > 0) {
			console.log('Selected roles:', selectedRoles.value);
		}
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `User ${e.states.email.value} invited successfully${selectedRoles.value.length > 0 ? ` with ${selectedRoles.value.length} role(s)` : ''}`;
	success.value = true;

	// Reset form and go back to group details
	e.reset();
	selectedRoles.value = [];

	// Wait a bit to show the success message, then navigate back
	setTimeout(() => {
		goBack();
	}, 2000);
};

const formatDate = (timestamp: bigint): string => {
	return new Date(Number(timestamp) * 1000).toLocaleDateString('en-GB', {
		day: 'numeric',
		month: 'short',
		year: 'numeric'
	});
};

// Initialize data on component mount
onMounted(async () => {
	// Load group data if not already loaded
	if (!group.value) {
		try {
			await groupStore.populate([groupId.value]);
		} catch (e) {
			errorsStore.showGRPC(e);
			goBack();
			return;
		}
	}

	if (group.value) {
		list();
	}
});
</script>

<template>
	<div class="flex flex-col h-full">
		<!-- Page Header -->
		<div class="flex items-center gap-4 p-6 border-b border-surface-200 dark:border-surface-700">
			<Button icon="pi pi-arrow-left" severity="secondary" text rounded class="w-10 h-10" @click="goBack"
				v-tooltip.bottom="'Back to group'" />
			<div class="flex items-center gap-3">
				<div
					class="flex items-center justify-center w-10 h-10 rounded-lg bg-green-100 dark:bg-green-400/30 border border-green-200 dark:border-green-400/20">
					<i class="pi pi-user-plus text-lg text-green-600 dark:text-green-300"></i>
				</div>
				<div class="flex flex-col">
					<h2 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Invite User</h2>
					<span class="text-sm text-surface-500 dark:text-surface-400">
						Send an invitation to join {{ group?.group?.name || 'this group' }}
					</span>
				</div>
			</div>
		</div>

		<!-- Success Message -->
		<Message v-if="success && message" severity="success" class="m-6 mb-0">{{ message }}</Message>

		<!-- User Information Form -->
		<div class="p-6 border-b border-surface-200 dark:border-surface-700 bg-surface-50 dark:bg-surface-800">
			<Form v-slot="$form" :resolver="resolver" :initialValues="initialValues" @submit="inviteUser"
				class="flex items-end gap-6">
				<FormField v-slot="$field" name="email" class="flex-1">
					<label for="invite-email" class="block text-color text-base font-medium mb-2">Email Address</label>
					<IconField icon-position="left" class="w-full">
						<InputIcon class="pi pi-envelope" />
						<InputText id="invite-email" name="email" placeholder="Enter user email address" type="email"
							class="w-full" />
					</IconField>
					<Message v-if="$field?.invalid" severity="error" size="small" variant="simple" class="mt-1">{{
						$field.error?.message
					}}
					</Message>
				</FormField>
				<div class="flex items-center gap-3">
					<span class="text-sm text-surface-500 dark:text-surface-400 whitespace-nowrap">
						{{ selectedRoles.length }} role(s) selected
					</span>
					<Button label="Cancel" outlined @click="goBack" />
					<Button label="Send Invitation" type="submit" />
				</div>
			</Form>
		</div>

		<!-- Role Selection Table -->
		<DataTable :value="views" :lazy="true" :loading="loading" :paginator="true" :rows="properties.rows"
			:totalRecords="Number(total)" :first="properties.first" v-model:filters="properties.filters"
			:scrollable="true" scrollHeight="calc(100vh - 20rem)" @page="onPage" @sort="onSort" @filter="onFilter"
			dataKey="id" filterDisplay="menu" :globalFilterFields="['name', 'created_at']" tableStyle="min-width: 50rem"
			:rowsPerPageOptions="[10, 25, 50]"
			paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
			currentPageReportTemplate="{first} - {last} ({totalRecords})" pt:header:class="!p-0">

			<template #header>
				<div
					class="flex justify-between items-center gap-6 p-4 bg-surface-50 dark:bg-surface-800 border-b border-surface-200 dark:border-surface-700">
					<div class="flex items-center gap-4">
						<div class="flex items-center gap-3">
							<div
								class="flex items-center justify-center w-10 h-10 rounded-lg bg-purple-100 dark:bg-purple-400/30 border border-purple-200 dark:border-purple-400/20">
								<i class="pi pi-shield text-lg text-purple-600 dark:text-purple-300"></i>
							</div>
							<div class="flex flex-col">
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Select Roles
								</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">Choose one or more roles to
									assign to the user</span>
							</div>
						</div>
						<div class="flex items-center gap-3 ml-6">
							<IconField icon-position="left" class="relative">
								<InputIcon class="pi pi-search text-surface-400" />
								<InputText type="text"
									class="w-96 pl-10 pr-4 py-2 border border-surface-300 dark:border-surface-600 rounded-lg bg-white dark:bg-surface-900 focus:border-primary-500 dark:focus:border-primary-400 focus:ring-2 focus:ring-primary-200 dark:focus:ring-primary-800 transition-all duration-200"
									placeholder="Search roles by name..." v-model="search" @input="onSearch" />
							</IconField>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10"
							@click="list()" v-tooltip.bottom="'Refresh roles'" />
					</div>
				</div>
			</template>

			<Column field="select" header="" style="width: 10%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex justify-center">
						<Checkbox :modelValue="isRoleSelected(ulid(data.role.iD))"
							@update:modelValue="toggleRoleSelection(ulid(data.role.iD))" binary />
					</div>
				</template>
			</Column>

			<Column field="name" header="Name" sortable style="width: 40%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex items-center gap-3">
						<div class="relative">
							<Avatar :label="data.role.name.charAt(0).toUpperCase()" size="large" shape="circle"
								class="bg-purple-100 dark:bg-purple-400/30 text-purple-600 dark:text-purple-300 border-2 border-purple-200 dark:border-purple-400/20" />
						</div>
						<div class="flex flex-col">
							<span class="font-semibold text-surface-900 dark:text-surface-0">
								{{ data.role.name }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<Column field="readPermissions" header="Read" style="width: 12%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-eye text-blue-500 dark:text-blue-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{data.permissions.filter(p => p.command === Command.C_read).length}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="editPermissions" header="Edit" style="width: 12%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-pencil text-orange-500 dark:text-orange-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{data.permissions.filter(p => p.command === Command.C_edit).length}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="writePermissions" header="Write" style="width: 12%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-file-edit text-green-500 dark:text-green-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{data.permissions.filter(p => p.command === Command.C_write).length}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="created_at" sortable style="width: 14%">
				<template #header>
					<span class="flex-1 text-right font-bold">Created</span>
				</template>
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex flex-col gap-1 justify-end">
						<div class="flex items-center gap-2 justify-end">
							<i class="pi pi-calendar text-surface-500 dark:text-surface-400"></i>
							<span class="font-medium text-surface-700 dark:text-surface-200">
								{{ formatDate(data.role.createdAt) }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<template #empty>
				<div class="text-center p-4">
					<i class="pi pi-shield text-4xl text-gray-400 mb-4"></i>
					<p class="text-gray-500">No roles found.</p>
				</div>
			</template>
		</DataTable>
	</div>
</template>

<style scoped></style>
