<script setup lang="ts">
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer, LMarker, LIcon } from "@vue-leaflet/vue-leaflet";
import { ref, toRefs, computed } from "vue";
import type { Position } from "@pkg/geometry/position";
import MarkerCreate from "@/components/map/MarkerCreate.vue";
import { useReverseGeocoding } from "@/composables/useReverseGeocoding";
import { useErrorsStore } from "@/stores/errors";
import catpaw from '@/assets/cat-paw.svg?url';

const errorsStore = useErrorsStore();
const { error, message } = toRefs(errorsStore);

const zoom = ref(3); // Default zoom level
const cursor = ref<Position>();
const visible = ref(false);
const center = ref([47.41322, -1.219482]);
const address = ref<string>('');

// Reverse geocoding composable
const { reverseGeocode, formatAddress, loading: geocodingLoading, error: geocodingError } = useReverseGeocoding();

// Define max bounds to prevent showing dark areas at extreme latitudes
const maxBounds = ref([
	[-85, -180], // Southwest coordinates
	[85, 180]    // Northeast coordinates
]);

const openMarkerPanel = async (event: any) => {
	// Display temporary marker at clicked position
	const { lat, lng } = event.latlng;
	// Convert to the Position format used by the backend (assuming micro-degrees)
	cursor.value = {
		lat: lat,
		lng: lng,
	};

	// Center map on clicked position (bounds will be enforced by maxBounds)
	center.value = [lat, lng];

	// Open MarkerPanel with the clicked position
	visible.value = true;

	// Perform reverse geocoding using the regular lat/lng numbers
	try {
		const geocodingResult = await reverseGeocode(lat, lng);
		if (geocodingResult) {
			address.value = formatAddress(geocodingResult);
		} else {
			address.value = `${lat.toFixed(6)}, ${lng.toFixed(6)}`;
		}
	} catch (err) {
		// message.value = 'Failed to reverse geocode location: ' + err;
		// error.value = true;

		address.value = `${lat.toFixed(6)}, ${lng.toFixed(6)}`;
	}
};

const zoomUpdate = (z: number) => {
	zoom.value = z;
};

</script>

<template>
	<div class="h-full w-full relative">
		<l-map ref="map" :zoom="zoom" :min-zoom="3" :center="center" :max-bounds="maxBounds" class="h-full w-full"
			@click="openMarkerPanel" v-on:update:zoom="zoomUpdate">
			<l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
				name="map"></l-tile-layer>

			<!-- Render cursor marker -->
			<l-marker v-if="cursor" :lat-lng="[cursor.lat, cursor.lng]">
				<l-icon :icon-url="decodeURI(catpaw)" :icon-size="[42, 42]" :icon-anchor="[42 / 2, 42]" />
			</l-marker>
		</l-map>

		<!-- MarkerPanel - Now outside the map, positioned absolutely -->
		<MarkerCreate :visible="visible" :address="address" :coordinates="cursor" :geocoding-loading="geocodingLoading"
			:geocoding-error="geocodingError" @close="visible = false" />
	</div>
</template>
<style>
.leaflet-container {
	background-color: rgba(255, 0, 0, 0.0);
	height: 100%;
	width: 100%;
	/* Fallback minimum height */
	min-height: 400px;
}

.leaflet-layer,
.leaflet-control-zoom-in,
.leaflet-control-zoom-out,
.leaflet-control-attribution {
	filter: invert(100%) hue-rotate(180deg) brightness(95%) contrast(90%);
}

.leaflet-control {
	z-index: 0 !important
}

.leaflet-pane {
	z-index: 0 !important
}

.leaflet-top,
.leaflet-bottom {
	z-index: 0 !important
}
</style>
