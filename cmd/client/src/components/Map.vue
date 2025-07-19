<script setup lang="ts">
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer } from "@vue-leaflet/vue-leaflet";

let zoom = 3;
</script>

<template>
	<div class="overflow-auto flex-1 flex flex-col gap-8 transition-all pb-20">
		<l-map ref="map" v-model:zoom="zoom" :center="[47.41322, -1.219482]">
			<l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
				name="MAP"></l-tile-layer>
		</l-map>
	</div>
</template>
<style>
.leaflet-container {
	background-color: rgba(255, 0, 0, 0.0);
}

.leaflet-layer,
.leaflet-control-zoom-in,
.leaflet-control-zoom-out,
.leaflet-control-attribution {
	filter: invert(100%) hue-rotate(180deg) brightness(95%) contrast(90%);
}
</style>
