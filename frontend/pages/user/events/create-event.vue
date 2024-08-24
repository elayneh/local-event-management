<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <div
      class="w-full max-w-4xl mx-auto p-6 bg-gray-300 rounded-lg shadow-lg mt-24"
    >
      <p v-if="error" class="text-red-600 text-center mb-4">
        Error loading categories. Please try again later.
      </p>

      <div
        v-if="loadingCategories"
        class="flex justify-center items-center h-screen"
      >
        <div
          class="w-16 h-16 border-4 border-t-4 border-blue-600 border-solid rounded-full animate-spin"
        >
          <faSpinner class="text-blue-600">Loading</faSpinner>
        </div>
      </div>

      <div class="relative mt-8">
        <!-- Added top margin here -->
        <h2 class="text-3xl font-semibold text-gray-800 mb-6 text-center">
          Create Your Event
        </h2>
        <!-- Added title here -->

        <DynamicForm
          v-if="!loadingCategories && !error && !isEventCreated"
          :schema="firstLevelEventFormSchema"
          class="bg-white p-8 rounded-lg shadow-md"
          :validation-schema="firstLevelSchemaValidation"
          @submit="submitFirstLevelEvent"
        >
          <button
            type="submit"
            class="mt-4 bg-green-500 hover:bg-green-600 text-white font-semibold py-3 px-6 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300"
          >
            Save Event Data
          </button>
        </DynamicForm>
      </div>

      <div class="relative mt-8">
        <DynamicForm
          v-if="!error && !isEventCreated"
          @submit="submitSecondLevelEvent"
          class="bg-white p-8 rounded-lg shadow-md"
        >
          <div class="mb-6 flex items-center">
            <label class="block text-gray-700 text-lg font-medium mr-4"
              >Add Event Tags</label
            >
            <button
              type="button"
              @click="showTagSelector = true"
              class="bg-gray-400 hover:bg-gray-500 text-white font-semibold py-2 px-4 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-300"
            >
              <span class="text-white text-lg">+</span>
            </button>
            <p class="ml-4 text-gray-600">
              <span
                v-for="tag in selectedTags"
                :key="tag.id"
                class="bg-gray-200 text-gray-700 px-2 py-1 rounded-full mr-2"
              >
                #{{ tag.label }}
              </span>
            </p>
          </div>

          <div
            v-if="showTagSelector"
            class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center"
          >
            <div
              class="bg-white p-6 rounded-lg shadow-lg w-3/4 md:w-1/2 lg:w-1/3"
            >
              <h2 class="text-2xl font-semibold mb-4">Select Tags</h2>
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
                    class="mr-2 text-indigo-600 form-checkbox"
                  />
                  <label :for="`tag-${tag.id}`" class="text-gray-700">{{
                    tag.label
                  }}</label>
                </div>
              </div>
              <div class="flex justify-end mt-4">
                <button
                  @click="showTagSelector = false"
                  class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg mr-2"
                >
                  Done
                </button>
                <button
                  @click="clearSelectedTags"
                  class="bg-red-600 hover:bg-red-700 text-white font-semibold py-2 px-4 rounded-lg"
                >
                  Clear
                </button>
              </div>
            </div>
          </div>

          <div class="flex gap-4 mt-8">
            <button
              @click="isEventCreated = false"
              class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-300"
            >
              Previous
            </button>
            <button
              type="submit"
              class="ml-auto bg-green-500 hover:bg-green-600 text-white font-semibold py-2 px-4 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300"
            >
              Submit Event
            </button>
          </div>
        </DynamicForm>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from "vue";
import { useQuery, useMutation } from "@vue/apollo-composable";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import insertEvent from "~/graphql/mutations/events/insert.gql";
import insertEventTags from "~/graphql/mutations/event_tags/insert.gql";

import * as yup from "yup";
import { toast } from "vue3-toastify";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";

const now = new Date().toISOString().slice(0, 16);

const isEventCreated = ref(false);

