<script setup lang="ts">
import { useEntityStore } from '@/stores/entity';
import { computed, ref, toRefs } from 'vue';
import { marked } from "marked";
import { useAuthStore } from '@/stores/auth';
import { DeleteEntityReq, UpdateEntityReq } from '@internal/user/dto/entity';

const props = defineProps<{
	onDelete: () => void;
}>();

const authStore = useAuthStore();
const store = useEntityStore();
const {
	selected: selected,
} = toRefs(store);

const e = computed(() => selected.value.at(0));

const mdDescription = computed(() => {
	return marked.parse(e.value?.description || `*no description*`);
});

const dialogDelete = ref(false);
const closeDelete = () => {
	dialogDelete.value = false;
};

const confirmDelete = async () => {
	await store.delete_(DeleteEntityReq.create({ iD: e.value?.iD }));

	selected.value = [];
	await authStore.refreshToken();
	props.onDelete();

	dialogDelete.value = false;
};

const name = ref(false);
const description = ref(false);

const updateName = async function () {
	if (name.value) {
		await store.update(UpdateEntityReq.create({
			iD: e.value?.iD,
			name: { value: e.value?.name },
		}));
	}

	name.value = !name.value
};

const updateDescription = async function () {
	if (description.value) {
		await store.update(UpdateEntityReq.create({
			iD: e.value?.iD,
			description: { value: e.value?.description },
		}));
	}

	description.value = !description.value
};
</script>

<template>
	<!-- Info -->
	<v-card v-if="e" class="px-6 py-6 justify-center rounded" variant="flat">
		<v-avatar class="mb-4 justify-center" size="96" :color="!e.avatarURL ? 'primary' : ''">
			<img v-if="e.avatarURL" :src="e.avatarURL" alt="Avatar">
			<span v-else-if="!e.avatarURL" class=" mx-auto text-center text-h5">
				{{ e.name?.at(0)?.toUpperCase() }}
			</span>
		</v-avatar>
		<v-text-field class="justify-center text-h3" v-model="e.name" :variant="!name ? 'plain' : 'underlined'"
			:readonly="!name">
			<template v-slot:prepend-inner>
				<v-icon color="primary" size="large" @click="updateName"
					:icon="!name ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
			</template>
		</v-text-field>
		<!-- eslint-disable vue/no-v-html vue/no-v-text-v-html-on-component -->
		<v-textarea v-if="description" class="p-6 m-6" variant="solo" v-model="e.description" :readonly="false">
			<template v-slot:prepend-inner>
				<v-icon color="primary" size="large" @click="updateDescription"
					:icon="!description ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
			</template>
		</v-textarea>
		<v-text-field v-if="!description" class="p-6 m-6 font-italic" variant="solo" no-resize :readonly="true">
			Edit description
			<template v-slot:prepend-inner>
				<v-icon color="primary" size="large" @click="updateDescription"
					:icon="!description ? 'mdi-pencil-circle-outline' : 'mdi-arrow-right-bold-circle-outline'"></v-icon>
			</template>
		</v-text-field>
		<div class="px-4" v-if="!description" v-html="mdDescription"></div>
		<!--eslint-enable-->
		<v-card-actions>
			<v-divider></v-divider>
			<v-dialog v-model="dialogDelete" max-width="800px">
				<template v-slot:activator="{ props }">
					<v-btn variant="tonal" prepend-icon="mdi-trash-can" color="error" v-bind="props">
						Delete
						<template v-slot:prepend>
							<v-icon color="error"></v-icon>
						</template>
					</v-btn>
				</template>
				<v-card class="px-6 py-6 rounded" variant="elevated">
					<v-card-title class="text-h5">Are you sure you want to delete this entity
						?</v-card-title>
					<v-card-actions>
						<v-spacer></v-spacer>
						<v-btn variant="text" @click="closeDelete">Cancel</v-btn>
						<v-btn color="error" variant="text" @click="confirmDelete">OK</v-btn>
						<v-spacer></v-spacer>
					</v-card-actions>
				</v-card>
			</v-dialog>
		</v-card-actions>
	</v-card>
</template>
<style scoped></style>
