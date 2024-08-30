<
<template>
  <div>
    <footer class="bg-gray-800 text-white py-12 mt-12">
      <div
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 px-6 lg:px-24"
      >
        <div>
          <h3 class="text-lg font-semibold mb-4">Quick Links</h3>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="#" class="hover:text-gray-300">Home</NuxtLink>
            </li>
            <li>
              <NuxtLink to="#" class="hover:text-gray-300">About Us</NuxtLink>
            </li>
            <li>
              <NuxtLink to="#" class="hover:text-gray-300">Events</NuxtLink>
            </li>
            <li>
              <NuxtLink to="#" class="hover:text-gray-300">Contact</NuxtLink>
            </li>
          </ul>
        </div>

        <div class="mx-auto">
          <h3 class="text-lg font-semibold mb-4">Newsletter</h3>
          <p class="text-gray-400 mb-4">
            Subscribe to our newsletter for the latest updates.
          </p>
          <form @submit.prevent="console.log('Subscribed')">
            <input
              type="email"
              placeholder="Enter your email"
              class="w-full p-2 mb-4 text-gray-800 rounded"
            />
            <button
              type="submit"
              class="w-full bg-blue-600 hover:bg-blue-700 text-white p-2 rounded transition-colors duration-200"
            >
              Subscribe
            </button>
          </form>
        </div>

        <div class="mx-auto">
          <h3 class="text-lg font-semibold mb-4">Follow Us</h3>
          <div class="flex-col-1">
            <NuxtLink to="#" class="hover:text-gray-300">
              <p>LinkedIn</p>
            </NuxtLink>
            <NuxtLink to="#" class="hover:text-gray-300">
              <p>WhatsApp</p>
            </NuxtLink>
            <NuxtLink to="#" class="hover:text-gray-300">
              <p>Facebook</p>
            </NuxtLink>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup></script>
>
