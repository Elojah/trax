<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { Entity } from '@internal/user/entity';
import { ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';
import { ulid } from '@/utils/ulid';

const form = ref<VForm | null>(null);
const valid = ref(null as boolean | null)

const authStore = useAuthStore();
const {
	profile: profile,
} = toRefs(authStore);


const dialog = ref(false);
const close = () => {
	dialog.value = false;
	entity.value = Entity.create({})
};

const blankEntity = Entity.create({});

const save = () => {
	// TODO: save new entity
	dialog.value = false;
	entity.value = blankEntity;
};

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDeleteEntity = () => {
	// TODO: delete entity
	dialogDelete.value = false;
};

const entities: Entity[] = [
	Entity.create({ name: 'test1', avatarURL: 'https://gravatar.com/avatar/a7bb3266897ad708becc0a5eaff0b557?s=400&d=robohash&r=x' }),
	Entity.create({ name: 'test2', avatarURL: 'https://gravatar.com/avatar/a7bb3266897ad708becc0a5eaff0b557?s=400&d=robohash&r=x' }),
	Entity.create({ name: 'test3', avatarURL: 'https://gravatar.com/avatar/a7bb3266897ad708becc0a5eaff0b557?s=400&d=robohash&r=x' }),
]
const entity = ref(Entity.create({ name: 'test_init' }) as Entity);

const editEntity = (entity: Entity) => {
};
const deleteEntity = (entity: Entity) => {
};

</script>

<template>
	<v-toolbar class="bar-bg rounded-t-xl px-6" color="transparent" floating>
		<v-icon color="primary" size="large" icon="mdi-domain"></v-icon>
		<v-toolbar-title class="text-h6"> ENTITIES </v-toolbar-title>
		<v-dialog v-model="dialog" max-width="500px">
			<template v-slot:activator="{ props }">
				<v-btn class="rounded-xl" variant="outlined" color="primary" prepend-icon="mdi-plus-box" v-bind="props">
					New entity
					<template v-slot:prepend>
						<v-icon color="primary"></v-icon>
					</template>
				</v-btn>
			</template>
			<v-card class="px-6 py-6 rounded-xl" variant="elevated">
				<v-card-title>
					<span class="text-h6">New entity</span>
				</v-card-title>
				<v-card-text>
					<v-container>
						<v-row>
							<v-col cols="12" md="4">
								<v-text-field v-model="entity.name" label="Name"></v-text-field>
							</v-col>
							<v-col cols="12" md="4">
								<v-text-field v-model="entity.avatarURL" label="Avatar URL"></v-text-field>
							</v-col>
						</v-row>
					</v-container>
				</v-card-text>
				<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn variant="text" @click="close">
						Cancel
					</v-btn>
					<v-btn color="primary" variant="text" @click="save">
						Save
					</v-btn>
				</v-card-actions>
			</v-card>
		</v-dialog>
		<v-dialog v-model="dialogDelete" max-width="500px">
			<v-card>
				<v-card-title class="text-h5">Are you sure you want to delete this entity ?</v-card-title>
				<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn variant="text" @click="closeDelete">Cancel</v-btn>
					<v-btn color="danger" variant="text" @click="confirmDeleteEntity">OK</v-btn>
					<v-spacer></v-spacer>
				</v-card-actions>
			</v-card>
		</v-dialog>
	</v-toolbar>
	<v-row>
		<v-col class="mx-auto" cols="4">
			<v-list lines="three" bg-color="rgba(0, 0, 0, 0)">
				<v-list-item v-for="(entity, i) in entities" :key="ulid(entity.iD)" :prepend-avatar="entity.avatarURL"
					variant="plain" @click="() => { }">
					<v-list-item-title class="text-h6">{{ entity.name }}</v-list-item-title>
					<v-list-item-subtitle>{{ 'admin' }}</v-list-item-subtitle>
					<v-divider class="my-2"></v-divider>
				</v-list-item>
			</v-list>
		</v-col>
		<v-divider vertical></v-divider>
		<v-col class="mx-auto" cols="8">
			<v-card-item class="justify-center mb-2" prepend-icon="mdi-account-plus">
				<v-card-title>
					Entity
				</v-card-title>
				<template v-slot:prepend>
					<v-icon color="primary"></v-icon>
				</template>
			</v-card-item>
			<v-divider color="primary"></v-divider>
		</v-col>
	</v-row>
</template>
<style scoped>
.bar-bg {
	background: url('@/assets/img/bar-background.svg') no-repeat bottom center;
	background-size: cover;
}
</style>
