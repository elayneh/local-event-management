<script setup>
import { toast } from "vue3-toastify";
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useMutation } from "@vue/apollo-composable";
import { verifyEmail } from "~/graphql/mutations/users/verifyEmail.gql";
import { useAuthStore } from "~/stores";

const authenticationStore = useAuthStore();
const route = useRoute();
const router = useRouter();
const confirmationToken = route.query.token;
const success = ref(false);
const error = ref("");
const loading = ref(false);

const {
  mutate: verifyEmailMutation,
  onDone,
  onError,
} = useMutation(verifyEmail);

const confirmEmail = async () => {
  if (!confirmationToken) {
    error.value = "Invalid or missing token";
    return;
  }

  loading.value = true;

  try {
    const { data } = await verifyEmailMutation({ confirmationToken });
    if (data?.update_users?.affected_rows > 0) {
      success.value = true;
    } else {
      error.value = "Failed to confirm email. Token might be invalid.";
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
  >
    <div class="max-w-lg w-full bg-white shadow-lg rounded-lg p-8">
      <h1 class="text-2xl font-semibold mb-4">Email Confirmation</h1>
      <p>
        Do you try to register with your email to the local event management
        site? If so, confirm the registration.
      </p>
      <p v-if="error" class="text-red-500 mb-4">{{ error }}</p>
      <p v-if="success" class="text-green-500 mb-4">
        Your email has been successfully confirmed!
      </p>
      <button
        v-if="!success && !loading"
        @click="confirmEmail"
        class="w-full bg-blue-500 text-white font-semibold py-2 rounded hover:bg-blue-600 transition"
      >
        Confirm Email
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
