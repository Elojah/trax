<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { useAuthStore } from '@/stores/auth';
import { ulid } from '@/utils/ulid';
import { ListEntityReq } from '@internal/user/dto/entity';
import type { ReadonlyHeaders } from '@/utils/headers';

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const pageOptions = [
	{
		value: 10,
		title: "10"
	},
	{
		value: 50,
		title: "50"
	},
	{
		value: 100,
		title: "100"
	},
]

const authStore = useAuthStore();
const store = useEntityStore();
const {
	entities: entities,
	total: total,
	selected: selected,
} = toRefs(store);

const name = ref('');
const avatarURL = ref('');
const description = ref('');

const loading = ref(false);
const search = ref('');

const headers: ReadonlyHeaders = [
	{
		title: 'Name',
		key: 'name',
		align: 'start',
		sortable: true,
	},
	{
		title: 'Created',
		key: 'created_at',
		align: 'end',
		sortable: true,
	},
];

const viewIDs = ref<string[]>([])

const views = computed(() => {
	return viewIDs.value.map((entityID: string) => entities.value?.get(entityID));
});

const select = (_: any, row: any) => {
	selected.value = [];
	row.toggleSelect({ value: row.item, selectable: true });
};

const list = async (options: any = { page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] }) => {
	loading.value = true;

	const { page, itemsPerPage, sortBy } = options;
	const newIDs = await store.list(ListEntityReq.create({
		own: true,
		search: search.value,
		paginate: {
			start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
			end: BigInt(page * itemsPerPage),
			sort: sortBy?.at(0)?.key ?? '',
			order: sortBy?.at(0)?.order === 'asc' ? true : false,
		}
	}));

	viewIDs.value = newIDs;

	loading.value = false;
};

defineExpose({
	list,
});

// #MARK:Entity dialogs
const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const dialogCreate = ref(false);
const closeCreate = () => {
	dialogCreate.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';
};

const create = async () => {
	await store.create(name.value, avatarURL.value, description.value);
	dialogCreate.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';

	await authStore.refreshToken();
	await list({ page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] })
};

const short = (description: string): string => {
	return description.length > 64 ? description.substring(0, 64) + '...' :
		description || 'No description'
}
</script>

<template>
	<v-container class="px-0">
		<v-row>
			<v-col cols="10">
				<v-text-field class="table-color-background" v-model="search" label="Search"
					prepend-inner-icon="mdi-magnify" variant="outlined" hide-details single-line>
				</v-text-field>
			</v-col>
			<v-col cols="2" class="d-flex align-center justify-end">
				<v-dialog v-model="dialogCreate" max-width="800px">
					<template v-slot:activator="{ props }">
						<v-btn variant="tonal" prepend-icon="mdi-plus-box" color="primary" size="large" v-bind="props">
							New
							<template v-slot:prepend>
								<v-icon color="primary"></v-icon>
							</template>
						</v-btn>
					</template>
					<v-card class="px-6 py-6 rounded" variant="elevated">
						<v-form ref="form" v-model="valid" lazy-validation>
							<v-card-title class="d-flex justify-center">
								<span class="text-h5">Create entity</span>
							</v-card-title>
							<v-divider></v-divider>
							<v-card-text>
								<v-container>
									<v-row>
										<v-col cols="6">
											<v-text-field v-model="name" :rules="nameRules" label="Name"></v-text-field>
										</v-col>
										<v-col cols="6">
											<v-text-field v-model="avatarURL" label="Avatar URL"></v-text-field>
										</v-col>
									</v-row>
									<v-row>
										<v-col cols="12">
											<v-textarea v-model="description" label="Description"></v-textarea>
										</v-col>
									</v-row>
								</v-container>
							</v-card-text>
							<v-card-actions>
								<v-spacer></v-spacer>
								<v-btn color="error" variant="text" @click="closeCreate">
									Cancel
								</v-btn>
								<v-btn color="primary" variant="text" @click="create">
									Create
								</v-btn>
							</v-card-actions>
						</v-form>
					</v-card>
				</v-dialog>
			</v-col>
		</v-row>
	</v-container>
	<v-data-table-server class="px-2 overflow-y-auto flex-grow-1" :headers="headers" fixed-footer fixed-header
		items-per-page-text="" :items-per-page-options="pageOptions" :items="views" :items-length="Number(total)"
		:loading="loading" :search="search" item-value="iD" @update:options="list" v-model="selected"
		@click:row="select" return-object item-selectable select-strategy="single">
		<template v-slot:item="{ item, columns, isSelected, index, props: itemProps }">
			<v-hover v-slot="{ isHovering, props: hoverProps }">
				<tr v-if="item" v-bind="{ ...itemProps, ...hoverProps }" :key="ulid(item.iD)">
					<td :colspan="columns.length" class="cursor-pointer px-1 py-1">
						<v-card class="justify-center border-primary" :class="{
							'row-hovered': isHovering,
							'row-even': index % 2 === 0,
							'row-odd': index % 2 !== 0,
							'row-selected': isSelected({ value: item, selectable: true }),
							'text-primary': isSelected({ value: item, selectable: true }),
						}" :title="item.name" :subtitle="short(item.description)">
							<template v-slot:prepend>
								<v-avatar class="mr-4" size="32" :color="!item.avatarURL ? 'primary' : ''">
									<img v-if="item.avatarURL" :src="item.avatarURL" alt="Avatar">
									<span v-else class=" mx-auto text-center text-h5">
										{{ item?.name?.at(0)?.toUpperCase() }}
									</span>
								</v-avatar>
							</template>
							<v-card-actions>
								<v-divider></v-divider>
								<p class="font-italic font-weight-light">
									{{ new Date(Number(item.createdAt) * 1000).toLocaleDateString('en-GB') }}
								</p>
							</v-card-actions>
						</v-card>
					</td>
				</tr>
			</v-hover>
		</template>
	</v-data-table-server>
</template>
<style scoped>
.table-color-background {
	background-color: #212121;
}

.row-odd {
	transition: background-color .2s ease-in-out;
	background-color: #616161;
}

.row-odd:not(.row-hovered) {
	background-color: #121212;
}

.row-even {
	transition: background-color .2s ease-in-out;
	background-color: #616161;
}

.row-even:not(.row-hovered) {
	background-color: #424242;
}

.row-selected {
	background-color: rgb(33, 47, 59) !important;
}

.cursor-pointer {
	cursor: pointer;
}
</style>
