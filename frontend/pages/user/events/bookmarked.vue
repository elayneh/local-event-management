<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <HomepageImage />
    <div
      v-if="!bookmarkedEvents.length"
      class="text-center m-12 p-24 bg-red-300 text-red-800 rounded-lg shadow-lg"
    >
      <p class="text-lg font-semibold mb-4">No bookmarked events found.</p>
      <p class="mb-4">It seems you haven't bookmarked any events yet.</p>
      <NuxtLink
        to="/user"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
      >
        Go to Home and Bookmark Events
      </NuxtLink>
    </div>

    <div v-else>
      <h2 class="text-2xl font-bold mb-4 text-center">Bookmarked Events</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center">
        <div v-for="event in bookmarkedEvents" :key="event.id">
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
    </div>

    <TagsCategoriesComponent :categories="categories" :tags="tags" />
    <FeaturesComponent />
    <CustomFooter />
  </div>
</template>

<script setup>
import CustomEventCard from "@/components/CustomEventCard.vue";
import TagsCategoriesComponent from "~/components/TagsCategoriesComponent.vue";
import FeaturesComponent from "~/components/FeaturesComponent.vue";
import CustomFooter from "~/components/CustomFooter.vue";
import { ref, computed, onMounted } from "vue";
import useUserFetchData from "~/composables/useUserFetchData";
import useFetchData from "~/composables/useFetchData";
import { useAuthStore } from "~/stores";


const isAuthnticated = useAuthStore();
const user_id = isAuthnticated.id;

const { bookmarkedEvents } = useUserFetchData(user_id);
const { tags, categories } = useFetchData();
const visibleEvents = ref([]);
const itemsPerPage = 3;
const currentPage = ref(1);

const updateVisibleEvents = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleEvents.value = bookmarkedEvents.value.slice(
    startIndex,
    startIndex + itemsPerPage
  );
};

const loadMoreEvents = () => {
  if (hasMoreEvents.value) {
    currentPage.value += 1;
    updateVisibleEvents();
  }
};

const hasMoreEvents = computed(
  () => bookmarkedEvents.value.length > visibleEvents.value.length
);

onMounted(() => {
  updateVisibleEvents(); // Ensure events are loaded on mount
});

definePageMeta({ layout: "authenticated" });
</script>
