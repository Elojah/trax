<script setup lang="ts">
import type { Permission, } from '@internal/user/role';
import { Command, Resource } from '@internal/user/role';

const props = defineProps<{
	colspan: number;
	permissions: Permission[] | undefined;
}>();

const permissionMap = ((perms: Permission[] | undefined) => {
	if (!perms) {
		return new Map<number, Map<number, boolean>>();
	}

	const map = new Map<number, Map<number, boolean>>();
	perms?.forEach((perm) => {
		const resource = perm.resource;
		const command = perm.command;

		if (!map.has(resource)) {
			map.set(resource, new Map<number, boolean>());
		}

		map.get(resource)?.set(command, true);
	});
	return map;
})(props.permissions);


</script>
<template>
	<tr>
		<td :colspan="props.colspan">
			<v-card outlined class="main-color-background">
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
								<v-checkbox color="info"
									:v-model="permissionMap.get(Resource.R_asset)?.get(Command.C_read)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="success"
									:v-model="permissionMap.get(Resource.R_asset)?.get(Command.C_create)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="warning"
									:v-model="permissionMap.get(Resource.R_asset)?.get(Command.C_update)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="error"
									:v-model="permissionMap.get(Resource.R_asset)?.get(Command.C_delete)"></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								entity
							</td>
							<td>
								<v-checkbox color="info"
									:v-model="permissionMap.get(Resource.R_entity)?.get(Command.C_read)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="success"
									:v-model="permissionMap.get(Resource.R_entity)?.get(Command.C_create)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="warning"
									:v-model="permissionMap.get(Resource.R_entity)?.get(Command.C_update)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="error"
									:v-model="permissionMap.get(Resource.R_entity)?.get(Command.C_delete)"></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								operation
							</td>
							<td>
								<v-checkbox color="info"
									:v-model="permissionMap.get(Resource.R_operation)?.get(Command.C_read)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="success"
									:v-model="permissionMap.get(Resource.R_operation)?.get(Command.C_create)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="warning"
									:v-model="permissionMap.get(Resource.R_operation)?.get(Command.C_update)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="error"
									:v-model="permissionMap.get(Resource.R_operation)?.get(Command.C_delete)"></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								role
							</td>
							<td>
								<v-checkbox color="info"
									:v-model="permissionMap.get(Resource.R_role)?.get(Command.C_read)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="success"
									:v-model="permissionMap.get(Resource.R_role)?.get(Command.C_create)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="warning"
									:v-model="permissionMap.get(Resource.R_role)?.get(Command.C_update)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="error"
									:v-model="permissionMap.get(Resource.R_role)?.get(Command.C_delete)"></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								user
							</td>
							<td>
								<v-checkbox color="info"
									:v-model="permissionMap.get(Resource.R_user)?.get(Command.C_read)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="success"
									:v-model="permissionMap.get(Resource.R_user)?.get(Command.C_create)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="warning"
									:v-model="permissionMap.get(Resource.R_user)?.get(Command.C_update)"></v-checkbox>
							</td>
							<td>
								<v-checkbox color="error"
									:v-model="permissionMap.get(Resource.R_user)?.get(Command.C_delete)"></v-checkbox>
							</td>
						</tr>
					</tbody>
				</v-table>
				<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn color="primary" variant="text" @click="null">
						Edit
					</v-btn>
					<v-btn color="error" variant="text" @click="null">
						Delete
					</v-btn>
				</v-card-actions>
			</v-card>
		</td>
	</tr>
</template>
