<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useRoleStore } from '@/stores/role';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { ListRoleReq, RolePermission, CreateRoleReq, UpdateRoleReq, DeleteRoleReq } from '@internal/user/dto/role';

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
import { useConfirm } from 'primevue/useconfirm';

const props = defineProps<{
	groupId: string;
}>();

const authStore = useAuthStore();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const roleStore = useRoleStore();
const { roles, total } = toRefs(roleStore);
const confirmRole = useConfirm();

const loading = ref(false);
const search = ref('');
const viewIDs = ref<string[]>([]);
const permissions = ref()

const group = computed(() => {
	return groups.value.get(props.groupId) || null;
});

const views = computed(() => {
	return viewIDs.value.map((roleId: string) => roles.value?.get(roleId));
});

const properties = ref<DataTableProps>({});

// Dialog states
const dialogCreateRole = ref(false);
const dialogManageRole = ref(false);
const selectedRole = ref<RolePermission | null>(null);

// Form validation
const resolver = zodResolver(
	z.object({
		name: z.string().min(1, { message: 'Name is required.' }),
	})
);

const initialValues = ref({
	name: '',
});

const list = async (props: DataTableProps = {
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
		const page = Math.floor((props.first ?? 0) / (props.rows ?? 1)) + 1;
		const sortBy = props.sortField ? [{
			key: props.sortField,
			order: props.sortOrder === 1 ? 'asc' : 'desc'
		}] : [{ key: 'created_at', order: 'desc' }];

		const newIDs = await roleStore.list(ListRoleReq.create({
			groupIDs: [group.value.group?.iD],
			search: search.value,
			paginate: {
				start: BigInt(((page - 1) * (props.rows ?? 10)) + 1),
				end: BigInt(page * (props.rows ?? 10)),
				sort: sortBy?.at(0)?.key ?? '',
				order: sortBy?.at(0)?.order === 'asc' ? true : false,
			}
		}));

		viewIDs.value = newIDs;

		if (props) {
			properties.value = props;
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

const openCreateRole = () => {
	initialValues.value = {
		name: '',
	};
	dialogCreateRole.value = true;
};

const openManageRole = (role: RolePermission) => {
	selectedRole.value = role;
	initialValues.value = {
		name: role.role?.name || '',
	};
	dialogManageRole.value = true;
};

const create = async (e: FormSubmitEvent) => {
	if (!e.valid || !group.value) {
		logger.error('Create role form is invalid', e);
		return;
	}

	try {
		await roleStore.create(
			group.value.group?.iD!,
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
	await list();
};

const update = async (e: FormSubmitEvent) => {
	if (!e.valid || !selectedRole.value?.role) {
		logger.error('Update role form is invalid', e);
		return;
	}

	try {
		await roleStore.update(UpdateRoleReq.create({
			iD: selectedRole.value.role.iD,
			name: { value: e.states.name.value },
			description: { value: e.states.description.value },
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
	await list();
};

const deleteRole = (role: RolePermission) => {
	if (!role.role) return;

	confirmRole.require({
		message: `Are you sure you want to delete the role "${role.role.name}"? This action cannot be undone.`,
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
			if (!role.role) return;

			try {
				await roleStore.delete_(DeleteRoleReq.create({ iD: role.role.iD }));
			} catch (e) {
				errorsStore.showGRPC(e);
				return;
			}

			message.value = `Role "${role.role.name}" deleted successfully`;
			success.value = true;
			await authStore.refreshToken();
			await list();
		}
	});
};

const formatDate = (timestamp: bigint): string => {
	return new Date(Number(timestamp) * 1000).toLocaleDateString('en-GB', {
		day: 'numeric',
		month: 'short',
		year: 'numeric'
	});
};

const short = (description: string): string => {
	return description && description.length > 64 ? description.substring(0, 64) + '...' :
		description || 'No description';
};

// Initialize data on component mount
onMounted(() => {
	if (group.value) {
		list();
	}
});

// Watch for group changes
watch(() => props.groupId, () => {
	if (group.value) {
		list();
	}
});
</script>

<template>
	<div class="flex flex-col h-full">
		<!-- Success Message -->
		<Message v-if="success && message" severity="success" class="mb-4">{{ message }}</Message>

		<DataTable :value="views" :lazy="true" :loading="loading" :paginator="true" :rows="properties.rows"
			:totalRecords="Number(total)" :first="properties.first" v-model:filters="properties.filters"
			:scrollable="true" scrollHeight="calc(100vh - 16rem)" @page="onPage" @sort="onSort" @filter="onFilter"
			dataKey="id" filterDisplay="menu" :globalFilterFields="['name', 'created_at']" tableStyle="min-width: 50rem"
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
									placeholder="Search roles by name..." v-model="search" @input="onSearch" />
							</IconField>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10"
							@click="list()" v-tooltip.bottom="'Refresh roles'" />
						<Button label="" icon="pi pi-plus" outlined severity="primary" class="font-medium"
							@click="openCreateRole" />
					</div>
				</div>
			</template>

			<Column field="name" header="Name" sortable style="width: 30%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex items-center gap-3">
						<div class="relative">
							<Avatar :label="data.role.name.charAt(0).toUpperCase()" size="large" shape="circle"
								class="bg-purple-100 dark:bg-purple-400/30 text-purple-600 dark:text-purple-300 border-2 border-purple-200 dark:border-purple-400/20" />
						</div>
						<div class="flex flex-col">
							<span
								class="font-semibold text-surface-900 dark:text-surface-0 cursor-pointer hover:text-primary-600 dark:hover:text-primary-400 transition-colors duration-200"
								@click="openManageRole(data)" v-tooltip.top="'Edit role'">
								{{ data.role.name }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<Column field="permissions" header="Permissions" style="width: 55%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.permissions" class="flex items-center gap-2">
						<i class="pi pi-shield text-surface-500 dark:text-surface-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{ data.permissions.length }}
						</span>
					</div>
				</template>
			</Column>

			<Column field="created_at" header="Created" sortable style="width: 15%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex flex-col gap-1">
						<div class="flex items-center gap-2">
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
				<div class="flex flex-col gap-6">
					<PermissionTable :disabled="true" :permissions="selectedRole?.permissions" />
				</div>
				<div class="flex justify-between items-center">
					<Button label="Delete Role" icon="pi pi-trash" severity="danger" outlined class="font-medium"
						@click="deleteRole(selectedRole!)" />
					<div class="flex gap-4">
						<Button label="Cancel" outlined @click="dialogManageRole = false" />
						<Button label="Update" type="submit" />
					</div>
				</div>
			</Form>
		</Dialog>
	</div>
</template>

<style scoped></style>
