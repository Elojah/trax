<script setup lang="ts">
import Button from 'primevue/button';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Menu from 'primevue/menu';
import MultiSelect from 'primevue/multiselect';
import Textarea from 'primevue/textarea';
import { ref, computed } from 'vue';

// Props
const props = withDefaults(defineProps<{
	visible: boolean;
}>(), {
	visible: false,
});

// Emits
const emit = defineEmits<{
	close: [];
}>();

// Computed class for visibility and animations
const panelClass = computed(() => ({
	'hidden': !props.visible,
	'animate-fadeinright': props.visible,
	'animate-fadeoutright': !props.visible && props.visible !== undefined,
}));

// Close panel function
const closePanel = () => {
	emit('close');
};

const tags = ref([
	{ name: 'Digital Painting', code: 'digital-painting' },
	{ name: 'Illustration', code: 'illustration' },
	{ name: 'Concept Art', code: 'concept-art' },
	{ name: 'Character Design', code: 'character-design' },
	{ name: 'Landscape', code: 'landscape' },
	{ name: 'Abstract', code: 'abstract' },
	{ name: 'Printable Poster', code: 'printable-poster' },
	{ name: 'Wall Art', code: 'wall-art' },
	{ name: 'Fantasy', code: 'fantasy' },
	{ name: 'Photography', code: 'photography' },
	{ name: '3D Model', code: '3d-model' }
]);

const selectedTags = ref([
	{ name: 'Printable Poster', code: 'printable-poster' },
	{ name: 'Wall Art', code: 'wall-art' }
]);

const uploadedFiles = ref([
	{
		id: 1,
		name: 'Celestial Sphere',
		size: '1.5 MB',
		image: 'https://fqjltiegiezfetthbags.supabase.co/storage/v1/render/image/public/block.images/blocks/slideover/upload-1.png'
	},
	{
		id: 2,
		name: 'Rainbow Flow',
		size: '17.3MB',
		image: 'https://fqjltiegiezfetthbags.supabase.co/storage/v1/render/image/public/block.images/blocks/slideover/upload-2.png'
	},
	{
		id: 3,
		name: 'Whimsical Carousel',
		size: '11.8MB',
		image: 'https://fqjltiegiezfetthbags.supabase.co/storage/v1/render/image/public/block.images/blocks/slideover/upload-3.png'
	},
	{
		id: 4,
		name: 'Folded Elegance',
		size: '3.5MB',
		image: 'https://fqjltiegiezfetthbags.supabase.co/storage/v1/render/image/public/block.images/blocks/slideover/upload-4.png'
	}
]);

const menuItems = ref([
	{
		label: 'Preview',
		icon: 'pi pi-eye'
	},
	{
		label: 'Download',
		icon: 'pi pi-download'
	},
	{
		label: 'Share',
		icon: 'pi pi-share-alt'
	},
	{
		label: 'Set as Featured',
		icon: 'pi pi-star'
	},
	{
		separator: true
	},
	{
		label: 'Edit',
		icon: 'pi pi-pencil'
	},
	{
		label: 'Delete',
		icon: 'pi pi-trash',
		class: 'text-red-500'
	}
]);

const menus = ref(Array(uploadedFiles.value.length).fill(null));

const toggleMenu = (event: any, index: number) => {
	if (menus.value[index]) {
		menus.value[index].toggle(event);
	}
};
</script>

