<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <Sidebar />

    <div class="flex-1 bg-gray-100 p-6 overflow-auto">
      <NuxtPage />
    </div>
  </div>
</template>

<script setup>
// Import your sidebar component
import Sidebar from "~/components/Sidebar.vue";
</script>