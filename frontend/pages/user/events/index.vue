<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <HomepageImage />
    <h2 class="text-2xl font-bold mb-4 text-center">Your Events</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in visibleEvents" :key="event.id">
        <NuxtLink :to="`/events/${event.id}`">
          <CustomEventCard :event="event" />
        </NuxtLink>
      </div>
    </div>
    <div v-if="userEventsLoading" class="text-center">
      <p>Loading...</p>
    </div>
    <div v-if="userEventsError" class="text-red-500">
      <p>Error loading events: {{ userEventsError.message }}</p>
    </div>

    <div v-if="hasMoreEvents" class="text-center mt-6">
      <button
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
        @click="loadMoreEvents"
      >
        Load More
      </button>
    </div>

    <div class="bg-violet-200 mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6">
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

    <div class="mt-12">
      <div
        class="flex flex-col lg:flex-row items-center lg:items-start justify-between bg-gradient-to-r from-purple-300 via-blue-300 to-pink-300 p-6 rounded-lg shadow-lg"
      >
        <div class="mb-6 lg:mb-0 lg:w-2/3 bg-white p-6 rounded-lg shadow-md">
          <h2 class="text-2xl font-bold mb-6 text-gray-800">
            Explore Amazing Features
          </h2>
          <div class="space-y-4">
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-blue-500 text-white rounded-full p-3"
              >
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

            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-green-500 text-white rounded-full p-3"
              >
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

            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-yellow-500 text-white rounded-full p-3"
              >
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

            <div class="flex items-center">
              <div class="flex-shrink-0 bg-red-500 text-white rounded-full p-3">
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
                    d="M12 4v16m8-8H4"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Quick and Easy Event Creation
                </h3>
                <p class="text-gray-500">
                  Set up new events in just a few clicks.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import CustomEventCard from "~/components/CustomEventCard.vue";
import HomepageImage from "~/components/HomepageImage.vue";
import useUserFetchData from "~/composables/useUserFetchData.js";
import { useAuthStore } from "~/stores";

const user_id = useAuthStore().id;

const { userEvents, userEventsLoading, userEventsError, categories, tags } =
  useUserFetchData(user_id);

const pageSize = 10;
const currentPage = ref(1);

const visibleEvents = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return userEvents.value.slice(start, end);
});

const hasMoreEvents = computed(() => {
  return userEvents.value.length > currentPage.value * pageSize;
});

const loadMoreEvents = () => {
  if (hasMoreEvents.value) {
    currentPage.value += 1;
  }
};
definePageMeta({ layout: "authenticated" });
</script>
