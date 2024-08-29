<template>
  <div
    class="flex flex-col items-center justify-center bg-gray-300 p-4 rounded"
  >
    <CustomModel title="Delete Category" content="Confirm delete category?">
      <h2 class="text-white text-xl font-semibold mb-4">{{ title }}</h2>
      <p class="text-white mb-4">{{ content }}</p>

      <div class="flex gap-4">
        <button
          @click="deleteHandler"
          class="bg-red-500 text-white px-4 py-2 rounded-full"
        >
          Yes
        </button>
        <button
          @click="navigateTo('/admin/categories')"
          class="bg-gray-700 text-white px-4 py-2 rounded-full"
        >
          No
        </button>
      </div>
    </CustomModel>
  </div>
</template>

<script setup>
import { useMutation } from "@vue/apollo-composable";
import { useRoute, useRouter } from "vue-router";
import deleteCategory from "~/graphql/mutations/categories/delete.gql";

const router = useRouter();
const route = useRoute();

const { mutate: deleteCategoryMutation } = useMutation(deleteCategory);

const deleteHandler = () => {
  const categoryId = route.params.id;
  console.log("CategoryID: ", categoryId);

  deleteCategoryMutation({ id: categoryId })
    .then(({ data }) => {
      console.log("Category deleted", data);
      router.push("/admin/categories");
    })
    .catch((err) => {
      console.error("Error deleting category:", err);
    });
};
definePageMeta({ layout: "admin-dashboard" });

const navigateTo = (path) => {
  router.push(path);
};
</script>
