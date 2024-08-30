<template>
  <div class="flex flex-col h-full bg-slate-500 text-white w-64">
    <!-- Header -->
    <div class="flex items-center justify-center py-6">
      <h2 class="text-2xl font-bold">Admin Dashboard</h2>
    </div>

    <!-- Navigation Links -->
    <nav class="flex flex-col mt-4 space-y-2 px-4">
      <NuxtLink
        v-for="link in links"
        :key="link.name"
        :to="link.path"
        class="flex items-center px-4 py-2 hover:bg-gray-700 rounded transition-colors duration-200"
        :class="{ 'bg-gray-700': $route.path === link.path }"
      >
        <span class="mr-3">
          <font-awesome-icon
            v-if="link.iconComponent"
            :icon="link.iconComponent"
          />
          <span v-else>{{ link.icon }}</span>
        </span>
        <span>{{ link.name }}</span>
      </NuxtLink>
    </nav>

    <!-- Logout Button -->
    <div class="mt-auto px-4 py-4">
      <button
        @click="openLogoutModal"
        class="w-full flex items-center justify-center px-4 py-2 bg-red-500 hover:bg-red-600 text-white font-bold rounded transition duration-200"
      >
        <font-awesome-icon icon="sign-out-alt" class="mr-2" />
        Logout
      </button>
    </div>

    <!-- Logout Confirmation Modal -->
    <LogoutModal
      v-if="isLogoutModalOpen"
      @confirm="confirmLogout"
      @cancel="closeLogoutModal"
    />
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import Cookies from "js-cookie";
import { useAuthStore } from "@/stores";
import LogoutModal from "@/components/LogoutModal.vue";

const isLogoutModalOpen = ref(false);

const links = ref([
  {
    name: "Dashboard",
    path: "/admin",
    iconComponent: ["fas", "tachometer-alt"],
  },
  { name: "Users", path: "/admin/users", iconComponent: ["fas", "user"] },
  {
    name: "Events",
    path: "/admin/events",
    iconComponent: ["fas", "calendar-alt"],
  },
  {
    name: "Categories",
    path: "/admin/categories",
    iconComponent: ["fas", "tag"],
  },
  { name: "Tags", path: "/admin/tags", iconComponent: ["fas", "tag"] },
  { name: "Settings", path: "/admin/settings", iconComponent: ["fas", "cog"] },
]);

const router = useRouter();
const authenticationStore = useAuthStore();

const openLogoutModal = () => {
  isLogoutModalOpen.value = true;
};

const closeLogoutModal = () => {
  isLogoutModalOpen.value = false;
};

const confirmLogout = () => {
  Cookies.remove("token");
  authenticationStore.setToken(null);
  authenticationStore.setId(null);
  authenticationStore.setUser(null);
  authenticationStore.setRole(null);
  closeLogoutModal();
  router.push("/");
};
</script>
