<script setup lang="ts">
import type { Permission, } from '@internal/user/role';
import PermissionTable from '@/components/permission/Table.vue';
import { ref } from 'vue';
import { useRoleStore } from '@/stores/role';
import { UpdateRoleReq } from '@internal/user/dto/role';

const props = defineProps<{
	colspan: number;
	permissions: Permission[] | undefined;
	roleID: Uint8Array | undefined
}>();

const store = useRoleStore();

const ps = ref();

const reset = () => {
	ps.value?.reset();
}

const edit = async () => {
	const values = ps.value?.values.map(ps.value.unhash);

	const req = UpdateRoleReq.create({
		iD: props.roleID,
		permissions: values,
	});

	await store.update(req);
}
</script>

<template>
	<tr>
		<td :colspan="props.colspan">
			<v-card outlined>
				<PermissionTable :permissions="props.permissions" ref="ps"></PermissionTable>
				<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn color="primary" variant="text" @click="edit">
						Edit
					</v-btn>
					<v-btn color="warning" variant="text" @click="reset">
						Reset
					</v-btn>
					<v-btn color="error" variant="text" @click="null">
						Delete
					</v-btn>
				</v-card-actions>
			</v-card>
		</td>
	</tr>
</template>
