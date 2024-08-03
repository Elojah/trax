<script setup lang="ts">
import PermissionTable from '@/components/permission/Table.vue';
import { ref, toRefs } from 'vue';
import { useRoleStore } from '@/stores/role';
import { useUserStore } from '@/stores/user';
import { useErrorsStore } from '@/stores/errors';
import { DeleteRoleReq, RolePermission, UpdateRoleReq } from '@internal/user/dto/role';

const props = defineProps<{
	colspan: number;
	item: RolePermission | undefined,
}>();

const store = useRoleStore();
const {
	delete_: delete_,
} = store;

const userStore = useUserStore();

const errorsStore = useErrorsStore()
const { success, message } = toRefs(errorsStore)

const ps = ref();

const reset = () => {
	ps.value?.reset();
}

const edit = async () => {
	const values = ps.value?.values.map(ps.value.unhash);

	let ok = true;
	try {
		const req = UpdateRoleReq.create({
			iD: props.item?.role?.iD,
			permissions: values,
		});

		await store.update(req);
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false;
	}
	if (ok) {
		message.value = `Role ${props.item?.role?.name} updated successfully`;
		success.value = true;
	}
}

const admin = props.item?.role?.name === 'admin'

const deleteRole = async () => {
	let ok = true;
	try {
		const req = DeleteRoleReq.create({
			iD: props.item?.role?.iD,
		})
		await delete_(req);
	} catch (e) {
		errorsStore.showGRPC(e)
		ok = false;
	}
	if (ok) {
		message.value = `Role ${props.item?.role?.name} deleted successfully`;
		success.value = true;
		// async refresh users to remove role
		userStore.deleteRoleGlobal(props.item?.role?.iD!);
	}
}

</script>

<!-- TODO: Add edit name -->
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
