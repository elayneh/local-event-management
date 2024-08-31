<template>
  <div
    class="fixed inset-0 bg-gray-800 bg-opacity-75 z-100 overflow-y-auto"
    v-if="isVisible"
  >
    <div
      class="relative max-w-3xl mx-auto mt-12 p-8 bg-white shadow-md rounded-lg overflow-y-auto h-[calc(100vh-3rem)]"
    >
      <button
        @click="closeModal"
        class="absolute top-4 right-4 p-2 bg-gray-200 hover:bg-red-100 rounded-full shadow-md transition-all duration-300 ease-in-out transform hover:scale-110 hover:shadow-lg text-gray-600 hover:text-gray-800"
      >
        <font-awesome-icon :icon="['fas', 'times']" class="text-lg" />
      </button>

      <h2 class="text-lg font-semibold mb-4">Filter Events</h2>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">Payment</label>
        <div class="flex space-x-4">
          <label class="inline-flex items-center">
            <input
              type="radio"
              value="true"
              v-model="isFree"
              class="form-radio"
            />
            <span class="ml-2">True</span>
          </label>
          <label class="inline-flex items-center">
            <input
              type="radio"
              value="false"
              v-model="isFree"
              class="form-radio"
            />
            <span class="ml-2">False</span>
          </label>
        </div>
      </div>

      <div class="mb-5">
        <label class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
          >Category</label
        >
        <div class="relative">
          <div
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300 cursor-pointer"
            @click="toggleCategoryDropdown"
          >
            <div class="flex flex-wrap gap-2">
              <span
                v-if="selectedCategory"
                class="bg-green-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
              >
                {{ selectedCategory.name }}
                <button
                  @click="removeCategory"
                  class="ml-2 text-white hover:text-gray-200"
                >
                  <button
                    @click="closeModal"
                    class="absolute top-4 right-4 p-2 bg-gray-200 hover:bg-gray-300 rounded-full shadow-md transition-all duration-300 ease-in-out transform hover:scale-110 hover:shadow-lg text-gray-600 hover:text-gray-800"
                  >
                    <font-awesome-icon
                      :icon="['fas', 'times']"
                      class="text-lg"
                    />
                  </button>
                </button>
              </span>
              <span v-else class="text-gray-500">Select Category</span>
            </div>
          </div>
          <div
            v-if="categoryDropdownOpen"
            class="absolute z-10 mt-2 w-full bg-white border border-gray-300 rounded-lg shadow-lg mb-10"
          >
            <div class="px-4 py-2 border-b">
              <input
                type="text"
                v-model="categorySearchQuery"
                placeholder="Search categories..."
                class="w-full px-3 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300"
              />
            </div>
            <div class="max-h-40 overflow-y-auto">
              <div
                v-for="option in filteredCategoryOptions"
                :key="option.name"
                class="px-4 py-2 flex items-center cursor-pointer hover:bg-gray-100"
                @click="selectCategory(option)"
              >
                <input
                  type="radio"
                  :value="option.name"
                  v-model="selectedCategoryValue"
                  class="mr-2"
                />
                {{ option.name }}
              </div>
            </div>
          </div>
        </div>
        <p v-if="categoryError" class="text-red-500 text-sm mt-2">
          {{ categoryError }}
        </p>
      </div>

      <div class="mb-5">
        <label class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
          >Tags</label
        >
        <div class="relative">
          <div
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300 cursor-pointer"
            @click="toggleDropdown"
          >
            <div class="flex flex-wrap gap-2">
              <span
                v-if="selectedTags.length > 0"
                v-for="tag in selectedTags"
                :key="tag.value"
                class="bg-blue-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
              >
                {{ tag.name }}
                <button
                  @click="closeModal"
                  class="absolute top-4 right-4 p-2 bg-gray-200 hover:bg-gray-300 rounded-full shadow-md transition-all duration-300 ease-in-out transform hover:scale-110 hover:shadow-lg text-gray-600 hover:text-gray-800"
                >
                  <font-awesome-icon :icon="['fas', 'times']" class="text-lg" />
                </button>
              </span>
              <span v-else class="text-gray-500">Select Tags</span>
            </div>
          </div>
          <div
            v-if="dropdownOpen"
            class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center"
          >
            <div
              class="bg-white p-6 rounded-lg shadow-lg w-3/4 md:w-1/2 lg:w-1/3"
            >
              <h2 class="text-2xl font-semibold mb-4">Select Tags</h2>
              <div class="px-4 py-2 border-b">
                <input
                  type="text"
                  v-model="searchQuery"
                  placeholder="Search tags..."
                  class="w-full px-3 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300"
                />
              </div>
              <div class="mb-4">
                <div v-if="tagOptions?.value.length > 0">
                  <div
                    v-for="option in filteredTagOptions"
                    :key="option.name"
                    class="px-4 py-2 flex items-center cursor-pointer hover:bg-gray-100"
                    @click="selectTag(option)"
                  >
                    <input
                      type="checkbox"
                      :id="`tag-${tag.id}`"
                      :value="option.name"
                      v-model="selectedTagValues"
                      class="mr-2 text-indigo-600 form-checkbox"
                    />
                    <label :for="`tag-${tag.id}`" class="text-gray-700">{{
                      option.label
                    }}</label>
                  </div>
                </div>
              </div>
              <div class="flex justify-end gap-4 mt-4">
                <button
                  @click="showTagSelector = false"
                  class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg"
                >
                  Done
                </button>
                <button
                  @click="clearSelectedTags"
                  class="bg-red-600 hover:bg-red-700 text-white font-semibold py-2 px-4 rounded-lg"
                >
                  Cancel
                </button>
              </div>
            </div>
          </div>
        </div>
        <p v-if="tagsError" class="text-red-500 text-sm mt-2">
          {{ tagsError }}
        </p>
      </div>

      <div class="mb-4 mt-20">
        <label class="block text-sm font-medium mb-2">Start Date</label>
        <input
          v-model="start_date"
          type="datetime-local"
          class="w-full p-3 border rounded"
        />
      </div>
      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">End Date</label>
        <input
          v-model="end_date"
          type="datetime-local"
          class="w-full p-3 border rounded"
        />
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">Venue</label>
        <div class="relative">
          <input
            v-model="venue"
            type="text"
            placeholder="Enter venue"
            class="w-full p-3 border rounded"
          />
          <div class="mt-2">
            <MapComponent @location-selected="handleLocationSelected" />
          </div>
        </div>
      </div>

      <div
        class="mx-12 flex justify-center items-center justify-between space-between"
      >
        <button
          @click="applyFilters"
          class="bg-green-600 text-white px-6 py-3 rounded-full shadow-lg hover:bg-green-700"
        >
          Filter
        </button>
        <button
          @click="resetFilters"
          class="bg-red-500 text-white px-6 py-3 rounded-full shadow-lg hover:bg-red-700"
        >
          Reset
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from "vue";
import MapComponent from "@/components/LeafletMap.vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import { defineEmits, defineProps } from "vue";

