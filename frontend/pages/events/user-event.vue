<template>
  <div class="bg-gradient-to-r from-blue-100 via-green-100 to-yellow-1">
    <div class="w-full flex justify-center items-center pt-24">
      <div class="items-center">
        <img src="/assets/images/home.png" />
      </div>
      <div class="items-center">
        <img src="/assets/images/home.png" />
      </div>
    </div>
    <h2 class="text-2xl font-bold mb-4 text-center">Latest Events</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in visibleEvents" :key="event.id">
        <CustomEventCard :event="event" />
      </div>
    </div>

    <div v-if="hasMoreEvents" class="text-center mt-6">
      <button
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
        @click="loadMoreEvents"
      >
        Load More
      </button>
    </div>

    <!-- Categories and Tags Section -->
    <div class="bg-violet-200 mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Categories -->
      <div>
        <h2 class="text-xl font-bold mb-2">Categories</h2>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(category, index) in categories"
            :key="index"
            class="bg-blue-100 text-blue-700 px-3 py-1 rounded-full text-sm font-semibold"
          >
            {{ category.name }}
          </span>
        </div>
      </div>

      <!-- Tags -->
      <div>
        <h2 class="text-xl font-bold mb-2">Tags</h2>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(tag, index) in tags"
            :key="index"
            class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-sm font-semibold"
          >
            {{ tag.name }}
          </span>
        </div>
      </div>
    </div>

    <!-- Interactive Features and Create Event Section -->
    <div class="mt-12">
      <div
        class="flex flex-col lg:flex-row items-center lg:items-start justify-between bg-gradient-to-r from-purple-300 via-blue-300 to-pink-300 p-6 rounded-lg shadow-lg"
      >
        <!-- Features Info Card -->
        <div class="mb-6 lg:mb-0 lg:w-2/3 bg-white p-6 rounded-lg shadow-md">
          <h2 class="text-2xl font-bold mb-6 text-gray-800">
            Explore Amazing Features
          </h2>
          <div class="space-y-4">
            <!-- Feature 1 -->
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-blue-500 text-white rounded-full p-3"
              >
                <!-- Icon for customization -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 8v4l3 3m6-7v10a2 2 0 01-2 2H6a2 2 0 01-2-2V7a2 2 0 012-2h7m-2-4h4m-2 4v4m-2 4l3 3"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Customize Your Event Listings
                </h3>
                <p class="text-gray-500">
                  Tailor your events with flexible settings and options.
                </p>
              </div>
            </div>

            <!-- Feature 2 -->
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-green-500 text-white rounded-full p-3"
              >
                <!-- Icon for community -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 20h5V10a2 2 0 00-2-2h-3a2 2 0 00-2 2v2m-8 10h5m-3-10a2 2 0 00-2-2H4a2 2 0 00-2 2v2a2 2 0 002 2h5a2 2 0 012-2v-2"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Connect with a Vibrant Community
                </h3>
                <p class="text-gray-500">
                  Engage with users and share experiences.
                </p>
              </div>
            </div>

            <!-- Feature 3 -->
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-yellow-500 text-white rounded-full p-3"
              >
                <!-- Icon for management -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8 17v5m8-5v5M12 14V9m0-2a4 4 0 00-4 4v2h8V9a4 4 0 00-4-4z"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Manage Attendance Seamlessly
                </h3>
                <p class="text-gray-500">
                  Keep track of participants and event logistics.
                </p>
              </div>
            </div>

            <!-- Feature 4 -->
            <div class="flex items-center">
              <div class="flex-shrink-0 bg-red-500 text-white rounded-full p-3">
                <!-- Icon for analytics -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 19V6a1 1 0 011-1h1a1 1 0 011 1v13a1 1 0 01-1 1h-1a1 1 0 01-1-1zm-4 0v-7a1 1 0 011-1h1a1 1 0 011 1v7a1 1 0 01-1 1H8a1 1 0 01-1-1zm8 0v-9a1 1 0 011-1h1a1 1 0 011 1v9a1 1 0 01-1 1h-1a1 1 0 01-1-1z"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Track Event Analytics
                </h3>
                <p class="text-gray-500">
                  Gain insights with real-time data and reports.
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="lg:w-1/3 flex justify-center lg:justify-end">
          <NuxtLink to="/user/events/create-event"
            ><button
              class="bg-blue-600 text-white px-6 py-3 rounded-full font-semibold hover:bg-blue-700 transition-colors duration-200 mt-6 lg:mt-0"
            >
              Create Your Event Here
            </button></NuxtLink
          >
        </div>
      </div>
    </div>
    <CustomFooter />
  </div>
</template>

<script setup>
import CustomEventCard from "@/components/CustomEventCard.vue";
import { ref, watch, computed } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";

const events = ref([]);
const visibleEvents = ref([]);
const { result, loading, error } = useQuery(getEvents);

const itemsPerPage = 9;
const currentPage = ref(1);

watch(result, (newResult) => {
  if (newResult?.events) {
    events.value = newResult.events;
    updateVisibleEvents();
  }
});

watch(loading, (newLoading) => {
  console.log("Loading:", newLoading);
});

watch(error, (newError) => {
  console.error("Error:", newError);
});

const updateVisibleEvents = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleEvents.value = events.value.slice(
    startIndex,
    startIndex + itemsPerPage
  );
};

const loadMoreEvents = () => {
  if (hasMoreEvents) {
    currentPage.value += 1;
    updateVisibleEvents();
  }
};

const hasMoreEvents = computed(
  () => events.value.length > visibleEvents.value.length
);

definePageMeta({ layout: "authenticated" });
</script>
