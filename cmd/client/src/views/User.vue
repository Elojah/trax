<script setup lang="ts">
import User from '@/components/user/Me.vue';
</script>

<template>
	<v-main class="bg">
		<v-container fluid>
			<v-row>
				<v-col class="mx-auto px-12 py-12" cols="12" md="4">
					<User />
				</v-col>
			</v-row>
		</v-container>
	</v-main>
</template>

<style scoped>
.bg {
	background: url('@/assets/img/main-background.svg') !important;
}
</style>
