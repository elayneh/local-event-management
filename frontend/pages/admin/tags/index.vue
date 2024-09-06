<script setup>
import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getTags from "~/graphql/queries/tags/getTags.gql";

const tags = ref([]);

const { onResult } = useQuery(getTags);

onResult((response) => {
  const newResult = response.data?.tags;
  if (newResult) {
    tags.value = newResult;
  }
});

definePageMeta({ layout: "admin-dashboard" });
</script>

<template>
  <div class="p-6">
    <div v-if="tags.length === 0" class="text-center text-gray-500">
      <h2>No tags found</h2>
    </div>

    <div v-else class="overflow-x-auto">
      <table
        class="min-w-full bg-white border border-gray-200 shadow-lg rounded-lg"
      >
        <thead class="bg-gray-200">
          <tr>
            <th class="text-left py-3 px-4 text-sm font-semibold text-gray-700">
              Tag Name
            </th>
            <th class="text-left py-3 px-4 text-sm font-semibold text-gray-700">
              Description
            </th>
            <th
              class="text-center py-3 px-4 text-sm font-semibold text-gray-700"
            >
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="tag in tags"
            :key="tag.id"
            class="border-b border-gray-200 hover:bg-gray-100 transition-all duration-300"
          >
            <td class="py-3 px-4 text-gray-800">{{ tag.name }}</td>
            <td class="py-3 px-4 text-gray-600">{{ tag.description }}</td>
            <td class="py-3 px-4 text-center">
              <NuxtLink
                :to="`/admin/tags/update/${tag.id}`"
                class="bg-blue-500 px-4 py-2 rounded-full text-white inline-block m-4 hover:bg-blue-600"
              >
                Update
              </NuxtLink>
              <NuxtLink
                :to="`/admin/tags/delete/${tag.id}`"
                class="bg-red-500 px-4 py-2 rounded-full text-white inline-block hover:bg-red-600"
              >
                Delete
              </NuxtLink>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="mt-6 flex justify-center">
      <NuxtLink
        to="/admin/tags/add"
        class="bg-green-500 text-white px-6 py-3 rounded-full shadow-lg hover:bg-green-600 transition-all duration-300"
      >
        Add New Tag
      </NuxtLink>
    </div>
  </div>
</template>
