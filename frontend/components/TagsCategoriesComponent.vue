<template>
  <div>
    <div v-if="!categories.length && !tags.length" >
    </div>

    <div
      class="bg-violet-200 mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6 mx-12"
      v-else
    >
      <div class="px-4">
        <h2 class="text-xl font-bold mb-2 text-center">Categories</h2>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(category, index) in categories"
            :key="index"
            class="bg-blue-100 text-blue-700 px-3 py-1 rounded-full text-sm font-semibold"
          >
            {{ category.name }}
          </span>
        </div>
      </div>

      <div>
        <h2 class="text-xl font-bold mb-2 text-center">Tags</h2>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(tag, index) in tags"
            :key="index"
            class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-sm font-semibold"
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
