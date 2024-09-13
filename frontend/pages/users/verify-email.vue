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
    const { data } = await verifyEmailMutation({
      confirmationToken: confirmationToken,
    });

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
    class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50 z-50 animate-fade-in"
    role="dialog"
    aria-modal="true"
  >
    <div
      class="relative max-w-lg w-full bg-white shadow-lg rounded-lg p-8 transform transition-transform duration-500 ease-in-out animate-scale-up h-80 flex flex-col"
      role="document"
    >
      <button
        @click="goToHome"
        class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 focus:outline-none transform transition-transform duration-300 hover:scale-110"
        aria-label="Close"
      >
        <i class="fas fa-times text-2xl"></i>
      </button>

      <div class="flex-grow">
        <h1
          class="text-2xl font-semibold mb-4 text-center animate-fade-in animate-delay-200"
        >
          Email Confirmation
        </h1>

        <p class="mb-4 text-center animate-fade-in animate-delay-400">
          Do you want to register your email for the local event management
          site? If so, confirm your registration.
        </p>
      </div>

      <div class="mt-auto flex justify-center">
        <div
          v-if="!success && !loading"
          @click="confirmEmail"
          class="w-16 h-16 flex items-center justify-center bg-green-500 text-white font-semibold rounded-full hover:bg-green-600 transition-transform duration-300 transform hover:scale-105 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed animate-fade-in animate-delay-800"
          :disabled="loading"
          role="button"
          aria-label="Confirm Email"
        >
          <i class="fas fa-check text-2xl"></i>
        </div>

        <div
          v-if="loading"
          class="w-16 h-16 flex items-center justify-center bg-green-500 text-white font-semibold rounded-full cursor-not-allowed animate-fade-in animate-delay-800"
          disabled
        >
          <i class="fas fa-spinner fa-spin text-xl"></i>
          <!-- Loading spinner -->
        </div>

        <div
          v-if="success"
          @click="goToHome"
          class="w-16 h-16 flex items-center justify-center bg-green-500 text-white font-semibold rounded-full hover:bg-green-600 transition-transform duration-300 transform hover:scale-105 cursor-pointer animate-fade-in animate-delay-1000"
          role="button"
          aria-label="Go to Home"
        >
          <i class="fas fa-home text-2xl"></i>
        </div>
      </div>
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

.animate-delay-1000 {
  animation-delay: 1000ms;
}
</style>
