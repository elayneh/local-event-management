<template>
  <div>
    <NuxtLink to="/products/add">
      <button>This is the home page for the authenticated users</button>
    </NuxtLink>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "authenticated",
});
</script>
