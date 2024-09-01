<template>
  <div
    class="fixed inset-0 bg-gray-800 bg-opacity-75 z-50 overflow-y-auto"
    v-if="isVisible"
  >
    <div
      class="relative max-w-3xl mx-auto mt-12 p-8 bg-white shadow-lg rounded-lg overflow-y-auto h-[calc(100vh-3rem)]"
    >
      <button
        @click="closeModal"
        class="absolute top-4 right-4 p-2 bg-gray-200 hover:bg-red-100 rounded-full shadow-md transition-transform transform hover:scale-110 hover:shadow-lg text-gray-600 hover:text-gray-800"
      >
        <font-awesome-icon :icon="['fas', 'times']" class="text-lg" />
      </button>

      <h2 class="text-xl font-semibold mb-6 text-center">Filter Events</h2>

      <div class="mb-6">
        <label class="block text-sm font-medium mb-2">Payment</label>
        <div class="flex space-x-6">
          <label class="inline-flex items-center">
            <input
              type="radio"
              value="true"
              v-model="isFree"
              class="form-radio text-blue-500"
            />
            <span class="ml-2 text-gray-700">Free</span>
          </label>
          <label class="inline-flex items-center">
            <input
              type="radio"
              value="false"
              v-model="isFree"
              class="form-radio text-blue-500"
            />
            <span class="ml-2 text-gray-700">Paid</span>
          </label>
        </div>
      </div>

      <div class="mb-6">
        <label class="block text-sm font-medium mb-2">Start Date</label>
        <input
          v-model="start_date"
          type="datetime-local"
          class="w-full p-3 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>
      <div class="mb-6">
        <label class="block text-sm font-medium mb-2">End Date</label>
        <input
          v-model="end_date"
          type="datetime-local"
          class="w-full p-3 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>

      <div class="mb-6">
        <label class="block text-sm font-medium mb-2">Venue</label>
        <div class="relative">
          <input
            v-model="venue"
            type="text"
            placeholder="Enter venue"
            class="w-full p-3 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <div class="mt-2">
            <MapComponent @location-selected="handleLocationSelected" />
          </div>
        </div>
      </div>

      <div class="flex justify-between items-center space-x-4">
        <button
          @click="applyFilters"
          class="bg-green-500 text-white px-6 py-3 rounded-full shadow-lg hover:bg-green-700 transition-colors duration-200"
        >
          Filter
        </button>
        <button
          @click="resetFilters"
          class="bg-red-500 text-white px-6 py-3 rounded-full shadow-lg hover:bg-red-600 transition-colors duration-200"
        >
          Reset
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";
import MapComponent from "@/components/LeafletMap.vue";

const venue = ref("");
const start_date = ref("");
const end_date = ref("");
const isFree = ref(false);
const mapLocation = ref({ lat: null, lng: null });

const selectedTagValues = ref([]);
const selectedCategoryValue = ref(null);

const props = defineProps({
  isVisible: Boolean,
});

const emit = defineEmits(["item-selected", "close-modal"]);

const handleLocationSelected = (event) => {
  mapLocation.value = event.detail.location;
};

const emitFilterData = () => {
  const whereClause = generateWhereClause();
  const event = new CustomEvent("apply-filters", { detail: whereClause });
  window.dispatchEvent(event);
};

const generateWhereClause = () => {
  const where = {};
  if (venue.value) {
    where.venue = { _ilike: `%${venue.value}%` };
  }
  if (start_date.value || end_date.value) {
    where.start_date = {};
    if (start_date.value) {
      where.start_date._gte = start_date.value;
    }
    where.end_date = {};
    if (end_date.value) {
      where.end_date._lte = end_date.value;
    }
  }
  if (selectedCategoryValue.value) {
    where.categories = { _eq: selectedCategoryValue.value };
  }
  if (selectedTagValues.value.length > 0) {
    where.tags = { _in: `{${selectedTagValues.value}}` };
  }
  if (isFree.value) {
    where.is_free = { _eq: isFree.value };
  }
  if (mapLocation.value.lat && mapLocation.value.lng) {
    where.location = {
      _ilike: `%${mapLocation.value.lat},${mapLocation.value.lng}%`,
    };
  }

  return where;
};

const applyFilters = () => {
  emitFilterData();
  closeModal();
};

const resetFilters = () => {
  venue.value = "";
  start_date.value = "";
  end_date.value = "";
  selectedCategoryValue.value = "";
  selectedTagValues.value = [];
  isFree.value = false;
  mapLocation.value = { lat: null, lng: null };
  emitFilterData();
};

function closeModal() {
  emit("close-modal");
}

onMounted(() => {
  window.addEventListener("location-selected", handleLocationSelected);
});

onUnmounted(() => {
  window.removeEventListener("location-selected", handleLocationSelected);
});
</script>
