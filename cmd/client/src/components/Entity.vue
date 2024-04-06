<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { Entity } from '@internal/user/entity';
import { computed, onMounted, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const entityStore = useEntityStore();
const {
	entity: entity,
	entities: entities,
	total: total,
	newEntity: newEntity,
} = toRefs(entityStore);

const itemsPerPage = ref(10);
const loading = ref(true);
const search = ref('');

const headers = [
	{
		text: 'Name',
		key: 'name',
		align: 'start',
		sortable: true,
	},
	{
		text: 'Actions',
		key: 'actions',
		align: 'end',
		sortable: false,
	},
];

const tableEntities = computed(() => {
	return entities.value?.map((entity) => {
		return {
			...entity,
			actions: [
				{
					icon: 'mdi-pencil',
					handler: () => editEntity(entity),
				},
				{
					icon: 'mdi-delete',
					handler: () => deleteEntity(entity),
				},
			],
		};
	});
});

const listEntity = async (options: any) => {
	loading.value = true;
	const { page, itemsPerPage } = options;
	await entityStore.listEntity({
		start: BigInt((page - 1) * itemsPerPage), // page starts at 1
		end: BigInt(page * itemsPerPage),
		sort: 'created_at',
		order: false,
	});
	console.log('listEntity', entities.value);
	console.log('total', total.value);


	loading.value = false;
};

const dialog = ref(false);
const close = () => {
	dialog.value = false;
	entity.value = Entity.create({})
};

const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const create = async () => {
	await entityStore.createEntity();
	dialog.value = false;
	entity.value = {} as Entity;
};

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDeleteEntity = () => {
	// TODO: delete entity
	dialogDelete.value = false;
};

const editEntity = (entity: Entity) => {
};
const deleteEntity = (entity: Entity) => {
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
					<v-card-title>
						<span class="text-h6">New entity</span>
					</v-card-title>
					<v-card-text>
						<v-container>
							<v-row>
								<v-col cols="12" md="4">
									<v-text-field v-model="newEntity.name" :rules="nameRules"
										label="Name"></v-text-field>
								</v-col>
								<v-col cols="12" md="4">
									<v-text-field v-model="newEntity.avatarURL" label="Avatar URL"></v-text-field>
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
		<v-col class="mx-auto" cols="4">
			<v-data-table-server v-model:items-per-page="itemsPerPage" :headers="headers" :items="tableEntities"
				:items-length="Number(total)" :loading="loading" :search="search" item-value="name"
				@update:options="listEntity">
			</v-data-table-server>
		</v-col>
		<v-divider vertical></v-divider>
		<v-col class="mx-auto" cols="8">
			<v-card-item class="justify-center mb-2" prepend-icon="mdi-account-plus">
				<v-card-title>
					Entity
				</v-card-title>
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-card-item>
			<v-divider></v-divider>
		</v-col>
	</v-row>
</template>
<style scoped>
.bar-bg {
	background: url('@/assets/img/bar-background.svg') no-repeat bottom center;
	background-size: cover;
}
</style>
