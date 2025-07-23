<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { FilterMatchMode, FilterOperator } from '@primevue/core/api';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import Table from '@/components/group/Table.vue';

// [Info, Roles, User]
</script>

<template>
	<div class="min-h-screen flex flex-col relative flex-auto">
		<div
			class="flex justify-between items-center px-8 py-4 bg-surface-0 dark:bg-surface-950 relative lg:static border-b border-surface gap-8">
			<div class="flex flex-1 gap-4">
				<a v-styleclass="{
					selector: '#app-sidebar-9',
					enterFromClass: 'hidden',
					enterActiveClass: 'animate-fadeinleft',
					leaveToClass: 'hidden',
					leaveActiveClass: 'animate-fadeoutleft',
					hideOnOutsideClick: true
				}" class="cursor-pointer flex items-center justify-center lg:hidden text-surface-700 dark:text-surface-100">
					<i class="pi pi-bars !text-xl" />
				</a>
				<IconField icon-position="left">
					<InputIcon class="pi pi-search text-surface-700 dark:text-surface-100" />
					<InputText type="text" class="border-0 w-full sm:w-80" placeholder="Search" />
				</IconField>
			</div>
		</div>
		<div class="p-8 flex flex-col flex-auto">
			<Table />
		</div>
	</div>
</template>
<style scoped></style>
