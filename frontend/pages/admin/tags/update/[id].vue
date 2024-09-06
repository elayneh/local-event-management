<script setup>
import { ref, onMounted } from "vue";
import * as yup from "yup";
import { useMutation } from "@vue/apollo-composable";
import { toast } from "vue3-toastify";
import { Form, Field } from "vee-validate";
import { useRoute, useRouter } from "vue-router";
import getSingleTag from "~/graphql/queries/tags/getSingleTag.gql";
import UpdateTagsByPk from "~/graphql/mutations/tags/update.gql";

const route = useRoute();
const router = useRouter();
const tagId = route.params.id;

const editableTag = ref({});

const schema = yup.object({
  name: yup.string().required("Tag name is required"),
  description: yup
    .string()
    .required("Description is required")
    .min(10, "Description must be at least 10 characters long"),
});

// Fetch the tag details on component mount
const { mutate: getTagMutation } = useMutation(getSingleTag);
onMounted(async () => {
  try {
    const { data } = await getTagMutation({ id: tagId });
    const tagData = data?.tags_by_pk;

    if (tagData) {
      editableTag.value = tagData;
    } else {
      console.error("Tag data not found");
    }
  } catch (error) {
    console.error("Error fetching tag data:", error);
    toast.error("Failed to load tag data");
  }
});

const { mutate: updateTagMutation } = useMutation(UpdateTagsByPk);

const saveTagChanges = async (values) => {
  try {
    const { data } = await updateTagMutation({
      id: tagId,
      name: values.name,
      description: values.description,
    });

    if (data) {
      toast.success("Tag updated successfully");
      router.push("/admin/tags");
    } else {
      toast.error("Failed to update the tag");
    }
  } catch (err) {
    toast.error("Error updating the tag, try again");
    console.error("Error updating tag: ", err);
  }
};

const cancelEdit = () => {
  router.push("/admin/tags");
};

definePageMeta({
  layout: "admin-dashboard",
});
</script>

<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 w-full min-h-screen flex justify-center items-start"
  >
    <div
      class="bg-white shadow-xl rounded-lg p-8 w-full max-w-4xl mt-24 mx-8 transform transition-transform duration-300 ease-in-out hover:scale-105"
    >
      <div class="flex justify-end">
        <button @click="navigateTo('/admin/tags')" aria-label="Close">
          <font-awesome-icon :icon="['fas', 'times']" />
        </button>
      </div>
      <div class="flex flex-col justify-center items-center">
        <div
          class="w-16 h-16 bg-gray-700 rounded-full flex items-center justify-center focus:outline-none focus:ring-2 focus:ring-gray-500"
        >
          <font-awesome-icon icon="tag" class="text-gray-500 w-full h-full" />
        </div>
        <div class="text-center mt-4">
          <h1 class="text-2xl font-bold">Update Tag</h1>
        </div>
        <div class="mt-12 w-full">
          <Form
            :validation-schema="schema"
            @submit="saveTagChanges"
            v-slot="{ errors }"
          >
            <div class="pl-8">
              <label class="text-gray-800 font-semibold">Tag Name</label>
              <Field
                name="name"
                v-model="editableTag.name"
                class="border border-gray-300 rounded p-2 w-full mt-2"
                type="text"
              />
              <span class="text-red-500">{{ errors.name }}</span>
            </div>
            <div class="pl-8 mt-6">
              <label class="text-gray-800 font-semibold">Description</label>
              <Field
                name="description"
                v-model="editableTag.description"
                class="border border-gray-300 rounded p-2 w-full mt-2"
                type="text"
              />
              <span class="text-red-500">{{ errors.description }}</span>
            </div>
            <div class="mt-12 flex justify-center items-center">
              <button
                type="submit"
                class="bg-green-600 text-white px-6 py-2 rounded-lg shadow hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 mr-32"
              >
                Save
              </button>
              <button
                type="button"
                @click="cancelEdit"
                class="bg-red-600 text-white px-6 py-2 rounded-lg shadow hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
              >
                Cancel
              </button>
            </div>
          </Form>
        </div>
      </div>
    </div>
  </div>
</template>
