<template>
  <div class="mb-4">
    <label :for="name" class="block text-gray-700 font-bold mb-2">
      {{ label }}
    </label>
    <Field
      v-bind="fieldProps"
      :name="name"
      :type="type"
      :as="as"
      :placeholder="placeholder"
      :class="[
        'w-full px-4 py-2 border rounded-lg',
        { 'border-red-500': hasError },
      ]"
      v-model="modelValue"
    />
    <ErrorMessage :name="name" class="text-red-500 text-sm mt-1" />
  </div>
</template>

<script setup>
import { computed, toRefs } from "vue";
import { Field, ErrorMessage } from "vee-validate";

const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  label: {
    type: String,
    required: true,
  },
  type: {
    type: String,
    default: "text",
  },
  as: {
    type: String,
    default: "input", // Can be 'input', 'textarea', etc.
  },
  placeholder: {
    type: String,
    default: "",
  },
  modelValue: {
    type: String,
    default: "",
  },
  fieldProps: {
    type: Object,
    default: () => ({}),
  },
});

const { name, modelValue, fieldProps } = toRefs(props);

// Check if the field has an error
const hasError = computed(
  () =>
    !!fieldProps.value.errors && fieldProps.value.errors[name.value]?.length > 0
);
</script>
