<script setup lang="ts">
import Button from 'primevue/button';
import { toRefs, ref } from "vue";

const search = ref('');
</script>

<template>
	<div class="overflow-auto flex-1 flex flex-col gap-8 transition-all pb-20">
		<div class="flex flex-col md:flex-row md:items-center md:justify-between py-6 px-8 gap-4 pb-0">
			<div class="flex flex-col gap-1">
				<h1 class="text-lg font-medium text-surface-900 dark:text-surface-0 leading-tight">Dashboard</h1>
				<p class="text-surface-500 leading-tight">Solar metrics</p>
			</div>
			<div class="flex items-center gap-2">
				<IconField>
					<InputIcon class="pi pi-search" />
					<InputText v-model="search" placeholder="Search" size="small" class="md:w-56 w-full" />
				</IconField>
				<Button icon="pi pi-bell" severity="secondary" outlined class="!h-10 shrink-0" />
			</div>
		</div>

		<hr class="border-t border-surface-200 dark:border-surface-700" />
		<div class="flex gap-8 flex-col px-8 w-full"></div>
	</div>
</template>
