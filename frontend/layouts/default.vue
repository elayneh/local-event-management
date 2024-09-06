<template>
  <div class="flex flex-col h-screen">
    <header class="fixed top-0 left-0 w-full bg-transparent">
      <nav class="flex pt-4 justify-between items-center bg-gray-1">
        <ul class="flex gap-4 items-center ml-8 p-0 m-0">
          <li class="flex items-center gap-2 p-0 m-0">
            <NuxtLink to="/" class="flex items-center p-0 m-0">
              <img
                src="~/assets/images/logo.png"
                alt="Logo"
                class="h-16 w-auto transition-transform transform hover:scale-110"
              />
            </NuxtLink>
          </li>
        </ul>

        <div class="flex space-x-4 gap-4 justify-center items-center mr-24">
          <ul
            class="flex justify-center items-center text-white bg-gray-950 rounded-lg p-2 transition delay-150 duration-300 ease-in-out hover:delay-150 hover:-translate-y-2 hover:skew-x-1 hover:bg-gray-500 hover:scale-105 hover:shadow-lg hover:shadow-gray-800 hover:text-yellow-300 delay-50 active:scale-95 active:bg-gray-700 active:shadow-inner active:text-yellow-400 focus:outline-none focus:ring-2 focus:ring-gray-700 focus:ring-offset-2 focus:ring-offset-gray-950"
          >
            <NuxtLink to="/">Home</NuxtLink>
          </ul>
          <ul
            class="flex justify-center items-center text-white bg-gray-950 rounded-lg p-2 transition delay-150 duration-300 ease-in-out hover:delay-150 hover:-translate-y-2 hover:skew-x-1 hover:bg-gray-500 hover:scale-105 hover:shadow-lg hover:shadow-gray-800 hover:text-yellow-300 delay-50 active:scale-95 active:bg-gray-700 active:shadow-inner active:text-yellow-400 focus:outline-none focus:ring-2 focus:ring-gray-700 focus:ring-offset-2 focus:ring-offset-gray-950"
          >
            <NuxtLink to="/about"> About Us </NuxtLink>
          </ul>
          <ul
            class="flex justify-center items-center text-white bg-gray-950 rounded-lg p-2 transition delay-150 duration-300 ease-in-out hover:delay-150 hover:-translate-y-2 hover:skew-x-1 hover:bg-gray-500 hover:scale-105 hover:shadow-lg hover:shadow-gray-800 hover:text-yellow-300 delay-50 active:scale-95 active:bg-gray-700 active:shadow-inner active:text-yellow-400 focus:outline-none focus:ring-2 focus:ring-gray-700 focus:ring-offset-2 focus:ring-offset-gray-950"
          >
            <NuxtLink to="/services"> Services </NuxtLink>
          </ul>
          <ul
            class="flex justify-center items-center text-white bg-green-500 rounded-lg p-2 transition delay-150 duration-300 ease-in-out hover:delay-150 hover:-translate-y-2 hover:skew-x-1 hover:bg-gray-500 hover:scale-105 hover:shadow-lg hover:shadow-gray-800 hover:text-yellow-300 delay-50 active:scale-95 active:bg-gray-700 active:shadow-inner active:text-yellow-400 focus:outline-none focus:ring-2 focus:ring-gray-700 focus:ring-offset-2 focus:ring-offset-gray-950"
          >
            <NuxtLink to="/user/events/create-event">
              <button @click="isUserAuthenticated">Create event</button>
            </NuxtLink>
          </ul>
          <ul
            class="flex justify-center items-center bg-blue-500 rounded-lg p-2 justify-center items-center transition delay-150 duration-300 ease-in-out hover:delay-150 hover:-translate-y-2 hover:skew-x-1 hover:bg-blue-700 hover:scale-105 hover:shadow-lg hover:shadow-gray-800 hover:text-yellow-300 delay-50 active:scale-95 active:bg-gray-700 active:shadow-inner active:text-yellow-400 focus:outline-none focus:ring-2 focus:ring-gray-700 focus:ring-offset-2 focus:ring-offset-gray-950"
          >
            <li class="text-white">
              <NuxtLink to="/users/register"
                ><Button>Register</Button></NuxtLink
              >
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

const handleClickOutside = (event) => {
  if (!event.target.closest(".relative")) {
    eventDropDown.value = false;
  }
};

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>
