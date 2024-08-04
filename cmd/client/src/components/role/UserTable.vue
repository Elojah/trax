<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { useEntityStore } from '@/stores/entity';
import { ulid } from '@/utils/ulid';
import { computed, ref, toRefs } from 'vue';
import type { ReadonlyHeaders } from '@/utils/headers';
import UserTable from '@/components/user/Table.vue';
import { useRoleStore } from '@/stores/role';
import { useUserStore } from '@/stores/user';
import type { U } from '@internal/user/user';

const props = defineProps<{
	roleID: Uint8Array | undefined;
	colspan: number;
}>();

const userStore = useUserStore();

const authStore = useAuthStore();
const {
	claims: claims,
} = toRefs(authStore);

const entityStore = useEntityStore();
const {
	selected: selectedEntities,
} = toRefs(entityStore);
const selectedEntityID = computed(() => ulid(selectedEntities.value.at(0)?.iD));

const roleStore = useRoleStore();
const {
	users: users,
} = toRefs(roleStore);

const errorsStore = useErrorsStore();
const { success, message } = toRefs(errorsStore);

const us = computed(() => { return users.value.get(ulid(props.roleID)) })

const headers: ReadonlyHeaders = [
	{
		title: 'User',
		key: 'name',
		align: 'start',
		sortable: true,
	},
	{
		title: 'Created',
		key: 'created_at',
		align: 'end',
		sortable: true,
	},
	{
		title: ' ',
		key: 'delete',
		align: 'end',
		sortable: false,
	},
];

const deleteUserRole = async (userID: Uint8Array) => {
	let ok = true;
	try {
		await roleStore.deleteUser(props.roleID!, userID);
	} catch (e) {
		errorsStore.showGRPC(e);
		ok = false;
	}
	if (ok) {
		message.value = `Role deleted successfully to user`;
		success.value = true;
	}
};

const dialogAddUser = ref(false);

const closeAddUser = () => {
	dialogAddUser.value = false;
};

// User table properties

const addUserLoading = ref(false);

const addUser = async (u: U) => {
	addUserLoading.value = true;
	let ok = true;
	try {
		await roleStore.addUser(props?.roleID!, u?.iD);
	} catch (e) {
		errorsStore.showGRPC(e);
		ok = false;
	}
	if (ok) {
		message.value = `Role added successfully to user`;
		success.value = true;
	}
	addUserLoading.value = false;
};

</script>
<template>
	<tr v-if="us">
		<td :colspan="props.colspan">
			<v-data-table density="compact" :headers="headers" :items="us" class="">
				<template v-slot:item="{ item }">
	<tr v-if="item" :key="ulid(item.iD)">
		<td class="px-1 py-1">
			<p class="font-weight-bold">
				{{ item.lastName }} {{ item.firstName }}
			</p>
		</td>
		<td class="px-1 py-1">
			<div class="d-flex flex-row-reverse">
				<p class="font-italic font-weight-light">
					{{ new Date(Number(item.createdAt) * 1000).toLocaleDateString('en-GB') }}
				</p>
			</div>
		</td>
		<td class="px-1 py-1">
			<div class="d-flex flex-row-reverse">
				<v-btn variant="tonal" prepend-icon="mdi-trash-can" color="error" v-bind="props"
					@click="deleteUserRole(item.iD)">
					Remove
					<template v-slot:prepend>
						<v-icon color="error"></v-icon>
					</template>
				</v-btn>
			</div>
		</td>
	</tr>
</template>
<template v-slot:bottom>
	<div class="d-flex justify-center">
		<v-dialog v-model="dialogAddUser" max-width="1200px">
			<template v-slot:activator="{ props }">
				<v-btn variant="tonal" class="my-4" size="large" prepend-icon="mdi-plus-box" color="primary" v-bind="props">
					Add users
					<template v-slot:prepend>
						<v-icon color="primary"></v-icon>
					</template>
				</v-btn>
			</template>
			<v-sheet class="d-flex flex-column pa-4 fill-height fill-width" height="50vh">
				<UserTable :role-i-d="props.roleID" :addUser="addUser"></UserTable>
				<v-divider></v-divider>
				<v-btn color="error" variant="tonal" @click="closeAddUser">
					Close
				</v-btn>
			</v-sheet>
		</v-dialog>
	</div>
</template>
<template v-slot:top></template>
</v-data-table>
</td>
</tr>
</template>
