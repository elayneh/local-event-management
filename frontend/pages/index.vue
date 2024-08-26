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
    <h2 class="text-2xl font-bold mb-4 text-center text-">Latest Events</h2>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in events" :key="event.id">
        <NuxtLink :to="`/events/${event.id}`">
          <CustomEventCard :event="event" />
        </NuxtLink>
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

    <TagsCategoriesComponent :categories="categories" :tags="tags" />
    <FeaturesComponent />
    <CustomFooter />
  </div>
</template>

<script setup>
import CustomEventCard from "@/components/CustomEventCard.vue";
import { ref, computed } from "vue";
import useFetchData from "~/composables/useFetchData";

import { useAuthStore } from "~/stores";
const isAuthnticated = useAuthStore();

const userId = isAuthnticated.id;
const { events, categories, tags } = useFetchData(userId);

const visibleEvents = ref([]);

const itemsPerPage = 3;
const currentPage = ref(1);

const updateVisibleEvents = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleEvents.value = useFetchData.events.value.slice(
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
</script>
