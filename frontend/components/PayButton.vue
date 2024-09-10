<template>
  <div>
    <button @click="payForEvent">Pay for Event</button>
  </div>
</template>

<script setup>
const payForEvent = async () => {
  try {
    const response = await fetch("/pay", {
      method: "POST",
    });
    if (response.redirected) {
      window.location.href = response.url;
    } else {
      console.error("Failed to initialize payment");
    }
  } catch (error) {
    console.error("Error:", error);
  }
};
</script>
