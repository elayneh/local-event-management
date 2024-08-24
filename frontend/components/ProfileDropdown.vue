<template>
  <div
    class="dropdown-container bg-gradient-to-r from-blue-50 via-green-50 to-yellow-50"
  >
    <ul class="py-1">
      <li>
        <NuxtLink
          to="/profile"
          class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
          >Profile</NuxtLink
        >
      </li>
      <li>
        <NuxtLink
          to="/analysis"
          class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
          >Analysis</NuxtLink
        >
      </li>
      <li>
        <NuxtLink
          to="/settings"
          class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
          >Settings</NuxtLink
        >
      </li>
      <li>
        <button
          @click="openLogoutModal"
          class="w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100"
        >
          Logout
        </button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref } from "vue";

// Emit event to parent component to handle the logout modal
const emit = defineEmits(["open-logout-modal"]);

const openLogoutModal = () => {
  emit("open-logout-modal");
};
</script>
