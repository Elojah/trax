<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { Entity } from '@internal/user/entity';
import { computed, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { marked } from "marked";
import type { VDataTable } from 'vuetify/components/VDataTable';
import { useAuthStore } from '@/stores/auth';
import { useRoleStore } from '@/stores/role';
import type { Role } from '@internal/user/role';
import { ulid } from '@/utils/ulid';

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const authStore = useAuthStore();

const entityStore = useEntityStore();
const {
	entities: entities,
	total: entityTotal,
} = toRefs(entityStore);

const roleStore = useRoleStore();
const {
	roles: roles,
	total: roleTotal,
} = toRefs(roleStore);

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
		sortable: false,
	},
	{
		title: 'Created',
		key: 'created_at',
		align: 'end',
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

const selectedEntity = ref<Entity[]>([]);

type RoleEntity = Role & { entity: Entity | undefined };

const selectedRole = ref<RoleEntity[]>([]);

const selected = computed(() => {
	return selectedEntity.value?.at(0) || selectedRole.value?.at(0)?.entity;
});

const viewEntityIDs = ref<string[]>([])

const viewEntities = computed(() => {
	return viewEntityIDs.value.map((entityID: string) => entities.value?.get(entityID));
});

const viewRoleIDs = ref<string[]>([])

const viewRoles = computed(() => {
	return viewRoleIDs.value.map((roleID: string) => {
		const role = roles.value?.get(roleID);
		return {
			...role,
			entity: entities.value?.get(ulid(role?.entityID!)),
		}
	});
});

const select = (_: any, row: any) => {
	selectedEntity.value = [];
	selectedRole.value = [];
	row.toggleSelect({ value: row.item, selectable: true });
};

const mdDescription = computed(() => {
	return marked.parse(selected?.value?.description || `*no description*`);
});

const listEntity = async (options: any) => {
	loading.value = true;

	const { page, itemsPerPage, sortBy } = options;
	const newIDs = await entityStore.listEntity(null, search.value, {
		start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
		end: BigInt(page * itemsPerPage),
		sort: sortBy?.at(0)?.key ?? '',
		order: sortBy?.at(0)?.order === 'asc' ? true : false,
	});

	viewEntityIDs.value = newIDs;

	loading.value = false;
};


const listRole = async (options: any) => {
	loading.value = true;
	const { page, itemsPerPage, sortBy } = options;
	const [newRoleIDs, newEntityIDs] = await roleStore.listRole(null, search.value, {
		start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
		end: BigInt((page * itemsPerPage)),
		sort: sortBy?.at(0)?.key ?? '',
		order: sortBy?.at(0)?.order === 'asc' ? true : false,
	});

	await entityStore.populateEntity(newEntityIDs);

	viewRoleIDs.value = newRoleIDs;

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
		<v-col class="ml-8 mt-8" cols="4">
			<v-row class="main-color-background">
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
			<v-row class="px-3 main-color-background">
				<v-text-field v-model="search" label="Search" prepend-inner-icon="mdi-magnify" variant="outlined"
					hide-details single-line></v-text-field>
			</v-row>
			<v-row class="main-color-background">
				<v-data-table-server v-if="tableView === 0" class="transparent-background" :headers="headers"
					fixed-footer min-height="50vh" max-height="100vh" items-per-page-text=""
					:items-per-page-options="pageOptions" :items="viewEntities" :items-length="Number(entityTotal)"
					:loading="loading" :search="search" item-value="name" item-key="iD" @update:options="listEntity"
					v-model="selectedEntity" @click:row="select" return-object item-selectable select-strategy="single">
					<template v-slot:item="{ item, isSelected, index, props }">
						<tr v-if="item" class="mt-4 mb-4 cursor-pointer" v-bind="props" :key="item.name" :class="[(index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'),
						(isSelected({ value: item, selectable: true }) ? 'active-bg' : '')]">
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
					:items-per-page-options="pageOptions" :items="viewRoles" :items-length="Number(roleTotal)"
					:loading="loading" :search="search" item-value="name" item-key="iD" @update:options="listRole"
					v-model="selectedRole" @click:row="select" return-object item-selectable select-strategy="single">
					<template v-slot:item="{ item, isSelected, index, props }">
						<tr class="mt-4 mb-4 cursor-pointer" v-bind="props" :key="item.name" :class="[(index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'),
						(isSelected({ value: item, selectable: true }) ? 'active-bg' : '')]">
							<td class="d-flex align-center">
								<span class="text-h6">{{ item.name }}</span>
							</td>
							<td>
								<v-avatar class="mr-4" size="32" :color="!item.entity?.avatarURL ? 'primary' : ''">
									<img v-if="item.entity?.avatarURL" :src="item.entity?.avatarURL" alt="Avatar">
									<span v-else class=" mx-auto text-center text-h5">
										{{ item.entity?.name?.at(0)?.toUpperCase() }}
									</span>
								</v-avatar>
								<span class="text-h6">{{ item.entity?.name }}</span>
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
				<v-card class="mt-8 px-6 py-6 justify-center rounded-xl main-color-background" variant="flat"
					v-if="selected" :title="selected?.name">
					<template v-slot:prepend>
						<v-avatar size="32" :color="!selected?.avatarURL ? 'primary' : ''">
							<img v-if="selected?.avatarURL" :src="selected?.avatarURL" alt="Avatar">
							<span v-else-if="!selected?.avatarURL" class=" mx-auto text-center text-h5">
								{{ selected?.name?.at(0)?.toUpperCase() }}
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
.active-bg {
	background-color: #950c75 !important;
}

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