const venue = ref("");
const start_date = ref("");
const end_date = ref("");
const minPriceRange = ref(0);
const maxPriceRange = ref(5000);
const isFree = ref(false);
const mapLocation = ref({ lat: null, lng: null });

const searchQuery = ref("");
const selectedTags = ref([]);
const selectedTagValues = ref([]);
const dropdownOpen = ref(false);
const selectedCategory = ref(null);
const selectedCategoryValue = ref(null);
const categoryDropdownOpen = ref(false);
const categorySearchQuery = ref("");
const tagsError = ref("");
const categoryError = ref("");

const tagOptions = ref([]);

const categoryOptions = ref([]);

const props = defineProps({
  isVisible: Boolean,
});

const emit = defineEmits(["item-selected", "close-modal"]);

const { onResult: tagResult, refetch: refetchTags } = useQuery(getTags, {
  limit: 100,
  offset: 0,
  order_by: [{ created_at: "desc" }],
});
const { onResult: categoryResult, refetch: refetchCategories } = useQuery(
  getCategories,
  {
    limit: 100,
    offset: 0,
    order_by: [{ created_at: "desc" }],
  }
);
tagResult((data) => {
  if (data) {
    tagOptions.value = data.tags;
  }
});
categoryResult((result) => {
  if (result.data) {
    categoryOptions.value = result.data.categories;
  }
});

