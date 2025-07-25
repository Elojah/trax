<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useUserStore } from '@/stores/user';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
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
import Dialog from 'primevue/dialog';
import ConfirmDialog from 'primevue/confirmdialog';
import Message from 'primevue/message';
import Avatar from 'primevue/avatar';
import Card from 'primevue/card';
import Skeleton from 'primevue/skeleton';
import Textarea from 'primevue/textarea';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

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
const userStore = useUserStore();
const { users, total } = toRefs(userStore);
const confirmUser = useConfirm();

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

// Dialog states
const dialogInviteUser = ref(false);

// Form validation
const resolver = zodResolver(
	z.object({
		email: z.string().email({ message: 'Please enter a valid email address.' }),
	})
);

const initialValues = ref({
	email: '',
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

		const newIDs = await userStore.list(ListUserReq.create({
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

const openInviteUser = () => {
	initialValues.value = {
		email: '',
	};
	dialogInviteUser.value = true;
};

const inviteUser = async (e: FormSubmitEvent) => {
	if (!e.valid || !group.value) {
		logger.error('Invite user form is invalid', e);
		return;
	}

	try {
		// TODO: Implement user invitation logic
		// await userStore.invite(e.states.email.value, group.value.iD);
		console.log('Inviting user:', e.states.email.value, 'to group:', group.value.group?.name);
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `User ${e.states.email.value} invited successfully`;
	success.value = true;
	dialogInviteUser.value = false;
	e.reset();
	await authStore.refreshToken();
	await list();
};

const removeUser = (user: U) => {
	if (!user) return;

	confirmUser.require({
		message: `Are you sure you want to remove "${user.email}" from this group? This action cannot be undone.`,
		header: 'Remove User',
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
				// TODO: Implement user removal logic
				// await userStore.removeFromGroup(user.iD, group.value.iD);
				console.log('Removing user:', user.email, 'from group');
			} catch (e) {
				errorsStore.showGRPC(e);
				return;
			}

			message.value = `User "${user.email}" removed successfully`;
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
			dataKey="id" filterDisplay="menu" :globalFilterFields="['email', 'firstName', 'lastName', 'created_at']"
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
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Users</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">Manage group members</span>
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
						<Button label="" icon="pi pi-plus" outlined severity="primary" class="font-medium"
							@click="openInviteUser" />
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
							<span class="font-semibold text-surface-900 dark:text-surface-0">
								{{ getUserDisplayName(data) }}
							</span>
							<span class="text-sm text-surface-500 dark:text-surface-400">
								{{ data.email }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<Column field="firstName" header="Name" sortable style="width: 25%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex flex-col gap-1">
						<span class="text-surface-700 dark:text-surface-200">
							{{ data.firstName && data.lastName ? `${data.firstName} ${data.lastName}` : 'Not provided'
							}}
						</span>
					</div>
				</template>
			</Column>

			<Column field="created_at" header="Joined" sortable style="width: 20%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex flex-col gap-1">
						<div class="flex items-center gap-2">
							<i class="pi pi-calendar text-surface-500 dark:text-surface-400"></i>
							<span class="font-medium text-surface-700 dark:text-surface-200">
								{{ formatDate(data.createdAt) }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<Column field="actions" header="Actions" style="width: 20%">
				<template #body="{ data }: { data: U }">
					<div v-if="data" class="flex items-center gap-2">
						<Button icon="pi pi-trash" severity="danger" outlined size="small" class="p-2"
							@click="removeUser(data)" v-tooltip.top="'Remove from group'" />
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

		<!-- Invite User Dialog -->
		<Dialog v-model:visible="dialogInviteUser" append-to="body" modal
			:breakpoints="{ '960px': '75vw', '640px': '80vw' }" :style="{ width: '40rem' }" :draggable="false"
			:resizable="false" :show-header="false" class="shadow-sm rounded-2xl"
			:pt="{ content: '!p-6', footer: '!pb-6 !px-6' }">
			<div class="flex justify-between items-start gap-4">
				<div class="flex gap-4">
					<div
						class="flex items-center justify-center w-9 h-9 bg-green-100 dark:bg-green-400/30 text-green-600 dark:text-green-300 rounded-3xl border border-green-200 dark:border-green-400/20 shrink-0">
						<i class="pi pi-users !text-xl !leading-none" />
					</div>
				</div>
				<Button icon="pi pi-times" text rounded severity="secondary" class="w-10 h-10 !p-2"
					@click="dialogInviteUser = false" />
			</div>
			<Form v-slot="$form" :resolver="resolver" :initialValues="initialValues" @submit="inviteUser"
				class="flex flex-col gap-6">
				<div class="flex items-start gap-4">
					<div class="flex-1 flex flex-col gap-2">
						<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Invite
							user
						</h1>
						<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">Send an invitation
							to
							join this group.</span>
					</div>
				</div>
				<div class="flex flex-col gap-6">
					<FormField v-slot="$field" name="email" class="flex flex-col gap-2">
						<label for="invite-email" class="text-color text-base">Email</label>
						<IconField icon-position="left" class="w-full">
							<InputIcon class="pi pi-envelope" />
							<InputText id="invite-email" name="email" placeholder="Enter user email" type="email"
								class="w-full" />
						</IconField>
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message
						}}
						</Message>
					</FormField>
				</div>
				<div class="flex justify-end gap-4">
					<Button label="Cancel" outlined @click="dialogInviteUser = false" />
					<Button label="Send Invitation" type="submit" />
				</div>
			</Form>
		</Dialog>
	</div>
</template>

<style scoped></style>
