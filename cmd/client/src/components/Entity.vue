<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { Entity } from '@internal/user/entity';
import { computed, ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { marked } from "marked";
import type { VDataTable } from 'vuetify/components/VDataTable';
import { useAuthStore } from '@/stores/auth';
import { useRoleStore } from '@/stores/role';
import { useUserStore } from '@/stores/user';
import type { Role } from '@internal/user/role';
import type { U } from '@internal/user/user';
import { ulid } from '@/utils/ulid';
import { ListRoleReq } from '@internal/user/dto/role';
import { ListEntityReq } from '@internal/user/dto/entity';
import { en } from 'vuetify/locale';
import { ListUserReq } from '@internal/user/dto/user';

// #MARK:Common
// ______________________________________________________
const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const authStore = useAuthStore();

// [Info, Roles, User]
const tableView = ref(0)

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

// #MARK:Entity
// ______________________________________________________
const entityStore = useEntityStore();
const {
	entities: entities,
	total: entityTotal,
} = toRefs(entityStore);


const name = ref('');
const avatarURL = ref('');
const description = ref('');

const loadingEntity = ref(true);
const searchEntity = ref('');

const headersEntity: ReadonlyHeaders = [
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

const selectedEntities = ref<Entity[]>([]);
const selectedEntity = computed(() => selectedEntities.value.at(0));

const viewEntityIDs = ref<string[]>([])

const viewEntities = computed(() => {
	return viewEntityIDs.value.map((entityID: string) => entities.value?.get(entityID));
});

const selectEntity = (_: any, row: any) => {
	selectedEntities.value = [];
	row.toggleSelect({ value: row.item, selectable: true });
};

const mdDescription = computed(() => {
	return marked.parse(selectedEntities.value.at(0)?.description || `*no description*`);
});

const listEntity = async (options: any) => {
	loadingEntity.value = true;

	const { page, itemsPerPage, sortBy } = options;
	const newIDs = await entityStore.listEntity(ListEntityReq.create({
		userIDs: true,
		search: searchEntity.value,
		pagination: {
			start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
			end: BigInt(page * itemsPerPage),
			sort: sortBy?.at(0)?.key ?? '',
			order: sortBy?.at(0)?.order === 'asc' ? true : false,
		}
	}));

	viewEntityIDs.value = newIDs;

	loadingEntity.value = false;
};

// #MARK:Entity dialogs
const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const dialogCreateEntity = ref(false);
const closeCreateEntity = () => {
	dialogCreateEntity.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';
};

const createEntity = async () => {
	await entityStore.createEntity(name.value, avatarURL.value, description.value);
	dialogCreateEntity.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';

	await authStore.refreshToken();
	await listEntity({ page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] })
};

const dialogDeleteEntity = ref(false);
const closeDeleteEntity = () => {
	dialogDeleteEntity.value = false;
};

const confirmDeleteEntity = () => {
	// TODO: delete entity
	dialogDeleteEntity.value = false;
};

const deleteEntity = () => {
};

const nameEdit = ref(false);
const descriptionEdit = ref(false);

const updateName = async function () {
	if (nameEdit.value) {
		await entityStore.updateEntity(selectedEntity.value?.iD, selectedEntity.value?.name, undefined, undefined);
	}

	nameEdit.value = !nameEdit.value
};

const updateDescription = async function () {
	if (descriptionEdit.value) {
		await entityStore.updateEntity(selectedEntity.value?.iD, undefined, selectedEntity.value?.description, undefined);
	}

	descriptionEdit.value = !descriptionEdit.value
};


// #MARK:Role
// ______________________________________________________
const roleStore = useRoleStore();
const {
	roles: roles,
	total: roleTotal,
} = toRefs(roleStore);

const loadingRole = ref(true);
const searchRole = ref('');

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

const selectedRoles = ref<Role[]>([]);

const selectedRole = computed(() => selectedRoles.value.at(0));

const viewRoleIDs = ref<string[]>([])

const viewRoles = computed(() => {
	return viewRoleIDs.value.map((roleID: string) => { return roles.value?.get(roleID) });
});

const selectRole = (_: any, row: any) => {
	selectedRoles.value = [];
	row.toggleSelect({ value: row.item, selectable: true });
};

const listRole = async (options: any) => {
	loadingRole.value = true;
	const { page, itemsPerPage, sortBy } = options;
	const [newRoleIDs] = await roleStore.listRole(ListRoleReq.create({
		entityIDs: selectedEntity.value ? [selectedEntity.value.iD!] : [],
		search: searchRole.value,
		pagination: {
			start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
			end: BigInt((page * itemsPerPage)),
			sort: sortBy?.at(0)?.key ?? '',
			order: sortBy?.at(0)?.order === 'asc' ? true : false,
		}
	}));

	viewRoleIDs.value = newRoleIDs;

	loadingRole.value = false;
};

// #MARK:Role dialogs
const dialogCreateRole = ref(false);
const closeCreateRole = () => {
	dialogCreateRole.value = false;
	name.value = '';
};

const createRole = async () => {
	await roleStore.createRole(name.value);
	dialogCreateRole.value = false;
	name.value = '';

	await authStore.refreshToken();
	await listRole({ page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] })
};

