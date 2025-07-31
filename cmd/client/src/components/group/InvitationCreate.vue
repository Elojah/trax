<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useRoleStore } from '@/stores/role';

// Internal utilities and types
import { parse, ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { formatDate } from '@/utils/date';
import { createPagination } from '@/utils/requests';
import { back } from '@/utils/router';
import { useTable } from '@/composables';
import { ListRoleReq } from '@internal/user/dto/role';
import type { RolePermission } from '@internal/user/dto/role';
import { Command } from '@internal/user/role';

// PrimeVue UI Components
import DataTable, {
	type DataTableProps,
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
import { useUserStore } from '@/stores/user';
import { useInvitationStore } from '@/stores/invitation';

const route = useRoute();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const roleStore = useRoleStore();
const { roles, total } = toRefs(roleStore);
const invitationStore = useInvitationStore();

const groupId = route.params.groupId as string

const selectedRoles = ref<Map<string, boolean>>(new Map());

const group = computed(() => groups.value.get(groupId) || null);

// Create table composable with request creation function
const createRequest = (props: DataTableProps, search: string) => {
	if (!group.value) {
		return ListRoleReq.create({});
	}

	return ListRoleReq.create({
		groupIDs: [group.value?.group?.iD],
		search,
		paginate: createPagination(props)
	});
};

const table = useTable({
	createRequest,
	listMethod: roleStore.list,
	itemsMap: roles,
	total
});

const { views, handlers, search, loading: tableLoading } = table;

// Keep separate loading state for group loading
const loading = ref(false);

const loadGroup = async () => {
	loading.value = true;

	try {
		await groupStore.populate([groupId]);
	} catch (e) {
		errorsStore.showGRPC(e);
		back();
		return;
	}

	loading.value = false;
};


// Role validation error
const roleValidationError = ref<string | null>(null);

// Form validation
const resolver = zodResolver(
	z.object({
		email: z.string().email({ message: 'Please enter a valid email address.' }),
	})
);

const initialValues = ref({
	email: '',
});

// Validate roles selection
const validateRoles = (): boolean => {
	if (selectedRoles.value.size === 0) {
		roleValidationError.value = 'Please select at least one role for the user.';
		return false;
	}
	roleValidationError.value = null;
	return true;
};

// Watch for changes in selectedRoles to clear validation error
watch(selectedRoles, () => {
	if (selectedRoles.value.size > 0 && roleValidationError.value) {
		roleValidationError.value = null;
	}
}, { deep: true });

const toggleRoleSelection = (roleId: string) => {
	if (selectedRoles.value.has(roleId)) {
		selectedRoles.value.delete(roleId);
	} else {
		selectedRoles.value.set(roleId, true);
	}
};

const isRoleSelected = (roleId: string): boolean => {
	return selectedRoles.value.has(roleId);
};

const createInvitation = async (e: FormSubmitEvent) => {
	// Validate roles first
	const isRolesValid = validateRoles();

	if (!e.valid || !group.value || !isRolesValid) {
		logger.error('Invite user form is invalid', e);
		return;
	}

	try {
		await invitationStore.create(e.states.email.value, group.value?.group?.iD!, Array.from(selectedRoles.value.keys()).map(id => parse(id)));
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `User ${e.states.email.value} invited successfully${selectedRoles.value.size > 0 ? ` with ${selectedRoles.value.size} role(s)` : ''}`;
	success.value = true;

	// Reset form and go back to group details
	e.reset();
	selectedRoles.value = new Map();
	roleValidationError.value = null;

	// Wait a bit to show the success message, then navigate back
	setTimeout(() => {
		back();
	}, 2000);
};

// Initialize data on component mount
onMounted(async () => {
	// Load group data if not already loaded
	if (!group.value) {
		await loadGroup()
	}

	table.list();
});
</script>

<template>
	<div class="flex flex-col h-full">
		<!-- Page Header -->
		<div class="flex items-center gap-4 px-8 py-4 border-b border-surface-200 dark:border-surface-700">
			<Button icon="pi pi-arrow-left" severity="secondary" text rounded class="w-10 h-10" @click="back"
				v-tooltip.bottom="'Back'" />
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
			<Form v-slot="$form" :resolver="resolver" :initialValues="initialValues" @submit="createInvitation"
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
						{{ selectedRoles.size }} role(s) selected
					</span>
					<Button label="Cancel" outlined @click="back" />
					<Button label="Send Invitation" type="submit" />
				</div>
			</Form>
			<!-- Role validation error message -->
			<Message v-if="roleValidationError" severity="error" size="small" variant="simple" class="mt-2 mx-6">
				{{ roleValidationError }}
			</Message>
		</div>

		<!-- Role Selection Table -->
		<DataTable :value="views" :lazy="true" :loading="tableLoading" :paginator="true"
			:rows="table.properties.value.rows" :totalRecords="Number(total)" :first="table.properties.value.first"
			v-model:filters="table.properties.value.filters" :scrollable="true" scrollHeight="calc(100vh - 20rem)"
			@page="handlers.onPage" @sort="handlers.onSort" @filter="handlers.onFilter" dataKey="id"
			filterDisplay="menu" :globalFilterFields="['name', 'created_at']" tableStyle="min-width: 50rem"
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
									placeholder="Search roles by name..." v-model="search" @input="handlers.onSearch" />
							</IconField>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10"
							@click="table.list()" v-tooltip.bottom="'Refresh roles'" />
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
