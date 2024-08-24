<template>
  <div>
    <h2 class="text-xl font-semibold mb-4 text-center text-slate-500">
      #All Widgets
    </h2>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <NuxtLink to="/admin/users">
        <div class="bg-white shadow rounded-lg p-4">
          <h3 class="text-lg font-medium">Total Users</h3>
          <p class="text-2xl font-bold">{{ users.length }}</p>
        </div>
      </NuxtLink>
      <NuxtLink to="/events">
        <div class="bg-white shadow rounded-lg p-4">
          <h3 class="text-lg font-medium">Total Events</h3>
          <p class="text-2xl font-bold">{{ events.length }}</p>
        </div>
      </NuxtLink>
      <NuxtLink to="/admin/tags">
        <div class="bg-white shadow rounded-lg p-4">
          <h3 class="text-lg font-medium">Tags</h3>
          <p class="text-2xl font-bold">{{ tags.length }}</p>
        </div>
      </NuxtLink>
      <NuxtLink to="/admin/categories">
        <div class="bg-white shadow rounded-lg p-4">
          <h3 class="text-lg font-medium">Categories</h3>
          <p class="text-2xl font-bold">{{ categories.length }}</p>
        </div>
      </NuxtLink>
    </div>
    <div v-if="loading" class="mt-4">Loading...</div>
    <div v-if="error" class="mt-4 text-red-500">Error: {{ error.message }}</div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getUsers from "~/graphql/queries/users/getUsers.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import getEvents from "~/graphql/queries/events/getEvents.gql";

const users = ref([]);
const events = ref([]);
const categories = ref([]);
const tags = ref([]);

const {
  result: usersResult,
  loading: usersLoading,
  error: usersError,
} = useQuery(getUsers);
const {
  result: eventsResult,
  loading: eventsLoading,
  error: eventsError,
} = useQuery(getEvents);
const {
  result: categoriesResult,
  loading: categoriesLoading,
  error: categoriesError,
} = useQuery(getCategories);
const {
  result: tagsResult,
  loading: tagsLoading,
  error: tagsError,
} = useQuery(getTags);

// Set up watchers to update reactive data
watch(usersResult, (newResult) => {
  if (newResult?.users) {
    users.value = newResult.users;
  }
});

watch(eventsResult, (newResult) => {
  if (newResult?.events) {
    events.value = newResult.events;
  }
});

watch(categoriesResult, (newResult) => {
  if (newResult?.categories) {
    categories.value = newResult.categories;
  }
});

watch(tagsResult, (newResult) => {
  if (newResult?.tags) {
    tags.value = newResult.tags;
  }
});

watch(
  [usersLoading, eventsLoading, categoriesLoading, tagsLoading],
  (newLoading) => {
    console.log("Loading:", newLoading);
  }
);

watch([usersError, eventsError, categoriesError, tagsError], (newError) => {
  console.error("Error:", newError);
});

definePageMeta({ layout: "admin-dashboard" });
</script>
