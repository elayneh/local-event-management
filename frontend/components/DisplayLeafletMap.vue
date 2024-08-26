<template>
  <div class="map-container">
    <div id="map" class="map"></div>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from 'vue';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css'; // Import Leaflet CSS

const props = defineProps({
  latitude: {
    type: Number,
    required: true,
  },
  longitude: {
    type: Number,
    required: true,
  },
  address: {
    type: String,
    required: false,
  },
});

const map = ref(null);

onMounted(() => {
  if (!map.value) {
    // Initialize the map after the component is mounted
    map.value = L.map('map').setView([props.latitude, props.longitude], 13);

    // Set up the map tiles
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      maxZoom: 19,
    }).addTo(map.value);

    // Add the marker
    L.marker([props.latitude, props.longitude])
      .addTo(map.value)
      .bindPopup(props.address || 'Event Location')
      .openPopup();
  }
});

// Watch for prop changes if the latitude/longitude can dynamically change
watch([() => props.latitude, () => props.longitude], ([newLat, newLng]) => {
  if (map.value) {
    map.value.setView([newLat, newLng], 13);
    L.marker([newLat, newLng])
      .addTo(map.value)
      .bindPopup(props.address || 'Event Location')
      .openPopup();
  }
});
</script>

<style>
/* Non-scoped styles */
.map-container {
  position: relative;
  height: 100vh; /* Full viewport height for better visibility */
  width: 100%;
  border-radius: 0.5rem; /* Rounded corners */
  overflow: hidden; /* Hide overflow */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); /* Shadow effect */
}

.map {
  height: 100%;
  width: 100%;
}
</style>
