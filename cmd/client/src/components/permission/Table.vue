<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useGroupStore } from '@/stores/group';
import { Permission } from '@internal/user/role';
import { ulid } from '@/utils/ulid';
import { Command, Resource } from '@internal/user/role';
import { computed, ref, toRefs } from 'vue';
import Checkbox from 'primevue/checkbox';

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
const selectedGroupID = computed(() => ulid(selectedGroups.value.at(0)?.group?.iD));

// Define the resource and command labels for display
const resourceLabels: Record<Resource, string> = {
	[Resource.R_asset]: 'Asset',
	[Resource.R_group]: 'Group',
	[Resource.R_operation]: 'Operation',
	[Resource.R_role]: 'Role',
	[Resource.R_user]: 'User',
};

const commandLabels: Record<Command, string> = {
	[Command.C_read]: 'Read',
	[Command.C_edit]: 'Edit',
	[Command.C_write]: 'Write',
};

// Get all resources and commands for the matrix
const resources = Object.values(Resource).filter(v => typeof v === 'number') as Resource[];
const commands = Object.values(Command).filter(v => typeof v === 'number') as Command[];

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
		return new Map();
	}

	const result: Map<string, boolean> = new Map();
	perms?.forEach((perm) => {
		result.set(hash(perm.resource, perm.command), true);
	});

	return result;
}

const values = ref<Map<string, boolean>>(init(props.permissions));

const clear = () => {
	values.value = new Map();
}

const reset = () => {
	values.value = init(props.permissions);
}

// Check if a specific permission is checked
const isChecked = (resource: Resource, command: Command): boolean => {
	return values.value.has(hash(resource, command));
}

// Toggle a specific permission
const togglePermission = (resource: Resource, command: Command) => {
	if (props.disabled) {
		return;
	}

	const key = hash(resource, command);
	values.value.set(key, true)
}

// Get all selected permissions as Permission objects
const selected = (): Permission[] => {
	const selectedPermissions: Permission[] = [];
	for (const [key, checked] of values.value.entries()) {
		if (checked) {
			selectedPermissions.push(unhash(key));
		}
	}
	return selectedPermissions;
}

defineExpose({
	clear,
	reset,
	unhash,
	selected,
});

</script>
<template>
	<div class="permissions-table p-4">
		<div class="overflow-x-auto">
			<table class="w-full border-collapse bg-surface-0 dark:bg-surface-900 rounded-lg shadow-sm">
				<thead>
					<tr class="bg-surface-50 dark:bg-surface-800">
						<th
							class="text-left p-3 font-semibold text-surface-700 dark:text-surface-200 border-b border-surface-200 dark:border-surface-700 rounded-tl-lg">
							Resource
						</th>
						<th v-for="(command, index) in commands" :key="command" :class="[
							'text-center p-3 font-semibold text-surface-700 dark:text-surface-200 border-b border-surface-200 dark:border-surface-700 min-w-24',
							index === commands.length - 1 ? 'rounded-tr-lg' : ''
						]">
							{{ commandLabels[command] }}
						</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(resource, rowIndex) in resources" :key="resource"
						class="hover:bg-surface-50 dark:hover:bg-surface-800 transition-colors duration-200">
						<td
							class="p-3 font-medium text-surface-900 dark:text-surface-100 border-b border-surface-100 dark:border-surface-800">
							{{ resourceLabels[resource] }}
						</td>
						<td v-for="command in commands" :key="`${resource}-${command}`"
							class="text-center p-3 border-b border-surface-100 dark:border-surface-800">
							<div class="flex justify-center">
								<Checkbox :model-value="isChecked(resource, command)" :disabled="disabled"
									:binary="true" @update:model-value="togglePermission(resource, command)"
									class="permission-checkbox" />
							</div>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>
</template>

<style scoped>
.permissions-table .permission-checkbox {
	scale: 1.1;
}

.permissions-table .permission-checkbox:not([disabled]):hover {
	scale: 1.2;
	transition: scale 0.2s ease-in-out;
}
</style>
