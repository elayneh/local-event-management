<template>
  <div
    :class="[
      'bg-white shadow-lg rounded-lg overflow-hidden flex flex-col',
      'transition-transform duration-300 ease-in-out transform hover:scale-105',
      customClass,
    ]"
    class="w-full max-w-md mx-auto p-4 hover:bg-indigo-300"
  >
    <img
      :src="computedImage"
      alt="Event Image"
      class="w-full h-48 object-cover rounded-lg"
    />

    <div class="text-black text-center py-4">
      <h2 class="text-xl font-semibold">{{ event.title }}</h2>
    </div>

    <div class="flex-1 p-4">
      <p class="text-gray-700 mb-2">
        <strong>Description:</strong> {{ truncatedDescription }}
      </p>
      <p class="text-gray-700 mb-2">
        <strong>Address:</strong> {{ event.address || "N/A" }}
      </p>
      <p class="text-gray-700 mb-2">
        <strong>Category ID:</strong> {{ event.category_id || "N/A" }}
      </p>
      <p class="text-gray-700 mb-2">
        <strong>Created At:</strong> {{ formatDate(event.created_at) || "N/A" }}
      </p>
      <p class="text-gray-700 mb-2">
        <strong>End Date:</strong> {{ formatDate(event.end_date) || "N/A" }}
      </p>
      <p class="text-gray-700 mb-2">
        <span :class="freeClass">{{ isFreeText }}</span>
      </p>
    </div>

    <div class="p-4 rounded-lg border-gray-200 flex justify-between bg-gray-50">
      <!-- Follow Button -->
      <button
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
        @click="followEvent"
      >
        Follow
      </button>

      <!-- Centered Bookmark Button -->
      <button
        class="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600 transition-colors duration-200 mx-4 flex-grow text-center"
        @click="bookmarkEvent"
      >
        +Bookmark
      </button>

      <!-- Purchase Button (Only shown if not free) -->
      <button
        v-if="!isFree"
        class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition-colors duration-200"
        @click="purchaseEvent"
      >
        Purchase
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import defaultImage from "~/assets/images/home.png";

const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
  customClass: {
    type: String,
    default: "",
  },
});

const computedImage = computed(() => props.event?.image || defaultImage);

const followEvent = () => {
  console.log("Follow event:", props.event?.id);
};

const purchaseEvent = () => {
  console.log("Purchase event:", props.event?.id);
};

const bookmarkEvent = () => {
  console.log("Bookmark event:", props.event?.id);
};

// Truncate description to a fixed length
const truncatedDescription = computed(() => {
  const maxLength = 100; // Set a fixed length for description
  return props.event?.description?.length > maxLength
    ? props.event.description.slice(0, maxLength) + "..."
    : props.event?.description || "N/A";
});

// Check if the event is free
const isFree = computed(() => props.event?.is_free);

// Define text for "Is Free"
const isFreeText = computed(() => (isFree.value ? "Free" : "Paid"));

// Define class for "Is Free" text
const freeClass = computed(() =>
  isFree.value
    ? "bg-gray-300 text-white px-2 py-1 rounded"
    : "bg-gray-800 text-white px-2 py-1 rounded"
);

const formatDate = (dateString) => {
  if (!dateString) return "N/A";
  const date = new Date(dateString);
  return date.toLocaleDateString(); // Customize date format as needed
};
</script>

<style scoped>
/* Add any additional scoped styles here */
</style>
