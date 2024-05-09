<script setup lang="ts">
import EntityTable from '@/components/entity/Table.vue';
import EntityDetails from '@/components/entity/Details.vue';
import RoleTable from '@/components/role/Table.vue';
import UserTable from '@/components/user/Table.vue';
import { ref } from 'vue';

// [Info, Roles, User]
const view = ref(0)

const tableEntities = ref();
</script>

<template>
	<v-row>
		<v-col cols="4">
			<EntityTable ref="tableEntities"></EntityTable>
		</v-col>
		<v-divider vertical></v-divider>
		<v-col class="mx-auto" cols="6">
			<v-col class="d-flex justify-center" cols="12">
				<v-btn-toggle v-model="view">
					<v-btn class="main-color-background" size="large" variant="outlined" append-icon="mdi-domain">
						Info
						<template v-slot:append>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
					<v-btn class="main-color-background" size="large" variant="outlined" append-icon="mdi-account-cog">
						Roles
						<template v-slot:append>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
					<v-btn class="main-color-background" size="large" variant="outlined"
						append-icon="mdi-account-multiple">
						Users
						<template v-slot:append>
							<v-icon color="primary"></v-icon>
						</template>
					</v-btn>
				</v-btn-toggle>
			</v-col>
			<!-- Info -->
			<EntityDetails v-if="view === 0" :onDelete="() => { tableEntities?.value?.list() }"></EntityDetails>
			<!-- Roles -->
			<RoleTable v-if="view === 1"></RoleTable>
			<!-- User -->
			<UserTable v-if="view === 2"></UserTable>
			<v-divider></v-divider>
		</v-col>
	</v-row>
</template>
<style scoped>
.active-bg {
	background-color: #950c75 !important;
}

.main-color-background {
	background-color: #263238;
}

.transparent-background {
	background-color: transparent;
}

.row-odd {
	transition: background-color .2s ease-in-out;
	background-color: rgba(55, 71, 79, 0.3);
}

.row-odd:not(.row-hovered) {
	background-color: rgba(33, 33, 33, 0.3);
}

.row-even {
	transition: background-color .2s ease-in-out;
	background-color: rgba(55, 71, 79, 0.3);
}

.row-even:not(.row-hovered) {
	background-color: rgba(0, 229, 255, 0.3);
}

.row-selected {
	background-color: rgba(149, 12, 117, 0.6) !important;
}

.cursor-pointer {
	cursor: pointer;
}
</style>
