<template>
  <div class="flex h-screen items-center justify-center">
    <CustomCard title="File Upload">
      <template #body>
        <DynamicForm
          ref="DynamicForm"
          :schema="fileUploadSchema"
          :submitHandler="submitHandler"
        >
          <div>
            <div class="flex items-center justify-center gap-6 mt-4">
              <button
                type="submit"
                class="bg-green-300 hover:bg-green-500 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              >
                Upload
              </button>
              <button
                type="button"
                @click="navigateTo('/home')"
                class="bg-gray-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
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
import CustomCard from "~/components/CustomCard.vue";
import DynamicForm from "~/components/DynamicForm.vue";
import { useAuthStore } from "~/stores";
import fileUpload from "~/graphql/mutations/fileUpload.gql";
import { toast } from "vue3-toastify";
import * as JsCookie from "js-cookie";
import { reactive } from "vue";

const Cookies = JsCookie.default;

const authenticationStore = useAuthStore();
const { mutate, onDone, result, loading, onError } = authentication(fileUpload);

const fileUploadSchema = reactive({
  fields: [
    {
      name: "fileUpload",
      as: "input",
      placeholder: "Upload file",
      type: "file",
    },
  ],
});

function submitHandler(values) {
  console.log("File uploaded: ", values.fileUpload);
  const fileUploadData = {
    fileUploaded: values.fileUpload,
  };
  mutate(fileUploadData);
}

onDone((result) => {
  const { id, role, token } = result.data.fileUpload;
  toast.success("File uploaded successfully", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
  Cookies.set("token", token, { expires: 3 });
  authenticationStore.setToken(token);
  authenticationStore.setId(token);
  authenticationStore.setUser(id);
  authenticationStore.setRole(role);
  navigateTo("/home");
});
onError((error) => {
  console.log("Error: ", error.message);
  toast.error("Something went wrong, try again", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});

definePageMeta({
  layout: "authenticated",
});
</script>
