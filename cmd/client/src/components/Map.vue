<script setup lang="ts">
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer } from "@vue-leaflet/vue-leaflet";

let zoom = 3;
</script>

<template>
	<v-row class="fill-height">
		<v-col cols="12 pa-0 ma-0">
			<l-map ref="map" v-model:zoom="zoom" :center="[47.41322, -1.219482]">
				<l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
					name="OpenStreetMap"></l-tile-layer>
			</l-map>
		</v-col>
	</v-row>
</template>
<style>
.bg {
	background: url('@/assets/img/main-background.svg') !important;
}

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
