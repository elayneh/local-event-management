<script setup>
import { onMounted } from "vue";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

const latitude = 9.027593829423443;
const longitude = 38.75139713287354;

onMounted(() => {
  const map = L.map("map").setView([latitude, longitude], 13);

  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    maxZoom: 19,
    attribution:
      '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
  }).addTo(map);

  L.marker([latitude, longitude])
    .addTo(map)
    .bindPopup("Event Location")
    .openPopup();
});
</script>

<style scoped>
#map {
  height: 100%;
  width: 100%;
}
</style>

<template>
  <div id="map" class="h-screen w-full"></div>
</template>
