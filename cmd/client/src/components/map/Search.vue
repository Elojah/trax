<script setup lang="ts">
import { ref, watch, nextTick, onMounted, onUnmounted } from 'vue';
import { useGeocoding, type GeocodingResult } from '@/composables/useGeocoding';

// Props
interface Props {
	showResults?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
	showResults: true,
});

// Emits
const emit = defineEmits<{
	locationSelected: [result: GeocodingResult];
}>();

// State
const searchQuery = ref('');
const showDropdown = ref(false);
const searchInput = ref<HTMLInputElement>();
const searchContainer = ref<HTMLDivElement>();

// Geocoding composable
const { loading, error, results, geocode, getShortDisplayName, clear } = useGeocoding();

// Debounced search
let searchTimeout: NodeJS.Timeout;

const handleSearch = () => {
	clearTimeout(searchTimeout);

	if (!searchQuery.value.trim()) {
		showDropdown.value = false;
		clear();
		return;
	}

	searchTimeout = setTimeout(async () => {
		await geocode(searchQuery.value);
		if (results.value.length > 0 && props.showResults) {
			showDropdown.value = true;
		}
	}, 300); // 300ms debounce
};

const selectLocation = (result: GeocodingResult) => {
	searchQuery.value = getShortDisplayName(result);
	showDropdown.value = false;
	emit('locationSelected', result);
};

const handleKeydown = (event: KeyboardEvent) => {
	if (event.key === 'Escape') {
		showDropdown.value = false;
		searchInput.value?.blur();
	}
};

const handleClickOutside = (event: Event) => {
	if (searchContainer.value && !searchContainer.value.contains(event.target as Node)) {
		showDropdown.value = false;
	}
};

// Click outside handling
onMounted(() => {
	document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
	document.removeEventListener('click', handleClickOutside);
	clearTimeout(searchTimeout);
});

// Watch for input changes
watch(searchQuery, handleSearch);
</script>
<template>
	<div ref="searchContainer" class="absolute top-4 left-1/2 transform -translate-x-1/2 z-10 w-96 max-w-[60%]">
		<div
			class="border border-surface-200 dark:border-surface-700 rounded-xl flex items-center overflow-hidden bg-surface-0/80 dark:bg-surface-900/80 backdrop-blur-sm">
			<div class="flex-1 p-3 flex items-center gap-2 border-r border-surface-200 dark:border-surface-700">
				<i class="pi pi-search !text-base !leading-none text-surface-400"
					:class="{ 'pi-spin pi-spinner': loading }" />
				<input ref="searchInput" v-model="searchQuery"
					class="flex-1 bg-transparent w-full outline-none placeholder:text-surface-400 text-surface-950 dark:text-surface-0"
					placeholder="Search for places..." @keydown="handleKeydown" />
			</div>
		</div>

		<!-- Search Results Dropdown -->
		<div v-if="showDropdown && results.length > 0"
			class="absolute top-full left-0 right-0 mt-2 bg-surface-0 dark:bg-surface-900 border border-surface-200 dark:border-surface-700 rounded-xl shadow-lg overflow-hidden z-20">
			<div class="max-h-64 overflow-y-auto">
				<div v-for="result in results" :key="result.place_id"
					class="p-3 hover:bg-surface-50 dark:hover:bg-surface-800 cursor-pointer border-b border-surface-100 dark:border-surface-800 last:border-b-0"
					@click="selectLocation(result)">
					<div class="flex items-start gap-2">
						<i class="pi pi-map-marker text-surface-400 mt-1 text-sm" />
						<div class="flex-1 min-w-0">
							<div class="text-surface-950 dark:text-surface-0 font-medium truncate">
								{{ result.name || getShortDisplayName(result) }}
							</div>
							<div class="text-surface-500 text-sm truncate">
								{{ result.display_name }}
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<!-- Error Message -->
		<div v-if="error"
			class="absolute top-full left-0 right-0 mt-2 p-3 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl text-red-700 dark:text-red-300 text-sm">
			{{ error }}
		</div>
	</div>
</template>
<style scoped></style>
