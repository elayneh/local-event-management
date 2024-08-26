<template>
  <div class="container mx-auto p-6 bg-gray-100">
    <h1 class="text-3xl font-semibold mb-4">Location Form</h1>

    <form @submit.prevent="handleSubmit" class="bg-white p-6 rounded-lg shadow-md">
      <div class="mb-4">
        <LeafletMap
          :zoom="zoom"
          :center="center"
          @click="handleMapClick"
        />
        <LMarker v-if="marker" :lat-lng="marker" />
      </div>

      <div class="mb-4">
        <label for="location" class="block text-gray-700 text-sm font-bold mb-2">
          Location Address:
        </label>
        <input
          v-model="address"
          type="text"
          id="location"
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          placeholder="Address will be auto-filled"
          readonly
        />
      </div>

      <button
        type="submit"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline"
      >
        Submit
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import LeafletMap from '~/components/LeafletMap.vue';

// Data and refs
const zoom = 13;
const center = ref([9.0306, 38.7506]); // Center set to Addis Ababa, Piassa
const marker = ref(null);
const address = ref('');

// Handle map click event
const handleMapClick = async (coords) => {
  marker.value = [coords.lat, coords.lng];
  
  // Use a geocoding service to convert lat/lng to address
  const geocodeUrl = `https://nominatim.openstreetmap.org/reverse?lat=${coords.lat}&lon=${coords.lng}&format=json`;
  
  try {
    const response = await fetch(geocodeUrl);
    const data = await response.json();
    address.value = data.display_name || 'Address not found';
  } catch (error) {
    console.error('Error fetching address:', error);
    address.value = 'Error retrieving address';
  }
};

// Handle form submission
const handleSubmit = () => {
  if (marker.value) {
    const locationData = {
      address: address.value,
      latitude: marker.value[0],
      longitude: marker.value[1],
    };

    console.log('Form submitted with:', locationData);
    // Handle form submission (e.g., send data to the server)
  } else {
    console.error('Please select a location on the map.');
  }
};
</script>

<style scoped>
/* Add any additional styling here */
</style>
