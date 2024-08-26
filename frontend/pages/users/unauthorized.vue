<template>
  <div
    class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-r from-blue-100 via-green-100 to-yellow-100"
  >
    <div
      class="w-2/5 bg-white p-8 rounded-lg shadow-md text-center"
    >
      <h2 class="text-3xl font-bold text-red-600 mb-4">Unauthorized</h2>
      <p class="text-gray-600 mb-6">
        You do not have permission to access this page.
      </p>
      <button
        @click="goToLogin"
        class="bg-indigo-600 text-white py-2 px-4 rounded-lg hover:bg-indigo-700 transition-colors"
      >
        Go to Login
      </button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

function goToLogin() {
  router.push("/users/login");
}
</script>
