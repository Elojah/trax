<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs, watch } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { useAuthStore } from '@/stores/auth';
import { ulid } from '@/utils/ulid';
import { useUserStore } from '@/stores/user';
import { ListUserReq } from '@internal/user/dto/user';
import RoleTable from '@/components/user/RoleTable.vue';
import type { ReadonlyHeaders } from '@/utils/headers';

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

const selectedEntity = computed(() => selectedEntities.value.at(0));

// #MARK:User
// ______________________________________________________
const store = useUserStore();
const {
	users: users,
	total: total,
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

	loading.value = false;
};

// refresh list when selected entity changes
watch(selectedEntity, async () => {
	await list();
});

const name = ref('');

// #MARK:User dialogs
const dialogInvite = ref(false);
const closeInvite = () => {
	dialogInvite.value = false;
	name.value = '';
};

const invite = async () => {
	await store.invite(name.value);
	dialogInvite.value = false;
	name.value = '';

	await authStore.refreshToken();
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

</script>

<template>
	<v-col class="d-flex justify-end align-center rounded-t-lg table-color-background" cols="12">
		<v-dialog v-model="dialogInvite" max-width="800px">
			<template v-slot:activator="{ props }">
				<v-btn variant="tonal" prepend-icon="mdi-plus-box" color="primary" v-bind="props">
					New
					<template v-slot:prepend>
						<v-icon color="primary"></v-icon>
					</template>
				</v-btn>
			</template>
			<v-sheet class="px-1 rounded-lg" outlined color="primary">
				<v-card class="px-6 py-6 rounded-lg" variant="elevated">
					<v-form ref="form" v-model="valid" lazy-validation>
						<v-card-title>
							<span class="text-h6">Invite user</span>
						</v-card-title>
						<v-card-text>
							<v-container>
								<v-row>
									<v-col cols="6">
										<v-text-field v-model="name" :rules="nameRules" label="Name"></v-text-field>
									</v-col>
									<v-col cols="6">
										<v-text-field label="Avatar URL"></v-text-field>
									</v-col>
								</v-row>
								<v-row>
									<v-col cols="12">
										<v-textarea label="Description"></v-textarea>
									</v-col>
								</v-row>
							</v-container>
						</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn color="error" variant="text" @click="closeInvite">
								Cancel
							</v-btn>
							<v-btn color="primary" variant="text" @click="invite">
								Invite
							</v-btn>
						</v-card-actions>
					</v-form>
				</v-card>
			</v-sheet>
		</v-dialog>
	</v-col>
	<v-text-field class="table-color-background px-1" v-model="search" label="Search" prepend-inner-icon="mdi-magnify"
		variant="outlined" hide-details single-line>
	</v-text-field>
	<v-data-table-server class="px-6 rounded-0" :headers="headers" fixed-footer min-height="50vh" max-height="100vh"
		items-per-page-text="" :items-per-page-options="pageOptions" :items="views" :items-length="Number(total)"
		:loading="loading" :search="search" item-value="user.iD" @update:options="list" @click:row="expand"
		return-object>
		<template v-slot:item="{ item, internalItem, columns, isExpanded, index, props: itemProps }">
			<v-hover v-slot="{ isHovering, props: hoverProps }">
				<tr v-if="item" v-bind="{ ...itemProps, ...hoverProps }" :key="ulid(item.iD)">
					<td :colspan="columns.length" class="cursor-pointer px-1 py-1">
						<v-card class="justify-center" variant="flat" :class="{
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
			<RoleTable v-if="item" :colspan="columns.length" :user-i-d="item.iD">
			</RoleTable>
		</template>
	</v-data-table-server>
	<v-col cols="12" class="p-8 table-color-background rounded-b-lg"></v-col>
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
