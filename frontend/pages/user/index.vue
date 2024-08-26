<template>
  <div class="bg-gradient-to-r from-blue-100 via-green-100 to-yellow-100">
    <div class="w-full flex justify-center items-center pt-24 space-x-4">
      <img src="/assets/images/home.png" alt="Home Image 1" />
      <img src="/assets/images/home.png" alt="Home Image 2" />
    </div>

    <h2 class="text-2xl font-bold mb-4 text-center">Latest Events</h2>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in events" :key="event.id">
        <NuxtLink :to="`/events/${event.id}`">
          <CustomEventCard :event="event" />
        </NuxtLink>
      </div>
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

<script setup>
import CustomEventCard from "@/components/CustomEventCard.vue";
import TagsCategoriesComponent from "@/components/TagsCategoriesComponent.vue";
import FeaturesComponent from "@/components/FeaturesComponent.vue";
import CustomFooter from "@/components/CustomFooter.vue";
import { ref, computed, watchEffect } from "vue";
import useFetchData from "~/composables/useFetchData";
import { useAuthStore } from "~/stores";
const isAuthnticated = useAuthStore();

const userId = isAuthnticated.id;
const { events, categories, tags } = useFetchData(userId);

const visibleEvents = ref([]);
const itemsPerPage = 3;
const currentPage = ref(1);

watchEffect(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleEvents.value = events.value.slice(
    startIndex,
    startIndex + itemsPerPage
  );
});

const loadMoreEvents = () => {
  if (hasMoreEvents.value) {
    currentPage.value += 1;
  }
};

const hasMoreEvents = computed(
  () => events.value.length > visibleEvents.value.length
);

definePageMeta({ layout: "authenticated" });
</script>
