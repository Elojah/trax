<script setup lang="ts">
import { useGroupStore } from '@/stores/group';
import { computed, ref, toRefs, onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { ulid } from '@/utils/ulid';
import { ListGroupReq } from '@internal/user/dto/group';
import { Group } from '@internal/user/group';
import { grpccodes } from '@/utils/errors';
import { logger } from "@/config";

import DataTable, { type DataTableFilterEvent, type DataTablePageEvent, type DataTableProps, type DataTableSortEvent } from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import Avatar from 'primevue/avatar';
import Dialog from 'primevue/dialog';
import Textarea from 'primevue/textarea';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import z from 'zod';
import type { FormEmits, FormSubmitEvent } from '@primevue/forms';


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

// const select = (_: any, row: any) => {
// 	selected.value = [];
// 	row.toggleSelect({ value: row.item, selectable: true });
// };

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

const visible = ref(false);
const close = () => {
	visible.value = false;
	// name.value = '';
	// avatarURL.value = '';
	// description.value = '';
};

const create = async (e: FormSubmitEvent) => {
	logger.info('Create group form is invalid', e.states);
	if (!e.valid) {
		logger.error('Create group form is invalid', e);
		return;
	}

	let ok = true;
	try {
		await store.create(e.states.name.value, e.states['avatar-url'].value, e.states.description.value);
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false
	}

	if (ok) {
		message.value = `Group ${e.states.name.value} created successfully`;
		success.value = true;
	}

	visible.value = false;
	e.reset()

	await authStore.refreshToken();
	await list();
};

const short = (description: string): string => {
	return description.length > 64 ? description.substring(0, 64) + '...' :
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
		v-model:selection="selected" @page="onPage" @sort="onSort" @filter="onFilter" dataKey="id" filterDisplay="menu"
		:globalFilterFields="['name', 'created_at']" tableStyle="min-width: 50rem"
		:rowsPerPageOptions="[10, 25, 50, 100]"
		paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
		currentPageReportTemplate="{first} - {last} ({totalRecords})">

		<template #header>
			<div class="flex justify-between items-center gap-4">
				<IconField icon-position="left">
					<InputIcon class="pi pi-search text-surface-400" />
					<InputText type="text" class="w-full sm:w-80" placeholder="Search group" v-model="search"
						@input="onSearch" />
				</IconField>
				<Button label="Create Group" icon="pi pi-plus-circle" @click="visible = true" />
			</div>
		</template>

		<Column field="name" header="Name" sortable style="width: 25%">
			<template #body="{ data }: { data: Group }">
				<div class="flex items-center gap-2">
					<Avatar v-if="data.avatarURL" :image="data.avatarURL" size="large" shape="circle" />
					<Avatar v-else :label="data.name.charAt(0).toUpperCase()" size="large" shape="circle" />
					<span class="font-semibold">{{ data.name }}</span>
				</div>
			</template>
		</Column>

		<Column field="description" header="Description" style="width: 50%">
			<template #body="{ data }: { data: Group }">
				<span>{{ short(data.description) }}</span>
			</template>
		</Column>

		<Column field="created_at" header="Created" sortable style="width: 20%">
			<template #body="{ data }: { data: Group }">
				<span>{{ new Date(Number(data.createdAt) * 1000).toLocaleDateString('en-GB') }}</span>
			</template>
		</Column>

		<Column header="" style="width: 5%">
			<template #body="{ data }: { data: Group }">
				<Button icon="pi pi-ellipsis-v" severity="secondary" text
					@click="(event) => {/* Add your menu logic here */ }" aria-haspopup="true" />
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
	<Dialog v-model:visible="visible" append-to="body" modal :breakpoints="{ '960px': '75vw', '640px': '80vw' }"
		:style="{ width: '40rem' }" :draggable="false" :resizable="false" :show-header="false"
		class="shadow-sm rounded-2xl" :pt="{ content: '!p-6', footer: '!pb-6 !px-6' }">
		<div class="flex gap-4">
			<div
				class="flex items-center justify-center w-9 h-9 bg-primary-100 dark:bg-primary-400/30 text-primary-500 dark:text-primary-200 rounded-3xl border border-primary-200 dark:border-primary-400/20 shrink-0">
				<i class="pi pi-users !text-xl !leading-none" />
			</div>
		</div>
		<Form v-slot="$form" :resolver @submit="create" class="flex flex-col gap-6">
			<div class="flex items-start gap-4">
				<div class="flex-1 flex flex-col gap-2">
					<h1 class="m-0 text-surface-900 dark:text-surface-0 font-semibold text-xl leading-normal">Create
						New group</h1>
					<span class="text-surface-500 dark:text-surface-400 text-base leading-normal">Manage your
						users seamlessly.</span>
				</div>
				<Button icon="pi pi-times" text rounded severity="secondary" class="w-10 h-10 !p-2"
					@click="visible = false" />
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
				<Button label="Cancel" outlined severity="secondary" @click="close" />
				<Button label="Create" type="submit" severity="secondary" />
			</div>
		</Form>
	</Dialog>
</template>
<style scoped></style>
