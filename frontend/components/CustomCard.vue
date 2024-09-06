<script setup>
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  image: {
    type: String,
    default: null,
  },
  customClass: {
    type: String,
    default: "",
  },
});
</script>

<template>
  <div
    :class="[
      customClass,
      'bg-white w-full max-w-md mx-auto shadow-lg rounded-lg overflow-hidden flex flex-col transform transition-transform duration-300 hover:scale-105 hover:shadow-2xl',
    ]"
  >
    <transition name="fade-slide">
      <img
        v-if="image"
        :src="image"
        alt="Card Image"
        class="w-full h-48 object-cover transition-opacity duration-500 ease-in-out"
      />
    </transition>

    <div class="text-white text-center py-4">
      <transition name="fade-in-down">
        <h2 class="text-2xl font-semibold text-blue-500">{{ title }}</h2>
      </transition>
    </div>
    <div class="flex-1 p-6">
      <transition name="fade-slide">
        <slot name="body"></slot>
      </transition>
    </div>
    <div class="p-6 border-t border-gray-200 flex justify-between">
      <transition name="fade-slide">
        <slot name="footer-left"></slot>
      </transition>
      <transition name="fade-slide">
        <slot name="footer-right"></slot>
      </transition>
    </div>
  </div>
</template>

<style>
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: opacity 0.5s ease, transform 0.5s ease;
}
.fade-slide-enter-from,
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

.fade-in-down-enter-active,
.fade-in-down-leave-active {
  transition: opacity 0.5s ease, transform 0.5s ease;
}
.fade-in-down-enter-from,
.fade-in-down-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}
</style>
