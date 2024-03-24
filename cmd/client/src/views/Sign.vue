<script setup lang="ts">
import Signup from '@/components/Signup.vue';
import Signin from '@/components/Signin.vue';
</script>

<template>
	<v-main class="bg">
		<v-container fluid>
			<v-row>
				<v-col class="mx-auto px-12 py-12" cols="12" md="4">
					<v-sheet class="px-1 rounded-xl" outlined color="success">
						<Signin />
					</v-sheet>
				</v-col>
				<v-divider vertical :thickness="1" color="success"></v-divider>
				<v-col class="mx-auto px-12 py-12" cols="12" md="4">
					<v-sheet class="px-1 rounded-xl" outlined color="success">
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
