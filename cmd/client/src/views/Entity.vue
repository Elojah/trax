<script setup lang="ts">
import Entity from '@/components/Entity.vue';
</script>

<template>
	<v-main class="bg">
		<v-container fluid>
			<v-sheet class="px-1 rounded-xl" outlined color="primary">
				<Entity />
			</v-sheet>
		</v-container>
	</v-main>
</template>

<style scoped>
.bg {
	background: url('@/assets/img/main-background.svg') !important;
}
</style>
