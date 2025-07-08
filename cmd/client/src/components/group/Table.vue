<script setup lang="ts">
import { useGroupStore } from '@/stores/group';
import { computed, ref, toRefs } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { ulid } from '@/utils/ulid';
import { ListGroupReq } from '@internal/user/dto/group';
import { grpccodes } from '@/utils/errors';

const valid = ref(null as boolean | null)

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
const store = useGroupStore();
const {
	groups: groups,
	total: total,
	selected: selected,
} = toRefs(store);

const errorsStore = useErrorsStore()
const { success, message } = toRefs(errorsStore)

const name = ref('');
const avatarURL = ref('');
const description = ref('');

const loading = ref(false);
const search = ref('');

const viewIDs = ref<string[]>([])

const views = computed(() => {
	return viewIDs.value.map((groupID: string) => groups.value?.get(groupID));
});

const select = (_: any, row: any) => {
	selected.value = [];
	row.toggleSelect({ value: row.item, selectable: true });
};

const list = async (options: any = { page: 1, itemsPerPage: 10, sortBy: [{ key: 'created_at', order: 'desc' }] }) => {
	loading.value = true;

	try {
		const { page, itemsPerPage, sortBy } = options;
		const newIDs = await store.list(ListGroupReq.create({
			own: true,
			search: search.value,
			paginate: {
				start: BigInt(((page - 1) * itemsPerPage) + 1), // page starts at 1, start starts at 1
				end: BigInt(page * itemsPerPage),
				sort: sortBy?.at(0)?.key ?? '',
				order: sortBy?.at(0)?.order === 'asc' ? true : false,
			}
		}));

		viewIDs.value = newIDs;
	} catch (e) {
		errorsStore.showGRPC(e)
	}

	loading.value = false;
};

defineExpose({
	list,
});

// #MARK:Group dialogs
const nameRules = [
	(v: string) => !!v || 'Required',
	(v: string) => (v && v.length >= 1) || 'Min 1 character',
];

const dialogCreate = ref(false);
const closeCreate = () => {
	dialogCreate.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';
};

const create = async () => {
	let ok = true;
	try {
		await store.create(name.value, avatarURL.value, description.value);
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false
	}

	if (ok) {
		message.value = `Group ${name.value} created successfully`;
		success.value = true;
	}

	dialogCreate.value = false;
	name.value = '';
	avatarURL.value = '';
	description.value = '';

	await authStore.refreshToken();
	await list();
};

const short = (description: string): string => {
	return description.length > 64 ? description.substring(0, 64) + '...' :
		description || 'No description'
}
</script>

<template>
</template>
<style scoped>
</style>
