<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useEntityStore } from '@/stores/entity';
import { ulid } from '@/utils/ulid';
import { computed, ref, toRefs } from 'vue';
import type { UserRoles } from '@internal/user/dto/user';
import type { ReadonlyHeaders } from '@/utils/headers';
import type { VForm } from 'vuetify/components';
import RoleTable from '@/components/role/Table.vue';
import { useRoleStore } from '@/stores/role';

const props = defineProps<{
	item: UserRoles | undefined;
	colspan: number;
}>();

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const authStore = useAuthStore();
const {
	claims: claims,
} = toRefs(authStore);

const entityStore = useEntityStore();
const {
	selected: selectedEntities,
} = toRefs(entityStore);
const selectedEntityID = computed(() => ulid(selectedEntities.value.at(0)?.iD));


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

const store = useRoleStore();
const {
	selected: selected,
} = toRefs(store);


const dialogAddRole = ref(false);

const closeAddRole = () => {
	dialogAddRole.value = false;
};

const addRole = async () => {
	console.log(selected)
};


</script>
<template>
	<tr>
		<td :colspan="props.colspan">
			<v-data-table density="compact" :headers="headers" :items="props?.item?.roles" class="">
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
				<v-btn variant="tonal" prepend-icon="mdi-trash-can" color="error" v-bind="props">
					Delete
					<template v-slot:prepend>
						<v-icon color="error"></v-icon>
					</template>
				</v-btn>
			</div>
		</td>
	</tr>
</template>
<template v-slot:bottom>
	<div class="d-flex justify-center px-1 py-1">
		<v-dialog v-model="dialogAddRole" max-width="800px">
			<template v-slot:activator="{ props }">
				<v-btn variant="tonal" size="large" prepend-icon="mdi-plus-box" color="primary" v-bind="props">
					Add roles
					<template v-slot:prepend>
						<v-icon color="primary"></v-icon>
					</template>
				</v-btn>
			</template>
			<v-sheet class="px-1 rounded-xl" outlined color="primary">
				<v-card class="px-6 py-6 rounded-xl" variant="elevated">
					<v-form ref="form" v-model="valid" lazy-validation>
						<v-card-title>
							<span class="text-h6">Add roles to user {{ props.item?.user?.email }}</span>
						</v-card-title>
						<v-card-text>
							<RoleTable :selection="true"></RoleTable>
						</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn variant="text" @click="closeAddRole">
								Cancel
							</v-btn>
							<v-btn color="primary" variant="text" @click="addRole">
								Add
							</v-btn>
						</v-card-actions>
					</v-form>
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
