<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs, watch, type ComputedRef } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { useAuthStore } from '@/stores/auth';
import { ulid } from '@/utils/ulid';
import { useRoleStore } from '@/stores/role';
import { ListRoleReq, RolePermission } from '@internal/user/dto/role';
import RoleDetails from '@/components/role/Details.vue';
import PermissionTable from '@/components/permission/Table.vue';
import type { ReadonlyHeaders } from '@/utils/headers';
import { useUserStore } from '@/stores/user';
import type { Role } from '@internal/user/role';

const props = withDefaults(defineProps<{
	userID: Uint8Array | undefined;
	addRole: ((role: RolePermission) => Promise<void>) | undefined;
}>(), {
	userID: undefined,
	addRole: undefined,
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

// to add role
const userStore = useUserStore();
const {
	roles: rolesUser,
} = toRefs(userStore);

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
	selected: selected,
	total: total,
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
	item.toggleExpand({ value: item.item });
};

const list = async (options: any = { page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] }) => {
	if (!selectedEntity.value) {
		viewIDs.value = [];

		return
	}

	loading.value = true;
	const { page, itemsPerPage, sortBy } = options;
	const [newRoleIDs] = await store.list(ListRoleReq.create({
		entityIDs: [selectedEntity.value.iD],
		search: search.value,
		paginate: {
			start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
			end: BigInt((page * itemsPerPage)),
			sort: sortBy?.at(0)?.key ?? '',
			order: sortBy?.at(0)?.order === 'asc' ? true : false,
		}
	}));

	viewIDs.value = newRoleIDs;

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
};

const permissions = ref();

const create = async () => {
	const values = permissions.value?.values.map(permissions.value.unhash);

	await store.create(selectedEntity?.value?.iD!, name.value, values);
	dialogCreate.value = false;
	name.value = '';
	permissions.value.clear();

	await authStore.refreshToken();
	await list()
};

const nameEdit = ref(false);

const updateName = async function () {
	if (nameEdit.value) {
		// await store.update(UpdateEntityReq.create({
		// 	iD: e.value?.iD,
		// 	name: { value: e.value?.name },
		// }));
	}

	nameEdit.value = !nameEdit.value
};

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDelete = () => {
	// TODO: delete entity
	dialogDelete.value = false;
};

const delete_ = () => {
};

// Optional add role to user id
const userRoles = computed(() => {
	return rolesUser.value.get(ulid(props.userID))?.roles?.reduce((acc: Map<string, boolean>, role: Role) => {
		acc.set(ulid(role?.iD), true);

		return acc;
	}, new Map<string, boolean>())
});

const loadingAddRole = ref(false);

const addRole = async (item: RolePermission) => {
	if (props.addRole) {
		loadingAddRole.value = true;
		await props.addRole(item);
		loadingAddRole.value = false;
	}
};

</script>

<template>
	<v-col class="px-6 pt-6 rounded-t table-color-background" cols="12">
		<v-row>
			<v-col cols="10">
				<v-text-field class="table-color-background px-1" v-model="search" label="Search"
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
								<v-btn color="primary" variant="text" @click="create">
									Create
								</v-btn>
							</v-card-actions>
						</v-form>
					</v-card>
				</v-dialog>
			</v-col>
		</v-row>
	</v-col>
	<v-data-table-server class="px-6 rounded-0" :headers="headers" fixed-footer min-height="50vh" max-height="100vh"
		items-per-page-text="" :items-per-page-options="pageOptions" :items="views" :items-length="Number(total)"
		:loading="loading" :search="search" item-value="role.iD" @update:options="list" @click:row="expand"
		return-object>
		<template v-slot:item="{ item, internalItem, columns, isExpanded, index, props: itemProps }">
			<v-hover v-slot="{ isHovering, props: hoverProps }">
				<tr v-if="item" v-bind="{ ...itemProps, ...hoverProps }" :key="ulid(item.role?.iD)">
					<td :colspan="columns.length" class="cursor-pointer px-1 py-1">
						<v-card class="justify-center" variant="flat" :class="{
							'row-hovered': isHovering,
							'row-expanded': isExpanded(internalItem),
							'row-even': index % 2 === 0,
							'row-odd': index % 2 !== 0,
						}" :title="item.role?.name" :subtitle="Number(item.permissions.length) + ' permission(s)'">
							<template v-slot:append>
								<v-icon v-if="isExpanded(internalItem)" icon="mdi-minus" size="x-large" color="primary">
								</v-icon>
								<v-icon v-else icon="mdi-plus" size="x-large" color="primary"> </v-icon>
							</template>
							<v-card-actions>
								<v-btn v-if="props.userID && !userRoles?.has(ulid(item.role?.iD))" variant="tonal"
									prepend-icon="mdi-plus-box" color="primary" v-bind="props" :loading="loadingAddRole"
									v-on:click.stop.prevent="addRole(item)">
									Add role
									<template v-slot:prepend>
										<v-icon color="primary"></v-icon>
									</template>
								</v-btn>
								<v-btn v-else-if="props.userID && userRoles?.has(ulid(item.role?.iD))" variant="tonal"
									disabled prepend-icon="mdi-check-bold" color="secondary" v-bind="props"
									v-on:click.stop.prevent="async () => { }">
									Added
									<template v-slot:prepend>
										<v-icon color="secondary"></v-icon>
									</template>
								</v-btn>
								<v-divider></v-divider>
								<p class="font-italic font-weight-light">
									{{ new Date(Number(item.role?.createdAt) * 1000).toLocaleDateString('en-GB') }}
								</p>
							</v-card-actions>
						</v-card>
					</td>
				</tr>
			</v-hover>
		</template>
		<template v-slot:expanded-row="{ columns, item }">
			<RoleDetails :colspan="columns.length" :item="item">
			</RoleDetails>
		</template>
	</v-data-table-server>
	<v-col cols="12" class="p-8 table-color-background rounded-b"></v-col>
</template>
<style scoped>
.table-color-background {
	background-color: #212121;
}

.row-odd {
	transition: background-color .2s ease-in-out;
	background-color: rgba(66, 66, 66, 0.5);
	background-color: #424242;
}

.row-odd:not(.row-hovered) {
	background-color: #212121;
}

.row-even {
	transition: background-color .2s ease-in-out;
	background-color: rgba(66, 66, 66, 0.5);
}

.row-even:not(.row-hovered) {
	background-color: #263238;
}

.row-expanded {
	background-color: rgba(0, 145, 234, 0.5) !important;
}

.cursor-pointer {
	cursor: pointer;
}
</style>
