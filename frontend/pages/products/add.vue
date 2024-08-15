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

async function fileToBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result.split(",")[1]);
    reader.onerror = (error) => reject(error);
  });
}

async function submitHandler(values) {
  try {
    const imageFile = values.image;
    const base64Image = await fileToBase64(imageFile);

    const productData = {
      name: values.name,
      description: values.description,
      price: String(values.price),
      category: values.category,
      image: base64Image,
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

const { mutate, onDone, loading, onError } =
  useAuthenticatedMutation(addproduct);

onDone((result) => {
  toast.success("Product added successfully", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
  navigateTo("/products");
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
