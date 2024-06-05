<script setup lang="ts">
import PermissionTable from '@/components/permission/Table.vue';
import { ref } from 'vue';
import { useRoleStore } from '@/stores/role';
import { RolePermission, UpdateRoleReq } from '@internal/user/dto/role';

const props = defineProps<{
	colspan: number;
	item: RolePermission | undefined,
}>();

const store = useRoleStore();

const ps = ref();

const reset = () => {
	ps.value?.reset();
}

const edit = async () => {
	const values = ps.value?.values.map(ps.value.unhash);

	const req = UpdateRoleReq.create({
		iD: props.item?.role?.iD,
		permissions: values,
	});

	await store.update(req);
}

const admin = props.item?.role?.name === 'admin'
</script>

<template>
	<tr>
		<td :colspan="props.colspan">
			<v-card outlined class="mb-4">
				<PermissionTable :permissions="props.item?.permissions" ref="ps" :disabled="admin"></PermissionTable>
				<v-card-actions v-if="!admin">
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
