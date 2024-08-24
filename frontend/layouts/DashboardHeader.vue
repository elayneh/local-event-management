<template>
  <div>
    <header class="bg-white shadow px-6 py-4 flex justify-between items-center">
      <h1 class="text-lg font-semibold">Admin Dashboard</h1>
      <div class="flex items-center space-x-4">
        <button class="text-gray-600 hover:text-gray-800">Notifications</button>
        <button class="text-gray-600 hover:text-gray-800">Profile</button>
      </div>
    </header>
    <slot />
  </div>
</template>
