<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg max-w-md w-full">
      <h2 class="text-2xl font-bold mb-6 text-center text-gray-700">
        Update Password
      </h2>
      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <label
            for="new-password"
            class="block text-sm font-medium text-gray-600"
            >New Password</label
          >
          <input
            id="new-password"
            v-model="form.newPassword"
            type="password"
            class="w-full p-3 mt-1 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-indigo-200"
            placeholder="Enter your new password"
            required
          />
          <p v-if="errors.newPassword" class="text-red-500 text-xs mt-1">
            {{ errors.newPassword }}
          </p>
        </div>

        <div class="mb-6">
          <label
            for="confirm-password"
            class="block text-sm font-medium text-gray-600"
            >Confirm Password</label
          >
          <input
            id="confirm-password"
            v-model="form.confirmPassword"
            type="password"
            class="w-full p-3 mt-1 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-indigo-200"
            placeholder="Confirm your new password"
            required
          />
          <p v-if="errors.confirmPassword" class="text-red-500 text-xs mt-1">
            {{ errors.confirmPassword }}
          </p>
        </div>

        <div class="text-center">
          <button
            type="submit"
            class="w-full py-3 bg-indigo-500 text-white font-bold rounded-lg hover:bg-indigo-600 transition duration-300"
          >
            Update Password
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import ResetPassword from "~/graphql/mutations/users/resetPassword.gql";
import { useRoute } from "vue-router";

// Extract the route parameters
const route = useRoute();
const resetToken = route.query.reset_token || "";

// Reactive form state
const form = ref({
  newPassword: "",
  confirmPassword: "",
});

// Reactive error state
const errors = ref({
  newPassword: "",
  confirmPassword: "",
});

const {
  mutate: resetPasswordMutation,
  onDone,
  onError,
} = useMutation(ResetPassword);

// Validate the form fields
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

// Handle form submission
const handleSubmit = async () => {
  if (validateForm()) {
    try {
      const response = await resetPasswordMutation({
        reset_token: resetToken,
        new_password: form.value.newPassword,
      });
    } catch (error) {
      console.error("Password reset request failed:", error);
    }
  }
};

// Handle mutation completion
onDone((response) => {
  console.log("Mutation completed:", response.data);
});

// Handle mutation errors
onError((error) => {
  console.error("Mutation error:", error);
});
</script>
