<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { Entity } from '@internal/user/entity';
import { computed, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { marked } from "marked";
import type { VDataTable } from 'vuetify/components/VDataTable';
import { useAuthStore } from '@/stores/auth';

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const authStore = useAuthStore();

const entityStore = useEntityStore();
const {
	entity: entity,
	entities: entities,
	total: total,
} = toRefs(entityStore);

const name = ref('');
const avatarURL = ref('');
const description = ref('');

const loading = ref(true);
const search = ref('');
const tableView = ref(0)

type ReadonlyHeaders = VDataTable['$props']['headers']

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

const headersByRole: ReadonlyHeaders = [
	{
		title: 'Role',
		key: 'role',
		align: 'start',
		sortable: true,
	},
	{
		title: 'Name',
		key: 'name',
		align: 'start',
		sortable: true,
	},
];

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

const tableEntities = computed(() => {
	return entities.value?.map((entity) => {
		return {
			...entity,
		};
	});
});

const mdDescription = computed(() => {
	return marked.parse(entity.value?.description ?? `*no description*`);
});

const listEntity = async (options: any) => {
	loading.value = true;
	const { page, itemsPerPage, sortBy } = options;
	await entityStore.listEntity(search.value, {
		start: BigInt((page - 1) * itemsPerPage), // page starts at 1
		end: BigInt(page * itemsPerPage),
		sort: sortBy?.at(0)?.key ?? '',
		order: sortBy?.at(0)?.order === 'asc' ? true : false,
	});

	loading.value = false;
};

const dialog = ref(false);
const close = () => {
	dialog.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';
};

const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const create = async () => {
	await entityStore.createEntity(name.value, avatarURL.value, description.value);
	dialog.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';

	await authStore.refreshToken();
	await listEntity({ page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] })
};

const displayEntity = (_: any, row: { item: Entity }) => {
	entity.value = row.item;
};

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDeleteEntity = () => {
	// TODO: delete entity
	dialogDelete.value = false;
};

const editEntity = () => {
};
const deleteEntity = () => {
};

</script>

<template>
	<v-dialog v-model="dialogDelete" max-width="500px">
		<v-card>
			<v-card-title class="text-h5">Are you sure you want to delete this entity ?</v-card-title>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn variant="text" @click="closeDelete">Cancel</v-btn>
				<v-btn color="danger" variant="text" @click="confirmDeleteEntity">OK</v-btn>
				<v-spacer></v-spacer>
			</v-card-actions>
		</v-card>
	</v-dialog>

	<v-row>
		<v-col class="ml-8 mt-8 main-color-background" cols="4">
			<v-row>
				<v-col cols="9">
					<v-btn-toggle v-model="tableView">
						<v-btn size="large" variant="outlined" append-icon="mdi-domain">
							By Name
							<template v-slot:append>
								<v-icon color="primary"></v-icon>
							</template>
						</v-btn>
						<v-btn size="large" variant="outlined" append-icon="mdi-account-cog">
							By Role
							<template v-slot:append>
								<v-icon color="primary"></v-icon>
							</template>
						</v-btn>
					</v-btn-toggle>
				</v-col>
				<v-col class="d-flex justify-end align-center" cols="3">
					<v-dialog v-model="dialog" max-width="800px">
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
													<v-text-field v-model="name" :rules="nameRules"
														label="Name"></v-text-field>
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
										<v-btn variant="text" @click="close">
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
			</v-row>
			<v-row>
				<v-text-field v-model="search" label="Search" prepend-inner-icon="mdi-magnify" variant="outlined"
					hide-details single-line></v-text-field>
			</v-row>
			<v-row>
				<v-data-table-server v-if="tableView === 0" class="transparent-background" :headers="headers"
					fixed-footer min-height="50vh" max-height="100vh" items-per-page-text=""
					:items-per-page-options="pageOptions" :items="tableEntities" :items-length="Number(total)"
					:loading="loading" :search="search" item-value="name" @update:options="listEntity"
					@click:row="displayEntity" select-strategy="single" item-selectable="true">
					<template v-slot:item="{ item, index, props }">
						<tr class="mt-4 mb-4 cursor-pointer" v-bind="props"
							v-bind:class="index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'">
							<td class="d-flex align-center">
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
					</template>
				</v-data-table-server>
				<v-data-table-server v-if="tableView === 1" class="transparent-background" :headers="headersByRole"
					fixed-footer min-height="50vh" max-height="100vh" items-per-page-text=""
					:items-per-page-options="pageOptions" :items="tableEntities" :items-length="Number(total)"
					:loading="loading" :search="search" item-value="name" @update:options="listEntity"
					@click:row="displayEntity" select-strategy="single" item-selectable="true">
					<template v-slot:item="{ item, index, props }">
						<tr class="mt-4 mb-4 cursor-pointer" v-bind="props"
							v-bind:class="index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'">
							<td class="d-flex align-center">
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
					</template>
				</v-data-table-server>
			</v-row>
		</v-col>
		<v-divider vertical></v-divider>
		<v-col class="mx-auto" cols="6">
			<v-sheet class="px-1 rounded-xl" outlined color="primary">
				<v-card class="mt-8 px-6 py-6 justify-center rounded-xl main-color-background justify-center"
					variant="flat" v-if="entity" :title="entity.name">
					<template v-slot:prepend>
						<v-avatar size="32" :color="!entity.avatarURL ? 'primary' : ''">
							<img v-if="entity.avatarURL" :src="entity.avatarURL" alt="Avatar">
							<span v-else-if="!entity.avatarURL" class=" mx-auto text-center text-h5">
								{{ entity?.name?.at(0)?.toUpperCase() }}
							</span>
						</v-avatar>
					</template>
					<!-- eslint-disable vue/no-v-html vue/no-v-text-v-html-on-component -->
					<v-card-text class="p-6 m-6" variant="tonal" v-html="mdDescription">
					</v-card-text>
					<!--eslint-enable-->
					<v-card-actions>
						<v-spacer></v-spacer>
						<v-btn color="primary" @click="editEntity()">Edit</v-btn>
						<v-btn color="error" @click="deleteEntity()">Delete</v-btn>
					</v-card-actions>
				</v-card>
			</v-sheet>
			<v-divider></v-divider>
		</v-col>
	</v-row>
</template>
<style scoped>
.main-color-background {
	background-color: #263238;
}

.transparent-background {
	background-color: transparent;
}

.row-bg-odd {
	background-color: rgba(33, 33, 33, 0.3);
}

.row-bg-odd:hover {
	background-color: rgba(55, 71, 79, 0.3);
}

.row-bg-even {
	background-color: rgba(0, 229, 255, 0.3);
}

.row-bg-even:hover {
	background-color: rgba(55, 71, 79, 0.3);
}

.cursor-pointer {
	cursor: pointer;
}
</style>
