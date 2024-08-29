<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <HomepageImage />
    <div
      v-if="!followingEvents.length"
      class="text-center m-12 p-24 bg-red-300 text-red-800 rounded-lg shadow-lg"
    >
      <p class="text-lg font-semibold mb-4">No following events found.</p>
      <p class="mb-4">It seems you haven't followed any events yet.</p>
      <NuxtLink
        to="/user"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
      >
        Go to Home and Follow Your Favorable Events
      </NuxtLink>
    </div>

    <div v-else>
      <h2 class="text-2xl font-bold mb-4 text-center">Following Events</h2>
      <div
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center"
      >
        <div v-for="event in followingEvents" :key="event.id">
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
import HomepageImage from "~/components/HomepageImage.vue";
const isAuthnticated = useAuthStore();

const userId = isAuthnticated.id;
const { followingEvents } = useUserFetchData(userId);
const { tags, categories } = useFetchData();

const visibleEvents = ref([]);
const itemsPerPage = 3;
const currentPage = ref(1);

const updateVisibleEvents = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleEvents.value = followingEvents.value.slice(
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
  () => followingEvents.value.length > visibleEvents.value.length
);

onMounted(() => {
  updateVisibleEvents();
});

definePageMeta({ layout: "authenticated" });
</script>
