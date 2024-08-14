<template>
  <Form @submit="submitHandler" class="bg-white px-8 pt-6 pb-8 mb-4">
    <div
      v-for="{
        as,
        name,
        label,
        placeholder,
        children,
        ...attrs
      } in schema.fields"
      :key="name"
      class="mb-4"
    >
      <label :for="name" class="block text-gray-700 text-sm font-bold mb-2">
        {{ label }}
      </label>
      <Field
        :as="as"
        :id="name"
        :name="name"
        :placeholder="placeholder"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        v-bind="attrs"
      >
        <template v-if="children && children.length">
          <component
            v-for="({ tag, text, ...childAttrs }, idx) in children"
            :key="idx"
            :is="tag"
            v-bind="childAttrs"
          >
            {{ text }}
          </component>
        </template>
      </Field>
      <ErrorMessage :name="name" class="text-red-500 text-xs mt-1" />
    </div>
    <slot />
  </Form>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";

export default {
  name: "DynamicForm",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  props: {
    schema: {
      type: Object,
      required: true,
    },
    submitHandler: {
      type: Function,
      required: true,
    },
  },
};
</script>
