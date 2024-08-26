<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <div
      class="w-full max-w-4xl mx-auto p-6 bg-gray-300 rounded-lg shadow-lg mt-24"
    >
      <p v-if="error" class="text-red-600 text-center mb-4">
        Error loading categories. Please try again later.
      </p>

      <div v-if="loadingCategories" class="flex justify-center items-center">
        <div
          class="w-16 h-16 border-4 border-t-4 border-blue-600 border-solid rounded-full animate-spin"
        >
          <faSpinner class="text-blue-600">Loading</faSpinner>
        </div>
      </div>

      <div class="relative mt-8">
        <h2 class="text-3xl font-semibold text-gray-800 mb-6 text-center">
          Create Your Event
        </h2>

        <DynamicForm
          v-if="!error"
          :schema="firstLevelEventFormSchema"
          class="bg-white p-8 rounded-lg shadow-md"
          :validation-schema="firstLevelSchemaValidation"
          @submit="submitFirstLevelEvent"
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
                <button
                  @click="removeTag(tag)"
                  class="text-red-500 hover:text-red-700 ml-2"
                >
                  x
                </button>
              </span>
            </p>
          </div>

          <!-- Tag Selector Modal -->
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
                  <label :for="`tag-${tag.id}`" class="text-gray-700">
                    {{ tag.label }}
                  </label>
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

          <div class="mb-5">
            <div class="relative">
              <input
                @click="openLocationModal"
                type="text"
                v-model="locationText"
                placeholder="Select event location from the map"
                class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300 cursor-pointer"
                readonly
              />
            </div>
            <p v-if="locationError" class="text-red-500 text-sm mt-2">
              {{ locationError }}
            </p>
          </div>

          <button
            type="submit"
            class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300"
            :disabled="createEventLoading"
            :class="{ 'opacity-50 cursor-not-allowed': createEventLoading }"
          >
            {{ createEventLoading ? "Creating Event..." : "Create Event" }}
          </button>
        </DynamicForm>
      </div>

      <!-- Location Selection Modal -->
      <div
        v-if="showModal"
        class="fixed inset-0 z-50 flex items-center justify-center bg-gray-800 bg-opacity-75"
      >
        <div class="container mx-auto p-6 bg-gray-100">
          <h1 class="text-3xl font-semibold mb-4">Location Form</h1>

          <div class="bg-white p-6 rounded-lg shadow-md">
            <div class="p-4 flex justify-between items-center border-b">
              <h3 class="text-xl font-semibold">Select Event Location</h3>
              <button
                @click="closeLocationModal"
                class="text-gray-600 hover:text-gray-800"
              >
                X
              </button>
            </div>
            <div class="mb-4">
              <LeafletMap
                :zoom="zoom"
                :center="center"
                @click="handleMapClick"
              />
              <LMarker v-if="marker" :lat-lng="marker" />
            </div>

            <div class="mb-4">
              <label
                for="location"
                class="block text-gray-700 text-sm font-bold mb-2"
              >
                Location Address:
              </label>
              <input
                v-model="address"
                type="text"
                id="location"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Address will be auto-filled"
                readonly
              />
            </div>

            <button
              @click="submitHandler"
              class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline"
            >
              Submit
            </button>
          </div>
        </div>
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
import ImageUploadMutation from "~/graphql/mutations/image_upload.gql";

import * as yup from "yup";
import { toast } from "vue3-toastify";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";
import LeafletMap from "~/components/LeafletMap.vue";
const now = new Date().toISOString().slice(0, 16);

const showModal = ref(false);
const selectedLocation = ref(null);
const selectedLocationName = ref(null);
const locationText = ref("");
const searchQuery = ref("");
const selectedTags = ref([]);
const selectedTagValues = ref([]);
const selectedCategoryValues = ref([]);
const categorySearchQuery = ref("");
const tagsError = ref("");
const categoryError = ref("");
const locationError = ref("");
const zoom = 13;
const center = ref([9.0306, 38.7506]); // Center set to Addis Ababa, Piassa
const marker = ref(null);
const address = ref("");
const locationData = ref(null);

// Handle map click event
const handleMapClick = async (coords) => {
  marker.value = [coords.lat, coords.lng, coords.address];
  locationText.value = marker.value[2];
  locationData.value = {
    address: coords.address,
    latitude: coords.lat,
    longitude: coords.lng,
  };
};

// Handle form submission
const submitHandler = () => {
  closeLocationModal();
};
///
const isEventCreated = ref(false);

