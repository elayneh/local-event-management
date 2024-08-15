<template>
  <div class="flex h-screen items-center justify-center">
    <CustomCard title="Login">
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
                Login
              </button>
              <button
                type="button"
                @click="clearForm"
                class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              >
                Clear
              </button>
            </div>
            <div>
              <p class="mt-10 text-gray-700 text-center">
                Don't have an account?
                <NuxtLink
                  class="text-blue-500 hover:text-blue-700 font-semibold ml-1"
                  to="/authentication/register"
                  ><span>Register</span></NuxtLink
                >
              </p>
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
import login from "~/graphql/mutations/user/login.gql";
import { toast } from "vue3-toastify";
import * as JsCookie from "js-cookie";

const Cookies = JsCookie.default;

const authenticationStore = useAuthStore();
const { mutate, onDone, loading, onError } = authentication(login);

const formSchema = {
  fields: [
    {
      name: "email",
      as: "input",
      placeholder: "Email",
      rules: yup
        .string()
        .email("Invalid email address")
        .required("Email is required"),
    },
    {
      name: "password",
      as: "input",
      type: "password",
      placeholder: "Password",
      rules: yup
        .string()
        .min(6, "Password must be at least 6 characters")
        .required("Password is required"),
    },
  ],
};

function submitHandler(values) {
  const loginCredentials = {
    email: values.email,
    password: values.password,
  };
  console.log("Login form data: ", loginCredentials);
  mutate(loginCredentials);
  onDone((result) => {
    toast.success("User logged in successfully", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    Cookies.set("token", result.data.login.token, { expires: 3 });
    authenticationStore.setToken(result.data.login.token);
    authenticationStore.setId(result.data.login.id);
    authenticationStore.setUser(result.data.login.id);
    authenticationStore.setRole(result.data.login.role);
    navigateTo("/products");
  });
  onError((error) => {
    console.log("Error: ", error.message);
    toast.error("Something want wrong try again", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  });
}
</script>
