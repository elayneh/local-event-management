<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Adding animation classes */
.animate-fadeIn {
  animation: fadeIn 1s ease-out;
}

.animate-fadeInUp {
  animation: fadeInUp 1s ease-out;
}

.animation-delay-1 {
  animation-delay: 0.2s;
}

.animation-delay-2 {
  animation-delay: 0.4s;
}

.animation-delay-3 {
  animation-delay: 0.6s;
}

.animation-delay-4 {
  animation-delay: 0.8s;
}

.animation-delay-5 {
  animation-delay: 1s;
}

.animation-delay-6 {
  animation-delay: 1.2s;
}
</style>

<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <header class="text-center py-24">
      <h1 class="text-5xl font-extrabold mb-4 text-yellow-600 animate-pulse">
        About Minab
      </h1>
      <p class="text-xl text-gray-800 font-medium animate-fadeIn">
        Your go-to platform for seamless event management
      </p>
    </header>

    <section class="max-w-10xl mx-auto p-10">
      <h2
        class="text-4xl font-bold text-yellow-600 animate-fadeInUp flex justify-center"
      >
        Our Mission
      </h2>
      <p
        class="text-gray-800 text-lg leading-relaxed animate-fadeInUp animation-delay-1"
      >
        At <strong>Minab Local Event Mangement Platform</strong>, we are
        dedicated to providing a seamless and efficient event management
        experience for both organizers and attendees. Whether you're planning a
        small gathering or a large-scale conference, our platform offers the
        tools and resources you need to make your event a success.
      </p>
    </section>

    <section class="max-w-10xl mx-auto px-10">
      <h2 class="text-4xl font-bold text-yellow-600 mb-6 animate-fadeInUp">
        Our Key Features
      </h2>
      <section class="max-w-8xl p-10">
        <h2 class="text-3xl font-semibold text-gray-800">Features</h2>
        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-8">
          <div
            class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl animate-fadeInUp animation-delay-1"
          >
            <h3 class="text-2xl font-bold text-gray-900 mb-4">
              Comprehensive Event Management
            </h3>
            <p class="text-gray-700">
              Manage all aspects of your event from one central platform,
              including scheduling, ticketing, and attendee management. Tailor
              every detail to ensure a successful and memorable experience.
            </p>
          </div>

          <div
            class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl animate-fadeInUp animation-delay-2"
          >
            <h3 class="text-2xl font-bold text-gray-900 mb-4">
              Interactive Event Discovery
            </h3>
            <p class="text-gray-700">
              Explore events based on your interests. Filter by location, date,
              category, and price, or search by title, description, and tags.
              Discover new experiences effortlessly.
            </p>
          </div>

          <div
            class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl animate-fadeInUp animation-delay-3"
          >
            <h3 class="text-2xl font-bold text-gray-900 mb-4">
              Secure Ticketing and Payments
            </h3>
            <p class="text-gray-700">
              Book your spot with confidence. Our platform ensures secure
              payment processing, making ticket purchasing simple and
              worry-free.
            </p>
          </div>

          <div
            class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl animate-fadeInUp animation-delay-4"
          >
            <h3 class="text-2xl font-bold text-gray-900 mb-4">
              Real-Time Analytics
            </h3>
            <p class="text-gray-700">
              Gain insights into your event's performance with real-time data on
              ticket sales, attendee engagement, and more, helping you optimize
              and grow.
            </p>
          </div>

          <div
            class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl animate-fadeInUp animation-delay-5"
          >
            <h3 class="text-2xl font-bold text-gray-900 mb-4">
              Customizable Event Pages
            </h3>
            <p class="text-gray-700">
              Create beautiful, personalized event pages that match your brand
              and capture the essence of your event. Stand out and make a
              lasting impression.
            </p>
          </div>

          <div
            class="bg-white p-8 rounded-lg shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl animate-fadeInUp animation-delay-6"
          >
            <h3 class="text-2xl font-bold text-gray-900 mb-4">
              Social Media Integration
            </h3>
            <p class="text-gray-700">
              Promote your events effortlessly with built-in social media
              sharing tools and integrations. Reach your audience where they
              are.
            </p>
          </div>
        </div>
      </section>
    </section>

    <CustomFooter />
  </div>
</template>
