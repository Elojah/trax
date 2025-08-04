<script setup lang="ts">
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer } from "@vue-leaflet/vue-leaflet";

let zoom = 3;
</script>

<template>
	<div class="h-full w-full">
		<l-map ref="map" v-model:zoom="zoom" :center="[47.41322, -1.219482]" class="h-full w-full">
			<l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
				name="map"></l-tile-layer>
		</l-map>
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
