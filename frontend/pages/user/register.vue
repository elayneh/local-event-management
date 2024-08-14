<template>
  <div class="flex h-screen items-center justify-center">
    <CustomCard title="Sign Up">
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
                Register
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
                Already have an account?
                <NuxtLink
                  class="text-blue-500 hover:text-blue-700 font-semibold ml-1"
                  to="/authentication/login"
                  ><span>Login</span></NuxtLink
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
import register from "~/graphql/mutations/user/register.gql";
import { toast } from "vue3-toastify";
import * as JsCookie from "js-cookie";
import { reactive } from "vue";

const Cookies = JsCookie.default;

const authenticationStore = useAuthStore();
const { mutate, onDone, result, loading, onError } = authentication(register);

const formSchema = reactive({
  fields: [
    {
      name: "first_name",
      as: "input",
      placeholder: "First Name",
      rules: yup.string().required("First name is required"),
    },
    {
      name: "last_name",
      as: "input",
      placeholder: "Last Name",
      rules: yup.string().required("Last name is required"),
    },
    {
      name: "role",
      as: "input",
      placeholder: "Role",
      rules: yup.string(),
    },
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
    {
      placeholder: "Confirm Password",
      name: "cpassword",
      as: "input",
      type: "password",
      rules: yup
        .string()
        .min(6, "Passwords must match")
        .required("Confirm Password is required"),
    },
  ],
});

function submitHandler(values) {
  console.log("Add product form data: ", values);
  const formData = {
    first_name: values.first_name,
    last_name: values.last_name,
    email: values.email,
    password: values.password,
    role: values.role,
  };
  mutate(formData);
}

onDone((result) => {
  const { id, role, token } = result.data.register;
  toast.success("User registered successfully", {
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
</script>
