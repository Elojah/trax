<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import type { VDataTable } from 'vuetify/components/VDataTable';
import { useAuthStore } from '@/stores/auth';
import { ulid } from '@/utils/ulid';
import {  ListEntityReq } from '@internal/user/dto/entity';

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

type ReadonlyHeaders = VDataTable['$props']['headers']

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

const loading = ref(true);
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
		userIDs: true,
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
</script>

<template>
	<v-sheet class="px-1 rounded-xl" outlined color="primary">
		<v-col class="d-flex justify-end align-center rounded-t-xl main-color-background" cols="12">
			<v-dialog v-model="dialogCreate" max-width="800px">
				<template v-slot:activator="{ props }">
					<v-btn variant="tonal" prepend-icon="mdi-plus-box" color="primary" v-bind="props">
						New
						<template v-slot:prepend>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
				</template>
				<v-sheet class="px-1 rounded-xl" outlined color="primary">
					<v-card class="px-6 py-6 rounded-xl" variant="elevated">
						<v-form ref="form" v-model="valid" lazy-validation>
							<v-card-title>
								<span class="text-h6">New entity</span>
							</v-card-title>
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
								<v-btn variant="text" @click="closeCreate">
									Cancel
								</v-btn>
								<v-btn color="primary" variant="text" @click="create">
									Create
								</v-btn>
							</v-card-actions>
						</v-form>
					</v-card>
				</v-sheet>
			</v-dialog>
		</v-col>
		<v-text-field class="main-color-background" v-model="search" label="Search"
			prepend-inner-icon="mdi-magnify" variant="outlined" hide-details single-line>
		</v-text-field>
		<v-data-table-server class="main-color-background table-spacing rounded-0" :headers="headers" fixed-footer
			min-height="50vh" max-height="100vh" items-per-page-text="" :items-per-page-options="pageOptions"
			:items="views" :items-length="Number(total)" :loading="loading" :search="search"
			item-value="name" item-key="iD" @update:options="list" v-model="selected"
			@click:row="select" return-object item-selectable select-strategy="single">
			<template v-slot:item="{ item, isSelected, index, props: itemProps }">
				<v-hover v-slot="{ isHovering, props: hoverProps }">
					<tr v-if="item" v-bind="{ ...itemProps, ...hoverProps }" class="cursor-pointer py-8"
						:key="ulid(item.iD)" :class="{
							'row-hovered': isHovering,
							'row-selected': isSelected({ value: item, selectable: true }),
							'row-even': index % 2 === 0,
							'row-odd': index % 2 !== 0,
						}">
						<td>
							<v-avatar class="mr-4" size="32" :color="!item.avatarURL ? 'primary' : ''">
								<img v-if="item.avatarURL" :src="item.avatarURL" alt="Avatar">
								<span v-else class=" mx-auto text-center text-h5">
									{{ item?.name?.at(0)?.toUpperCase() }}
								</span>
							</v-avatar>
							<span class="text-h6">{{ item.name }}</span>
						</td>
						<td class="text-caption text-right">{{ new Date(Number(item.createdAt) *
							1000).toLocaleDateString('en-GB')
							}}
						</td>
					</tr>
				</v-hover>
			</template>
		</v-data-table-server>
		<v-col cols="12" class="p-8 main-color-background rounded-b-xl"></v-col>
	</v-sheet>
</template>
<style scoped>
.main-color-background {
	background-color: #263238;
}

.row-odd {
	transition: background-color .2s ease-in-out;
	background-color: rgba(55, 71, 79, 0.3);
}

.row-odd:not(.row-hovered) {
	background-color: rgba(33, 33, 33, 0.3);
}

.row-even {
	transition: background-color .2s ease-in-out;
	background-color: rgba(55, 71, 79, 0.3);
}

.row-even:not(.row-hovered) {
	background-color: rgba(0, 229, 255, 0.3);
}

.row-selected {
	background-color: rgba(149, 12, 117, 0.6) !important;
}

.cursor-pointer {
	cursor: pointer;
}
</style>