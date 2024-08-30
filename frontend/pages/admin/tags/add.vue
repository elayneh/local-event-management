<template>
  <div class="flex h-screen items-center justify-center">
    <CustomCard title="Add Tag">
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
                Add Tag
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
import addtag from "~/graphql/mutations/tags/add.gql";
import { toast } from "vue3-toastify";
import { reactive } from "vue";

const {
  mutate: insertTagMutation,
  onDone,
  loading,
  onError,
} = useAuthenticatedMutation(addtag);
const formSchema = reactive({
  fields: [
    {
      name: "name",
      as: "input",
      placeholder: "Tag Name",
      rules: yup.string().required("Tag name is required"),
    },
    {
      name: "description",
      as: "textarea",
      placeholder: "Tag Description",
      rules: yup.string().required("Tag description is required"),
    },
  ],
});

async function submitHandler(values) {
  try {
    const tagData = {
      name: values.name,
      description: values.description,
    };

    insertTagMutation(tagData);
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      
    });
  }
}

onDone((result) => {
  toast.success("Tag added successfully", {
 
  });
  navigateTo("/tags");
});

onError((error) => {
  console.log("Error: ", error.message);
  toast.error("Something went wrong, try again", {
   
  });
});

definePageMeta({
  layout: "admin-dashboard",
});
</script>
