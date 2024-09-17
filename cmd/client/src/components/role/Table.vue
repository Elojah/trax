<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs, watch } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { useAuthStore } from '@/stores/auth';
import { ulid } from '@/utils/ulid';
import { useRoleStore } from '@/stores/role';
import { ListRoleReq, RolePermission } from '@internal/user/dto/role';
import PermissionTable from '@/components/permission/Table.vue';
import type { ReadonlyHeaders } from '@/utils/headers';
import { useErrorsStore } from '@/stores/errors';
import RoleDetails from '@/components/role/Details.vue';
import UserTable from '@/components/user/Table.vue';

const props = withDefaults(defineProps<{
	showActionUserID: Uint8Array | undefined;
	filterByUserID: Uint8Array | undefined;
}>(), {
	showActionUserID: undefined,
	filterByUserID: undefined,
});

// #MARK:Common
// ______________________________________________________
const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
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

const authStore = useAuthStore();

const errorsStore = useErrorsStore()
const { success, message } = toRefs(errorsStore)

const entityStore = useEntityStore();
const {
	selected: selectedEntities,
} = toRefs(entityStore);

const selectedEntity = computed(() => selectedEntities.value.at(0));

// #MARK:Role
// ______________________________________________________
const store = useRoleStore();
const {
	roles: roles,
	total: total,
	rolesByUser: rolesByUser,
} = toRefs(store);

const loading = ref(false);
const search = ref('');

const headers: ReadonlyHeaders = [
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

const viewIDs = ref<string[]>([])

const views = computed(() => {
	return viewIDs.value.map((roleID: string) => { return roles.value?.get(roleID) });
});

const expand = (_: any, item: any) => {
	if (props.showActionUserID) {
		return
	}
	item.toggleExpand({ value: item.item });
};

const list = async (options: any = { page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] }) => {
	if (!selectedEntity.value) {
		viewIDs.value = [];

		return
	}

	loading.value = true;
	const { page, itemsPerPage, sortBy } = options;
	try {
		const newRoleIDs = await store.list(ListRoleReq.create({
			entityIDs: [selectedEntity.value.iD],
			userID: props.filterByUserID,
			search: search.value,
			paginate: {
				start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
				end: BigInt((page * itemsPerPage)),
				sort: sortBy?.at(0)?.key ?? '',
				order: sortBy?.at(0)?.order === 'asc' ? true : false,
			}
		}));

		viewIDs.value = newRoleIDs;
	} catch (e) {
		errorsStore.showGRPC(e)
	}

	loading.value = false;
};

// refresh list when selected entity changes
watch(selectedEntity, async () => {
	await list();
});

const name = ref('');

// #MARK:Role dialogs
const dialogCreate = ref(false);
const closeCreateRole = () => {
	dialogCreate.value = false;
	name.value = '';
	permissions.value.clear();
};

const permissions = ref();

const create = async () => {
	const values = permissions.value?.values.map(permissions.value.unhash);

	let ok = true;
	try {
		await store.create(selectedEntity?.value?.iD!, name.value, values);
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false;
	}
	if (ok) {
		message.value = `Role ${name.value} created successfully`;
		success.value = true;
	}

	dialogCreate.value = false;
	name.value = '';
	permissions.value.clear();

	await authStore.refreshToken();
	await list();
};


// Optional show action user ID
const userRoles = computed(() => {
	return rolesByUser.value.get(ulid(props.showActionUserID))
});

const loadingAddRole = ref(false);

const addRoleToUser = async (item: RolePermission) => {
	if (props.showActionUserID) {
		loadingAddRole.value = true;
		let ok = true;
		try {
			await store.addUser(item.role?.iD!, props.showActionUserID);
		} catch (e) {
			errorsStore.showGRPC(e);
			ok = false;
		}
		if (ok) {
			message.value = `Role added successfully to user`;
			success.value = true;
		}

		loadingAddRole.value = false;
	}
};

const loadingRemoveRole = ref(false);

