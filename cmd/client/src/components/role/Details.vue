<script setup lang="ts">
import type { Permission, } from '@internal/user/role';
import PermissionTable from '@/components/permission/Table.vue';

const props = defineProps<{
	colspan: number;
	permissions: Permission[] | undefined;
}>();
</script>

<template>
	<tr>
		<td :colspan="props.colspan">
			<v-card outlined class="main-color-background">
				<PermissionTable :permissions="props.permissions"></PermissionTable>
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
