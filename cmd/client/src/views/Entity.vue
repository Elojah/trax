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
	<v-main class="bg">
		<v-container fluid>
			<v-row>
				<v-col cols="5">
					<EntityTable ref="tableEntities"></EntityTable>
				</v-col>
				<v-col class="mx-auto" cols="7">
					<v-col class="d-flex justify-center" cols="12">
						<v-btn-toggle v-model="view">
							<v-btn class="main-color-background" size="large" variant="outlined"
								append-icon="mdi-domain">
								Info
								<template v-slot:append>
									<v-icon color="primary"></v-icon>
								</template>
							</v-btn>
							<v-btn class="main-color-background" size="large" variant="outlined"
								append-icon="mdi-account-cog">
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
					<v-divider></v-divider>
					<!-- Info -->
					<v-sheet v-if="view === 0">
						<EntityDetails :onDelete="() => { tableEntities?.value?.list() }">
						</EntityDetails>
					</v-sheet>
					<!-- Roles -->
					<v-sheet v-if="view === 1">
						<RoleTable></RoleTable>
					</v-sheet>
					<!-- User -->
					<v-sheet v-if="view === 2">
						<UserTable></UserTable>
					</v-sheet>
					<v-divider></v-divider>
				</v-col>
			</v-row>
		</v-container>
	</v-main>
</template>

<style scoped>
.bg {
	background: url('@/assets/img/main-background.svg') !important;
}

.main-color-background {
	background-color: #263238;
}
</style>