const firstLevelSchemaValidation = yup.object({
  title: yup.string().required("Event title is required"),
  description: yup.string().required("Event description is required"),
  category_id: yup.string().required("Category is required"),
  isFree: yup.string().required("Please specify if the event is free"),
  price: yup.number(),
  // .when('isFree', {
  //   is: 'false',
  //   then: yup.number().required("Event price is required").min(0, "Price cannot be negative"),
  //   otherwise: yup.number().notRequired(),
  // }),

  venue: yup.string().required("Venue is required"),
  address: yup.string().required("Address is required"),
  start_date: yup
    .date()
    .min(now, "Start date cannot be in the past")
    .required("Start date is required"),
  end_date: yup
    .date()
    .min(yup.ref("start_date"), "End date cannot be before start date")
    .required("End date is required"),
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

const {
  mutate: insertEventMutation,
  onError: onEventError,
  onDone: onEventDone,
} = useMutation(insertEvent);
const {
  mutate: insertEventTagsMutation,
  onError: onTagsError,
  onDone: onTagsDone,
} = useMutation(insertEventTags);

const eventId = ref(null);

const submitFirstLevelEvent = async (values) => {
  try {
    const eventData = {
      address: values.address,
      category_id: values.category_id,
      description: values.description,
      end_date: values.end_date,
      is_free: values.isFree,
      price: values.isFree === "false" ? values.price : 0,
      title: values.title,
      start_date: values.start_date,
      venue: values.venue,
    };

    const result = await insertEventMutation({
      variables: { input: eventData },
    });
    eventId.value = result.data.insert_events.returning[0].id;
    isEventCreated.value = true;
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

const submitSecondLevelEvent = async (values) => {
  try {
    if (!eventId.value) return;

    const tagsData = selectedTags.value.map((tag) => ({
      event_id: eventId.value,
      tag_id: tag.id,
    }));

    await insertEventTagsMutation({ variables: { objects: tagsData } });
    toast.success("Tags added successfully", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  } catch (error) {
    console.error("Error:", error);
    toast.error("Failed to add tags", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

const clearSelectedTags = () => {
  selectedTags.value = [];
};
const firstLevelEventFormSchema = computed(() => ({
  fields: [
    {
      name: "title",
      label: "Event Title",
      placeholder: "Enter event title",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "description",
      label: "Event Description",
      placeholder: "Enter event description",
      type: "textarea",
      as: "textarea",
      attrs: {
        required: true,
      },
    },
    {
      name: "category_id",
      label: "Category",
      placeholder: "Select category",
      as: "select",
      options:
        categories.value.length > 0
          ? [{ value: "", label: "Select a category" }, ...categories.value]
          : [{ value: "", label: "No category found" }],
      attrs: {
        required: true,
      },
    },
    {
      name: "isFree",
      label: "Is the event free?",
      as: "select",
      options: [
        { value: "", label: "Select True / False" },
        { value: "true", label: "True" },
        { value: "false", label: "False" },
      ],
      attrs: {
        required: true,
      },
    },
    {
      name: "price",
      label: "Event Price",
      placeholder: "Enter event price (leave empty if free)",
      type: "number",
      attrs: {
        min: 0,
        step: 0.01,
      },
    },
    {
      name: "address",
      label: "Address",
      placeholder: "Enter address",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "location",
      label: "Location",
      placeholder: "Enter location",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "venue",
      label: "Venue",
      placeholder: "Enter venue",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "start_date",
      label: "Start Date",
      placeholder: "Select start date",
      type: "datetime-local",
      attrs: {
        required: true,
      },
    },
    {
      name: "end_date",
      label: "End Date",
      placeholder: "Select end date",
      type: "datetime-local",
      attrs: {
        required: true,
      },
    },
  ],
}));

definePageMeta({ layout: "authenticated" });
</script>
<!-- 
<script setup>
import { ref, computed, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import insertEvent from "~/graphql/mutations/events/insert.gql";
import insertEventTags from "~/graphql/mutations/event_tags/insert.gql";

import * as yup from "yup";
import { toast } from "vue3-toastify";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";

const now = new Date().toISOString().slice(0, 16);

let isEventCreated = ref(false);
console.log("ISEventCreated: ", isEventCreated);

const firstLevelSchemaValidation = yup.object({
  title: yup.string().required("Event title is required"),
  description: yup.string().required("Event description is required"),
  category_id: yup.string().required("Category is required"),
  isFree: yup.string().required("Please specify if the event is free"),
  price: yup
    .number()
    .required("Event price is required")
    .min(0, "Price cannot be negative"),
  venue: yup.string().required("Venue is required"),
  address: yup.string().required("Address is required"),
  start_date: yup
    .date()
    .min(now, "Start date cannot be in the past")
    .required("Start date is required"),
  end_date: yup
    .date()
    .min(yup.ref("start_date"), "End date cannot be before start date")
    .required("End date is required"),
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

const {
  mutate: insertEventMutation,
  onError,
  onResult,
  loading,
  onDone,
} = useAuthenticatedMutation(insertEvent);

const eventId = ref(null);

const submitFirstLevelEvent = async (values) => {
  try {
    const eventData = {
      address: values.address,
      category_id: values.category_id,
      description: values.description,
      end_date: values.end_date,
      is_free: values.isFree,
      location: values.location,
      price: values.isFree === "false" ? values.price : 0,
      title: values.title,
      start_date: values.start_date,
      venue: values.venue,
    };

    insertEventMutation(eventData);
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

onDone((result) => {
  eventId.value = result.data.insert_events.returning[0].id;
  isEventCreated.value = true;
});

onError((error) => {
  console.log("Error: ", error.message);
  toast.error("Something went wrong, try again", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});

const submitSecondLevelEvent = async (values) => {
  try {
    const tagsData = selectedTags.value.map((tag) => ({
      event_id: eventId.value,

      tag_id: tag.id,
    }));

    console.log("Tags: ", tagsData);

    insertEventTagsMutation({ objects: imageData });
  } catch (error) {
    console.error("Error:", error);

    toast.error("Failed to add images and tags", {
      transition: toast.TRANSITIONS.FLIP,

      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

const firstLevelEventFormSchema = computed(() => ({
  fields: [
    {
      name: "title",
      label: "Event Title",
      placeholder: "Enter event title",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "description",
      label: "Event Description",
      placeholder: "Enter event description",
      type: "textarea",
      as: "textarea",
      attrs: {
        required: true,
      },
    },
    {
      name: "category_id",
      label: "Category",
      placeholder: "Select category",
      as: "select",
      options:
        categories.value.length > 0
          ? [{ value: "", label: "Select a category" }, ...categories.value]
          : [{ value: "", label: "No category found" }],
      attrs: {
        required: true,
      },
    },
    {
      name: "isFree",
      label: "Is the event free?",
      as: "select",
      options: [
        { value: "", label: "Select True / False" },
        { value: "true", label: "True" },
        { value: "false", label: "False" },
      ],
      attrs: {
        required: true,
      },
    },
    {
      name: "price",
      label: "Event Price",
      placeholder: "Enter event price (leave empty if free)",
      type: "number",
      attrs: {
        min: 0,
        step: 0.01,
      },
    },
    {
      name: "address",
      label: "Address",
      placeholder: "Enter address",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "location",
      label: "Location",
      placeholder: "Enter location",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "venue",
      label: "Venue",
      placeholder: "Enter venue",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "start_date",
      label: "Start Date",
      placeholder: "Select start date",
      type: "datetime-local",
      attrs: {
        required: true,
      },
    },
    {
      name: "end_date",
      label: "End Date",
      placeholder: "Select end date",
      type: "datetime-local",
      attrs: {
        required: true,
      },
    },
  ],
}));

const clearSelectedTags = () => {
  selectedTags.value = [];
};

definePageMeta({ layout: "authenticated" });
</script> -->
