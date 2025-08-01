<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useUserStore } from '@/stores/user';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { formatDate } from '@/utils/date';
import { createPagination } from '@/utils/requests';
import { useTable } from '@/composables';
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

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import { hasClaim } from '@/utils/claims';
import { Command, Resource } from '@internal/user/role';

const props = defineProps<{
	groupId: string;
}>();

const router = useRouter();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const userStore = useUserStore();
const { users, total } = toRefs(userStore);
const authStore = useAuthStore();
const { claims } = toRefs(authStore);

const group = groups.value.get(props.groupId) || null;

const writeUserPermission = computed(() => {
	return hasClaim(claims.value, group?.group?.iD!, Resource.R_user, Command.C_write);
});

// Create table composable with request creation function
const createRequest = (props: DataTableProps, search: string) => {
	if (!group) {
		return ListUserReq.create({});
	}

	return ListUserReq.create({
		groupIDs: [group.group?.iD],
		search,
		paginate: createPagination(props)
	});
};

const table = useTable({
	createRequest,
	listMethod: userStore.list,
	itemsMap: users,
	total
});

const { views, handlers, search, loading } = table;

const openInvitationCreate = () => {
	router.push({
		name: 'group-invite',
		params: {
			groupId: props.groupId
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
			filterDisplay="menu" :globalFilterFields="['email', 'name', 'created_at']" tableStyle="min-width: 50rem"
			:rowsPerPageOptions="[10, 25, 50, 100]"
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
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Users</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">Manage group members</span>
							</div>
						</div>
						<div class="flex items-center gap-3 ml-6">
							<IconField icon-position="left" class="relative">
								<InputIcon class="pi pi-search text-surface-400" />
								<InputText type="text"
									class="w-96 pl-10 pr-4 py-2 border border-surface-300 dark:border-surface-600 rounded-lg bg-white dark:bg-surface-900 focus:border-primary-500 dark:focus:border-primary-400 focus:ring-2 focus:ring-primary-200 dark:focus:ring-primary-800 transition-all duration-200"
									placeholder="Search users by name or email..." v-model="search"
									@input="handlers.onSearch" />
							</IconField>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10"
							@click="table.list()" v-tooltip.bottom="'Refresh users'" />
						<Button :disabled="!writeUserPermission" label="" icon="pi pi-plus" outlined severity="primary"
							class="font-medium" @click="openInvitationCreate" />
					</div>
				</div>
			</template>

			<Column field="user" header="User" sortable style="width: 35%">
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

			<Column field="email" header="Email" sortable style="width: 25%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex flex-col gap-1">
						<span class="text-surface-700 dark:text-surface-200">
							{{ data.email }}
						</span>
					</div>
				</template>
			</Column>

			<Column field="created_at" sortable style="width: 20%">
				<template #header>
					<span class="flex-1 text-right font-bold">Created</span>
				</template>
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex flex-col justify-end gap-1">
						<div class="flex items-center justify-end gap-2">
							<i class="pi pi-calendar text-surface-500 dark:text-surface-400"></i>
							<span class="font-medium text-surface-700 dark:text-surface-200">
								{{ formatDate(data.createdAt) }}
							</span>
						</div>
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
