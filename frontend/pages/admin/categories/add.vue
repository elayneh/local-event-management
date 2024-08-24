<template>
  <div class="flex h-screen items-center justify-center">
    <CustomCard title="Add Product">
      <template #body>
        <DynamicForm
          ref="DynamicForm"
          :schema="formSchema"
          :submitHandler="submitHandler"
        >
          <div>
            <div class="flex items-center justify-center gap-4 mt-4">
              <button
                type="submit"
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              >
                Add Category
              </button>
              <button
                type="button"
                @click="clearForm"
                class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              >
                Clear
              </button>
            </div>
          </div>
        </DynamicForm>
      </template>
    </CustomCard>
  </div>
</template>

<script setup>
import * as yup from "yup";
import CustomCard from "~/components/CustomCard.vue";
import DynamicForm from "~/components/DynamicForm.vue";
import addcategory from "~/graphql/mutations/categories/add.gql";
import { toast } from "vue3-toastify";
import { reactive } from "vue";

const formSchema = reactive({
  fields: [
    {
      name: "name",
      as: "input",
      placeholder: "Category Name",
      rules: yup.string().required("Category name is required"),
    },
    {
      name: "description",
      as: "textarea",
      placeholder: "Category Description",
      rules: yup.string().required("Category description is required"),
    },
  ],
});

async function submitHandler(values) {
  try {

    const categoryData = {
      name: values.name,
      description: values.description,
    };

    mutate(categoryData);
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
}

const { mutate, onDone, loading, onError } =
  useAuthenticatedMutation(addcategory);

onDone((result) => {
  toast.success("Category added successfully", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
  navigateTo("/categories");
});

onError((error) => {
  console.log("Error: ", error.message);
  toast.error("Something went wrong, try again", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});

definePageMeta({
  layout: "admin-dashboard",
});
</script>
