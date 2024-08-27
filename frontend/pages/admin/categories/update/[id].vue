<template>
  <div class="flex h-screen items-center justify-center">
    <CustomCard title="Update Product">
      <template #body>
        <DynamicForm
          ref="DynamicForm"
          :schema="formSchema"
          :submitHandler="updateHandler"
        >
          <div>
            <div class="flex items-center justify-center gap-4 mt-4">
              <button
                type="submit"
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              >
                Update
              </button>
              <button
                type="button"
                @click="cancelHandler"
                class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              >
                Cancel
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
import { ref, reactive, onMounted } from "vue";
import { toast } from "vue3-toastify";

import CustomCard from "~/components/CustomCard.vue";
import DynamicForm from "~/components/DynamicForm.vue";
import updateCategory from "~/graphql/mutations/categories/update.gql";
import getSingleCategory from "~/graphql/queries/categories/fetchCategoryById.gql";

const route = useRoute();
const categoryId = route.params.id;
const categoryData = ref(null);

const formSchema = reactive({
  fields: [
    {
      name: "name",
      as: "input",
      value: "", // The value will be updated with the fetched category data
      placeholder: "Category Name",
      rules: yup.string().required("Category name is required"),
    },
    {
      name: "description",
      as: "textarea",
      value: "", // The value will be updated with the fetched category data
      placeholder: "Category Description",
      rules: yup.string().required("Category description is required"),
    },
  ],
});

// Fetch category data on mount
async function fetchCategory() {
  try {
    const { result, error } = await useQuery(getSingleCategory, {
      id: categoryId,
    });

    if (error.value) {
      throw new Error("Failed to fetch category data");
    }

    if (result.value && result.value.category) {
      categoryData.value = result.value.category;

      // Update the form schema with the fetched data
      formSchema.fields[0].value = categoryData.value.name || "";
      formSchema.fields[1].value = categoryData.value.description || "";
    }
  } catch (error) {
    console.error("Error:", error.message);
    toast.error("Error fetching category data", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
}

onMounted(() => {
  fetchCategory();
});

const { mutate, onDone, onError } = useAuthenticatedMutation(updateCategory);

async function updateHandler(values) {
  try {
    const categoryPayload = {
      id: categoryId,
      name: values.name,
      description: values.description,
    };
    await mutate(categoryPayload);
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
}

onDone(() => {
  toast.success("Category updated successfully", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
  navigateTo("/admin/categories");
});

onError((error) => {
  console.error("Error:", error.message);
  toast.error("Something went wrong, try again", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});

definePageMeta({
  layout: "admin-dashboard",
});
</script>
