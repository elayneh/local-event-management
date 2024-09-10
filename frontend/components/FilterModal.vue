<script setup>
import { ref, computed, onMounted } from "vue";
import debounce from "lodash/debounce";
import MapComponent from "@/components/LeafletMap.vue";

const events = ref([]);
const offset = ref(0);
const limit = ref(10);
const search = ref("");
const venue = ref("");
const start_date = ref("");
const end_date = ref("");
const isFree = ref(false);
const mapLocation = ref({ lat: null, lng: null });
const isModalVisible = ref(false);
const hasMoreEvents = ref(true);

const filters = computed(() => {
  const where = {};
  if (search.value) {
    where._or = [
      { title: { _ilike: `%${search.value}%` } },
      { description: { _ilike: `%${search.value}%` } },
    ];
  }
  if (venue.value) {
    where.venue = { _ilike: `%${venue.value}%` };
  }
  if (start_date.value) {
    where.start_date = { _gte: start_date.value };
  }
  if (end_date.value) {
    where.end_date = { _lte: end_date.value };
  }
  if (isFree.value !== null) {
    where.is_free = { _eq: isFree.value };
  }
  if (mapLocation.value.lat && mapLocation.value.lng) {
    where.location = {
      _ilike: `%${mapLocation.value.lat},${mapLocation.value.lng}%`,
    };
  }
  return where;
});

const fetchEvents = async () => {
  const response = await fetch("https://your-graphql-endpoint", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer your-token`,
    },
    body: JSON.stringify({
      query: `
          query GetEvents($where: events_bool_exp, $limit: Int, $offset: Int) {
            events(where: $where, limit: $limit, offset: $offset) {
              id
              title
              description
              start_date
              end_date
              venue
            }
          }
        `,
      variables: {
        where: filters.value,
        limit: limit.value,
        offset: offset.value,
      },
    }),
  });
  const { data } = await response.json();
  if (offset.value === 0) {
    events.value = data.events;
  } else {
    events.value = [...events.value, ...data.events];
  }
  hasMoreEvents.value = data.events.length === limit.value;
};

const debouncedFetchEvents = debounce(fetchEvents, 300);

const applyFilters = () => {
  offset.value = 0;
  fetchEvents();
  closeModal();
};

const resetFilters = () => {
  venue.value = "";
  start_date.value = "";
  end_date.value = "";
  isFree.value = false;
  mapLocation.value = { lat: null, lng: null };
  offset.value = 0;
  fetchEvents();
};

const loadMoreEvents = async () => {
  offset.value += limit.value;
  await fetchEvents();
};

const closeModal = () => {
  isModalVisible.value = false;
};

const openModal = () => {
  isModalVisible.value = true;
};

const handleLocationSelected = (location) => {
  mapLocation.value = location;
};

onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <div>
    <input
      v-model="search"
      @input="debouncedFetchEvents"
      type="text"
      placeholder="Search events"
      class="mb-4 p-2 border rounded w-full"
    />

    <button
      @click="openModal"
      class="mb-4 px-4 py-2 bg-blue-500 text-white rounded-full"
    >
      Open Filters
    </button>

    <div
      v-for="event in events"
      :key="event.id"
      class="mb-4 p-4 border rounded"
    >
      <h3>{{ event.title }}</h3>
      <p>{{ event.description }}</p>
    </div>

    <button
      @click="loadMoreEvents"
      v-if="hasMoreEvents"
      class="mt-4 px-4 py-2 bg-green-500 text-white rounded-full"
    >
      Load More
    </button>

    <div
      class="fixed inset-0 bg-gray-800 bg-opacity-75 z-50 overflow-y-auto"
      v-if="isModalVisible"
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
  </div>
</template>
