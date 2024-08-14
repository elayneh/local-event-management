<template>
  <div
    class="bg-white w-full max-w-md mx-auto shadow-lg rounded-lg overflow-hidden"
  >
    <div class="bg-gray-800 text-white text-center py-4">
      <h2 class="text-2xl font-semibold">{{ title }}</h2>
    </div>
    <div class="p-6">
      <slot name="body"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "CustomCard",
  props: {
    title: {
      type: String,
      required: true,
    },
  },
};
</script>
