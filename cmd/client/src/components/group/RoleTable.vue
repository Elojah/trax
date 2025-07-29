<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useRoleStore } from '@/stores/role';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { formatDate } from '@/utils/date';
import { createPagination } from '@/utils/requests';
import { useTable } from '@/composables';
import { ListRoleReq, RolePermission, CreateRoleReq } from '@internal/user/dto/role';
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
import Dialog from 'primevue/dialog';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

// Table components
import PermissionTable from '@/components/permission/Table.vue';

// Form Validation
import { zodResolver } from '@primevue/forms/resolvers/zod';
import z from 'zod';
import { Form, FormField } from '@primevue/forms';
import type { FormSubmitEvent } from '@primevue/forms';

const props = defineProps<{
	groupId: string;
}>();

const router = useRouter();
const authStore = useAuthStore();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const roleStore = useRoleStore();
const { roles, total } = toRefs(roleStore);

const permissions = ref()

const group = groups.value.get(props.groupId) || null;

// Create table composable with request creation function
const createRequest = (props: DataTableProps, search: string) => {
	if (!group) {
		return ListRoleReq.create({});
	}

	return ListRoleReq.create({
		groupIDs: [group.group?.iD],
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

const { views, handlers, search, loading } = table;

// Dialog states
const dialogCreateRole = ref(false);

// Form validation
const resolver = zodResolver(
	z.object({
		name: z.string().min(1, { message: 'Name is required.' }),
	})
);

const initialValues = ref({
	name: '',
});

const openCreateRole = () => {
	initialValues.value = {
		name: '',
	};
	dialogCreateRole.value = true;
};

const navigateToRoleDetails = (role: RolePermission) => {
	if (!role.role) return;
	router.push({
		name: 'role-details',
		params: {
			groupId: props.groupId,
			roleId: ulid(role.role.iD)
		}
	});
};

const create = async (e: FormSubmitEvent) => {
	if (!e.valid || !group) {
		logger.error('Create role form is invalid', e);
		return;
	}

	try {
		await roleStore.create(
			group.group?.iD!,
			e.states.name.value,
			permissions.value?.selected() || [],
		);
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `Role ${e.states.name.value} created successfully`;
	success.value = true;
	dialogCreateRole.value = false;
	e.reset();

	// Clear permissions table for next use
	permissions.value?.clear();

	await authStore.refreshToken();
	await table.list();
};

// Initialize data on component mount
onMounted(() => {
	table.list();
});

</script>

<template>
	<div class="flex flex-col h-full">
		<!-- Success Message -->
		<Message v-if="success && message" severity="success" class="mb-4">{{ message }}</Message>

		<DataTable :value="views" :lazy="true" :loading="loading" :paginator="true" :rows="table.properties.value.rows"
			:totalRecords="Number(total)" :first="table.properties.value.first"
			v-model:filters="table.properties.value.filters" :scrollable="true" scrollHeight="calc(100vh - 16rem)"
			@page="handlers.onPage" @sort="handlers.onSort" @filter="handlers.onFilter" dataKey="id"
			filterDisplay="menu" :globalFilterFields="['name', 'created_at']" tableStyle="min-width: 50rem"
			:rowsPerPageOptions="[10, 25, 50, 100]"
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
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Roles</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">Manage group roles and
									permissions</span>
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
						<Button label="" icon="pi pi-plus" outlined severity="primary" class="font-medium"
							@click="openCreateRole" />
					</div>
				</div>
			</template>

			<Column field="name" header="Name" sortable style="width: 40%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex items-center gap-3">
						<div class="relative">
							<Avatar :label="data.role.name.charAt(0).toUpperCase()" size="large" shape="circle"
								class="bg-purple-100 dark:bg-purple-400/30 text-purple-600 dark:text-purple-300 border-2 border-purple-200 dark:border-purple-400/20" />
						</div>
						<div class="flex flex-col">
							<span
								class="font-semibold text-surface-900 dark:text-surface-0 cursor-pointer hover:text-primary-600 dark:hover:text-primary-400 transition-colors duration-200"
								@click="navigateToRoleDetails(data)" v-tooltip.top="'View role details'">
								{{ data.role.name }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<Column field="readPermissions" header="Read" style="width: 10%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-eye text-blue-500 dark:text-blue-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{data.permissions.filter(p => p.command === Command.C_read).length}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="editPermissions" header="Edit" style="width: 10%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-pencil text-orange-500 dark:text-orange-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{data.permissions.filter(p => p.command === Command.C_edit).length}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="writePermissions" header="Write" style="width: 10%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-file-edit text-green-500 dark:text-green-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{data.permissions.filter(p => p.command === Command.C_write).length}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="created_at" sortable style="width: 30%">
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

		<!-- Create Role Dialog -->
		<Dialog v-model:visible="dialogCreateRole" append-to="body" modal
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
					@click="dialogCreateRole = false" />
			</div>
			<Form v-slot="$form" :resolver="resolver" :initialValues="initialValues" @submit="create"
				class="flex flex-col gap-6">
				<div class="flex items-start gap-4">
					<div class="flex-1 flex flex-col gap-2">
						<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Create
							role
						</h1>
						<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">Create a new role
							for
							this group.</span>
					</div>
				</div>
				<div class="flex flex-col gap-6">
					<FormField v-slot="$field" name="name" class="flex flex-col gap-2">
						<label for="create-name" class="text-color text-base">Name</label>
						<IconField icon-position="left" class="w-full">
							<InputIcon class="pi pi-shield" />
							<InputText id="create-name" name="name" placeholder="Enter role name" type="text"
								class="w-full" />
						</IconField>
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message
						}}
						</Message>
					</FormField>
				</div>
				<div class="flex flex-col gap-6">
					<PermissionTable :disabled="false" :permissions="[]" ref="permissions" />
				</div>

				<div class="flex justify-end gap-4">
					<Button label="Cancel" outlined @click="dialogCreateRole = false" />
					<Button label="Create" type="submit" />
				</div>
			</Form>
		</Dialog>
	</div>
</template>

<style scoped></style>
