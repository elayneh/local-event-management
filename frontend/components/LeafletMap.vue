<template>
  <div ref="mapContainer" class="leaflet-map" />
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import L from "leaflet";
import "leaflet/dist/leaflet.css";
import axios from "axios";

// Props
const props = defineProps({
  zoom: {
    type: Number,
    default: 13,
  },
  center: {
    type: Array,
    default: () => [9.0306, 38.7506], // Center set to Addis Ababa, Piassa
  },
  onClick: {
    type: Function,
    default: () => {},
  },
});

const mapContainer = ref(null);
let map = null;

// Initialize map
const initializeMap = async () => {
  if (mapContainer.value) {
    map = L.map(mapContainer.value).setView(props.center, props.zoom);

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(map);

    try {
      map.on("click", async (event) => {
        const response = await axios.get(
          "https://nominatim.openstreetmap.org/reverse",
          {
            params: {
              lat: event.latlng.lat,
              lon: event.latlng.lng,
              format: "json",
            },
          }
        );

        // locationName.value = response.data.display_name;
        const { lat, lng } = event.latlng;
        props.onClick({ lat, lng, address: response.data.display_name });
      });
    } catch (error) {
      console.error("Failed to fetch place name: ", error);
    }
  }
};

// Watch for changes in props
watch(
  () => props.center,
  (newCenter) => {
    if (map) {
      map.setView(newCenter, map.getZoom());
    }
  }
);

onMounted(() => {
  initializeMap();
});
</script>

<style scoped>
.leaflet-map {
  height: 400px;
  width: 100%;
}
</style>
