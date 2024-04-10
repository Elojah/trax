<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { Entity } from '@internal/user/entity';
import { computed, onMounted, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { marked } from "marked";

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const entityStore = useEntityStore();
const {
	entity: entity,
	entities: entities,
	total: total,
	newEntity: newEntity,
} = toRefs(entityStore);

const mdDescription = computed(() => {
	return marked.parse(entity.value?.description ?? '');
});

const itemsPerPage = ref(10);
const loading = ref(true);
const search = ref('');

const headers = [
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

const tableEntities = computed(() => {
	return entities.value?.map((entity) => {
		return {
			...entity,
		};
	});
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
	newEntity.value = Entity.create({})
};

const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const create = async () => {
	await entityStore.createEntity();
	dialog.value = false;
	newEntity.value = {} as Entity;
	search.value = '0_0';
	search.value = ''; // Set empty search to trigger table reload
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
	<v-toolbar class="bar-bg px-6" floating>
		<v-icon color="primary" size="large" icon="mdi-domain"></v-icon>
		<v-toolbar-title class="text-h5 font-weight-black font-italic"> Entities </v-toolbar-title>
		<v-dialog v-model="dialog" max-width="800px">
			<template v-slot:activator="{ props }">
				<v-btn variant="tonal" prepend-icon="mdi-plus-box" v-bind="props">
					New entity
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
										<v-text-field v-model="newEntity.name" :rules="nameRules"
											label="Name"></v-text-field>
									</v-col>
									<v-col cols="6">
										<v-text-field v-model="newEntity.avatarURL" label="Avatar URL"></v-text-field>
									</v-col>
								</v-row>
								<v-row>
									<v-col cols="12">
										<v-text-area v-model="newEntity.description" label="Description"></v-text-area>
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
	</v-toolbar>
	<v-row>
		<v-col class="mx-auto pt-8" cols="4">
			<v-text-field v-model="search" label="Search" prepend-inner-icon="mdi-magnify" variant="outlined"
				hide-details single-line></v-text-field>
			<v-data-table-server class="transparent-background" v-model:items-per-page="itemsPerPage" :headers="headers"
				:items="tableEntities" :items-length="Number(total)" :loading="loading" :search="search"
				item-value="name" :options="{ sortBy: [{ key: 'created_at', order: 'desc' }] }"
				@update:options="listEntity" @click:row="displayEntity" select-strategy="single" item-selectable="true">
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
		</v-col>
		<v-divider vertical></v-divider>
		<v-col class="mx-auto" cols="8">
			<v-sheet class="px-1 rounded-xl" outlined color="primary">
				<v-card class="mt-8 px-6 py-6 justify-center rounded-xl main-color-background justify-center"
					v-if="entity" :title="entity.name">
					<template v-slot:prepend>
						<v-avatar size="32" :color="!entity.avatarURL ? 'primary' : ''">
							<img v-if="entity.avatarURL" :src="entity.avatarURL" alt="Avatar">
							<span v-else-if="!entity.avatarURL" class=" mx-auto text-center text-h5">
								{{ entity?.name?.at(0)?.toUpperCase() }}
							</span>
						</v-avatar>
					</template>
					<v-card-text v-model="mdDescription">
					</v-card-text>
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
.bar-bg {
	background: url('@/assets/img/bar-background.svg') no-repeat bottom center;
	background-size: cover;
}

.main-color-background {
	background-color: #263238;
}

.transparent-background {
	background-color: transparent;
}

.row-bg-odd {
	background-color: rgba(33, 33, 33, 0.7);
}

.row-bg-odd:hover {
	background-color: rgba(38, 50, 56, 0.7);
}

.row-bg-even {
	background-color: rgba(66, 66, 66, 0.7);
}

.row-bg-even:hover {
	background-color: rgba(55, 71, 79, 0.7);
}

.cursor-pointer {
	cursor: pointer;
}
</style>
