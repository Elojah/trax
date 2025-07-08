<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useGroupStore } from '@/stores/group';
import { Permission } from '@internal/user/role';
import { ulid } from '@/utils/ulid';
import { Command, Resource } from '@internal/user/role';
import { computed, ref, toRefs } from 'vue';

const props = defineProps<{
	permissions: Permission[] | undefined;
	disabled: boolean;
}>();

const authStore = useAuthStore();
const {
	claims: claims,
} = toRefs(authStore);

const groupStore = useGroupStore();
const {
	selected: selectedGroups,
} = toRefs(groupStore);
const selectedGroupID = computed(() => ulid(selectedGroups.value.at(0)?.iD));

const hash = (resource: Resource, command: Command): string => {
	return `${resource.toString()}_${command.toString()}`;
}

const unhash = (value: string): Permission => {
	const [resource, command] = value.split('_');

	return Permission.create({
		resource: Number(resource) as Resource,
		command: Number(command) as Command,
	});
}

const init = (perms: Permission[] | undefined) => {
	if (!perms) {
		return [];
	}

	const result: string[] = [];
	perms?.forEach((perm) => {
		const resource = perm.resource;
		const command = perm.command;
		result.push(hash(resource, command));
	});

	return result;
}

const values = ref<string[]>(init(props.permissions));

const clear = () => {
	values.value = [];
}

const reset = () => {
	values.value = init(props.permissions);
}

defineExpose({
	values,
	clear,
	reset,
	unhash,
});

</script>
<template>
</template>
