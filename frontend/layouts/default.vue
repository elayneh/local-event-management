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
        <div class="flex space-x-4 gap-4 justify-center items-center mr-24">
          <ul class="flex justify-center items-center text-gray-700">
            <NuxtLink to="/user/events/create-event">
              <button @click="isUserAuthenticated">Create event</button>
            </NuxtLink>
          </ul>
          <ul
            class="flex justify-center items-center bg-blue-500 rounded-lg px-4 justify-center items-center py-1"
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
    \
    <FilterModal
      :isOpen="isFilterModalOpen"
      @close-modal="closeFilterModal"
      @apply-filter="applyFilter"
    />
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
import FilterModal from "~/components/FilterModal.vue";

const isFilterModalOpen = ref(false);

const openFilterModal = () => {
  isFilterModalOpen.value = true;
};

const closeFilterModal = () => {
  isFilterModalOpen.value = false;
};

const applyFilter = (selectedFields) => {
  console.log("Apply filter with:", selectedFields);
  // Implement your filter logic here
  closeFilterModal();
};
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
