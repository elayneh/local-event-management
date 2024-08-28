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
        <div class="flex space-x-4 gap-4 justify-center items-center pr-24">
          <ul class="flex justify-center items-center">
            <NuxtLink to="/user/events/create-event">
              <button>Create event</button>
            </NuxtLink>
          </ul>
          <ul class="flex m-2">
            <span>All What You Have</span>
            <li class="relative items-center gap-2">
              <button @click.stop="toggleDropdown">
                <font-awesome-icon
                  :icon="['fas', 'caret-down']"
                  class="text-gray-600 pl-4"
                />
              </button>
              <ul
                v-if="eventDropDown"
                class="absolute left-0 top-full mt-2 w-48 bg-white rounded-lg shadow-lg z-10"
              >
                <li>
                  <NuxtLink
                    to="/user/events"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >All events</NuxtLink
                  >
                </li>
                <li>
                  <NuxtLink
                    to="/user/events/bookmarked"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >Bookmarks</NuxtLink
                  >
                </li>
                <li>
                  <NuxtLink
                    to="/user/events/following"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >Following</NuxtLink
                  >
                </li>
                <li>
                  <NuxtLink
                    to="/user/events/tickets"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >Purchased Events</NuxtLink
                  >
                </li>
                <li>
                  <NuxtLink
                    to="#"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >Option 5</NuxtLink
                  >
                </li>
                <li>
                  <NuxtLink
                    to="#"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >Option 6</NuxtLink
                  >
                </li>
              </ul>
            </li>
          </ul>
          <ul class="flex gap-4 items-center mr-20">
            <li class="relativep">
              <button
                @click="toggleProfileDropdown"
                class="w-10 h-10 rounded-full bg-gray-300 flex items-center justify-center focus:outline-none focus:ring-2 focus:ring-gray-400"
              >
                <img src="" alt="Profile" class="w-full h-full rounded-full" />
              </button>
              <div
                v-if="dropdownOpen"
                class="absolute mt-2 w-64 bg-white shadow-lg bg-neutral-300"
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
                    <button
                      @click="openLogoutModal"
                      class="w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >
                      Analysis
                    </button>
                  </li>
                  <li>
                    <button
                      @click="openLogoutModal"
                      class="w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >
                      Settings
                    </button>
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
import * as JsCookie from "js-cookie";
import { useAuthStore } from "~/stores";

const Cookies = JsCookie.default;

const authenticationStore = useAuthStore();
const dropdownOpen = ref(false);
const eventDropDown = ref(false);
const isLogoutModalOpen = ref(false);

const toggleDropdown = () => {
  eventDropDown.value = !eventDropDown.value;
};

const toggleProfileDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

const openLogoutModal = () => {
  isLogoutModalOpen.value = true;
};

const closeLogoutModal = () => {
  window.location.reload(); // This reloads the current page

  isLogoutModalOpen.value = false;
};

const confirmLogout = () => {
  Cookies.remove("token");
  authenticationStore.setToken(null);
  authenticationStore.setId(null);
  authenticationStore.setUser(null);
  authenticationStore.setRole(null);

  closeLogoutModal();
};
const handleClickOutside = (event) => {
  if (!event.target.closest(".relative")) {
    eventDropDown.value = false;
  }
  if (!event.target.closest(".relativep")) {
    dropdownOpen.value = false;
  }
};

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>