const dialogDeleteRole = ref(false);
const closeDeleteRole = () => {
	dialogDeleteRole.value = false;
};

const confirmDeleteRole = () => {
	// TODO: delete entity
	dialogDeleteRole.value = false;
};

const deleteRole = () => {
};

// #MARK:User
// ______________________________________________________
const userStore = useUserStore();
const {
	users: users,
	total: userTotal,
} = toRefs(userStore);

const loadingUser = ref(true);
const searchUser = ref('');

const headersByUser: ReadonlyHeaders = [
	{
		title: 'Email',
		key: 'email',
		align: 'start',
		sortable: true,
	},
	{
		title: 'Username',
		key: 'username',
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

type UserEntity = U & { entity: Entity | undefined };

const selectedUsers = ref<UserEntity[]>([]);

const selectedUser = computed(() => selectedUsers.value.at(0));

const viewUserIDs = ref<string[]>([])

const viewUsers = computed(() => {
	return viewUserIDs.value.map((userID: string) => {
		const user = users.value?.get(userID);
		return {
			...user,
			// entity: entities.value?.get(ulid(user?.entityID!)),
		}
	});
});

const selectUser = (_: any, row: any) => {
	selectedUsers.value = [];
	row.toggleSelect({ value: row.item, selectable: true });
};

const listUser = async (options: any) => {
	loadingUser.value = true;
	const { page, itemsPerPage, sortBy } = options;
	const newUserIDs = await userStore.listUser(ListUserReq.create({
		entityIDs: selectedEntity.value ? [selectedEntity.value.iD!] : [],
		search: searchUser.value,
		pagination: {
			start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
			end: BigInt((page * itemsPerPage)),
			sort: sortBy?.at(0)?.key ?? '',
			order: sortBy?.at(0)?.order === 'asc' ? true : false,
		}
	}));

	viewUserIDs.value = newUserIDs;

	loadingUser.value = false;
};

// #MARK:User dialogs
const dialogInviteUser = ref(false);
const closeInviteUser = () => {
	dialogInviteUser.value = false;
	name.value = '';
};

const inviteUser = async () => {
	await userStore.inviteUser(name.value);
	dialogInviteUser.value = false;
	name.value = '';

	await authStore.refreshToken();
	await listUser({ page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] })
};

const dialogDeleteUser = ref(false);
const closeDeleteUser = () => {
	dialogDeleteUser.value = false;
};

const confirmDeleteUser = () => {
	// TODO: delete entity
	dialogDeleteUser.value = false;
};

const deleteUser = () => {
};

</script>

<template>
	<v-dialog v-model="dialogDeleteEntity" max-width="500px">
		<v-card>
			<v-card-title class="text-h5">Are you sure you want to delete this entity ?</v-card-title>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn variant="text" @click="closeDeleteEntity">Cancel</v-btn>
				<v-btn color="danger" variant="text" @click="confirmDeleteEntity">OK</v-btn>
				<v-spacer></v-spacer>
			</v-card-actions>
		</v-card>
	</v-dialog>
	<v-row>
		<v-col cols="4">
			<v-sheet class="px-1 rounded-xl" outlined color="primary">
				<v-col class="d-flex justify-end align-center rounded-t-xl main-color-background" cols="12">
					<v-dialog v-model="dialogCreateEntity" max-width="800px">
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
										<v-btn variant="text" @click="closeCreateEntity">
											Cancel
										</v-btn>
										<v-btn color="primary" variant="text" @click="createEntity">
											Create
										</v-btn>
									</v-card-actions>
								</v-form>
							</v-card>
						</v-sheet>
					</v-dialog>
				</v-col>
				<v-text-field class="main-color-background" v-model="searchEntity" label="Search"
					prepend-inner-icon="mdi-magnify" variant="outlined" hide-details single-line>
				</v-text-field>
				<v-data-table-server class="main-color-background table-spacing rounded-0" :headers="headersEntity"
					fixed-footer min-height="50vh" max-height="100vh" items-per-page-text=""
					:items-per-page-options="pageOptions" :items="viewEntities" :items-length="Number(entityTotal)"
					:loading="loadingEntity" :search="searchEntity" item-value="name" item-key="iD"
					@update:options="listEntity" v-model="selectedEntities" @click:row="selectEntity" return-object
					item-selectable select-strategy="single">
					<template v-slot:item="{ item, isSelected, index, props }">
						<tr v-if="item" v-bind="props" class="cursor-pointer py-8" :key="item.name" :class="[(index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'),
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
				<v-col cols="12" class="p-8 main-color-background rounded-b-xl"></v-col>
			</v-sheet>
		</v-col>
		<v-divider vertical></v-divider>
		<v-col class="mx-auto" cols="6">
			<v-col class="d-flex justify-center" cols="12">
				<v-btn-toggle v-model="tableView">
					<v-btn size="large" variant="outlined" append-icon="mdi-domain">
						Info
						<template v-slot:append>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
					<v-btn size="large" variant="outlined" append-icon="mdi-account-cog">
						Roles
						<template v-slot:append>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
					<v-btn size="large" variant="outlined" append-icon="mdi-account-multiple">
						Users
						<template v-slot:append>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
				</v-btn-toggle>
			</v-col>
			<v-sheet class="px-1 rounded-xl" outlined color="primary" v-if="selectedEntity">
				<!-- Info -->
				<v-card v-if="tableView == 0" class="mt-8 px-6 py-6 justify-center rounded-xl main-color-background"
					variant="flat">
					<v-avatar class="mb-4 justify-center" size="96" :color="!selectedEntity.avatarURL ? 'primary' : ''">
						<img v-if="selectedEntity.avatarURL" :src="selectedEntity.avatarURL" alt="Avatar">
						<span v-else-if="!selectedEntity.avatarURL" class=" mx-auto text-center text-h5">
							{{ selectedEntity.name?.at(0)?.toUpperCase() }}
						</span>
					</v-avatar>
					<v-text-field class="justify-center text-h3" v-model="selectedEntity.name"
						:variant="!nameEdit ? 'plain' : 'underlined'" :readonly="!nameEdit">
						<template v-slot:append-inner>
							<v-icon color="primary" size="large" @click="updateName"
								:icon="!nameEdit ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
						</template>
					</v-text-field>
					<!-- eslint-disable vue/no-v-html vue/no-v-text-v-html-on-component -->
					<v-textarea v-if="descriptionEdit" class="p-6 m-6" variant="solo" no-resize
						:model-value="selectedEntity.description" :readonly="false">
						<template v-slot:append-inner>
							<v-icon color="primary" size="large" @click="updateDescription"
								:icon="!descriptionEdit ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
						</template>
					</v-textarea>
					<v-text-field v-if="!descriptionEdit" class="p-6 m-6" variant="solo" no-resize
						:v-html="mdDescription" :readonly="true">
						<template v-slot:append-inner>
							<v-icon color="primary" size="large" @click="updateDescription"
								:icon="!descriptionEdit ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
						</template>
					</v-text-field>
					<!--eslint-enable-->
					<v-card-actions>
						<v-spacer></v-spacer>
						<v-btn color="error" @click="deleteEntity()">Delete</v-btn>
					</v-card-actions>
				</v-card>
				<!-- Roles -->
				<v-col v-if="tableView == 1" class="d-flex justify-end align-center rounded-t-xl main-color-background"
					cols="12">
					<v-dialog v-model="dialogCreateRole" max-width="800px">
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
										<span class="text-h6">New role</span>
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
										<v-btn variant="text" @click="closeCreateRole">
											Cancel
										</v-btn>
										<v-btn color="primary" variant="text" @click="createRole">
											Create
										</v-btn>
									</v-card-actions>
								</v-form>
							</v-card>
						</v-sheet>
					</v-dialog>
				</v-col>
				<v-text-field v-if="tableView == 1" class="main-color-background" v-model="searchEntity" label="Search"
					prepend-inner-icon="mdi-magnify" variant="outlined" hide-details single-line>
				</v-text-field>
				<v-data-table-server v-if="tableView === 1" class="main-color-background rounded-0"
					:headers="headersByRole" fixed-footer min-height="50vh" max-height="100vh" items-per-page-text=""
					:items-per-page-options="pageOptions" :items="viewRoles" :items-length="Number(roleTotal)"
					:loading="loadingRole" :search="searchRole" item-value="name" item-key="iD"
					@update:options="listRole" v-model="selectedRoles" @click:row="selectRole" return-object
					item-selectable select-strategy="single">
					<template v-slot:item="{ item, isSelected, index, props }">
						<tr class="cursor-pointer" v-bind="props" :key="item.name" :class="[(index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'),
						(isSelected({ value: item, selectable: true }) ? 'active-bg' : '')]">
							<td class="d-flex align-center">
								<span class="text-h6">{{ item.name }}</span>
							</td>
							<td class="text-caption text-right">{{ new Date(Number(item.createdAt) *
								1000).toLocaleDateString('en-GB')
								}}
							</td>
						</tr>
					</template>
				</v-data-table-server>
				<v-col v-if="tableView == 1" cols="12" class="p-8 main-color-background rounded-b-xl"></v-col>
				<!-- User -->
				<v-col v-if="tableView == 2" class="d-flex justify-end align-center rounded-t-xl main-color-background"
					cols="12">
					<v-dialog v-model="dialogInviteUser" max-width="800px">
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
										<span class="text-h6">Invite user</span>
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
										<v-btn variant="text" @click="closeInviteUser">
											Cancel
										</v-btn>
										<v-btn color="primary" variant="text" @click="inviteUser">
											Invite
										</v-btn>
									</v-card-actions>
								</v-form>
							</v-card>
						</v-sheet>
					</v-dialog>
				</v-col>
				<v-text-field v-if="tableView == 2" class="main-color-background" v-model="searchUser" label="Search"
					prepend-inner-icon="mdi-magnify" variant="outlined" hide-details single-line>
				</v-text-field>
				<v-data-table-server v-if="tableView === 2" class="main-color-background rounded-0"
					:headers="headersByUser" fixed-footer min-height="50vh" max-height="100vh" items-per-page-text=""
					:items-per-page-options="pageOptions" :items="viewUsers" :items-length="Number(userTotal)"
					:loading="loadingUser" :search="searchUser" item-value="name" item-key="iD"
					@update:options="listUser" v-model="selectedUsers" @click:row="selectUser" return-object
					item-selectable select-strategy="single">
					<template v-slot:item="{ item, isSelected, index, props }">
						<tr class="cursor-pointer" v-bind="props" :key="item.email" :class="[(index % 2 === 0 ? 'row-bg-even' : 'row-bg-odd'),
						(isSelected({ value: item, selectable: true }) ? 'active-bg' : '')]">
							<td class="d-flex align-center">
								<span class="text-h6">{{ item.email }}</span>
							</td>
							<td class="d-flex align-center">
								<span class="text-h6">{{ item.lastName + ' ' + item.firstName }}</span>
							</td>
							<td class="text-caption text-right">{{ new Date(Number(item.createdAt) *
								1000).toLocaleDateString('en-GB')
								}}
							</td>
						</tr>
					</template>
				</v-data-table-server>
				<v-col v-if="tableView == 2" cols="12" class="p-8 main-color-background rounded-b-xl"></v-col>
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
