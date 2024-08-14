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
                Add Product
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
import addproduct from "~/graphql/mutations/products/add.gql";
import { toast } from "vue3-toastify";
import { reactive } from "vue";
import axios from "axios";

async function uploadImage(image) {
  try {
    if (image) {
      const formData = new FormData();
      formData.append("file", image); // Ensure the key is 'file'

      console.log("FormData content:", formData.get("file")); // Log to verify the content

      const response = await axios.post(
        "http://localhost:5050/uploadFile",
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data", // Ensure this is set
          },
        }
      );

      return response.data.url; // Return URL from the server response
    } else {
      throw new Error("No file provided");
    }
  } catch (error) {
    console.error(
      "Error:",
      error.response ? error.response.data : error.message
    );
    throw new Error("File upload failed: " + error.message);
  }
}

// Using the custom authentication hook with the addProduct mutation
const { mutate, onDone, loading, onError } =
  useAuthenticatedMutation(addproduct);

const formSchema = reactive({
  fields: [
    {
      name: "name",
      as: "input",
      placeholder: "Product Name",
      rules: yup.string().required("Product name is required"),
    },
    {
      name: "description",
      as: "textarea",
      placeholder: "Product Description",
      rules: yup.string().required("Product description is required"),
    },
    {
      name: "price",
      as: "input",
      type: "number",
      placeholder: "Price",
      rules: yup.number().required("Price is required"),
    },
    {
      name: "category",
      as: "input",
      placeholder: "Category",
      rules: yup.string().required("Category is required"),
    },
    {
      name: "image",
      as: "input",
      type: "file",
      placeholder: "Product Image",
      rules: yup.mixed().required("Product image is required"),
    },
  ],
});

// Function to handle form submission
async function submitHandler(values) {
  try {
    const imageUrl = await uploadImage(values.image);

    const productData = {
      name: values.name,
      description: values.description,
      price: parseFloat(values.price),
      category: values.category,
      image: imageUrl,
    };

    // Execute the mutation with the product data
    mutate(productData);
  } catch (error) {
    console.error("Error:", error);
    toast.error("Something went wrong, try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
}

// Handle successful mutation
onDone((result) => {
  toast.success("Product added successfully", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
  // Redirect or perform additional actions after product is added
  navigateTo("/products");
});

// Handle mutation errors
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
