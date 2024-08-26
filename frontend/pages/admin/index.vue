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

      <NuxtLink to="/admin/events">
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
import { ref, onMounted } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getUsers from "~/graphql/queries/users/getUsers.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import getEvents from "~/graphql/queries/events/getEvents.gql";

const users = ref([]);
const events = ref([]);
const categories = ref([]);
const tags = ref([]);

const loading = ref(true);
const error = ref(null);

const {
  result: usersResult,
  loading: usersLoading,
  error: usersError,
  onResult: onUsersResult,
} = useQuery(getUsers);
const {
  result: eventsResult,
  loading: eventsLoading,
  error: eventsError,
  onResult: onEventsResult,
} = useQuery(getEvents);
const {
  result: categoriesResult,
  loading: categoriesLoading,
  error: categoriesError,
  onResult: onCategoriesResult,
} = useQuery(getCategories);
const {
  result: tagsResult,
  loading: tagsLoading,
  error: tagsError,
  onResult: onTagsResult,
} = useQuery(getTags);

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

loading.value =
  usersLoading.value ||
  eventsLoading.value ||
  categoriesLoading.value ||
  tagsLoading.value;
error.value =
  usersError.value ||
  eventsError.value ||
  categoriesError.value ||
  tagsError.value;

definePageMeta({ layout: "admin-dashboard" });
</script>
