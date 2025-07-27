<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useUserStore } from '@/stores/user';

// Internal utilities and types
import { ulid, parse } from '@/utils/ulid';
import { logger } from "@/config";
import { ListUserReq } from '@internal/user/dto/user';
import type { U } from '@internal/user/user';

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
import Tag from 'primevue/tag';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

// Form Validation
import { useConfirm } from 'primevue/useconfirm';

const props = defineProps<{
	roleId: string;
	groupId: string;
}>();

const router = useRouter();
const authStore = useAuthStore();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const userStore = useUserStore();
const { users, total, usersByRole } = toRefs(userStore);
const confirm = useConfirm();

const loading = ref(false);
const search = ref('');
const viewIDs = ref<string[]>([]);

const group = computed(() => {
	return groups.value.get(props.groupId) || null;
});

const views = computed(() => {
	return viewIDs.value.map((userId: string) => users.value?.get(userId));
});

const properties = ref<DataTableProps>({});

// Get users who have this role
const usersWithRole = computed(() => {
	return usersByRole.value.get(props.roleId) || new Map<string, boolean>();
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

		const newIDs = await userStore.list(ListUserReq.create({
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

// Load users with this role
const loadUsersWithRole = async () => {
	if (!group.value || !props.roleId) return;

	try {
		await userStore.list(ListUserReq.create({
			groupIDs: [group.value.group?.iD],
			roleID: parse(props.roleId),
		}));
	} catch (e) {
		errorsStore.showGRPC(e);
	}
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

const addRole = async (user: U) => {
	if (!user) return;

	try {
		await userStore.addRole(
			user.iD,
			parse(props.roleId)
		);
		message.value = `Role assigned to user "${getUserDisplayName(user)}"`;
		success.value = true;
		await authStore.refreshToken();
		await loadUsersWithRole();
	} catch (e) {
		errorsStore.showGRPC(e);
	}
};

const removeRole = (user: U) => {
	if (!user) return;

	confirm.require({
		message: `Are you sure you want to remove this role from user "${getUserDisplayName(user)}"?`,
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
			if (!user) return;

			try {
				await userStore.deleteRole(
					user.iD,
					parse(props.roleId)
				);
				message.value = `Role removed from user "${getUserDisplayName(user)}"`;
				success.value = true;
				await authStore.refreshToken();
				await loadUsersWithRole();
			} catch (e) {
				errorsStore.showGRPC(e);
			}
		}
	});
};

const navigateToUserDetails = (user: U) => {
	if (!user) return;
	router.push({
		name: 'user-details',
		params: {
			groupId: props.groupId,
			userId: ulid(user.iD)
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
onMounted(() => {
	if (group.value) {
		list();
		loadUsersWithRole();
	}
});

// Watch for group changes
watch(() => props.groupId, () => {
	if (group.value) {
		list();
		loadUsersWithRole();
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
			dataKey="id" filterDisplay="menu" :globalFilterFields="['email', 'name', 'created_at']"
			tableStyle="min-width: 50rem" :rowsPerPageOptions="[10, 25, 50, 100]"
			paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
			currentPageReportTemplate="{first} - {last} ({totalRecords})" pt:header:class="!p-0">

			<template #header>
				<div
					class="flex justify-between items-center gap-6 p-4 bg-surface-50 dark:bg-surface-800 border-b border-surface-200 dark:border-surface-700">
					<div class="flex items-center gap-4">
						<div class="flex items-center gap-3">
							<div
								class="flex items-center justify-center w-10 h-10 rounded-lg bg-green-100 dark:bg-green-400/30 border border-green-200 dark:border-green-400/20">
								<i class="pi pi-users text-lg text-green-600 dark:text-green-300"></i>
							</div>
							<div class="flex flex-col">
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Role Users
								</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">Users in group and their
									role assignment status</span>
							</div>
						</div>
						<div class="flex items-center gap-3 ml-6">
							<IconField icon-position="left" class="relative">
								<InputIcon class="pi pi-search text-surface-400" />
								<InputText type="text"
									class="w-96 pl-10 pr-4 py-2 border border-surface-300 dark:border-surface-600 rounded-lg bg-white dark:bg-surface-900 focus:border-primary-500 dark:focus:border-primary-400 focus:ring-2 focus:ring-primary-200 dark:focus:ring-primary-800 transition-all duration-200"
									placeholder="Search users by name or email..." v-model="search" @input="onSearch" />
							</IconField>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10"
							@click="list()" v-tooltip.bottom="'Refresh users'" />
					</div>
				</div>
			</template>

			<Column field="user" header="User" sortable style="width: 40%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex items-center gap-3">
						<div class="relative">
							<Avatar v-if="data.avatarURL" :image="data.avatarURL" size="large" shape="circle"
								class="border-2 border-surface-200 dark:border-surface-700" />
							<Avatar v-else :label="getUserInitials(data)" size="large" shape="circle"
								class="bg-green-100 dark:bg-green-400/30 text-green-600 dark:text-green-300 border-2 border-green-200 dark:border-green-400/20" />
						</div>
						<div class="flex flex-col">
							<span
								class="font-semibold text-surface-900 dark:text-surface-0 cursor-pointer hover:text-primary-600 dark:hover:text-primary-400 transition-colors duration-200"
								@click="navigateToUserDetails(data)" v-tooltip.top="'View user details'">
								{{ getUserDisplayName(data) }}
							</span>
							<span class="text-sm text-surface-500 dark:text-surface-400">
								{{ data.email }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<Column field="name" header="Name" sortable style="width: 25%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex flex-col gap-1">
						<span class="text-surface-700 dark:text-surface-200">
							{{ data.firstName && data.lastName ? `${data.lastName} ${data.firstName}` : 'Not provided'
							}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="status" style="width: 25%">
				<template #header>
					<span class="flex-1 text-center font-bold">Status</span>
				</template>
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex items-center justify-center gap-2">
						<Tag v-if="usersWithRole.has(ulid(data.iD))" value="Assigned" severity="success" />
						<Tag v-else value="Unassigned" severity="secondary" />
					</div>
				</template>
			</Column>

			<Column field="actions" header="" style="width: 10%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex items-center gap-2 justify-end">
						<Button v-if="!usersWithRole.has(ulid(data.iD))" icon="pi pi-plus" severity="success" outlined
							size="small" @click="addRole(data)" v-tooltip.top="'Assign role'" />
						<Button v-else icon="pi pi-minus" severity="danger" outlined size="small"
							@click="removeRole(data)" v-tooltip.top="'Remove role'" />
					</div>
				</template>
			</Column>

			<template #empty>
				<div class="text-center p-4">
					<i class="pi pi-users text-4xl text-gray-400 mb-4"></i>
					<p class="text-gray-500">No users found.</p>
				</div>
			</template>
		</DataTable>
	</div>
</template>

<style scoped></style>
