<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useEntityStore } from '@/stores/entity';
import { ulid } from '@/utils/ulid';
import { computed, ref, toRefs } from 'vue';
import type { UserRoles } from '@internal/user/dto/user';
import type { ReadonlyHeaders } from '@/utils/headers';

const props = defineProps<{
	item: UserRoles | undefined;
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
		align: 'start',
		sortable: true,
	},
	{
		title: ' ',
		key: 'delete',
		align: 'end',
		sortable: false,
	},
];

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
			<p class="font-italic font-weight-light">
				{{ new Date(Number(item.createdAt) * 1000).toLocaleDateString('en-GB') }}
			</p>
		</td>
		<td class="px-1 py-1">
			<v-col class="d-flex justify-end">
				<v-btn variant="tonal" prepend-icon="mdi-trash-can" color="error" v-bind="props">
					Delete
					<template v-slot:prepend>
						<v-icon color="error"></v-icon>
					</template>
				</v-btn>
			</v-col>
		</td>
	</tr>
</template>
<template v-slot:bottom="{ columns }">
	<tr>
		<td :colspan="columns.length">
			<v-btn variant="tonal" prepend-icon="mdi-plus-box" color="primary" v-bind="props">
				New
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-btn>
		</td>
	</tr>
</template>
</v-data-table>
</td>
</tr>
</template>