<template>
	<div :class="panelClass"
		class="bg-surface-0 dark:bg-surface-900 absolute w-full sm:w-[36rem] h-full shadow-xl right-0 top-0 transition-all duration-300 z-50 overflow-hidden">
		<div class="flex flex-col h-full">
			<!-- Fixed Header -->
			<div class="px-6 py-4 border-b border-surface flex items-center justify-between flex-shrink-0">
				<div class="flex items-center gap-4">
					<div class="flex items-center gap-2">
						<svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" width="33" height="32"
							viewBox="0 0 33 32" fill="none">
							<path fill-rule="evenodd" clip-rule="evenodd"
								d="M7.09219 2.87829C5.94766 3.67858 4.9127 4.62478 4.01426 5.68992C7.6857 5.34906 12.3501 5.90564 17.7655 8.61335C23.5484 11.5047 28.205 11.6025 31.4458 10.9773C31.1517 10.087 30.7815 9.23135 30.343 8.41791C26.6332 8.80919 21.8772 8.29127 16.3345 5.51998C12.8148 3.76014 9.71221 3.03521 7.09219 2.87829ZM28.1759 5.33332C25.2462 2.06 20.9887 0 16.25 0C14.8584 0 13.5081 0.177686 12.2209 0.511584C13.9643 0.987269 15.8163 1.68319 17.7655 2.65781C21.8236 4.68682 25.3271 5.34013 28.1759 5.33332ZM32.1387 14.1025C28.2235 14.8756 22.817 14.7168 16.3345 11.4755C10.274 8.44527 5.45035 8.48343 2.19712 9.20639C2.0292 9.24367 1.86523 9.28287 1.70522 9.32367C1.2793 10.25 0.939308 11.2241 0.695362 12.2356C0.955909 12.166 1.22514 12.0998 1.50293 12.0381C5.44966 11.161 11.0261 11.1991 17.7655 14.5689C23.8261 17.5991 28.6497 17.561 31.9029 16.838C32.0144 16.8133 32.1242 16.7877 32.2322 16.7613C32.2441 16.509 32.25 16.2552 32.25 16C32.25 15.358 32.2122 14.7248 32.1387 14.1025ZM31.7098 20.1378C27.8326 20.8157 22.5836 20.5555 16.3345 17.431C10.274 14.4008 5.45035 14.439 2.19712 15.1619C1.475 15.3223 0.825392 15.5178 0.252344 15.7241C0.250782 15.8158 0.25 15.9078 0.25 16C0.25 24.8366 7.41344 32 16.25 32C23.6557 32 29.8862 26.9687 31.7098 20.1378Z"
								class="fill-surface-700 dark:fill-surface-200" />
						</svg>
						<svg xmlns="http://www.w3.org/2000/svg" width="10" height="3" viewBox="0 0 10 3" fill="none">
							<path
								d="M0.666667 1.5C0.666667 2.23638 1.26362 2.83333 2 2.83333C2.73638 2.83333 3.33333 2.23638 3.33333 1.5C3.33333 0.76362 2.73638 0.166667 2 0.166667C1.26362 0.166667 0.666667 0.76362 0.666667 1.5ZM6.66667 1.5C6.66667 2.23638 7.26362 2.83333 8 2.83333C8.73638 2.83333 9.33333 2.23638 9.33333 1.5C9.33333 0.76362 8.73638 0.166667 8 0.166667C7.26362 0.166667 6.66667 0.76362 6.66667 1.5ZM2 1.75H8V1.25H2V1.75Z"
								class="fill-surface-700 dark:fill-surface-400" />
						</svg>
						<img src="https://fqjltiegiezfetthbags.supabase.co/storage/v1/object/public/block.images/blocks/slideover/forms-avatar.png"
							alt="Profile Avatar" class="w-8 h-8 rounded-lg" />
					</div>
					<span class="text-surface-900 dark:text-surface-0 text-base font-medium">Robert Fox</span>
				</div>

				<Button @click="closePanel" icon="pi pi-times" rounded text severity="secondary" />
			</div>

			<!-- Scrollable Content -->
			<div class="flex-1 overflow-y-auto">
				<div class="flex flex-col">

					<div class="flex flex-col">
						<div class="px-6 py-3 bg-surface-50 dark:bg-surface-800">
							<span class="text-surface-900 dark:text-surface-0 font-medium">General</span>
						</div>
						<div class="p-6 flex flex-col gap-6">
							<InputText placeholder="Name" />
							<InputText placeholder="Price" />
							<Textarea placeholder="Description" rows="5" class="w-full" />
						</div>
					</div>

					<div class="flex flex-col">
						<div class="px-6 py-3 bg-surface-50 dark:bg-surface-800">
							<span class="text-surface-900 dark:text-surface-0 font-medium">Meta</span>
						</div>
						<div class="p-6 flex flex-col gap-6">
							<MultiSelect v-model="selectedTags" :options="tags" option-label="name" display="chip"
								class="w-full" placeholder="Select Tags" />

							<InputGroup class="w-full">
								<InputGroupAddon>
									<span class="text-surface-500 dark:text-surface-400">robert/</span>
								</InputGroupAddon>
								<InputText class="w-full" />
							</InputGroup>
						</div>
					</div>

					<div class="flex flex-col">
						<div class="px-6 py-3 bg-surface-50 dark:bg-surface-800">
							<span class="text-surface-900 dark:text-surface-0 font-medium">Media</span>
						</div>
						<div class="p-6 flex flex-col gap-4">
							<div
								class="w-full border border-dashed border-surface-200 dark:border-surface-700 rounded-lg flex flex-col gap-4 py-6 px-6 items-center justify-center">
								<div class="text-surface-900 dark:text-surface-0 text-center font-medium">Drop your
									images
									here or click to browse.</div>
								<div class="text-surface-500 dark:text-surface-400 text-center">1600 × 1200 (4:3)
									recommended, up to 20MB each.</div>
							</div>
							<span class="text-surface-500 dark:text-surface-400 text-sm">Add up to 5 images to your
								product.
								Used to represent your product during checkout, in email, social sharing and
								more.</span>
						</div>
					</div>

					<div class="flex flex-col flex-1">
						<div class="px-6 py-3 bg-surface-50 dark:bg-surface-800">
							<span class="text-surface-900 dark:text-surface-0 font-medium">Uploads</span>
						</div>
						<div class="p-6 flex flex-col gap-6">
							<ul class="flex flex-col gap-6">
								<li v-for="(file, index) in uploadedFiles" :key="file.id"
									class="flex justify-between items-center gap-4">
									<div class="flex flex-row gap-4">
										<div class="flex flex-col gap-2">
											<span class="text-surface-900 dark:text-surface-0 font-medium">{{
												file.name
											}}</span>
											<span class="text-surface-500 dark:text-surface-400">{{ file.size }}</span>
										</div>
									</div>

									<Button icon="pi pi-ellipsis-h !text-xl !leading-none" text rounded
										severity="secondary" @click="toggleMenu($event, index)" />
									<Menu :ref="(el) => (menus[index] = el)" :model="menuItems" :popup="true" />
								</li>
							</ul>
						</div>
					</div>
					<div class="flex flex-col gap-4">
						<Button @click="closePanel" outlined class="flex-1 py-2 px-3" label="Cancel" />
						<Button @click="closePanel" class="flex-1 py-2 px-3" label="Save" />
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped></style>
