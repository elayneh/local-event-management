<template>
  <div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
      <div v-if="events.length <= 0">
        <h2>No event found</h2>
      </div>
      <CustomCard
        v-for="event in events"
        :key="event.id"
        :title="event.name"
        :customClass="bg - gray - 700"
      >
        <template #body>
          <p class="text-gray-700">{{ event.description }}</p>
        </template>
        <template #footer-left>
          <div class="bg-blue-500 px-8 py-2 rounded-full">
            <NuxtLink :to="`/events/edit/${event.id}`">
              <button class="text-white">Edit</button>
            </NuxtLink>
          </div>
        </template>
        <template #footer-right>
          <div class="bg-red-500 px-4 py-2 rounded-full">
            <NuxtLink :to="`/events/delete/${event.id}`">
              <button class="text-white">Delete</button>
            </NuxtLink>
          </div>
        </template>
      </CustomCard>
    </div>
  </div>
</template>

<script setup>
import CustomCard from "~/components/CustomCard.vue";
import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";

const events = ref([]);
const { result, loading, error } = useQuery(getEvents);

definePageMeta({ layout: "admin-dashboard" });

watch(result, (newResult) => {
  if (newResult?.events) {
    events.value = newResult.events;
  }
});

watch(loading, (newLoading) => {
  console.log("Loading:", newLoading);
});

watch(error, (newError) => {
  console.error("Error:", newError);
});
</script>
