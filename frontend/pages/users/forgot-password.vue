<script setup>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import RequestPasswordReset from "~/graphql/mutations/users/requestPasswordReset.gql";
import { toast } from "vue3-toastify";

import { useRouter } from "vue-router";

const router = useRouter();

const email = ref("");

const {
  mutate: requestPasswordResetMutation,
  onDone,
  onError,
} = useMutation(RequestPasswordReset);

const requestPasswordReset = async () => {
  try {
    const response = await requestPasswordResetMutation({ email: email.value });
  } catch (error) {
    console.error("Password reset request failed:", error);
  }
};

const handleClose = () => {
  router.push("/users/login");
};

onDone((response) => {
  toast.success(response.data.requestPasswordReset.message, {
    position: "top-right",
    autoClose: 3000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
  });

  navigateTo("/users/login");
});

onError((error) => {
  toast.error("Error sending password reset email", {
    position: "top-right",
    autoClose: 3000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
  });
});
</script>

<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="absolute w-full animate-slide-in">
      <div
        class="flex flex-col items-center justify-center h-full p-32 text-center w-1/2"
      >
        <h1 class="text-5xl font-bold mb-4">Welcome Back!</h1>
        <p class="text-lg mb-6">
          Please enter your email to reset your password. We're here to help you
          get back on track.
        </p>
        <img
          src="~/assets/images/wellcome.jpg"
          alt="Welcome Image"
          class="w-32 h-32 object-cover rounded-full shadow-lg"
        />
      </div>
    </div>

    <div
      class="relative bg-white p-8 ml-64 rounded-lg shadow-xl w-full max-w-md z-20 transform rotate-1 translate-x-12 transition-transform duration-500 hover:translate-x-0 hover:rotate-0 border-2 border-gray-300 hover:border-gray-400"
    >
      <button
        @click="handleClose"
        class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 pr-4 pt-2"
      >
        <i class="fas fa-times text-2xl"></i>
      </button>
      <h2 class="text-2xl font-bold mb-6 text-center text-gray-700">
        Reset Password
      </h2>
      <form @submit.prevent="requestPasswordReset" class="space-y-4">
        <input
          v-model="email"
          type="email"
          placeholder="Enter your email"
          class="w-full p-4 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-300 hover:border-blue-500 hover:ring-blue-300"
          required
        />
        <button
          type="submit"
          class="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md transition-transform duration-300 transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          Reset Password
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped>
@keyframes slideIn {
  from {
    transform: translateX(-100%);
  }
  to {
    transform: translateX(0);
  }
}

.animate-slide-in {
  animation: slideIn 1s ease-out;
}
</style>