const filteredCategoryOptions = computed(() => {
  if (!categoryOptions.value) return [];
  if (categorySearchQuery.value) {
    return categoryOptions.value.filter((category) =>
      category.name
        .toLowerCase()
        .includes(categorySearchQuery.value.toLowerCase())
    );
  }
  return categoryOptions.value;
});

const toggleCategoryDropdown = () => {
  categoryDropdownOpen.value = !categoryDropdownOpen.value;
};

const selectCategory = (category) => {
  selectedCategory.value = category;
  selectedCategoryValue.value = category.name;
  categoryDropdownOpen.value = false;
};

const removeCategory = () => {
  selectedCategory.value = null;
  selectedCategoryValue.value = null;
};
const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

const toggleTag = (tag) => {
  const index = selectedTags.value.findIndex((t) => t.value === tag.name);
  if (index === -1) {
    selectedTags.value.push(tag);
    selectedTagValues.value.push(tag.name);
  } else {
    selectedTags.value.splice(index, 1);
    selectedTagValues.value.splice(
      selectedTagValues.value.indexOf(tag.name),
      1
    );
  }
};

const removeTag = (tag) => {
  selectedTags.value = selectedTags?.value.filter((t) => t.value !== tag.value);
  selectedTagValues.value = selectedTagValues.value.filter(
    (v) => v !== tag.name
  );
};

const filteredTagOptions = computed(() =>
  tagOptions.value.filter((option) =>
    option.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
);

const handleLocationSelected = (event) => {
  mapLocation.value = event.detail.location;
};

const emitFilterData = () => {
  const whereClause = generateWhereClause();
  const event = new CustomEvent("apply-filters", { detail: whereClause });
  window.dispatchEvent(event);
};

const generateWhereClause = () => {
  const where = {};
  if (venue.value) {
    where.venue = { _ilike: `%${venue.value}%` };
  }
  if (start_date.value || end_date.value) {
    where.start_date = {};
    if (start_date.value) {
      where.start_date._gte = start_date.value;
    }
    where.end_date = {};
    if (end_date.value) {
      where.end_date._lte = end_date.value;
    }
  }
  if (selectedCategoryValue.value) {
    where.categories = { _eq: selectedCategoryValue.value };
  }
  if (selectedTagValues.value.length > 0) {
    where.tags = { _in: `{${selectedTagValues.value}}` };
  }
  if (isFree.value) {
    where.is_free = { _eq: isFree.value };
  }
  if (mapLocation.value.lat && mapLocation.value.lng) {
    where.location = {
      _ilike: `%${(mapLocation.value.lat, mapLocation.value.lng)}%`,
    };
  }

  return where;
};

const clearSelectedTags = () => {
  selectedTagValues.value = false;
  selectedTags.value = [];
};

const applyFilters = () => {
  emitFilterData();
  closeModal();
};

const resetFilters = () => {
  venue.value = "";
  start_date.value = "";
  end_date.value = "";
  selectedCategoryValue.value = "";
  selectedTagValues.value = [];
  minPriceRange.value = 0;
  maxPriceRange.value = 5000;
  mapLocation.value = { lat: null, lng: null };
  emitFilterData();
};

function closeModal() {
  emit("close-modal");
}

onMounted(() => {
  window.addEventListener("location-selected", handleLocationSelected);
});

onUnmounted(() => {
  window.removeEventListener("location-selected", handleLocationSelected);
});
</script>
