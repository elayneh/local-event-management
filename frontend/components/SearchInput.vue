<template>
  <div class="relative">
    <input
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      @change="$emit('update:modelValue', $event.target.value)"
      class="w-full h-10 rounded-full pl-4 pr-10 bg-gray-300 flex items-center focus:outline-none focus:ring-1 focus:ring-gray-400"
      placeholder="Search..."
    />
    <font-awesome-icon
      :icon="['fas', 'search']"
      class="text-gray-600 absolute right-4"
    />
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';

// Props and Emits
const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
});
const emit = defineEmits(['update:modelValue']);
</script>

<style scoped>
/* You can add any additional styling here */
</style>
