<template>
  <div class="relative">
    <DynamicForm
      v-if="!error"
      :schema="secondLevelEventFormSchema"
      @submit="submitSecondLevelEvent"
      class="bg-white px-8 pt-6 pb-8 mb-4"
      :validation-schema="secondLevelSchemaValidation"
      >
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2"
          >Add Event Tags</label
        >
        <button
          type="button"
          @click="showTagSelector = true"
          class="bg-gray-300 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          <span class="text-white-500 h-full">+</span>
        </button>
        <p class="mt-2">
          <span v-for="tag in selectedTags" :key="tag.id" class="mr-2"
            >#{{ tag.label }}</span
          >
        </p>
      </div>
      <div
        v-if="showTagSelector"
        class="fixed inset-0 bg-gray-900 bg-opacity-50 flex justify-center items-center"
      >
        <div class="bg-white p-6 rounded-lg shadow-lg w-1/3">
          <h2 class="text-xl font-bold mb-4">Select Tags</h2>
          <div class="mb-4">
            <div
              v-for="tag in tags"
              :key="tag.id"
              class="flex items-center mb-2"
            >
              <input
                type="checkbox"
                :id="`tag-${tag.id}`"
                :value="tag"
                v-model="selectedTags"
                class="mr-2"
              />
              <label :for="`tag-${tag.id}`">{{ tag.label }}</label>
            </div>
          </div>
          <div class="flex justify-end">
            <button
              @click="showTagSelector = false"
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mr-2"
            >
              Done
            </button>
            <button
              @click="clearSelectedTags"
              class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
            >
              Clear
            </button>
          </div>
        </div>
      </div>
      <div class="flex gapr-4">
        <button
          @click="isEventCreated = false"
          class="absolute left-0 justify-center bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          Previous
        </button>
        <button
          type="submit"
          class="absolute right-8 bg-green-400 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          Submit Event
        </button>
      </div>
    </DynamicForm>
  </div>
</template>

<script setup>
import { ref, computed, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import insertEventTags from "~/graphql/mutations/event_tags/insert.gql";
import insertEventImage from "~/graphql/mutations/event_images/insert.gql";
import fileToBase64 from "~/composables/fileToBase64";

import * as yup from "yup";
import { toast } from "vue3-toastify";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";

const now = new Date().toISOString().slice(0, 16);

const secondLevelSchemaValidation = yup.object({
  featured_image: yup.mixed().notRequired(),
});

const {
  result: categoriesResult,
  loading: loadingCategories,
  error: categoriesError,
} = useQuery(getCategories);

const {
  result: tagsResult,
  loading: loadingTags,
  error: tagsError,
} = useQuery(getTags);

const categories = ref([]);
const tags = ref([]);
const selectedTags = ref([]);
const showTagSelector = ref(false);

watch(
  categoriesResult,
  (newResult) => {
    if (newResult?.categories) {
      categories.value = newResult.categories.map((category) => ({
        value: category.id,
        label: category.name,
      }));
    }
  },
  { immediate: true }
);

watch(
  tagsResult,
  (newResult) => {
    if (newResult?.tags) {
      tags.value = newResult.tags.map((tag) => ({
        id: tag.id,
        label: tag.name,
      }));
    }
  },
  { immediate: true }
);

const { mutate: insertEventTagsMutation, onDone: OnInsertEventTagsDone } =
  useAuthenticatedMutation(insertEventTags);

const { mutate: insertEventImageMutation, onDone: OnInsertEventImageDone } =
  useAuthenticatedMutation(insertEventImage);

const eventId = ref("a4a6067c-f491-422e-9f58-5956826b2fd6");

const submitSecondLevelEvent = async (values) => {
  try {
    const imageFile = values.image;
    const base64Image = await fileToBase64(imageFile);

    const imageData = {
      event_id: eventId.value,
      is_featured: true,
      image_url: base64Image,
    };

    const tagsData = selectedTags.value.map((tag) => ({
      event_id: eventId.value,
      tag_id: tag.id,
    }));
    console.log("Image: ", imageData);
    insertEventImageMutation({ objects: imageData });
  } catch (error) {
    console.error("Error:", error);
    toast.error("Failed to add images and tags", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

const secondLevelEventFormSchema = reactive({
  fields: [
    {
      name: "image",
      as: "input",
      type: "file",
      placeholder: "Featured Image",
      rules: yup.mixed().required("Product image is required"),
    },
  ],
});

// const clearSelectedTags = () => {
//   selectedTags.value = [];
// };

definePageMeta({ layout: "authenticated" });
</script>
