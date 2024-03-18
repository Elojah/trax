<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const redirectHome = function () {
	router.push({ name: 'home' });
}

</script>

<template>
	<v-navigation-drawer rail>
		<v-list density="compact" nav>
			<v-list-item prepend-icon="mdi-home" title="Home" value="home" @click="redirectHome">
				<template v-slot:prepend>
					<v-icon color="success"></v-icon>
				</template>
			</v-list-item>
		</v-list>
		<v-divider></v-divider>
	</v-navigation-drawer>
</template>

<style scoped></style>
