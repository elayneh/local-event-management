<script setup>
import { ref, computed, onMounted } from "vue";
import CustomEventCard from "@/components/CustomEventCard.vue";
import TagsCategoriesComponent from "@/components/TagsCategoriesComponent.vue";
import FeaturesComponent from "@/components/FeaturesComponent.vue";
import CustomFooter from "@/components/CustomFooter.vue";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import { useAuthStore } from "~/stores";
import useFetchData from "~/composables/useFetchData";

const user_id = useAuthStore().id;

const searchQuery = ref("");
const isModalVisible = ref(false);
const events = ref([]);
const limit = ref(10);
const offset = ref(0);
const hasMoreEvents = ref(true);

const filters = ref({
  isFree: null,
  start_date: "",
  end_date: "",
  venue: "",
  tags: [],
});

const { categories, tags } = useFetchData();

const openFilterModalHandler = () => {
  isModalVisible.value = true;
};

const filter = computed(() => {
  const query = {};

  if (searchQuery.value.trim()) {
    query._or = [
      { title: { _ilike: `%${searchQuery.value}%` } },
      { description: { _ilike: `%${searchQuery.value}%` } },
      { venue: { _ilike: `%${searchQuery.value}%` } },
    ];
  }

  if (filters.value.isFree !== null) {
    query.is_free = { _eq: filters.value.isFree };
  }
  if (filters.value.start_date) {
    query.start_date = { _gte: filters.value.start_date };
  }
  if (filters.value.end_date) {
    query.end_date = { _lte: filters.value.end_date };
  }

  if (Array.isArray(filters.value.tags) && filters.value.tags.length > 0) {
    query.tags = {
      tag: {
        _in: filters.value.tags,
      },
    };
  }

  return query;
});

const fetchEvents = async () => {
  try {
    await refetch();
  } catch (error) {
    console.error("Failed to fetch events:", error);
  }
};

const { onResult, refetch } = listQuery(getEvents, {
  filter,
  limit,
  offset,
});

onResult(({ data }) => {
  if (data?.events) {
    if (offset.value === 0) {
      events.value = data.events;
    } else {
      events.value.push(...data.events);
    }
    hasMoreEvents.value = data.events.length === limit.value;
  }
});

const loadMoreEvents = () => {
  offset.value += limit.value;
  fetchEvents();
};

const applyFilters = () => {
  isModalVisible.value = false;
  offset.value = 0;
  fetchEvents();
};

const resetFilters = () => {
  filters.value = {
    isFree: null,
    start_date: "",
    end_date: "",
    venue: "",
    tags: [],
  };
  applyFilters();
};

definePageMeta({ layout: "authenticated" });
</script>

