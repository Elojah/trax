<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs, watch } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { ulid } from '@/utils/ulid';
import { useRoleStore } from '@/stores/role';
import { ListRoleReq, RolePermission } from '@internal/user/dto/role';
import PermissionTable from '@/components/permission/Table.vue';
import { useErrorsStore } from '@/stores/errors';
import RoleDetails from '@/components/role/Details.vue';
import UserTable from '@/components/user/Table.vue';

const props = withDefaults(defineProps<{
	showActionUserID: Uint8Array | undefined;
	filterByUserID: Uint8Array | undefined;
	disableNewButton: boolean;
}>(), {
	showActionUserID: undefined,
	filterByUserID: undefined,
	disableNewButton: false,
});

// #MARK:Common
// ______________________________________________________
const valid = ref(null as boolean | null)

const createActionText = props.showActionUserID ? 'Add' : 'New';

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
</template>
<style scoped>
</style>
