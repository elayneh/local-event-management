<template>
  <div
    :class="[
      'bg-white w-full max-w-md mx-auto shadow-lg rounded-lg overflow-hidden flex flex-col',
      customClass,
    ]"
  >
    <img
      v-if="image"
      :src="image"
      alt="Card Image"
      class="w-full h-48 object-cover"
    />

    <div class="bg-gray-800 text-white text-center py-4">
      <h2 class="text-2xl font-semibold">{{ title }}</h2>
    </div>
    <div class="flex-1 p-6">
      <slot name="body"></slot>
    </div>
    <div class="p-6 border-t border-gray-200 flex justify-between">
      <slot name="footer-left">
        <!-- Default content for the left side if no slot is provided -->
      </slot>
      <slot name="footer-right">
        <!-- Default content for the right side if no slot is provided -->
      </slot>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  image: {
    type: String,
    default: null,
  },
  customClass: {
    type: String,
    default: "", // Default to empty string if not provided
  },
});
</script>
