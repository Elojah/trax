<script setup lang="ts">
import type { Permission } from '@internal/user/role';
import { Command, Resource } from '@internal/user/role';
import { ref } from 'vue';

const props = defineProps<{
	permissions: Permission[] | undefined;
}>();

const hash = (resource: Resource, command: Command): string => {
	return `${resource.toString()}_${command.toString()}`;
}

const unhash = (value: string): Permission => {
	const [resource, command] = value.split('_');

	return {
		resource: Number(resource) as Resource,
		command: Number(command) as Command,
	} as Permission;
}

const values = ref<string[]>(((perms: Permission[] | undefined) => {
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
})(props.permissions));


const clear = () => {
	values.value = [];
}

defineExpose({
	values,
	clear,
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
					Create
				</th>
				<th class="text-left">
					Update
				</th>
				<th class="text-left">
					Delete
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
						:value="hash(Resource.R_asset, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:value="hash(Resource.R_asset, Command.C_create)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:value="hash(Resource.R_asset, Command.C_update)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="error" v-model="values"
						:value="hash(Resource.R_asset, Command.C_delete)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					entity
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:value="hash(Resource.R_entity, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:value="hash(Resource.R_entity, Command.C_create)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:value="hash(Resource.R_entity, Command.C_update)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="error" v-model="values"
						:value="hash(Resource.R_entity, Command.C_delete)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					operation
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:value="hash(Resource.R_operation, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:value="hash(Resource.R_operation, Command.C_create)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:value="hash(Resource.R_operation, Command.C_update)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="error" v-model="values"
						:value="hash(Resource.R_operation, Command.C_delete)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					role
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:value="hash(Resource.R_role, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:value="hash(Resource.R_role, Command.C_create)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:value="hash(Resource.R_role, Command.C_update)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="error" v-model="values"
						:value="hash(Resource.R_role, Command.C_delete)"></v-checkbox-btn>
				</td>
			</tr>
			<tr>
				<td>
					user
				</td>
				<td>
					<v-checkbox-btn color="info" v-model="values"
						:value="hash(Resource.R_user, Command.C_read)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="success" v-model="values"
						:value="hash(Resource.R_user, Command.C_create)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="warning" v-model="values"
						:value="hash(Resource.R_user, Command.C_update)"></v-checkbox-btn>
				</td>
				<td>
					<v-checkbox-btn color="error" v-model="values"
						:value="hash(Resource.R_user, Command.C_delete)"></v-checkbox-btn>
				</td>
			</tr>
		</tbody>
	</v-table>
</template>
