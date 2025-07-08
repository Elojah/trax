<script setup lang="ts">
import { useGroupStore } from '@/stores/group';
import { computed, ref, toRefs } from 'vue';
import { marked } from "marked";
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { DeleteGroupReq, UpdateGroupReq } from '@internal/user/dto/group';

const props = defineProps<{
	onDelete: () => void;
}>();

const authStore = useAuthStore();
const store = useGroupStore();
const {
	selected: selected,
} = toRefs(store);

const errorsStore = useErrorsStore()
const { success, message } = toRefs(errorsStore)

const g = computed(() => selected.value.at(0));

const mdDescription = computed(() => {
	return marked.parse(g.value?.description || `*no description*`);
});

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDelete = async () => {
	let ok = true;
	try {
		await store.delete_(DeleteGroupReq.create({ iD: g.value?.iD }));
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false
	}

	if (ok) {
		message.value = `Group ${g.value?.name} deleted successfully`;
		success.value = true;
	}

	selected.value = [];
	await authStore.refreshToken();
	props.onDelete();

	dialogDelete.value = false;
};

const name = ref(false);
const description = ref(false);

const updateName = async function () {
	if (name.value) {
		let ok = true;
		try {
			await store.update(UpdateGroupReq.create({
				iD: g.value?.iD,
				name: { value: g.value?.name },
			}));
		} catch (e) {
			errorsStore.showGRPC(e)
			ok = false
		}
		if (ok) {
			message.value = `Group ${g.value?.name} updated name successfully`;
			success.value = true;
		}
	}

	name.value = !name.value
};

const updateDescription = async function () {
	if (description.value) {
		let ok = true;
		try {
			await store.update(UpdateGroupReq.create({
				iD: g.value?.iD,
				description: { value: g.value?.description },
			}));
		} catch (e) {
			errorsStore.showGRPC(e)
			ok = false
		}
		if (ok) {
			message.value = `Group ${g.value?.name} updated description successfully`;
			success.value = true;
		}
	}

	description.value = !description.value
};
</script>

<template>
</template>
<style scoped></style>
