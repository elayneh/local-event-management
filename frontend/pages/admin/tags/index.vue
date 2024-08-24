<template>
  <div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
      <div v-if="tags.length <= 0">
        <h2>No tag found</h2>
      </div>
      <CustomCard
        v-for="tag in tags"
        :key="tag.id"
        :title="tag.name"
        :customClass="bg - gray - 700"
      >
        <template #body>
          <p class="text-gray-700">{{ tag.description }}</p>
        </template>
        <template #footer-left>
          <div class="bg-blue-500 px-8 py-2 rounded-full">
            <NuxtLink :to="`/tags/edit/${tag.id}`">
              <button class="text-white">Edit</button>
            </NuxtLink>
          </div>
        </template>
        <template #footer-right>
          <div class="bg-red-500 px-4 py-2 rounded-full">
            <NuxtLink :to="`/tags/delete/${tag.id}`">
              <button class="text-white">Delete</button>
            </NuxtLink>
          </div>
        </template>
      </CustomCard>
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6 w-50 w-100"
      >
        <NuxtLink to="/tags/add">
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
import getTags from "~/graphql/queries/tags/getTags.gql";

const tags = ref([]);
const { result, loading, error } = useQuery(getTags);

definePageMeta({ layout: "admin-dashboard" });

watch(result, (newResult) => {
  if (newResult?.tags) {
    tags.value = newResult.tags;
  }
});

watch(loading, (newLoading) => {
  console.log("Loading:", newLoading);
});

watch(error, (newError) => {
  console.error("Error:", newError);
});
</script>
