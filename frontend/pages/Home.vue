<template>
  <div>
    <NuxtLink to="/uploadFile">
      <button>Upload File</button>
    </NuxtLink>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "authenticated",
});
</script>
