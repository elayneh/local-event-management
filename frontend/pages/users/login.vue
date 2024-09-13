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
const { onLogin } = useApollo();

function submitHandler(values) {
  const loginCredentials = {
    email: values.email,
    password: values.password,
  };
  mutate(loginCredentials);
  onDone((result) => {
    onLogin(result.data?.login?.token, "authClient");
    console.log("RESULT: ", result.data);
    const role = result.data.login.role;
    toast.success("User logged in successfully", {});
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
    toast.error("Invalid Email or Password", {});
  });
}
</script>
<template>
  <div class="flex min-h-screen relative overflow-hidden">
    <div class="absolute inset-0 z-0">
      <img
        src="~/assets/images/logo.png"
        class="w-full h-full object-cover"
        alt="Background Image"
      />
      <div class="absolute inset-0 bg-gray-800 bg-opacity-80"></div>
    </div>

    <div
      class="relative z-10 flex text-white flex items-center justify-center lg:mr-8 px-4 lg:px-8"
    >
      <div class="absolute right-0 z-30 animate-arrow-bounce">
        <div class="arrow-right"></div>
      </div>
      <div class="text-center">
        <h1 class="text-4xl font-bold mb-4 opacity-0 animate-fade-in">
          Welcome to Our Platform
        </h1>
        <p class="text-lg mb-6 opacity-0 animate-fade-in delay-500">
          Explore our features and enjoy a seamless experience. Sign in to
          access full features!!
        </p>
        <div class="opacity-0 animate-fade-in delay-1000">
          <button
            @click="navigateTo('/about')"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-lg transition transform hover:scale-105 shadow-lg"
          >
            Learn More
          </button>
        </div>
      </div>
    </div>

    <div class="door-wrapper flex items-center justify-center w-1/2">
      <div
        class="door relative z-20 transform-gpu door-animation p-8 lg:p-16 w-full max-w-lg shadow-lg rounded-lg -translate-x-5"
      >
        <CustomCard title="Login" class="bg-white p-6 rounded-lg shadow-lg">
          <template #body>
            <DynamicForm
              ref="DynamicForm"
              :schema="formSchema"
              :submitHandler="submitHandler"
            >
              <div>
                <p>
                  <NuxtLink
                    class="text-blue-500 hover:text-blue-700 font-semibold"
                    to="/users/forgot-password"
                  >
                    Forgot Password?
                  </NuxtLink>
                </p>
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
                    <p class="mb-2">
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
              </div>
            </DynamicForm>
          </template>
        </CustomCard>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    navigateTo(route) {
      this.$router.push(route);
    },
  },
};
</script>

<style scoped>
.door-wrapper {
  perspective: 800px;
}

.door {
  transform-origin: right center;
  transform: rotateY(90deg);
  transition: transform 1s ease-in-out;
}

.door-animation {
  animation: doorOpen 1.5s forwards ease-in-out;
}

@keyframes doorOpen {
  0% {
    transform: rotateY(90deg);
  }
  100% {
    transform: rotateY(0deg);
  }
}

/* Animated Arrow */
.arrow-right {
  width: 0;
  height: 0;
  border-top: 10px solid transparent;
  border-bottom: 10px solid transparent;
  border-left: 20px solid #fff;
}

@keyframes arrowBounce {
  0% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
  100% {
    transform: translateY(0);
  }
}

.animate-arrow-bounce {
  animation: arrowBounce 1s infinite;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

.animate-fade-in {
  animation: fadeIn 1s forwards;
}

.delay-500 {
  animation-delay: 0.5s;
}

.delay-1000 {
  animation-delay: 1s;
}
</style>
