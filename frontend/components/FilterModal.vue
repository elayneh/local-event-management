<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
  >
    <div
      class="bg-white p-6 rounded-lg shadow-lg max-w-lg w-full space-y-4"
    >
      <h2 class="text-xl font-semibold">Filter Events</h2>
      <div class="space-y-2">
        <!-- Loop through the event fields and display checkboxes -->
        <div v-for="field in eventFields" :key="field">
          <label class="flex items-center space-x-2">
            <input
              type="checkbox"
              :value="field"
              v-model="selectedFields"
              class="form-checkbox text-blue-500"
            />
            <span>{{ field }}</span>
          </label>
        </div>
      </div>
      <div class="flex justify-end space-x-4">
        <button
          @click="applyFilter"
          class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-600"
        >
          Apply Filter
        </button>
        <button
          @click="closeModal"
          class="bg-gray-300 text-gray-700 font-bold py-2 px-4 rounded hover:bg-gray-400"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// Props passed from the parent component
const props = defineProps({
  isOpen: Boolean,
});

// Emit event back to the parent component
const emit = defineEmits(["close-modal", "apply-filter"]);

// Dummy event fields
const eventFields = ref([
  "Title",
  "Location",
  "Date",
  "Category",
  "Organizer",
]);

const selectedFields = ref([]);

const applyFilter = () => {
  // Emit the selected fields back to the parent
  emit("apply-filter", selectedFields.value);
  closeModal();
};

const closeModal = () => {
  emit("close-modal");
};
</script>
