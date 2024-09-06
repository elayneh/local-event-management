<script setup>
import CustomEventCard from "@/components/CustomEventCard.vue";
import useFetchData from "~/composables/useFetchData";
import defaultImage from "~/assets/images/home.png";
import TagsCategoriesComponent from "@/components/TagsCategoriesComponent.vue";
import FeaturesComponent from "@/components/FeaturesComponent.vue";
import CustomFooter from "@/components/CustomFooter.vue";
import { ref, computed } from "vue";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import { useAuthStore } from "~/stores";
const events = ref([]);

const { categories, tags } = useFetchData();

const event_images = computed(() => {
  if (events?.value?.event_images) {
    const imagesArray = events.value.event_images
      .replace(/{|}/g, "")
      .split(",");
    return imagesArray.length > 0 ? imagesArray : [defaultImage];
  }
  return [defaultImage];
});
const user_id = useAuthStore().id;

const searchQuery = ref("");
const isModalVisible = ref(false);
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
</script>

<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
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
      <div class="flex justify-center items-center space-x-4">
        <ul class="m-2">
          <li class="relative flex items-center">
            <input
              class="w-96 h-10 rounded-full pl-4 pr-10 flex items-center focus:outline-none px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-300 ease-in-out"
              v-model="searchQuery"
              @input="fetchEvents"
              placeholder="Search..."
            />
            <font-awesome-icon
              :icon="['fas', 'search']"
              class="text-gray-600 absolute right-4"
            />
          </li>
        </ul>

        <ul class="m-2">
          <li class="relative">
            <button @click="openFilterModalHandler" class="flex items-center">
              <font-awesome-icon
                :icon="['fas', 'filter']"
                class="text-gray-600"
              />
              <span class="ml-2 text-gray-500">Filter</span>
            </button>
          </li>
        </ul>
      </div>
    </div>

    <div class="pt-64"></div>
    <h2 class="text-2xl font-bold mb-4 text-center">Latest Events</h2>

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
