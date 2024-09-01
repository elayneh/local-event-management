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
          @click="navigateTo('/admin/tags')"
          class="bg-gray-700 text-white px-4 py-2 rounded-full"
        >
          No
        </button>
      </div>
    </CustomModel>
  </div>
</template>

<script setup>
import { useRouter, useRoute } from "vue-router";
import { toast } from "vue3-toastify";
import deleteTag from "~/graphql/mutations/tags/delete.gql";

const tags = ref(null);
const router = useRouter();
const route = useRoute();

const { mutate: deleteTagMutation, onDone } =
  useAuthenticatedMutation(deleteTag);

const deleteHandler = async () => {
  const tagId = route.params.id;
  try {
    await deleteTagMutation(
      { id: tagId },
      {
        onDone: () => {
          tags.value = tags.value.filter((tag) => {
            tag.id !== tagId;
          });
        },
      }
    );

    toast.success("Tag deleted successfully", {
      position: "top-right",
      autoClose: 3000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
    });
  } catch (error) {
    console.error("Error deleting the tag:", error);
    toast.error("Failed to delete the tag", {
      position: "top-right",
      autoClose: 3000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
    });
  } finally {
    navigateTo("/admin/tags");
  }
};

const navigateTo = (path) => {
  router.push(path);
};
definePageMeta({ layout: "admin-dashboard" });
</script>
