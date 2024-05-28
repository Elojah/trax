<script setup lang="ts">
import type { Permission, } from '@internal/user/role';
import { Command, Resource } from '@internal/user/role';
import { ref } from 'vue';

const props = defineProps<{
	colspan: number;
	permissions: Permission[] | undefined;
}>();

const ps = ref<string[]>(((perms: Permission[] | undefined) => {
	if (!perms) {
		return [];
	}

	const result: string[] = [];
	perms?.forEach((perm) => {
		const resource = perm.resource;
		const command = perm.command;
		result.push(`${resource.toString()}_${command.toString()}`);
	});

	return result;
})(props.permissions));


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
								<v-checkbox-btn color="info" v-model="ps"
									:value="`${Resource.R_asset.toString()}_${Command.C_read.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="success" v-model="ps"
									:value="`${Resource.R_asset.toString()}_${Command.C_create.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="warning" v-model="ps"
									:value="`${Resource.R_asset.toString()}_${Command.C_update.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="error" v-model="ps"
									:value="`${Resource.R_asset.toString()}_${Command.C_delete.toString()}`"></v-checkbox-btn>
							</td>
						</tr>
						<tr>
							<td>
								entity
							</td>
							<td>
								<v-checkbox-btn color="info" v-model="ps"
									:value="`${Resource.R_entity.toString()}_${Command.C_read.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="success" v-model="ps"
									:value="`${Resource.R_entity.toString()}_${Command.C_create.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="warning" v-model="ps"
									:value="`${Resource.R_entity.toString()}_${Command.C_update.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="error" v-model="ps"
									:value="`${Resource.R_entity.toString()}_${Command.C_delete.toString()}`"></v-checkbox-btn>
							</td>
						</tr>
						<tr>
							<td>
								operation
							</td>
							<td>
								<v-checkbox-btn color="info" v-model="ps"
									:value="`${Resource.R_operation.toString()}_${Command.C_read.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="success" v-model="ps"
									:value="`${Resource.R_operation.toString()}_${Command.C_create.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="warning" v-model="ps"
									:value="`${Resource.R_operation.toString()}_${Command.C_update.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="error" v-model="ps"
									:value="`${Resource.R_operation.toString()}_${Command.C_delete.toString()}`"></v-checkbox-btn>
							</td>
						</tr>
						<tr>
							<td>
								role
							</td>
							<td>
								<v-checkbox-btn color="info" v-model="ps"
									:value="`${Resource.R_role.toString()}_${Command.C_read.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="success" v-model="ps"
									:value="`${Resource.R_role.toString()}_${Command.C_create.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="warning" v-model="ps"
									:value="`${Resource.R_role.toString()}_${Command.C_update.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="error" v-model="ps"
									:value="`${Resource.R_role.toString()}_${Command.C_delete.toString()}`"></v-checkbox-btn>
							</td>
						</tr>
						<tr>
							<td>
								user
							</td>
							<td>
								<v-checkbox-btn color="info" v-model="ps"
									:value="`${Resource.R_user.toString()}_${Command.C_read.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="success" v-model="ps"
									:value="`${Resource.R_user.toString()}_${Command.C_create.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="warning" v-model="ps"
									:value="`${Resource.R_user.toString()}_${Command.C_update.toString()}`"></v-checkbox-btn>
							</td>
							<td>
								<v-checkbox-btn color="error" v-model="ps"
									:value="`${Resource.R_user.toString()}_${Command.C_delete.toString()}`"></v-checkbox-btn>
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
