<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, onMounted, ref, toRefs, watch } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { ulid, zero } from '@/utils/ulid';
import { useUserStore } from '@/stores/user';
import { ListUserReq } from '@internal/user/dto/user';
import UserRoleTable from '@/components/user/RoleTable.vue';
import type { ReadonlyHeaders } from '@/utils/headers';
import RoleTable from '@/components/role/Table.vue';
import type { RolePermission } from '@internal/user/dto/role';

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

const entityStore = useEntityStore();
const {
	selected: selectedEntities,
} = toRefs(entityStore);

const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);

const selectedEntity = computed(() => selectedEntities.value.at(0));

// #MARK:User
// ______________________________________________________
const store = useUserStore();
const {
	users: users,
	roles: roles,
	total: total,
	resetRoleDry: resetRoleDry,
} = toRefs(store);

const loading = ref(false);
const search = ref('');

const headers: ReadonlyHeaders = [
	{
		title: 'Email',
		key: 'email',
		align: 'start',
		sortable: true,
	},
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
	return viewIDs.value.map((userID: string) => { return users.value?.get(userID); });
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
	try {
		const newUserIDs = await store.list(ListUserReq.create({
			entityIDs: [selectedEntity.value.iD],
			search: search.value,
			paginate: {
				start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
				end: BigInt((page * itemsPerPage)),
				sort: sortBy?.at(0)?.key ?? '',
				order: sortBy?.at(0)?.order === 'asc' ? true : false,
			}
		}));

		viewIDs.value = newUserIDs;
	} catch (e) {
		errorsStore.showGRPC(e)
	}

	loading.value = false;
};

// refresh list when selected entity changes
watch(selectedEntity, async () => {
	await list();
	await resetRoleDry.value(zero);
});

// reset roles for new user
onMounted(async () => {
	await resetRoleDry.value(zero);
});

const name = ref('');

// #MARK:User dialogs
const dialogInvite = ref(false);
const closeInvite = () => {
	dialogInvite.value = false;
	name.value = '';
};

const invite = async () => {
	let ok = true;
	try {
		await store.invite(email.value, addedRoles);
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false;
	}

	if (ok) {
		message.value = `User ${email.value} invited successfully`;
		success.value = true;
	}

	dialogInvite.value = false;
	email.value = '';
	roles.value.set(ulid(zero), { roles: [] });

	await list()
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

const addRoleNewUser = async (role: RolePermission) => {
	await store.addRoleDry(zero, role.role!);
};

const email = ref(null as string | null)
const emailRules = [
	(v: string) => !!v || "Required",
	(v: string) => /.+@.+\..+/.test(v) || "Email must be valid"
]

const addedRoles = computed(() => roles.value.get(ulid(zero))?.roles);

const deleteRole = async (roleID: Uint8Array) => {
	await store.deleteRoleDry(zero, roleID);
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
				<v-dialog v-model="dialogInvite" max-width="1200px">
					<template v-slot:activator="{ props }">
						<v-btn variant="tonal" prepend-icon="mdi-plus-box" color="primary" size="large" v-bind="props">
							New
							<template v-slot:prepend>
								<v-icon color="primary"></v-icon>
							</template>
						</v-btn>
					</template>
					<v-sheet class="d-flex flex-column pa-4 h-screen fill-height fill-width">
						<v-container class="px-0">
							<v-form ref="form" v-model="valid" lazy-validation>
								<v-row>
									<v-col cols="10">
										<v-text-field v-model="email" label="Email" :rules="emailRules"
											prepend-inner-icon="mdi-email" underlined required clearable>
										</v-text-field>
									</v-col>
									<v-col cols="2" class="d-flex align-center justify-end">
										<v-btn size="large" :disabled="!valid || addedRoles?.length === 0"
											variant="tonal" prepend-icon="mdi-account-circle" @click="invite">Invite
											<template v-slot:prepend>
												<v-icon color="primary"></v-icon>
											</template>
										</v-btn>
									</v-col>
								</v-row>
								<v-row v-if="addedRoles">
									<v-chip v-for="item in addedRoles" :key="ulid(item.iD)" class="ma-2" variant="tonal"
										closable label color="primary" @click:close="deleteRole(item.iD)">
										<v-icon icon="mdi-account-cog" start></v-icon>
										{{ item.name }}
									</v-chip>
								</v-row>
							</v-form>
						</v-container>
						<v-divider></v-divider>
						<RoleTable :user-i-d="zero" :add-role="addRoleNewUser"></RoleTable>
						<v-divider></v-divider>
						<v-btn color="error" variant="tonal" @click="closeInvite">
							Close
						</v-btn>
					</v-sheet>
				</v-dialog>
			</v-col>
		</v-row>
	</v-container>
	<v-data-table-server class="px-2 overflow-y-auto flex-grow-1" :headers="headers" fixed-footer fixed-header
		items-per-page-text="" :items-per-page-options="pageOptions" :items="views" :items-length="Number(total)"
		:loading="loading" :search="search" item-value="user.iD" @update:options="list" @click:row="expand"
		return-object>
		<template v-slot:item="{ item, internalItem, columns, isExpanded, index, props: itemProps }">
			<v-hover v-slot="{ isHovering, props: hoverProps }">
				<tr v-if="item" v-bind="{ ...itemProps, ...hoverProps }" :key="ulid(item.iD)">
					<td :colspan="columns.length" class="cursor-pointer px-1 py-1">
						<v-card class="justify-center" :class="{
							'row-hovered': isHovering,
							'row-expanded': isExpanded(internalItem),
							'row-even': index % 2 === 0,
							'row-odd': index % 2 !== 0,
						}" :title="item.email" :subtitle="item.lastName + ' ' + item.firstName">
							<template v-slot:append>
								<v-icon v-if="isExpanded(internalItem)" icon="mdi-minus" size="x-large" color="primary">
								</v-icon>
								<v-icon v-else icon="mdi-plus" size="x-large" color="primary"> </v-icon>
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
		<template v-slot:expanded-row="{ columns, item }">
			<UserRoleTable v-if="item" :colspan="columns.length" :user-i-d="item.iD">
			</UserRoleTable>
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
