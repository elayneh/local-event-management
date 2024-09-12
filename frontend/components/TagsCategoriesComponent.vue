<template>
  <div>
    <div v-if="!categories.length && !tags.length" class="text-center py-16">
      <p class="text-lg text-gray-500">No categories or tags available.</p>
    </div>

    <div
      class="bg-gradient-to-r from-violet-200 to-blue-200 mt-8 grid grid-cols-1 lg:grid-cols-2 gap-10 mx-6 lg:mx-12 p-6 rounded-lg shadow-lg"
      v-else
    >
      <div class="px-4">
        <h2 class="text-2xl font-extrabold mb-4 text-center text-violet-800">
          Categories
        </h2>
        <div class="flex flex-wrap gap-3 justify-center">
          <span
            v-for="(category, index) in categories"
            :key="index"
            class="bg-gradient-to-r from-blue-100 to-blue-200 text-blue-800 px-4 py-2 rounded-full text-sm font-bold shadow-md hover:shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
          >
            {{ category.name }}
          </span>
        </div>
      </div>

      <div class="px-4">
        <h2 class="text-2xl font-extrabold mb-4 text-center text-green-800">
          Tags
        </h2>
        <div class="flex flex-wrap gap-3 justify-center">
          <span
            v-for="(tag, index) in tags"
            :key="index"
            class="bg-gradient-to-r from-green-100 to-green-200 text-green-800 px-4 py-2 rounded-full text-sm font-bold shadow-md hover:shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
          >
            {{ tag.name }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  categories: {
    type: Array,
    default: () => [],
  },
  tags: {
    type: Array,
    default: () => [],
  },
});
</script>
