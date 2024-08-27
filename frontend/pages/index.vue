<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <div class="w-full flex justify-center pt-24 pb-12 space-x-6">
      <div
        class="flex items-center bg-white shadow-lg rounded-lg p-6 transform hover:scale-105 transition-transform duration-300"
      >
        <img
          :src="event_images[0]"
          alt="Image 2"
          class="w-full object-cover rounded-lg"
        />
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
import defaultImage from "~/assets/images/home.png";

const { events, categories, tags } = useFetchData();
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

const event_images = computed(() => {
  if (events?.value?.event_images) {
    const imagesArray = events.value.event_images
      .replace(/{|}/g, "")
      .split(",");
    return imagesArray.length > 0 ? imagesArray : [defaultImage];
  }
  return [defaultImage];
});

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
