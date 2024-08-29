<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <div class="fixed w-full flex justify-center items-center pt-2 z-20">
      <div v-if="openFilterModal"></div>
      <div class="flex justify-center items-center space-x-4">
        <ul class="m-2">
          <li class="relative flex items-center">
            <input
              v-model="searchQuery"
              @change="searchHandler"
              class="w-96 h-10 rounded-full pl-4 pr-10 bg-gray-300 flex items-center focus:outline-none focus:ring-1 focus:ring-gray-400"
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
    <HomepageImage />

    <h2 class="text-2xl font-bold mb-4 text-center">Latest Events</h2>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in events" :key="event.id">
        <CustomEventCard :event="event" />
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
import HomepageImage from "~/components/HomepageImage.vue";

const { events, categories, tags } = useFetchData();
console.log(events);
const openFilterModal = ref(false);
const openFilterModalHandler = () => {
  openFilterModal.value = !openFilterModal.value;
  console.log("openFilterModal: ", openFilterModal);
};

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
