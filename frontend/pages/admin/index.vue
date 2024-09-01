<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 text-center text-slate-700">
      #Dashboard Overview
    </h2>

    <!-- Grid of Widgets -->
   <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
  <NuxtLink to="/admin/users" class="group">
    <div
      class="bg-white text-gray-800 shadow-lg rounded-lg p-6 transform transition-transform group-hover:scale-105"
    >
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold">Total Users</h3>
        <div class="bg-indigo-700 p-2 rounded-full text-white">
          <i class="fas fa-user"></i>
        </div>
      </div>
      <p class="text-4xl font-extrabold mt-4">{{ users.length }}</p>
    </div>
  </NuxtLink>

  <NuxtLink to="/admin/events" class="group">
    <div
      class="bg-white text-gray-800 shadow-lg rounded-lg p-6 transform transition-transform group-hover:scale-105"
    >
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold">Total Events</h3>
        <div class="bg-green-700 p-2 rounded-full text-white">
          <i class="fas fa-calendar"></i>
        </div>
      </div>
      <p class="text-4xl font-extrabold mt-4">{{ events.length }}</p>
    </div>
  </NuxtLink>

  <NuxtLink to="/admin/tags" class="group">
    <div
      class="bg-white text-gray-800 shadow-lg rounded-lg p-6 transform transition-transform group-hover:scale-105"
    >
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold">Tags</h3>
        <div class="bg-gray-700 p-2 rounded-full text-white">
          <i class="fas fa-tags"></i>
        </div>
      </div>
      <p class="text-4xl font-extrabold mt-4">{{ tags.length }}</p>
    </div>
  </NuxtLink>

  <NuxtLink to="/admin/categories" class="group">
    <div
      class="bg-white text-gray-800 shadow-lg rounded-lg p-6 transform transition-transform group-hover:scale-105"
    >
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold">Categories</h3>
        <div class="bg-yellow-700 p-2 rounded-full text-white">
          <i class="fas fa-list"></i>
        </div>
      </div>
      <p class="text-4xl font-extrabold mt-4">{{ categories.length }}</p>
    </div>
  </NuxtLink>
</div>


    <!-- Loading and Error Messages -->
    <div v-if="loading" class="mt-4 text-center">
      <svg
        class="animate-spin h-8 w-8 text-blue-500 mx-auto"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M12 2a10 10 0 100 20 10 10 0 000-20z"
        />
      </svg>
      <p class="mt-2">Loading...</p>
    </div>
    <div v-if="error" class="mt-4 text-center text-red-500">
      Error: {{ error.message }}
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getUsers from "~/graphql/queries/users/getUsers.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import getEvents from "~/graphql/queries/events/getEvents.gql";

const users = ref([]);
const events = ref([]);
const categories = ref([]);
const tags = ref([]);

const limit = ref(10);
const offset = ref(0);
const where = ref(null);

const {
  result,
  loading: eventsLoading,
  error: eventsError,
  onResult: onEventsResult,
} = useQuery(getEvents, {
  limit: limit.value,
  offset: offset.value,
  where: where.value,
});

const {
  loading: usersLoading,
  error: usersError,
  onResult: onUsersResult,
} = useQuery(getUsers);

const {
  loading: categoriesLoading,
  error: categoriesError,
  onResult: onCategoriesResult,
} = useQuery(getCategories);
const {
  loading: tagsLoading,
  error: tagsError,
  onResult: onTagsResult,
} = useQuery(getTags);

// Update data when queries resolve
onUsersResult(({ data }) => {
  if (data) users.value = data.users;
});
onEventsResult(({ data }) => {
  if (data) events.value = data.events;
});
onCategoriesResult(({ data }) => {
  if (data) categories.value = data.categories;
});
onTagsResult(({ data }) => {
  if (data) tags.value = data.tags;
});

// Computed properties to determine loading and error states
const loading = computed(
  () =>
    usersLoading.value ||
    eventsLoading.value ||
    categoriesLoading.value ||
    tagsLoading.value
);
const error = computed(
  () =>
    usersError.value ||
    eventsError.value ||
    categoriesError.value ||
    tagsError.value
);

definePageMeta({ layout: "admin-dashboard" });
</script>