const firstLevelSchemaValidation = yup.object({
  title: yup.string().required("Event title is required"),
  description: yup.string().required("Event description is required"),
  category_id: yup.string().required("Category is required"),
  isFree: yup.string().required("Please specify if the event is free"),
  price: yup.number(),
  venue: yup.string().required("Venue is required"),
  start_date: yup
    .date()
    .min(now, "Start date cannot be in the past")
    .required("Start date is required"),
  end_date: yup
    .date()
    .min(yup.ref("start_date"), "End date cannot be before start date")
    .required("End date is required"),
  event_images: yup.mixed().required("At least one image is required"),
});
const {
  result: categoriesResult,
  loading: loadingCategories,
  error: categoriesError,
} = useQuery(getCategories);
const {
  result: tagsResult,
  loading: loadingTags,
  error: tagsErr,
} = useQuery(getTags);

const categories = ref([]);
const tags = ref([]);
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
  loading: createEventLoading,
} = useMutation(insertEvent);
const {
  mutate: insertEventTagsMutation,
  onError: onTagsError,
  onDone: onTagsDone,
} = useMutation(insertEventTags);

const {
  mutate: uploadImages,
  loading: uploadImageLoading,
  onDone: onUploadImageDone,
  onError: onUploadImageError,
} = useMutation(ImageUploadMutation);

const eventId = ref(null);

const openLocationModal = () => {
  showModal.value = true;
};

const closeLocationModal = () => {
  showModal.value = false;
};

const removeTag = (tag) => {
  selectedTags.value = selectedTags.value.filter((t) => t.value !== tag.value);
  selectedTagValues.value = selectedTagValues.value.filter(
    (v) => v !== tag.value
  );
};

const filteredTagOptions = computed(() =>
  tagOptions.filter((option) =>
    option.label.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
);

const filteredCategoryOptions = computed(() =>
  categoryOptions.filter((option) =>
    option.label.toLowerCase().includes(categorySearchQuery.value.toLowerCase())
  )
);

const validateForm = () => {
  locationError.value = locationData.value ? "" : "Location is required";
  console.log("location daeaqqqqqqqqqqq: ", locationData);

  return !locationError.value;
};

const submitFirstLevelEvent = async (values) => {
  const files = values.event_images;

  if (!files || files.length === 0) {
    toast.error("No files selected", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    console.log("Error with image files");
    return;
  }
  console.log("location daea: ", locationData);

  if (!validateForm()) return;
  try {
    const readFilesAsBase64 = (file) => {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onloadend = () => {
          if (reader.result) {
            resolve(reader.result.split(",")[1]); // Return only the base64 part
          } else {
            reject(new Error("Failed to read file"));
          }
        };
        reader.readAsDataURL(file);
      });
    };
    console.log("Files: ", files);

    const base64Files = await Promise.all(
      files.map((file) => readFilesAsBase64(file))
    );

    console.log("base 64: ", base64Files);

    const uploadImagesInput = { input: { files: base64Files } };

    console.log("Fininnnnnnnnnnn");
    const { data } = await uploadImages(uploadImagesInput);

    if (!data || !data.uploadImages || !data.uploadImages.imageUrls) {
      throw new Error("Failed to upload images");
    }

    const uploadedImages = data.uploadImages.imageUrls;

    console.log(
      "locationnnnnnnnnnn: ",
      uploadedImages,data
    );
    const eventData = {
      address: locationData.value.address,
      location: locationData.value
        ? `${locationData.value.latitude},${locationData.value.longitude}`
        : values.location,
      category_id: values.category_id,
      description: values.description,
      end_date: values.end_date,
      is_free: values.isFree,
      price: values.isFree === "false" ? values.price : 0,
      title: values.title,
      start_date: values.start_date,
      venue: values.venue,
      event_images: `{${uploadedImages.join(",")}}`,
      quantity: parseInt(values.quantity) || 1,
    };

    console.log("Form submitted with::;;;;;;:", eventData);
    const result = await insertEventMutation(eventData);

    eventId.value = result.data.insert_events.returning[0].id;
    isEventCreated.value = true;

    await submitSecondLevelEvent(selectedTags.value);
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

const submitSecondLevelEvent = async (tags) => {
  try {
    if (!eventId.value) return;

    tags.map(
      async (tag) =>
        await insertEventTagsMutation({
          tag_id: tag.id,
          eventId: eventId.value,
        })
    );
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
      as: "input",
      name: "quantity",
      label: "Total Capacity",
      placeholder: "Enter Total Holding capacity",
      type: "number",
      rules: yup
        .number()
        .required("Total Capacity is required field")
        .min(1, "Total Capacity cannot be  less than 1"),
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
    {
      as: "input",
      name: "event_images",
      label: "Event Images",
      placeholder: "Event Images",
      type: "file",
      multiple: true,
      rules: yup.mixed().required("Event image is required"),
    },
  ],
}));

definePageMeta({ layout: "authenticated" });
</script>
