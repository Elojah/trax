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
		<v-container fluid class="fill-height h-screen py-0">
			<v-row no-gutters>
				<v-col cols="5">
					<v-sheet color="transparent" class="d-flex flex-column px-4 h-screen fill-height fill-width">
						<EntityTable ref="tableEntities"></EntityTable>
					</v-sheet>
				</v-col>
				<v-col cols="7">
					<v-sheet color="transparent" class="d-flex flex-column px-4 h-screen fill-height fill-width">
						<v-container class="flex-shrink-1">
							<v-row>
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
							</v-row>
						</v-container>
						<v-divider></v-divider>
						<!-- Info -->
						<EntityDetails v-if="view === 0" :onDelete="() => { tableEntities?.value?.list() }">
						</EntityDetails>
						<!-- Roles -->
						<RoleTable v-if="view === 1"></RoleTable>
						<!-- User -->
						<UserTable v-if="view === 2"></UserTable>
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

.main-color-background {
	background-color: #263238;
}
</style>
