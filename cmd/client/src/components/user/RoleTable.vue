<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useRoleStore } from '@/stores/role';
import { useUserStore } from '@/stores/user';

// Internal utilities and types
import { parse, ulid } from '@/utils/ulid';
import { createPagination } from '@/utils/requests';
import { useTable } from '@/composables';
import { ListRoleReq } from '@internal/user/dto/role';
import type { RolePermission } from '@internal/user/dto/role';
import { Command, Resource } from '@internal/user/role';

// PrimeVue UI Components
import DataTable, {
	type DataTableProps,
} from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';
import Tag from 'primevue/tag';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

// Form Validation
import { useConfirm } from 'primevue/useconfirm';
import router from '@/router';
import { hasClaim } from '@/utils/claims';

const props = defineProps<{
	userId: string;
	groupId: string;
}>();

const authStore = useAuthStore();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const roleStore = useRoleStore();
const { roles, rolesByUser, total } = toRefs(roleStore);
const userStore = useUserStore();
const { users } = toRefs(userStore);
const { claims } = toRefs(authStore);
const confirm = useConfirm();

const group = groups.value.get(props.groupId);
const user = users.value.get(props.userId);

const userRoleIds = computed(() => {
	const userHasRoles = new Map<string, boolean>();
	rolesByUser.value.get(ulid(user?.iD))?.forEach((ok, roleId) => {
		if (ok) {
			userHasRoles.set(roleId, true);
		}
	});
	return userHasRoles;
});

const readRolePermission = computed(() => {
	return hasClaim(claims.value, group?.group?.iD!, Resource.R_role, Command.C_read);
});

const editRolePermission = computed(() => {
	return hasClaim(claims.value, group?.group?.iD!, Resource.R_role, Command.C_edit);
});

// Create request function for the table
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

// Use the table composable
const table = useTable({
	createRequest,
	listMethod: roleStore.list,
	itemsMap: roles,
	total
});

const { viewIDs, views, properties, loading, search, handlers, list } = table;

const loadRoles = async () => {
	if (!group) return;

	loading.value = true;

	try {
		await roleStore.list(ListRoleReq.create({
			userID: user?.iD,
			groupIDs: [group.group?.iD],
		}));
	} catch (e) {
		errorsStore.showGRPC(e);
	}

	loading.value = false;
};

const addRole = async (role: RolePermission) => {
	if (!role.role) return;

	try {
		await roleStore.addUser(
			role.role.iD,
			user?.iD!,
		);
		await userStore.addRole(
			user?.iD!,
			role.role.iD,
			true // dry run to update local state only
		);
		message.value = `Role "${role.role.name}" added to user`;
		success.value = true;
		await authStore.refreshToken();
	} catch (e) {
		errorsStore.showGRPC(e);
	}
};

const removeRole = (role: RolePermission) => {
	if (!role.role) return;

	confirm.require({
		message: `Are you sure you want to remove the role "${role.role.name}" from this user?`,
		header: 'Remove Role',
		icon: 'pi pi-exclamation-triangle',
		rejectProps: {
			label: 'Cancel',
			severity: 'secondary',
			outlined: true
		},
		acceptProps: {
			label: 'Remove',
			severity: 'danger'
		},
		accept: async () => {
			if (!role.role) return;

			try {
				await roleStore.deleteUser(
					role.role.iD,
					user?.iD!,
				);
				await userStore.deleteRole(
					user?.iD!,
					role.role.iD,
					true
				);
				message.value = `Role "${role.role.name}" removed from user`;
				success.value = true;
				await authStore.refreshToken();
			} catch (e) {
				errorsStore.showGRPC(e);
			}
		}
	});
};

// Initialize data on component mount
onMounted(async () => {
	await loadRoles();
	list();
});
</script>

<template>
	<div class="flex flex-col h-full">
		<!-- Success Message -->
		<Message v-if="success && message" severity="success" class="mb-4">{{ message }}</Message>

		<DataTable :value="views" :lazy="true" :loading="loading" :paginator="true" :rows="properties.rows"
			:totalRecords="Number(total)" :first="properties.first" v-model:filters="properties.filters"
			:scrollable="true" scrollHeight="calc(100vh - 16rem)" @page="handlers.onPage" @sort="handlers.onSort"
			@filter="handlers.onFilter" dataKey="id" filterDisplay="menu"
			:globalFilterFields="['name', 'status', 'created_at']" tableStyle="min-width: 50rem"
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
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">User Roles
								</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">Roles assigned to this
									user</span>
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
							@click="list()" v-tooltip.bottom="'Refresh roles'" />
					</div>
				</div>
			</template>

			<Column field="name" header="Role" sortable style="width: 40%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex items-center gap-3">
						<div class="relative">
							<Avatar :label="data.role.name.charAt(0).toUpperCase()" size="large" shape="circle"
								class="bg-purple-100 dark:bg-purple-400/30 text-purple-600 dark:text-purple-300 border-2 border-purple-200 dark:border-purple-400/20" />
						</div>
						<div class="flex flex-col">
							<span
								@click="router.push({ name: 'role-details', params: { groupId: ulid(group?.group?.iD), roleId: ulid(data.role.iD) } })"
								class="font-semibold text-surface-900 dark:text-surface-0 hover:text-primary-600 dark:hover:text-primary-400 cursor-pointer transition-colors duration-200">
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

			<Column field="status" style="width: 20%">
				<template #header>
					<span class="flex-1 text-center font-bold">Status</span>
				</template>
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex items-center justify-center gap-2">
						<Tag v-if="userRoleIds.has(ulid(data.role.iD))" value="Assigned" severity="success" />
						<Tag v-else value="Unassigned" severity="secondary" />
					</div>
				</template>
			</Column>

			<Column header="" style="width: 10%">
				<template #body="{ data }: { data: RolePermission }">
					<div v-if="data?.role" class="flex items-center gap-2 justify-end">
						<Button :disabled="!editRolePermission" v-if="!userRoleIds.has(ulid(data.role.iD))"
							icon="pi pi-plus" severity="success" outlined size="small" @click="addRole(data)"
							v-tooltip.top="'Add role'" />
						<Button :disabled="!editRolePermission" v-else icon="pi pi-minus" severity="danger" outlined
							size="small" @click="removeRole(data)" v-tooltip.top="'Remove role'" />
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
