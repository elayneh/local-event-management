<template>
  <div class="flex flex-col h-screen">
    <header class="fixed top-0 left-0 w-full z-20 bg-transparent">
      <nav class="flex pt-4 justify-between items-center bg-gray-1">
        <ul class="flex gap-4 items-center ml-20">
          <li class="flex items-center gap-2">
            <NuxtLink to="/">
              <i class="fas fa-home"> </i>
            </NuxtLink>
          </li>
        </ul>
        <div class="flex">
          <ul class="m-2">
            <li class="relative flex items-center">
              <input
                @click="search"
                class="w-96 h-10 rounded-full pl-4 pr-10 bg-gray-300 flex items-center focus:outline-none focus:ring-1 focus:ring-gray-400"
                placeholder="Search..."
              />
              <font-awesome-icon
                :icon="['fas', 'search']"
                class="text-gray-600 absolute right-4"
              />
            </li>
          </ul>

          <ul class="m-2 gap-">
            <li class="relative">
              <button
                @click="search"
                class="w-64 h-10 rounded bg-gray-300 flex items-center focus:outline-none focus:ring-1 focus:ring-gray-400"
              >
                <font-awesome-icon
                  :icon="['fas', 'filter']"
                  class="text-gray-600 pl-4"
                />
                <span class="ml-2 text-gray-500">filter</span>
              </button>
            </li>
          </ul>
        </div>
        <div class="flex space-x-4 gap-4 justify-center items-center mr-24">
          <ul class="flex justify-center items-center text-gray-700">
            <NuxtLink to="/user/events/create-event">
              <button @click="isUserAuthenticated">Create event</button>
            </NuxtLink>
          </ul>
          <ul
            class="flex justify-center items-center bg-blue-500 rounded-lg px-4 justify-center items-center py-1"
          >
            <li>
              <NuxtLink to="users/register"><Button>Register</Button></NuxtLink>
            </li>
          </ul>
        </div>
      </nav>
    </header>

    <main class="flex-1 overflow-y-auto pt-18">
      <div>
        <slot />
      </div>
    </main>

    <div
      v-if="isLogoutModalOpen"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
    >
      <div
        class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full text-center"
      >
        <h2 class="text-xl font-semibold mb-4">Confirm Logout</h2>
        <p class="mb-6">Are you sure you want to log out?</p>
        <div class="flex justify-center gap-4">
          <button
            @click="confirmLogout"
            class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
          >
            Yes
          </button>
          <button
            @click="closeLogoutModal"
            class="bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";

const dropdownOpen = ref(false);
const eventDropDown = ref(false);
const isLogoutModalOpen = ref(false);

const isUserAuthenticated = () => {};

const closeLogoutModal = () => {
  isLogoutModalOpen.value = false;
};

const confirmLogout = () => {
  closeLogoutModal();
};

// Handle clicks outside of the dropdown menu
const handleClickOutside = (event) => {
  if (!event.target.closest(".relative")) {
    eventDropDown.value = false;
  }
};

// Add and remove click event listeners
onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>
