<script setup lang="ts">
import PermissionTable from '@/components/permission/Table.vue';
import { ref } from 'vue';
import { useRoleStore } from '@/stores/role';
import { DeleteRoleReq, RolePermission, UpdateRoleReq } from '@internal/user/dto/role';

const props = defineProps<{
	colspan: number;
	item: RolePermission | undefined,
}>();

const store = useRoleStore();
const {
	delete_: delete_,
} = store;

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

const deleteRole = async () => {
	const req = DeleteRoleReq.create({
		iD: props.item?.role?.iD,
	})
	await delete_(req);
}

</script>

<template>
	<tr>
		<td :colspan="props.colspan">
			<v-card outlined class="mb-4">
				<PermissionTable :permissions="props.item?.permissions" ref="ps" :disabled="admin"></PermissionTable>
				<v-card-actions v-if="!admin">
					<v-spacer></v-spacer>
					<v-btn color="primary" variant="text" prepend-icon="mdi-pencil" @click="edit">
						Edit
						<template v-slot:prepend>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
					<v-btn color="warning" variant="text" prepend-icon="mdi-pencil-off" @click="reset">
						Reset
						<template v-slot:prepend>
							<v-icon color="warning"></v-icon>
						</template>
					</v-btn>
					<v-btn color="error" variant="text" prepend-icon="mdi-trash-can" @click="deleteRole">
						Delete
						<template v-slot:prepend>
							<v-icon color="error"></v-icon>
						</template>
					</v-btn>
				</v-card-actions>
			</v-card>
		</td>
	</tr>
</template>
