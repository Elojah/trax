<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	methods: {
		redirectHome() {
			this.$router.push({ name: 'home' });
		}
	}
});

</script>

<template>
	<v-navigation-drawer rail>
		<v-list density="compact" nav>
			<v-list-item prepend-icon="mdi-radar" title="Home" value="home" @click="redirectHome"></v-list-item>
		</v-list>
		<v-divider></v-divider>
	</v-navigation-drawer>
</template>

<style scoped></style>
