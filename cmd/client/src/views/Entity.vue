<script setup lang="ts">
import Entity from '@/components/Entity.vue';
</script>

<template>
	<v-main class="bg">
		<v-container fluid>
			<Entity />
		</v-container>
	</v-main>
</template>

<style scoped>
.bg {
	background: url('@/assets/img/main-background.svg') !important;
}
</style>
