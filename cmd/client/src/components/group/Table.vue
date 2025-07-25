<script setup lang="ts">
// Vue and Store imports
import { computed, ref, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useGroupStore } from '@/stores/group';

// Internal utilities and types
import { ulid } from '@/utils/ulid';
import { grpccodes } from '@/utils/errors';
import { logger } from "@/config";
import { DeleteGroupReq, ListGroupReq, UpdateGroupReq } from '@internal/user/dto/group';
import { Group } from '@internal/user/group';

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
import AvatarGroup from 'primevue/avatargroup';
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

const router = useRouter();
const authStore = useAuthStore();
const store = useGroupStore();
const {
	groups: groups,
	total: total,
	selected: selected,
} = toRefs(store);

const errorsStore = useErrorsStore()
const { success, message } = toRefs(errorsStore)

const loading = ref(false);
const search = ref('');

const viewIDs = ref<string[]>([])

const views = computed(() => {
	return viewIDs.value.map((groupID: string) => groups.value?.get(groupID));
});

const properties = ref<DataTableProps>({});

const select = () => { };

const list = async (props: DataTableProps = {
	first: 0,
	rows: 10,
	sortField: 'created_at',
	sortOrder: -1, // -1 for desc, 1 for asc
}) => {
	loading.value = true;

	try {
		// Use event data if provided (from DataTable lazy loading), otherwise use current p
		const page = Math.floor((props.first ?? 0) / (props.rows ?? 1)) + 1; // Convert first index to page number

		// Map PrimeVue sort order to your API format
		const sortBy = props.sortField ? [{
			key: props.sortField,
			order: props.sortOrder === 1 ? 'asc' : 'desc'
		}] : [{ key: 'created_at', order: 'desc' }];

		const newIDs = await store.list(ListGroupReq.create({
			own: true,
			search: search.value,
			paginate: {
				start: BigInt(((page - 1) * (props.rows ?? 10)) + 1), // page starts at 1, start starts at 1
				end: BigInt(page * (props.rows ?? 10)),
				sort: sortBy?.at(0)?.key ?? '',
				order: sortBy?.at(0)?.order === 'asc' ? true : false,
			}
		}));

		viewIDs.value = newIDs;

		// Update props if event is provided
		if (props) {
			properties.value = props
		}

	} catch (e) {
		errorsStore.showGRPC(e)
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
	// Trigger a new search with current parameters
	list({
		first: 0, // Reset to first page when searching
		rows: properties.value.rows ?? 10,
		sortField: properties.value.sortField,
		sortOrder: properties.value.sortOrder ?? -1,
		filters: properties.value.filters,
	});
};

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

const dialogCreateGroup = ref(false);

const close = () => {
	dialogCreateGroup.value = false;
};

const navigateToDetails = (group: Group) => {
	router.push({ name: 'group-details', params: { id: ulid(group.iD) } });
};

const create = async (e: FormSubmitEvent) => {
	if (!e.valid) {
		logger.error('Create group form is invalid', e);
		return;
	}

	try {
		await store.create(e.states.name.value, e.states["avatar-url"].value, e.states.description.value);
	} catch (e) {
		errorsStore.showGRPC(e)
		return
	}

	message.value = `Group ${e.states.name.value} created successfully`;
	success.value = true;
	dialogCreateGroup.value = false;
	e.reset()
	await authStore.refreshToken();
	await list();
};

const short = (description: string): string => {
	return description.length > 128 ? description.substring(0, 128) + '...' :
		description || 'No description'
}

// Initialize data on component mount
onMounted(() => {
	list();
});
</script>

<template>
	<DataTable :value="views" :lazy="true" :loading="loading" :paginator="true" :rows="properties.rows"
		:totalRecords="Number(total)" :first="properties.first" v-model:filters="properties.filters"
		v-model:selection="selected" @row-select="select" @page="onPage" @sort="onSort" @filter="onFilter" dataKey="id"
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
							class="flex items-center justify-center w-10 h-10 rounded-lg bg-blue-100 dark:bg-blue-400/30 border border-blue-200 dark:border-blue-400/20">
							<i class="pi pi-building text-lg text-blue-600 dark:text-blue-300"></i>
						</div>
						<div class="flex flex-col">
							<h2 class="text-lg font-semibold text-surface-900 dark:text-surface-0 m-0">Groups</h2>
							<span class="text-sm text-surface-500 dark:text-surface-400">Manage group, roles and
								users</span>
						</div>
					</div>
					<div class="flex items-center gap-3 ml-6">
						<IconField icon-position="left" class="relative">
							<InputIcon class="pi pi-search text-surface-400" />
							<InputText type="text"
								class="w-96 pl-10 pr-4 py-2 border border-surface-300 dark:border-surface-600 rounded-lg bg-white dark:bg-surface-900 focus:border-primary-500 dark:focus:border-primary-400 focus:ring-2 focus:ring-primary-200 dark:focus:ring-primary-800 transition-all duration-200"
								placeholder="Search groups by name or description..." v-model="search"
								@input="onSearch" />
						</IconField>
					</div>
				</div>
				<div class="flex items-center gap-3">
					<Button icon="pi pi-refresh" severity="secondary" outlined rounded class="w-10 h-10" @click="list()"
						v-tooltip.bottom="'Refresh groups'" />
					<Button label="Create group" icon="pi pi-plus" severity="primary" class="font-medium"
						@click="dialogCreateGroup = true" />
				</div>
			</div>
		</template>

		<Column field="name" header="Name" sortable style="width: 25%">
			<template #body="{ data }: { data: Group }">
				<div v-if="data" class="flex items-center gap-3">
					<div class="relative">
						<Avatar v-if="data.avatarURL" :image="data.avatarURL" size="large" shape="circle"
							class="border-2 border-surface-200 dark:border-surface-700" />
						<Avatar v-else :label="data.name.charAt(0).toUpperCase()" size="large" shape="circle"
							class="bg-primary-100 dark:bg-primary-400/30 text-primary-600 dark:text-primary-300 border-2 border-primary-200 dark:border-primary-400/20" />
					</div>
					<div class="flex flex-col">
						<span
							class="font-semibold text-surface-900 dark:text-surface-0 cursor-pointer hover:text-primary-600 dark:hover:text-primary-400 transition-colors duration-200"
							@click="navigateToDetails(data)" v-tooltip.top="'View group details'">
							{{ data.name }}
						</span>
					</div>
				</div>
			</template>
		</Column>

		<Column field="description" header="Description" style="width: 60%">
			<template #body="{ data }: { data: Group }">
				<div v-if="data" class="flex flex-col gap-1">
					<span class="text-surface-700 dark:text-surface-200 line-clamp-2">{{ short(data.description)
					}}</span>
					<div class="flex items-center gap-2 text-xs text-surface-500 dark:text-surface-400">
						<AvatarGroup>
							<Avatar v-for="member in ['Raton', 'Belette', 'Nyx', 'Test']" :key="member.charAt(0)"
								:label="member.charAt(0).toUpperCase()" size="small" shape="circle" />
							<Avatar label="+2" shape="circle" />
						</AvatarGroup>
					</div>
				</div>
			</template>
		</Column>

		<Column field="created_at" header="Created" sortable style="width: 15%">
			<template #body="{ data }: { data: Group }">
				<div v-if="data" class="flex flex-col gap-1">
					<div class="flex items-center gap-2">
						<i class="pi pi-calendar text-surface-500 dark:text-surface-400"></i>
						<span class="font-medium text-surface-700 dark:text-surface-200">
							{{ new Date(Number(data.createdAt) * 1000).toLocaleDateString('en-GB', {
								day: 'numeric',
								month: 'short',
								year: 'numeric'
							}) }}
						</span>
					</div>
					<span class="text-xs text-surface-500 dark:text-surface-400">
						{{ new Date(Number(data.createdAt) * 1000).toLocaleTimeString('en-GB', {
							hour: '2-digit',
							minute: '2-digit'
						}) }}
					</span>
				</div>
			</template>
		</Column>

		<template #empty>
			<div class="text-center p-4">
				<i class="pi pi-search text-4xl text-gray-400 mb-4"></i>
				<p class="text-gray-500">No groups found.</p>
			</div>
		</template>
	</DataTable>

	<!-- Create Group Dialog -->
	<Dialog v-model:visible="dialogCreateGroup" append-to="body" modal
		:breakpoints="{ '960px': '75vw', '640px': '80vw' }" :style="{ width: '40rem' }" :draggable="false"
		:resizable="false" :show-header="false" class="shadow-sm rounded-2xl"
		:pt="{ content: '!p-6', footer: '!pb-6 !px-6' }">
		<div class="flex justify-between items-start gap-4">
			<div class="flex gap-4">
				<div
					class="flex items-center justify-center w-9 h-9 bg-blue-100 dark:bg-blue-400/30 text-blue-600 dark:text-blue-300 rounded-3xl border border-blue-200 dark:border-blue-400/20 shrink-0">
					<i class="pi pi-building !text-xl !leading-none" />
				</div>
			</div>
			<Button icon="pi pi-times" text rounded severity="secondary" class="w-10 h-10 !p-2"
				@click="dialogCreateGroup = false" />
		</div>
		<Form v-slot="$form" :resolver :initialValues @submit="create" class="flex flex-col gap-6">
			<div class="flex items-start gap-4">
				<div class="flex-1 flex flex-col gap-2">
					<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Create
						group
					</h1>
					<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">You will be
						automatically
						assigned as the group's administrator.</span>
				</div>
			</div>
			<div class="flex flex-col gap-6">
				<FormField v-slot="$field" name="name" class="flex flex-col gap-2">
					<label for="cardholder" class="text-color text-base">Name</label>
					<IconField icon-position="left" class="w-full">
						<InputIcon class="pi pi-user" />
						<InputText id="name" name="name" placeholder="Enter group name" type="text" class="w-full" />
					</IconField>
					<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
						$field.error?.message
					}}
					</Message>
				</FormField>
				<FormField v-slot="$field" name="avatar-url" class="flex flex-col gap-2">
					<label for="avatar-url" class="text-color text-base">Avatar URL</label>
					<IconField icon-position="left" class="w-full">
						<InputIcon class="pi pi-image" />
						<InputText id="avatar-url" name="avatar-url" placeholder="Enter avatar URL" type="text"
							class="w-full" />
					</IconField>
					<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
						$field.error?.message
					}}
					</Message>
				</FormField>
				<FormField v-slot="$field" name="description" class="flex flex-col gap-2">
					<label for="description" class="text-color text-base">Description</label>
					<Textarea id="description" name="description" placeholder="Enter group description" rows="3"
						class="w-full" />
					<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
						$field.error?.message
					}}
					</Message>
				</FormField>
			</div>
			<div class="flex justify-end gap-4">
				<Button label="Cancel" outlined @click="close" />
				<Button label="Create" type="submit" />
			</div>
		</Form>
	</Dialog>

</template>
<style scoped></style>
