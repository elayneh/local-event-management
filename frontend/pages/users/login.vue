<template>
  <div
    class="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50"
  >
    <div class="bg-white p-6 rounded-lg shadow-lg w-full max-w-md">
      <CustomCard title="Login">
        <template #body>
          <DynamicForm
            ref="DynamicForm"
            :schema="formSchema"
            :submitHandler="submitHandler"
          >
            <div class="mt-24">
              <div class="flex items-center justify-center gap-4 mt-4">
                <button
                  type="submit"
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                >
                  Login
                </button>
                <button
                  type="button"
                  @click="navigateTo('/')"
                  class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                >
                  Cancel
                </button>
              </div>
              <div class="mt-10 text-gray-700 text-center">
                <p>
                  Don't have an account?
                  <NuxtLink
                    class="text-blue-500 hover:text-blue-700 font-semibold ml-1"
                    to="/users/register"
                  >
                    Register
                  </NuxtLink>
                </p>
              </div>
            </div>
          </DynamicForm>
        </template>
      </CustomCard>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import CustomCard from "~/components/CustomCard.vue";
import DynamicForm from "~/components/DynamicForm.vue";
import { useAuthStore } from "~/stores";
import login from "~/graphql/mutations/users/login.gql";
import { toast } from "vue3-toastify";
import * as JsCookie from "js-cookie";
import * as yup from "yup";

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
  mutate(loginCredentials);
  onDone((result) => {
    const role = result.data.login.role;
    toast.success("User logged in successfully", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    Cookies.set("token", result.data.login.token, { expires: 3 });
    authenticationStore.setToken(result.data.login.token);
    authenticationStore.setId(result.data.login.id);
    authenticationStore.setUser(result.data.login.id);
    authenticationStore.setRole(role);
    role === "user"
      ? navigateTo("/user")
      : role === "admin"
      ? navigateTo("/admin")
      : "/";
  });
  onError((error) => {
    console.log("Error: ", error.message);
    toast.error("Invalid Email or Password", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  });
}
</script>