<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 w-full pl-12 pr-12 moving-clouds-bg"
  >
    <div class="fixed left-96 flex-row justify-center items-center pt-2">
      <div v-if="isModalVisible">
        <div
          class="fixed inset-0 bg-gray-800 bg-opacity-75 z-50 overflow-y-auto"
        >
          <div
            class="relative max-w-xl mx-auto mt-12 p-8 bg-white shadow-lg rounded-lg h-[50vh] overflow-y-auto"
          >
            <button
              @click="closeModal"
              class="absolute top-4 right-4 p-2 bg-gray-200 hover:bg-red-100 rounded-full shadow-md transition-transform transform hover:scale-110 hover:shadow-lg text-gray-600 hover:text-gray-800"
            >
              <font-awesome-icon :icon="['fas', 'times']" class="text-lg" />
            </button>

            <h2 class="text-xl font-semibold mb-6 text-center">
              Filter Events
            </h2>

            <div class="mb-6">
              <label class="block text-sm font-medium mb-2">Payment</label>
              <div class="flex space-x-6">
                <label class="inline-flex items-center">
                  <input
                    type="radio"
                    value="true"
                    v-model="filters.isFree"
                    class="form-radio text-blue-500"
                  />
                  <span class="ml-2 text-gray-700">Free</span>
                </label>
                <label class="inline-flex items-center">
                  <input
                    type="radio"
                    value="false"
                    v-model="filters.isFree"
                    class="form-radio text-blue-500"
                  />
                  <span class="ml-2 text-gray-700">Paid</span>
                </label>
              </div>
            </div>

            <div class="mb-6">
              <label class="block text-sm font-medium mb-2">Start Date</label>
              <input
                v-model="filters.start_date"
                type="datetime-local"
                class="w-full p-3 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div class="mb-6">
              <label class="block text-sm font-medium mb-2">End Date</label>
              <input
                v-model="filters.end_date"
                type="datetime-local"
                class="w-full p-3 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div class="mb-6">
              <label class="block text-sm font-medium mb-2">Venue</label>
              <div class="relative">
                <input
                  v-model="filters.venue"
                  type="text"
                  placeholder="Enter venue"
                  class="w-full p-3 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
                <div class="mt-2">
                  <MapComponent @location-selected="handleLocationSelected" />
                </div>
              </div>
            </div>
            <div>
              <label
                v-for="tag in availableTags"
                :key="tag.id"
                class="inline-flex items-center mt-3"
              >
                <input
                  type="checkbox"
                  :value="tag.id"
                  v-model="filters.value.tags"
                  @change="fetchEvents"
                />
                <span class="ml-2 text-gray-700">{{ tag.name }}</span>
              </label>
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
      <div
        class="flex justify-center items-center gap-16 p-1 bg-gray-100 rounded-lg shadow-md"
      >
        <ul class="mx-2">
          <li class="relative flex items-center">
            <input
              class="w-96 h-12 rounded-full pl-6 pr-12 bg-gray-200 text-gray-800 placeholder-gray-500 border border-gray-300 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-300 ease-in-out"
              v-model="searchQuery"
              @input="fetchEvents"
              placeholder="Search..."
            />
            <font-awesome-icon
              :icon="['fas', 'search']"
              class="text-gray-500 absolute right-4"
            />
          </li>
        </ul>

        <ul class="m-2">
          <li class="relative">
            <button
              @click="openFilterModalHandler"
              class="flex items-center bg-blue-500 text-white px-6 py-3 rounded-lg shadow-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 transition-all duration-300 ease-in-out"
            >
              <font-awesome-icon :icon="['fas', 'filter']" class="text-white" />
              <span class="ml-4 font-semibold">Filter</span>
            </button>
          </li>
        </ul>
      </div>
    </div>

    <div class="flex items-center justify-center pt-24 mb-8">
      <div
        class="flex flex-col md:flex-row w-full h-auto md:w-3/4 items-center justify-center gap-8 px-4 transition-transform duration-1000 ease-in-out"
      >
        <div
          class="w-full md:w-1/2 flex items-center justify-center side-to-side-animation"
        >
          <img
            src="~/assets/images/eventimage1.png"
            alt="Home page image"
            class="w-full max-w-lg h-64 mt-12 md:mt-0 rounded-lg shadow-lg hover:scale-110 hover:rotate-6 hover:shadow-2xl hover:brightness-110 hover:contrast-125 transition-all duration-700 ease-in-out animate-pulse"
          />
        </div>

        <div
          class="w-full md:w-1/2 flex flex-col justify-center gap-6 text-center"
        >
          <h2
            class="text-2xl font-bold mb-2 transition-transform hover:scale-105 duration-1000 ease-in-out"
          >
            Welcome to Our Event
          </h2>
          <p
            class="text-white leading-relaxed text-lg transition-opacity duration-1000 hover:opacity-90"
          >
            Discover the most exciting and memorable events in your area.
            Whether you're looking for music, arts, or fun activities, we have
            it all!
          </p>
          <p
            class="text-white leading-relaxed text-lg transition-opacity duration-1000 hover:opacity-90"
          >
            Get ready for an experience like no other, with engaging talks,
            workshops, and live performances to spark your creativity.
          </p>
          <a
            href="#"
            class="align-right bg-pink-500 px-4 py-2 rounded-lg shadow-md hover:bg-blue-700 transition-colors duration-500 ease-in-out"
          >
            Learn More
          </a>
        </div>
      </div>
    </div>

    <h2
      class="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 text-center mb-12 animate-fadeInBounce"
    >
      Latest Events
    </h2>
    `
    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <CustomEventCard v-for="event in events" :key="event.id" :event="event" />
    </div>

    <div class="text-center mt-6">
      <button
        @click="loadMoreEvents"
        v-if="hasMoreEvents"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
      >
        Load More
      </button>
      <button
        v-else
        class="bg-blue-500 text-white px-4 py-2 rounded opacity-50 cursor-not-allowed"
      >
        No more events
      </button>
    </div>

    <TagsCategoriesComponent :categories="categories" :tags="tags" />
    <FeaturesComponent />
    <CustomFooter />
  </div>
</template>

<style scoped>
.side-to-side-animation {
  animation: sideToSideAnimation 5s ease-in-out infinite;
}

.moving-clouds-bg {
  background-image: url("~/assets/images/clouds2.png");
  background-size: cover;
  background-position: 0 0;
  background-repeat: repeat-x;
  animation: moveClouds 30s linear infinite;
}

@keyframes moveClouds {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 100% 0;
  }
}
</style>
