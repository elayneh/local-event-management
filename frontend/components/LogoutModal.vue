<template>
  <div
    class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
  >
    <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full text-center">
      <h2 class="text-xl font-semibold mb-4">Confirm Logout</h2>
      <p class="mb-6">Are you sure you want to log out?</p>
      <div class="flex justify-center gap-4">
        <button
          @click="$emit('confirm')"
          class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
        >
          Yes
        </button>
        <button
          @click="$emit('cancel')"
          class="bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps(["confirm", "cancel"]);
</script>
