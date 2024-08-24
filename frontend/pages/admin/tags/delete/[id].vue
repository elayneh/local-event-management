<template>
  <div
    class="flex flex-col items-center justify-center bg-gray-300 p-4 rounded"
  >
    <CustomModel title="Delete Tag" content="Confirm delete tag?">
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
          @click="navigateTo('/tags')"
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
import { useRouter, useRoute } from "#app";
import deleteTag from "~/graphql/mutations/tags/delete.gql";

const router = useRouter();
const route = useRoute();

const { mutate: deleteTagMutation } = useMutation(deleteTag);

const deleteHandler = () => {
  const tagId = route.params.id;
  console.log("TagID: ", tagId);

  deleteTagMutation({ id: tagId })
    .then(({ data }) => {
      console.log("Tag deleted", data);
      router.push("/tags");
    })
    .catch((err) => {
      console.error("Error deleting tag:", err);
    });
};
definePageMeta({ layout: "admin-dashboard" });

const navigateTo = (path) => {
  router.push(path);
};
</script>
