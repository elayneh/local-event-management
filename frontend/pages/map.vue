<template>
  <div id="map" class="h-screen w-full"></div>
</template>

<script setup>
import { onMounted } from 'vue';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css'; // Import Leaflet CSS

const latitude = 9.027593829423443; // Replace with your latitude
const longitude = 38.75139713287354; // Replace with your longitude

onMounted(() => {
  // Initialize the map
  const map = L.map('map').setView([latitude, longitude], 13);

  // Add tile layer
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  }).addTo(map);

  // Add marker
  L.marker([latitude, longitude])
    .addTo(map)
    .bindPopup('Event Location')
    .openPopup();
});
</script>

<style scoped>
#map {
  height: 100%; /* Ensure full viewport height */
  width: 100%;  /* Ensure full viewport width */
}
</style>
