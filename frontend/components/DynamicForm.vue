<template>
  <Form @submit="submitHandler" class="bg-white px-8 pt-6 pb-8 mb-4">
    <div v-for="field in schema.fields" :key="field.name" class="mb-4">
      <label
        :for="field.name"
        class="block text-gray-700 text-sm font-bold mb-2"
      >
        {{ field.label }}
      </label>

      <Field
        :as="field.as === 'select' ? 'select' : 'input'"
        :id="field.name"
        :name="field.name"
        :placeholder="field.placeholder"
        :type="field.type"
        :disabled="field.disabled"
        :multiple="field.multiple"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        v-bind="field.attrs"
        @change="handleFileChange($event, field)"
      >
        <template v-if="field.as === 'select'">
          <option
            v-for="option in field.options"
            :key="option.value"
            :value="option.value"
          >
            {{ option.label }}
          </option>
        </template>

        <template
          v-if="
            field.as !== 'select' && field.children && field.children.length
          "
        >
          <component
            v-for="({ tag, text, ...childAttrs }, idx) in field.children"
            :key="idx"
            :is="tag"
            v-bind="childAttrs"
          >
            {{ text }}
          </component>
        </template>
      </Field>

      <template v-if="field.as === 'select' && selectedItems.length">
        <div class="mt-2">
          <span
            v-for="item in selectedItems"
            :key="item.value"
            class="bg-blue-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
          >
            {{ item.label }}
            <button
              @click="removeItem(item)"
              class="ml-2 text-white hover:text-gray-200"
            >
              ×
            </button>
          </span>
        </div>
      </template>

      <ErrorMessage :name="field.name" class="text-red-500 text-xs mt-1" />
    </div>
    <slot />
  </Form>
</template>

<script setup>
import { Form, Field, ErrorMessage } from "vee-validate";
import { ref } from "vue";

const selectedItemValues = ref([]);
const selectedItems = ref([]);

const removeItem = (option) => {
  console.log(
    "Before removing:",
    selectedItems.value,
    selectedItemValues.value
  );
  selectedItems.value = selectedItems.value.filter(
    (current) => current.value !== option.value
  );
  selectedItemValues.value = selectedItemValues.value.filter(
    (val) => val !== option.value
  );
  console.log("After removing:", selectedItems.value, selectedItemValues.value);
};

const props = defineProps({
  schema: {
    type: Object,
    required: true,
    default: () => ({ fields: [] }),
  },
  submitHandler: {
    type: Function,
    required: true,
  },
});

const selectedFiles = ref({});

const handleFileChange = (event, field) => {
  if (field.type === "file" && field.multiple) {
    selectedFiles.value[field.name] = Array.from(event.target.files);
  }
};
</script>
