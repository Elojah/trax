<script setup lang="ts">
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer, LMarker } from "@vue-leaflet/vue-leaflet";
import { ref } from "vue";
import type { Position } from "@pkg/geometry/position";
import MarkerCreate from "@/components/map/MarkerCreate.vue";



let zoom = 3;
const cursor = ref<Position>();
const visible = ref(false);

const openMarkerPanel = (event: any) => {
	// Display temporary marker at clicked position
	const { lat, lng } = event.latlng;
	cursor.value = { lat, lng };

	// Open MarkerPanel with the clicked position
	visible.value = true;
};

</script>

<template>
	<div class="h-full w-full relative">
		<l-map ref="map" v-model:zoom="zoom" :center="[47.41322, -1.219482]" class="h-full w-full"
			@click="openMarkerPanel">
			<l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
				name="map"></l-tile-layer>

			<!-- Render cursor marker -->
			<l-marker v-if="cursor" :lat-lng="[cursor.lat, cursor.lng]">
			</l-marker>
		</l-map>

		<!-- MarkerPanel - Now outside the map, positioned absolutely -->
		<MarkerCreate :visible="visible" @close="visible = false" />
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
</style>
