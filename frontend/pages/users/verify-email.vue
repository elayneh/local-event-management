<script setup>
import { toast } from "vue3-toastify";
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useMutation } from "@vue/apollo-composable";
import verifyEmail from "~/graphql/mutations/users/verifyEmail.gql";
import { useAuthStore } from "~/stores";

const authenticationStore = useAuthStore();
const route = useRoute();
const router = useRouter();
const confirmationToken = route.query.token;
const success = ref(false);
const error = ref("");
const loading = ref(false);

const { mutate: verifyEmailMutation } = useMutation(verifyEmail);

const confirmEmail = async () => {
  if (!confirmationToken) {
    error.value = "Invalid or missing token";
    return;
  }

  loading.value = true;

  try {
    const { data } = await verifyEmailMutation({ confirmationToken: confirmationToken });
    
    // Check if the `verifyEmail` response has a `message` field indicating success
    if (data?.verifyEmail?.message === "Email confirmed successfully") {
      success.value = true;
      toast.success("Your email has been successfully confirmed!");
    } else {
      error.value = "Failed to confirm email. Token might be invalid.";
      toast.error("Failed to confirm email. Token might be invalid.");
    }
  } catch (err) {
    error.value = "An error occurred while confirming your email.";
    console.error("Error: ", err.message);
    toast.error("An error occurred while confirming your email.");
  } finally {
    loading.value = false;
  }
};

const goToHome = () => {
  router.push("/");
};
</script>

<template>
  <div
    class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50 z-50"
    role="dialog"
    aria-modal="true"
  >
    <div
      class="relative max-w-lg w-full bg-white shadow-lg rounded-lg p-8"
      role="document"
    >
      <button
        @click="goToHome"
        class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 focus:outline-none"
        aria-label="Close"
      >
        <i class="fas fa-times text-2xl"></i>
      </button>

      <h1 class="text-2xl font-semibold mb-4 text-center">
        Email Confirmation
      </h1>

      <p class="mb-4 text-center">
        Do you want to register your email for the local event management site?
        If so, confirm your registration.
      </p>

      <p v-if="error" class="text-red-500 text-center mb-4">{{ error }}</p>
      <p v-if="success" class="text-green-500 text-center mb-4">
        Your email has been successfully confirmed!
      </p>

      <button
        v-if="!success && !loading"
        @click="confirmEmail"
        class="w-full bg-blue-500 text-white font-semibold py-2 rounded hover:bg-blue-600 transition disabled:opacity-50 disabled:cursor-not-allowed"
        :disabled="loading"
      >
        Confirm Email
      </button>

      <button
        v-if="loading"
        class="w-full bg-blue-500 text-white font-semibold py-2 rounded cursor-not-allowed"
        disabled
      >
        <i class="fas fa-spinner fa-spin mr-2"></i> Confirming...
      </button>

      <button
        v-if="success"
        @click="goToHome"
        class="w-full bg-blue-500 text-white font-semibold py-2 rounded hover:bg-blue-600 transition"
      >
        Go to Home
      </button>
    </div>
  </div>
</template>
