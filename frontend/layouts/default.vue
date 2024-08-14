<template>
  <div>
    <header class="shadow-sm bg-gray-100">
      <nav class="container mx-auto p-4 flex justify-end items-center">
        <ul class="flex gap-4 items-center">
          <li>
            <NuxtLink to="/" :class="linkClass('/')" @click="markVisited('/')"
              >Home</NuxtLink
            >
          </li>
          <li>
            <NuxtLink
              to="/about"
              :class="linkClass('/about')"
              @click="markVisited('/about')"
              >About</NuxtLink
            >
          </li>
          <li>
            <NuxtLink
              class="btn"
              to="/user/register"
              :class="linkClass('/user/register')"
              @click="markVisited('/user/register')"
            >
              <button
                type="button"
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
              >
                Register
              </button>
            </NuxtLink>
          </li>
        </ul>
      </nav>
    </header>
    <div class="container mx-auto p-4">
      <slot />
    </div>
  </div>
</template>

<script setup>
const visitedLinks = ref([]);
const route = useRoute();

const isActive = (path) => route.path === path;

const linkClass = (path) => {
  return {
    "text-blue-500 font-bold": isActive(path),
    "text-gray-300": visitedLinks.value.includes(path),
  };
};

const markVisited = (path) => {
  if (!visitedLinks.value.includes(path)) {
    visitedLinks.value.push(path);
  }
};
</script>
