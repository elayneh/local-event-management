<template>
  <div class="max-w-lg mx-auto bg-white shadow-md rounded-lg p-6 mt-24">
    <h2 class="text-2xl font-bold mb-4">Update Tag</h2>

    <Form @submit="submitForm">
      <div class="mb-4">
        <label for="tag-name" class="block text-gray-700 font-bold mb-2">
          Tag Name
        </label>
        <Field
          name="name"
          type="text"
          class="w-full px-4 py-2 border rounded-lg"
          :class="{ 'border-red-500': errors.name }"
          placeholder="Enter tag name"
        />
        <ErrorMessage name="name" class="text-red-500 text-sm mt-1" />
      </div>

      <div class="mb-4">
        <label for="description" class="block text-gray-700 font-bold mb-2">
          Description
        </label>
        <Field
          name="description"
          as="textarea"
          class="w-full px-4 py-2 border rounded-lg"
          :class="{ 'border-red-500': errors.description }"
          placeholder="Enter description"
          rows="4"
        />
        <ErrorMessage name="description" class="text-red-500 text-sm mt-1" />
      </div>

      <button
        type="submit"
        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
      >
        Update Tag
      </button>
    </Form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useForm, ErrorMessage, Field, Form } from 'vee-validate';
import * as yup from 'yup';
import UpdateTagsByPk from '~/graphql/mutations/tags/update.gql';
import getSingleTag from '~/graphql/queries/tags/getSingleTag.gql';
import useAuthenticatedMutation from '~/composables/useAuthenticatedMutation';
import { useMutation } from '@vue/apollo-composable';
import { toast } from 'vue3-toastify';

// Define validation schema
const schema = yup.object({
  name: yup.string().required('Tag name is required'),
  description: yup
    .string()
    .required('Description is required')
    .min(10, 'Description must be at least 10 characters long'),
});

// Setup form
const { handleSubmit, setValues, errors } = useForm({
  validationSchema: schema,
});

// Apollo setup for queries and mutations
const route = useRoute();
const router = useRouter();
const tagId = route.params.id;
const editableTag = ref(null);

// Fetch the tag details
const { mutate: getTagMutation } = useMutation(getSingleTag);

onMounted(async () => {
  try {
    const { data } = await getTagMutation({ id: tagId });
    const tagData = data?.tags_by_pk;
    
    if (tagData) {
      editableTag.value = tagData;
      console.log('Fetched Tag Data:', tagData);
      
      // Set form values with a delay to ensure reactivity
      setTimeout(() => {
        setValues({
          name: tagData.name,
          description: tagData.description,
        });
        console.log('Form values set:', {
          name: tagData.name,
          description: tagData.description,
        });
      }, 100); // Delay to ensure reactivity
    } else {
      console.error('Tag data not found');
    }
  } catch (error) {
    console.error('Error fetching tag data:', error);
  }
});

// Mutation for updating the tag
const { mutate: updateTagMutation } = useAuthenticatedMutation(UpdateTagsByPk);

const submitForm = handleSubmit(async (values) => {
  console.log("Values: ", values)
  try {
    const { data } = await updateTagMutation({
      id: tagId,
      name: values.name,
      description: values.description,
    });
    
    if (data) {
      console.log('Update response:', data);
      toast.success('Tag updated successfully', {
        position: 'top-right',
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
      });
      router.push('/tags');
    } else {
      toast.error('Failed to update the tag', {
        position: 'top-right',
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
      });
    }
  } catch (error) {
    console.error('Error updating tag:', error);
    toast.error('An error occurred while updating the tag', {
      position: 'top-right',
      autoClose: 3000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
    });
  }
});
</script>
