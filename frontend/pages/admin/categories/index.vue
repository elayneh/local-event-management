<template>
  <div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
      <div v-if="categories.length <= 0">
        <h2>No category found</h2>
      </div>
      <CustomCard
        v-for="category in categories"
        :key="category.id"
        :title="category.name"
        :customClass="bg - gray - 700"
      >
        <template #body>
          <p class="text-gray-700">{{ category.description }}</p>
        </template>
        <template #footer-left>
          <div class="bg-blue-500 px-8 py-2 rounded-full">
            <NuxtLink :to="`/categories/update/${category.id}`">
              <button class="text-white">Edit</button>
            </NuxtLink>
          </div>
        </template>
        <template #footer-right>
          <div class="bg-red-500 px-4 py-2 rounded-full">
            <NuxtLink :to="`/categories/delete/${category.id}`">
              <button class="text-white">Delete</button>
            </NuxtLink>
          </div>
        </template>
      </CustomCard>
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6 w-50 w-100"
      >
        <NuxtLink to="/categories/add">
          <img src="./../../assets/images/add.png" width="50px" />
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import CustomCard from "~/components/CustomCard.vue";
import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getCategories from "~/graphql/queries/categories/getCategories.gql";

const categories = ref([]);

const { result, loading, error, onResult } = useQuery(getCategories);
onResult(({ data }) => {
  if (data) {
    categories.value = data.categories;
  }
});

definePageMeta({ layout: "admin-dashboard" });
</script>
