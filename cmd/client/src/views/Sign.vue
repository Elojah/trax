<script setup lang="ts">
import Signup from '@/components/sign/Signup.vue';
import Signin from '@/components/sign/Signin.vue';
</script>

<template>
	<v-main class="bg">
		<v-container fluid>
			<v-row>
				<v-col class="mx-auto px-12 py-12" cols="12" md="4">
					<v-sheet class="px-1 rounded-lg" outlined color="primary">
						<Signin />
					</v-sheet>
				</v-col>
				<v-divider vertical :thickness="1" color="primary"></v-divider>
				<v-col class="mx-auto px-12 py-12" cols="12" md="4">
					<v-sheet class="px-1 rounded-lg" outlined color="primary">
						<Signup />
					</v-sheet>
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
