<template>
  <div
    class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-r from-blue-100 via-green-100 to-yellow-100"
  >
    <CustomCard title="Sign Up">
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
              <p class="mt-5 text-gray-700 text-center">
                Already have an account?
                <NuxtLink
                  class="text-blue-500 hover:text-blue-700 font-semibold ml-1"
                  to="/users/login"
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
import register from "~/graphql/mutations/users/register.gql";
import { toast } from "vue3-toastify";
import * as JsCookie from "js-cookie";
import { reactive } from "vue";

const Cookies = JsCookie.default;

const authenticationStore = useAuthStore();
const {
  mutate: userRegistration,
  onDone,
  result,
  loading,
  onError,
} = authentication(register);

const formSchema = reactive({
  fields: [
    {
      name: "username",
      as: "input",
      placeholder: "Username",
      label: "Username",
      rules: yup.string().required("Username is required"),
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
        .oneOf([yup.ref("password")], "Passwords must match")
        .required("Confirm Password is required"),
    },
    {
      name: "role",
      as: "select",
      placeholder: "Role",
      options: [
        { label: "Select a role", value: "" },
        { label: "User", value: "user" },
      ],
      rules: yup
        .string()
        .oneOf(["user"], "Please select a valid role")
        .required("Role is required"),
    },
  ],
});

function submitHandler(values) {
  console.log("User form data: ", values);
  const formData = {
    username: values.username,
    email: values.email,
    password: values.password,
    role: values.role,
  };
  userRegistration(formData);
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
  role == "user"
    ? navigateTo("/user")
    : role == "admin"
    ? navigateTo("/admin")
    : "/";
});
onError((error) => {
  console.log("Error: ", error.message);
  toast.error("Something went wrong, try again", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});
</script>
