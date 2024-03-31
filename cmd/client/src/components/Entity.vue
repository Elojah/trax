<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { Entity } from '@internal/user/entity';
import { ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);

const authStore = useAuthStore();
const {
	profile: profile,
} = toRefs(authStore);

const valid = ref(null as boolean | null)

const dialog = ref(false);
const close = () => {
	dialog.value = false;
	entity.value = Entity.create({})
};

const blankEntity = Entity.create({});

const save = () => {
	// TODO: save new entity
	dialog.value = false;
	entity.value = blankEntity;
};

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDeleteEntity = () => {
	// TODO: delete entity
	dialogDelete.value = false;
};

const headers = [
	{ title: 'avatar_url', key: 'avatar_url', sortable: false },
	{ title: 'name', key: 'name', sortable: true },
	{ title: 'roles', key: 'roles', sortable: true },
	{ title: 'actions', key: 'actions', sortable: false },
]
const entities: Entity[] = [
	Entity.create({ name: 'test1' }),
	Entity.create({ name: 'test2' }),
	Entity.create({ name: 'test3' }),
]
const entity = ref(Entity.create({ name: 'test_init' }) as Entity);

const editEntity = (entity: Entity) => {
};
const deleteEntity = (entity: Entity) => {
};

</script>

<template>
	<v-data-table :headers="headers" tems="entities" :sort-by="[{ key: 'name', order: 'asc' }]" hide-default-header>
		<template v-slot:top>
			<v-toolbar flat>
				<v-toolbar-title>Entities</v-toolbar-title>
				<v-divider class="mx-4" inset vertical></v-divider>
				<v-spacer></v-spacer>
				<v-dialog v-model="dialog" max-width="500px">
					<template v-slot:activator="{ props }">
						<v-btn class="mb-2" color="primary" v-bind="props">
							New entity
						</v-btn>
					</template>
					<v-card>
						<v-card-title>
							<span class="text-h5">Title</span>
						</v-card-title>

						<v-card-text>
							<v-container>
								<v-row>
									<v-col cols="12" md="4" sm="6">
										<v-text-field v-model="entity.name" label="Name"></v-text-field>
									</v-col>
								</v-row>
							</v-container>
						</v-card-text>

						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn variant="text" @click="close">
								Cancel
							</v-btn>
							<v-btn color="primary" variant="text" @click="save">
								Save
							</v-btn>
						</v-card-actions>
					</v-card>
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
		</template>
		<template v-slot:[`item.actions`]="{ item }">
			<v-icon class="me-2" size="small" @click="editEntity(item)">
				mdi-pencil
			</v-icon>
			<v-icon size="small" @click="deleteEntity(item)">
				mdi-delete
			</v-icon>
		</template>
		<template v-slot:no-data>
			<v-btn color="primary">
				Reset
			</v-btn>
		</template>
	</v-data-table>
</template>
<style scoped></style>
