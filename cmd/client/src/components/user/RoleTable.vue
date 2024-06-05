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
		sortable: false,
	},
	{
		title: 'Created',
		key: 'created_at',
		align: 'end',
		sortable: true,
	},
];

</script>
<template>
	<v-data-table density="compact" :headers="headers" :items="props?.item?.roles">
	</v-data-table>
</template>
