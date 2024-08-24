<template>
  <div class="flex flex-col h-full bg-slate-500 text-white w-64">
    <div class="flex items-center justify-center py-6">
      <h2 class="text-2xl font-bold">Admin Dashboard</h2>
    </div>

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
  </div>
</template>

<script setup>
import { ref } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faTachometerAlt,
  faUser,
  faCalendarAlt,
  faTag,
  faCog,
} from "@fortawesome/free-solid-svg-icons";
import { library } from "@fortawesome/fontawesome-svg-core";

library.add(faTachometerAlt, faUser, faCalendarAlt, faTag, faCog);

const links = ref([
  {
    name: "Dashboard",
    path: "/admin",
    iconComponent: ["fas", "tachometer-alt"],
  },
  { name: "Users", path: "/admin/users", iconComponent: ["fas", "user"] },
  { name: "Events", path: "/events", iconComponent: ["fas", "calendar-alt"] },
  {
    name: "Categories",
    path: "/admin/categories",
    iconComponent: ["fas", "tag"],
  },
  { name: "Tags", path: "/admin/tags", iconComponent: ["fas", "tag"] },
  {
    name: "Settings",
    path: "/admin/settings",
    iconComponent: ["fas", "cog"],
  },
]);
</script>
