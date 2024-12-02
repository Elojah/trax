<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useEntityStore } from '@/stores/entity';
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

const entityStore = useEntityStore();
const {
	selected: selectedEntities,
} = toRefs(entityStore);
const selectedEntityID = computed(() => ulid(selectedEntities.value.at(0)?.iD));

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
	<v-table density="compact">
		<thead>
			<tr>
				<th>
				</th>
				<th class="text-left">
					Read
				</th>
				<th class="text-left">
					Edit
				</th>
				<th class="text-left">
					Write
				</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>
					asset
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_asset]]?.commands[Command[Command.C_read]] === undefined"
						:value="hash(Resource.R_asset, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_asset]]?.commands[Command[Command.C_edit]] === undefined"
						:value="hash(Resource.R_asset, Command.C_edit)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_asset]]?.commands[Command[Command.C_write]] === undefined"
						:value="hash(Resource.R_asset, Command.C_write)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					entity
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_entity]]?.commands[Command[Command.C_read]] === undefined"
						:value="hash(Resource.R_entity, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_entity]]?.commands[Command[Command.C_edit]] === undefined"
						:value="hash(Resource.R_entity, Command.C_edit)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_entity]]?.commands[Command[Command.C_write]] === undefined"
						:value="hash(Resource.R_entity, Command.C_write)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					operation
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_operation]]?.commands[Command[Command.C_read]] === undefined"
						:value="hash(Resource.R_operation, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_operation]]?.commands[Command[Command.C_edit]] === undefined"
						:value="hash(Resource.R_operation, Command.C_edit)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_operation]]?.commands[Command[Command.C_write]] === undefined"
						:value="hash(Resource.R_operation, Command.C_write)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					role
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_role]]?.commands[Command[Command.C_read]] === undefined"
						:value="hash(Resource.R_role, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_role]]?.commands[Command[Command.C_edit]] === undefined"
						:value="hash(Resource.R_role, Command.C_edit)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_role]]?.commands[Command[Command.C_write]] === undefined"
						:value="hash(Resource.R_role, Command.C_write)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					user
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_user]]?.commands[Command[Command.C_read]] === undefined"
						:value="hash(Resource.R_user, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_user]]?.commands[Command[Command.C_edit]] === undefined"
						:value="hash(Resource.R_user, Command.C_edit)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:disabled="props.disabled || claims?.entities[selectedEntityID]?.resources[Resource[Resource.R_user]]?.commands[Command[Command.C_write]] === undefined"
						:value="hash(Resource.R_user, Command.C_write)"></v-checkbox-btn>
				</td>
			</tr>
		</tbody>
	</v-table>
</template>