const removeRoleToUser = async (item: RolePermission) => {
	if (props.showActionUserID) {
		loadingRemoveRole.value = true;
		let ok = true;
		try {
			await store.deleteUser(item.role?.iD!, props.showActionUserID);
		} catch (e) {
			errorsStore.showGRPC(e);
			ok = false;
		}
		if (ok) {
			message.value = `Role removed successfully to user`;
			success.value = true;
		}

		loadingRemoveRole.value = false;
	}
};

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
							{{ props.showActionUserID ? 'Add' : 'New' }}
							<template v-slot:prepend>
								<v-icon color="primary"></v-icon>
							</template>
						</v-btn>
					</template>
					<v-card class="px-6 py-6 rounded" variant="elevated">
						<v-form ref="form" v-model="valid" lazy-validation>
							<v-card-title class="d-flex justify-center">
								<span class="text-h5">Create role</span>
							</v-card-title>
							<v-divider></v-divider>
							<v-card-text>
								<v-container>
									<v-row>
										<v-col cols="12">
											<v-text-field v-model="name" :rules="nameRules" label="Name"></v-text-field>
										</v-col>
									</v-row>
									<v-row>
										<PermissionTable :permissions="[]" ref="permissions" :disabled="false">
										</PermissionTable>
									</v-row>
								</v-container>
							</v-card-text>
							<v-card-actions>
								<v-spacer></v-spacer>
								<v-btn color="error" variant="text" @click="closeCreateRole">
									Cancel
								</v-btn>
								<v-btn color="primary" variant="text" :disabled="!valid" @click="create">
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
		max-height="100vh" items-per-page-text="" :items-per-page-options="pageOptions" :items="views"
		:items-length="Number(total)" :loading="loading" :search="search" item-value="role.iD" @update:options="list"
		@click:row="expand" return-object>
		<template v-slot:no-data>
			<div></div>
		</template>
		<template v-slot:item="{ item, internalItem, columns, isExpanded, index, props: itemProps }">
			<v-hover v-slot="{ isHovering, props: hoverProps }">
				<tr v-if="item" v-bind="{ ...itemProps, ...hoverProps }" :key="ulid(item.role?.iD)">
					<td :colspan="columns.length" class="px-1 py-1">
						<v-card :class="{
							'cursor-pointer': !props.showActionUserID,
							'row-hovered': isHovering && !props.showActionUserID,
							'row-even': index % 2 === 0,
							'row-odd': index % 2 !== 0,
							'row-expanded': isExpanded(internalItem),
							'text-primary': isExpanded(internalItem),
						}" :title="item.role?.name" :subtitle="Number(item.permissions.length) + ' permission(s)'">
							<template v-slot:prepend>
								<v-icon v-if="!props.showActionUserID && isExpanded(internalItem)" class="mr-4"
									icon="mdi-minus" size="x-large" color="primary">
								</v-icon>
								<v-icon v-else-if="!props.showActionUserID" class="mr-4" icon="mdi-plus" size="x-large"
									color="primary"> </v-icon>
								<v-divider vertical></v-divider>
							</template>
							<template v-slot:append>
								<v-divider vertical v-if="props.showActionUserID"></v-divider>
								<v-btn v-if="props.showActionUserID && !userRoles?.has(ulid(item.role?.iD))"
									variant="tonal" class="mr-4" prepend-icon="mdi-plus-box" color="primary"
									v-bind="props" :loading="loadingAddRole"
									v-on:click.stop.prevent="addRoleToUser(item)">
									Add
									<template v-slot:prepend>
										<v-icon color="primary"></v-icon>
									</template>
								</v-btn>
								<v-btn v-else-if="props.showActionUserID && userRoles?.has(ulid(item.role?.iD))"
									variant="tonal" class="mr-4" prepend-icon="mdi-trash-can" color="error"
									v-bind="props" :loading="loadingRemoveRole"
									v-on:click.stop.prevent="removeRoleToUser(item)">
									Remove
									<template v-slot:prepend>
										<v-icon color="error"></v-icon>
									</template>
								</v-btn>
								<v-divider vertical></v-divider>
								<p class="ml-4 font-italic font-weight-light">
									{{ new Date(Number(item.role?.createdAt) * 1000).toLocaleDateString('en-GB') }}
								</p>
							</template>
						</v-card>
					</td>
				</tr>
			</v-hover>
		</template>
		<template v-slot:expanded-row="{ columns, item }">
			<tr v-if="!props.showActionUserID">
				<td :colspan="columns.length">
					<v-sheet>
						<v-row class="my-4">
							<v-col cols="4">
								<v-sheet class="fill-height fill-width">
									<RoleDetails :item="item">
									</RoleDetails>
								</v-sheet>
							</v-col>
							<v-divider vertical></v-divider>
							<v-col cols="8">
								<v-sheet class="fill-height fill-width">
									<UserTable :filter-by-role-i-d="item?.role?.iD"
										:show-action-role-i-d="item?.role?.iD">
									</UserTable>
								</v-sheet>
							</v-col>
						</v-row>
					</v-sheet>
				</td>
			</tr>
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
	background-color: #37474F;
}

.row-even {
	transition: background-color .2s ease-in-out;
	background-color: #616161;
}

.row-even:not(.row-hovered) {
	background-color: #263238;
}

.row-expanded {
	background-color: rgb(33, 47, 59) !important;
}

.cursor-pointer {
	cursor: pointer;
}
</style>
