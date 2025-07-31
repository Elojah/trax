<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';
import { useInvitationStore } from '@/stores/invitation';

// Internal utilities and types
import { parse, ulid } from '@/utils/ulid';
import { logger } from "@/config";
import { formatDate } from '@/utils/date';
import { createPagination } from '@/utils/requests';
import { useTable } from '@/composables';

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
import Badge from 'primevue/badge';
import Tag from 'primevue/tag';

// PrimeVue Input Components
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

// Form Validation
import { useConfirm } from 'primevue/useconfirm';
import { DeleteInvitationReq, InvitationView, ListInvitationReq } from '@internal/user/dto/invitation';

const props = defineProps<{
	groupId: string;
}>();

const router = useRouter();
const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);
const groupStore = useGroupStore();
const { groups } = toRefs(groupStore);
const invitationStore = useInvitationStore();
const { invitations, total } = toRefs(invitationStore);
const confirm = useConfirm();

const group = groups.value.get(props.groupId) || null;
// Create request function for the table
const createRequest = (props: DataTableProps, search: string) => {
	return ListInvitationReq.create({
		search,
		paginate: createPagination(props),
		groupIDs: [group?.group?.iD]
	});
};

const table = useTable({
	createRequest,
	listMethod: invitationStore.list,
	itemsMap: invitations,
	total
});

const { views, handlers, search, loading } = table;

// Actions
const deleteInvitation = (invitation: InvitationView) => {
	if (!invitation.invitation) return;

	confirm.require({
		message: `Are you sure you want to delete the invitation for "${invitation.invitation.email}"? This action cannot be undone.`,
		header: 'Delete Invitation',
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
			if (!invitation.invitation) return;

			const email = invitation.invitation.email;

			try {
				await invitationStore.delete_(DeleteInvitationReq.create({
					iD: invitation.invitation.iD,
					groupID: group?.group?.iD
				}));
			} catch (e) {
				errorsStore.showGRPC(e);
				return;
			}

			message.value = `Invitation for "${email}" deleted successfully`;
			success.value = true;
			await table.list();
		}
	});
};

const resendInvitation = async (invitation: InvitationView) => {
	if (!invitation.invitation) return;

	const email = invitation.invitation.email;
	const invitationId = ulid(invitation.invitation.iD);

	try {
		await invitationStore.resendInvitation(invitationId);
	} catch (e) {
		errorsStore.showGRPC(e);
		return;
	}

	message.value = `Invitation resent to "${email}" successfully`;
	success.value = true;
	await table.list();
};


const openInvitationCreate = () => {
	router.push({
		name: 'group-invite',
		params: {
			groupId: props.groupId
		}
	});
};

// Initialize data on component mount
onMounted(async () => {
	await table.list();
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
			filterDisplay="menu" :globalFilterFields="['invitation.email']" tableStyle="min-width: 50rem"
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
								<i class="pi pi-envelope text-lg text-purple-600 dark:text-purple-300"></i>
							</div>
							<div class="flex flex-col">
								<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Invitations
								</h3>
								<span class="text-sm text-surface-500 dark:text-surface-400">
									Manage pending invitations for {{ group?.group?.name || 'this group' }}
								</span>
							</div>
						</div>
						<div class="flex items-center gap-3 ml-6">
							<IconField icon-position="left" class="relative">
								<InputIcon class="pi pi-search text-surface-400" />
								<InputText type="text"
									class="w-96 pl-10 pr-4 py-2 border border-surface-300 dark:border-surface-600 rounded-lg bg-white dark:bg-surface-900 focus:border-primary-500 dark:focus:border-primary-400 focus:ring-2 focus:ring-primary-200 dark:focus:ring-primary-800 transition-all duration-200"
									placeholder="Search invitations..." v-model="search" @input="handlers.onSearch" />
							</IconField>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10"
							@click="table.list()" v-tooltip.bottom="'Refresh invitations'" />
						<Button label="" icon="pi pi-plus" outlined severity="primary" class="font-medium"
							@click="openInvitationCreate" />

					</div>
				</div>
			</template>

			<!-- Email Column -->
			<Column field="invitation.email" header="Email" sortable class="min-w-64">
				<template #body="{ data }: { data: InvitationView }">
					<div class="flex items-center gap-3">
						<Avatar :label="data.invitation?.email?.charAt(0).toUpperCase()" shape="circle"
							class="bg-gradient-to-br from-blue-400 to-blue-600 text-white text-sm font-medium" />
						<div class="flex flex-col">
							<span class="text-surface-900 dark:text-surface-0 font-medium">
								{{ data.invitation?.email }}
							</span>
							<span class="text-surface-500 dark:text-surface-400 text-xs">
								Invited {{ formatDate(data.invitation?.createdAt) }}
							</span>
						</div>
					</div>
				</template>
			</Column>

			<!-- Roles Column -->
			<Column field="roles" header="Roles" class="min-w-48">
				<template #body="{ data }: { data: InvitationView }">
					<div class="flex flex-wrap gap-1">
						<Tag v-for="role in data.roleSample?.slice(0, 9)" :key="ulid(role.iD)" :value="role.name"
							severity="secondary" class="text-xs" />
						<Tag v-if="Number(data.roleCount) > 9" :value="`+${Number(data.roleCount) - 9}`"
							severity="primary" class="text-xs" />
						<span v-if="!data.roleSample || data.roleSample.length === 0"
							class="text-surface-400 dark:text-surface-500 text-sm">No roles</span>
					</div>
				</template>
			</Column>

			<!-- Role Count Column -->
			<Column field="roleCount" header="Role Count" sortable class="min-w-32">
				<template #body="{ data }: { data: InvitationView }">
					<Badge :value="Number(data.roleCount) || 0" severity="info" />
				</template>
			</Column>

			<!-- Status Column -->
			<Column header="Status" class="min-w-32">
				<template #body="{ data }: { data: InvitationView }">
					<Tag value="Pending" severity="warning" />
				</template>
			</Column>

			<!-- Created Date Column -->
			<Column field="invitation.createdAt" header="Invited" sortable class="min-w-40">
				<template #body="{ data }: { data: InvitationView }">
					<div class="flex items-center gap-2">
						<i class="pi pi-calendar text-surface-400 text-sm"></i>
						<span class="text-surface-600 dark:text-surface-300 text-sm">
							{{ formatDate(data.invitation?.createdAt) }}
						</span>
					</div>
				</template>
			</Column>

			<!-- Actions Column -->
			<Column header="" class="min-w-48">
				<template #body="{ data }: { data: InvitationView }">
					<div class="flex items-center justify-end gap-2">
						<Button icon="pi pi-send" severity="info" size="small" outlined @click="resendInvitation(data)"
							v-tooltip.bottom="'Resend invitation'" />
						<Button icon="pi pi-trash" severity="danger" size="small" outlined
							@click="deleteInvitation(data)" v-tooltip.bottom="'Delete invitation'" />
					</div>
				</template>
			</Column>

			<!-- Empty State -->
			<template #empty>
				<div class="text-center p-8">
					<i class="pi pi-envelope text-4xl text-surface-400 mb-4"></i>
					<h3 class="text-xl font-semibold text-surface-900 dark:text-surface-0 mb-2">No Invitations</h3>
					<p class="text-surface-500 dark:text-surface-400">
						There are no pending invitations for this group.
					</p>
				</div>
			</template>
		</DataTable>
	</div>
</template>

<style scoped></style>
