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
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        v-bind="field.attrs"
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
      <ErrorMessage :name="field.name" class="text-red-500 text-xs mt-1" />
    </div>
    <slot />
  </Form>
</template>
<script setup>
import { Form, Field, ErrorMessage } from "vee-validate";

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
</script>
