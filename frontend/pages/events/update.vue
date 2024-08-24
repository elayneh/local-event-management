<template>
  <DynamicForm
    :schema="eventFormSchema"
    @submit="submitEvent"
    class="bg-white px-8 pt-6 pb-8 mb-4"
  >
    <button
      type="submit"
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
    >
      Submit Event
    </button>
  </DynamicForm>
</template>

<script setup>
import { ref } from "vue";
import DynamicForm from "./DynamicForm.vue"; // Assuming your dynamic form is saved as DynamicForm.vue

const submitEvent = (values) => {
  console.log("Event Submitted:", values);
};

const eventFormSchema = ref({
  fields: [
    {
      name: "title",
      label: "Event Title",
      placeholder: "Enter event title",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "description",
      label: "Event Description",
      placeholder: "Enter event description",
      type: "textarea",
      as: "textarea",
      attrs: {
        required: true,
      },
    },
    {
      name: "category_id",
      label: "Category",
      placeholder: "Select category",
      as: "select",
      options: [
        { value: "", label: "Select a category" },
        { value: "music", label: "Music" },
        { value: "food", label: "Food" },
        { value: "sports", label: "Sports" },
      ],
      attrs: {
        required: true,
      },
    },
    {
      name: "price",
      label: "Event Price",
      placeholder: "Enter event price (leave empty if free)",
      type: "number",
      attrs: {
        min: 0,
        step: 0.01,
      },
    },
    {
      name: "is_free",
      label: "Is the event free?",
      as: "select",
      options: [
        { value: "false", label: "No" },
        { value: "true", label: "Yes" },
      ],
      attrs: {
        required: true,
      },
    },
    {
      name: "venue",
      label: "Venue",
      placeholder: "Enter venue name",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "address",
      label: "Address",
      placeholder: "Enter address",
      type: "text",
      attrs: {
        required: true,
      },
    },
    {
      name: "start_date",
      label: "Start Date",
      placeholder: "Select start date",
      type: "datetime-local",
      attrs: {
        required: true,
      },
    },
    {
      name: "end_date",
      label: "End Date",
      placeholder: "Select end date",
      type: "datetime-local",
      attrs: {
        required: true,
      },
    },
    {
      name: "featured_image",
      label: "Featured Image",
      placeholder: "Upload a featured image",
      type: "file",
      attrs: {
        accept: "image/*",
      },
    },
    {
      name: "tags",
      label: "Tags",
      placeholder: "Enter tags (comma separated)",
      type: "text",
    },
  ],
});
</script>
