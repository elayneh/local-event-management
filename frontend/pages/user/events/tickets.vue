<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <HomepageImage />
    <div
      v-if="!tickets.length"
      class="text-center m-12 p-24 bg-red-300 text-red-800 rounded-lg shadow-lg"
    >
      <p class="text-lg font-semibold mb-4">No tickets found.</p>
      <p class="mb-4">It seems you haven't bought any tickets yet.</p>
      <NuxtLink
        to="/user"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
      >
        Go to Home and Buy Tickets
      </NuxtLink>
    </div>

    <div v-else>
      <h2 class="text-2xl font-bold mb-4 text-center">Tickets</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div v-for="ticket in visibleTickets" :key="ticket.id">
          <CustomTicketCard :ticket="ticket" />
        </div>
      </div>

      <div v-if="hasMoreTickets" class="text-center mt-6">
        <button
          class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
          @click="loadMoreTickets"
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
import TagsCategoriesComponent from "~/components/TagsCategoriesComponent.vue";
import FeaturesComponent from "~/components/FeaturesComponent.vue";
import CustomFooter from "~/components/CustomFooter.vue";
import { ref, computed, onMounted } from "vue";
import useFetchData from "~/composables/useFetchData";
import { useAuthStore } from "~/stores";
import CustomTicketCard from "~/components/CustomTicketCard.vue";
import HomepageImage from "~/components/HomepageImage.vue";

const isAuthenticated = useAuthStore();

// Fetch user ID
const userId = isAuthenticated.id;

// Use composable to fetch data
const { tickets, tags, categories } = useFetchData(userId);

// Pagination control
const visibleTickets = ref([]);
const itemsPerPage = 3;
const currentPage = ref(1);

// Function to update visible tickets for current page
const updateVisibleTickets = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleTickets.value = tickets.value.slice(
    startIndex,
    startIndex + itemsPerPage
  );
};

// Function to load more tickets on button click
const loadMoreTickets = () => {
  if (hasMoreTickets.value) {
    currentPage.value += 1;
    updateVisibleTickets();
  }
};

// Check if more tickets are available
const hasMoreTickets = computed(
  () => tickets.value.length > visibleTickets.value.length
);

// Initial load
onMounted(() => {
  updateVisibleTickets();
});

definePageMeta({ layout: "authenticated" });
</script>
