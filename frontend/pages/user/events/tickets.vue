
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

const userId = isAuthenticated.id;

const { tickets, tags, categories } = useFetchData(userId);

const visibleTickets = ref([]);
const itemsPerPage = 3;
const currentPage = ref(1);

const updateVisibleTickets = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleTickets.value = tickets.value.slice(
    startIndex,
    startIndex + itemsPerPage
  );
};

const loadMoreTickets = () => {
  if (hasMoreTickets.value) {
    currentPage.value += 1;
    updateVisibleTickets();
  }
};

const hasMoreTickets = computed(
  () => tickets.value.length > visibleTickets.value.length
);

onMounted(() => {
  updateVisibleTickets();
});

definePageMeta({ layout: "authenticated" });
</script>


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
    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
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
