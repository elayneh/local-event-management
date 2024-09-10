<template>
  <div class="flex flex-col h-screen">
    <header class="fixed top-0 left-0 w-full bg-transparent">
      <nav class="flex pt-4 justify-between items-center bg-gray-1">
        <ul class="flex gap-4 items-center ml-8 p-0 m-0">
          <li class="flex items-center gap-2 p-0 m-0">
            <NuxtLink
              to="/"
              @click.native="scrollToTop"
              class="flex items-center p-0 m-0"
            >
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
            class="bg-gradient-to-r from-gray-400 via-gray-500 to-gray-600 text-white font-bold uppercase tracking-wide py-3 px-6 rounded-full shadow-lg hover:shadow-xl transform hover:-translate-y-1 hover:scale-105 hover:bg-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 transition-all duration-500 ease-in-out"
          >
            <NuxtLink to="/" @click.native="scrollToTop">Home</NuxtLink>
          </ul>
          <ul
            class="bg-gradient-to-r from-gray-400 via-gray-500 to-gray-600 text-white font-bold uppercase tracking-wide py-3 px-6 rounded-full shadow-lg hover:shadow-xl transform hover:-translate-y-1 hover:scale-105 hover:bg-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 transition-all duration-500 ease-in-out"
          >
            <NuxtLink to="/about"> About Us </NuxtLink>
          </ul>
          <ul
            class="bg-gradient-to-r from-gray-400 via-gray-500 to-gray-600 text-white font-bold uppercase tracking-wide py-3 px-6 rounded-full shadow-lg hover:shadow-xl transform hover:-translate-y-1 hover:scale-105 hover:bg-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 transition-all duration-500 ease-in-out"
          >
            <NuxtLink to="/services"> Services </NuxtLink>
          </ul>
          <ul
            class="bg-gradient-to-r from-green-400 via-green-500 to-green-600 text-white font-bold uppercase tracking-wide py-3 px-6 rounded-full shadow-lg hover:shadow-xl transform hover:-translate-y-1 hover:scale-105 hover:bg-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 transition-all duration-500 ease-in-out"
          >
            <NuxtLink
              to="/user/events/create-event"
              @click="isUserAuthenticated"
              >Create event
            </NuxtLink>
          </ul>
          <ul
            class="bg-gradient-to-r from-blue-500 via-blue-500 to-blue-600 text-white font-bold py-3 px-6 rounded-lg shadow-lg hover:shadow-xl transform hover:-translate-y-1 hover:scale-105 hover:bg-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 transition-all duration-500 ease-in-out"
          >
            <li class="text-white">
              <NuxtLink to="/users/register">Register</NuxtLink>
            </li>
          </ul>
        </div>
      </nav>
    </header>

    <main ref="mainContent" class="flex-1 overflow-y-auto pt-18">
      <div>
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";

const dropdownOpen = ref(false);
const eventDropDown = ref(false);
const isLogoutModalOpen = ref(false);
const mainContent = ref(null);

const scrollToTop = () => {
  if (mainContent.value) {
    mainContent.value.scrollTo({ top: 0, behavior: "smooth" });
  }
};

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
