<script setup>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import ResetPassword from "~/graphql/mutations/users/resetPassword.gql";
import { toast } from "vue3-toastify";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();
const resetToken = route.query.reset_token || "";

const form = ref({
  newPassword: "",
  confirmPassword: "",
});

const errors = ref({
  newPassword: "",
  confirmPassword: "",
});

const { mutate: resetPasswordMutation } = useMutation(ResetPassword);

const validateForm = () => {
  errors.value = {
    newPassword: "",
    confirmPassword: "",
  };

  let valid = true;

  if (!form.value.newPassword) {
    errors.value.newPassword = "New password is required";
    valid = false;
  }

  if (!form.value.confirmPassword) {
    errors.value.confirmPassword = "Please confirm your password";
    valid = false;
  }

  if (form.value.newPassword !== form.value.confirmPassword) {
    errors.value.confirmPassword = "Passwords do not match";
    valid = false;
  }

  return valid;
};

const handleSubmit = async () => {
  if (validateForm()) {
    try {
      const response = await resetPasswordMutation({
        reset_token: resetToken,
        new_password: form.value.newPassword,
      });
      toast.success(response.data.updatePassword.message, {
        position: "top-right",
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
      });
      router.push("/users/login");
    } catch (error) {
      console.error("Password reset request failed:", error);
      toast.error("Error updating password", {
        position: "top-right",
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
      });
    }
  }
};

const handleClose = () => {
  router.push("/");
};
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-gray-100 moving-clouds-bg animate-fade-in"
  >
    <div
      class="relative bg-white p-8 rounded-lg shadow-lg max-w-md w-full transform transition-transform duration-300 animate-scale-up"
    >
      <button
        @click="handleClose"
        class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 transition-transform duration-300 hover:scale-110"
        aria-label="Close"
      >
        <i class="fas fa-times text-2xl"></i>
      </button>
      <h2
        class="text-2xl font-bold mb-6 text-center text-gray-700 animate-fade-in animate-delay-200"
      >
        Update Password
      </h2>
      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <label
            for="new-password"
            class="block text-sm font-medium text-gray-600 animate-fade-in animate-delay-400"
            >New Password</label
          >
          <input
            id="new-password"
            v-model="form.newPassword"
            type="password"
            class="w-full p-3 mt-1 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-indigo-200 transition-transform duration-300 hover:scale-105"
            placeholder="Enter your new password"
            required
          />
          <p
            v-if="errors.newPassword"
            class="text-red-500 text-xs mt-1 animate-fade-in animate-delay-600"
          >
            {{ errors.newPassword }}
          </p>
        </div>

        <div class="mb-6">
          <label
            for="confirm-password"
            class="block text-sm font-medium text-gray-600 animate-fade-in animate-delay-400"
            >Confirm Password</label
          >
          <input
            id="confirm-password"
            v-model="form.confirmPassword"
            type="password"
            class="w-full p-3 mt-1 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-indigo-200 transition-transform duration-300 hover:scale-105"
            placeholder="Confirm your new password"
            required
          />
          <p
            v-if="errors.confirmPassword"
            class="text-red-500 text-xs mt-1 animate-fade-in animate-delay-600"
          >
            {{ errors.confirmPassword }}
          </p>
        </div>

        <div class="text-center">
          <button
            type="submit"
            class="w-full py-3 bg-indigo-500 text-white font-bold rounded-lg hover:bg-indigo-600 transition-transform duration-300 hover:scale-105 animate-fade-in animate-delay-800"
          >
            Update Password
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes scale-up {
  from {
    transform: scale(0.9);
  }
  to {
    transform: scale(1);
  }
}

@keyframes moveClouds {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 100% 0;
  }
}

.animate-fade-in {
  animation: fade-in 0.5s ease-in-out;
}

.animate-scale-up {
  animation: scale-up 0.5s ease-in-out;
}

.animate-delay-200 {
  animation-delay: 200ms;
}

.animate-delay-400 {
  animation-delay: 400ms;
}

.animate-delay-600 {
  animation-delay: 600ms;
}

.animate-delay-800 {
  animation-delay: 800ms;
}

.moving-clouds-bg {
  background-image: url("~/assets/images/clouds.png");
  background-size: cover;
  background-position: 0 0;
  background-repeat: repeat-x;
  animation: moveClouds 30s linear infinite;
}
</style>
