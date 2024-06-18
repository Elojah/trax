<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useEntityStore } from '@/stores/entity';
import { ulid } from '@/utils/ulid';
import { computed, ref, toRefs, watch } from 'vue';
import type { ReadonlyHeaders } from '@/utils/headers';
import RoleTable from '@/components/role/Table.vue';
import { useUserStore } from '@/stores/user';
import type { RolePermission } from '@internal/user/dto/role';

const props = defineProps<{
	userID: Uint8Array | undefined;
	colspan: number;
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

const userStore = useUserStore();
const {
	roles: roles,
} = toRefs(userStore);

const user = computed(() => { return roles.value.get(ulid(props.userID)) })

const name = `${user.value?.user?.lastName} ${user.value?.user?.firstName}`;

const headers: ReadonlyHeaders = [
	{
		title: 'Name',
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

const deleteUserRole = async (roleID: Uint8Array) => {
	await userStore.deleteRole(props.userID!, roleID);
};

const dialogAddRole = ref(false);

const closeAddRole = () => {
	dialogAddRole.value = false;
};

// Role table properties

const addRoleLoading = ref(false);

const addRole = async (role: RolePermission) => {
	addRoleLoading.value = true;
	await userStore.addRole(props?.userID!, role.role?.iD!);
	addRoleLoading.value = false;
};

</script>
<template>
	<tr v-if="user">
		<td :colspan="props.colspan">
			<v-data-table density="compact" :headers="headers" :items="user.roles" class="">
				<template v-slot:item="{ item }">
	<tr v-if="item" :key="item.name">
		<td class="px-1 py-1">
			<p class="font-weight-bold">
				{{ item.name }}
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
	<div class="d-flex justify-center px-1 py-6">
		<v-dialog v-model="dialogAddRole" max-width="1200px">
			<template v-slot:activator="{ props }">
				<v-btn variant="tonal" size="large" prepend-icon="mdi-plus-box" color="primary" v-bind="props">
					Add roles
					<template v-slot:prepend>
						<v-icon color="primary"></v-icon>
					</template>
				</v-btn>
			</template>
			<v-sheet class="px-1 rounded-lg" outlined color="primary">
				<v-card class="px-6 py-6 rounded-lg" variant="elevated">
					<v-card-title>
						<span class="text-h6">Add roles to {{ name }}</span>
					</v-card-title>
					<v-card-text>
						<RoleTable :user-i-d="props.userID" :addRole="addRole"></RoleTable>
					</v-card-text>
					<v-card-actions>
						<v-spacer></v-spacer>
						<v-btn color="error" variant="text" @click="closeAddRole">
							Close
						</v-btn>
					</v-card-actions>
				</v-card>
			</v-sheet>
		</v-dialog>
	</div>
</template>
<template v-slot:top></template>
</v-data-table>
</td>
</tr>
</template>
