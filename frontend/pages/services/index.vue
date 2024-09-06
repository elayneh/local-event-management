<script>
export default {
  name: "ServicesPage",
};
</script>

<style scoped>
@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes slide-in-left {
  from {
    transform: translateX(-50%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes slide-in-right {
  from {
    transform: translateX(50%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.animate-fade-in {
  animation: fade-in 1s ease-out forwards;
}

.animate-slide-in-left {
  animation: slide-in-left 1s ease-out forwards;
}

.animate-slide-in-right {
  animation: slide-in-right 1s ease-out forwards;
}

.animate-bounce {
  animation: bounce 2s infinite;
}

.animate-pulse {
  animation: pulse 2s infinite;
}
</style>

<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <header class="text-center py-12">
      <h1 class="text-5xl font-bold mb-4 animate-fade-in animate-pulse">
        Our Services
      </h1>
      <p class="text-lg animate-slide-in-left">
        Delivering excellence in every project
      </p>
    </header>

    <section class="max-w-7xl mx-auto py-12">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        <div
          class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl hover:bg-yellow-50"
        >
          <div class="flex items-center justify-center mb-6">
            <span class="text-yellow-500 text-4xl animate-bounce"
              >&#127881;</span
            >
          </div>
          <h3 class="text-2xl font-semibold text-gray-900 mb-4">
            Explore Events
          </h3>
          <p class="text-gray-700 mb-6">
            Discover a wide range of events tailored to your interests. Filter
            by location, date, category, price, and more to find the perfect
            event.
          </p>
          <NuxtLink
            to="/events/"
            class="text-yellow-500 hover:text-yellow-600 transition duration-300"
          >
            Learn More &rarr;
          </NuxtLink>
        </div>

        <div
          class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl hover:bg-yellow-50"
        >
          <div class="flex items-center justify-center mb-6">
            <span class="text-yellow-500 text-4xl animate-bounce"
              >&#128221;</span
            >
          </div>
          <h3 class="text-2xl font-semibold text-gray-900 mb-4">
            Create and Manage Events
          </h3>
          <p class="text-gray-700 mb-6">
            Have an event idea? Easily create, edit, and manage your events.
            Upload images, set locations, define categories, and more with our
            intuitive event management tools.
          </p>
          <NuxtLink
            to="/user/events/create-event"
            class="text-yellow-500 hover:text-yellow-600 transition duration-300"
          >
            Learn More &rarr;
          </NuxtLink>
        </div>

        <div
          class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl hover:bg-yellow-50"
        >
          <div class="flex items-center justify-center mb-6">
            <span class="text-yellow-500 text-4xl animate-bounce"
              >&#128179;</span
            >
          </div>
          <h3 class="text-2xl font-semibold text-gray-900 mb-4">
            Reserve, Follow, and Buy Tickets
          </h3>
          <p class="text-gray-700 mb-6">
            Stay engaged with events that interest you. Bookmark your favorites,
            follow event updates, and securely purchase tickets all in one
            place.
          </p>
          <NuxtLink
            to="/user/events"
            class="text-yellow-500 hover:text-yellow-600 transition duration-300"
          >
            Learn More &rarr;
          </NuxtLink>
        </div>
      </div>
    </section>

    <section class="bg-yellow-500 text-gray-900 py-12">
      <div class="max-w-7xl mx-auto text-center">
        <h2 class="text-4xl font-semibold mb-6 animate-fade-in">
          Ready to Make Your Event Unforgettable?
        </h2>
        <p class="text-lg mb-8 animate-slide-in-right">
          Get in touch with us to start your project today. Our team of experts
          is here to help you succeed.
        </p>
        <NuxtLink
          to="/contact"
          class="bg-gray-900 text-white font-bold py-4 px-8 rounded-full hover:bg-gray-800 transition duration-300 animate-pulse"
        >
          Contact Us
        </NuxtLink>
      </div>
    </section>
    <CustomFooter />
  </div>
</template>
