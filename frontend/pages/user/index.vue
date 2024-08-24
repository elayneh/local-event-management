<template>
  <div class="bg-gradient-to-r from-blue-100 via-green-100 to-yellow-1">
    <div class="w-full flex justify-center items-center pt-24">
      <div class="items-center">
        <img src="/assets/images/home.png" />
      </div>
      <div class="items-center">
        <img src="/assets/images/home.png" />
      </div>
    </div>
    <h2 class="text-2xl font-bold mb-4 text-center">Latest Events</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in visibleEvents" :key="event.id">
        <CustomEventCard :event="event" />
      </div>
    </div>

    <div v-if="hasMoreEvents" class="text-center mt-6">
      <button
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-200"
        @click="loadMoreEvents"
      >
        Load More
      </button>
    </div>

    <div class="bg-violet-200 mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6 mx-12">
      <div class="px-4">
        <h2 class="text-xl font-bold mb-2 text-center">Categories</h2>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(category, index) in categories"
            :key="index"
            class="bg-blue-100 text-blue-700 px-3 py-1 rounded-full text-sm font-semibold"
          >
            {{ category.name }}
          </span>
        </div>
      </div>

      <!-- Tags -->
      <div>
        <h2 class="text-xl font-bold mb-2 text-center">Tags</h2>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(tag, index) in tags"
            :key="index"
            class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-sm font-semibold"
          >
            {{ tag.name }}
          </span>
        </div>
      </div>
    </div>

    <!-- Interactive Features and Create Event Section -->
    <div class="m-12">
      <div
        class="flex flex-col lg:flex-row items-center lg:items-start justify-between bg-gradient-to-r from-purple-300 via-blue-300 to-pink-300 p-6 rounded-lg shadow-lg"
      >
        <!-- Features Info Card -->
        <div class="mb-6 lg:mb-0 lg:w-2/3 bg-white p-6 rounded-lg shadow-md">
          <h2 class="text-2xl font-bold mb-6 text-gray-800">
            Explore Amazing Features
          </h2>
          <div class="space-y-4">
            <!-- Feature 1 -->
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-blue-500 text-white rounded-full p-3"
              >
                <!-- Icon for customization -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 8v4l3 3m6-7v10a2 2 0 01-2 2H6a2 2 0 01-2-2V7a2 2 0 012-2h7m-2-4h4m-2 4v4m-2 4l3 3"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Customize Your Event Listings
                </h3>
                <p class="text-gray-500">
                  Tailor your events with flexible settings and options.
                </p>
              </div>
            </div>

            <!-- Feature 2 -->
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-green-500 text-white rounded-full p-3"
              >
                <!-- Icon for community -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 20h5V10a2 2 0 00-2-2h-3a2 2 0 00-2 2v2m-8 10h5m-3-10a2 2 0 00-2-2H4a2 2 0 00-2 2v2a2 2 0 002 2h5a2 2 0 012-2v-2"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Connect with a Vibrant Community
                </h3>
                <p class="text-gray-500">
                  Engage with users and share experiences.
                </p>
              </div>
            </div>

            <!-- Feature 3 -->
            <div class="flex items-center">
              <div
                class="flex-shrink-0 bg-yellow-500 text-white rounded-full p-3"
              >
                <!-- Icon for management -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8 17v5m8-5v5M12 14V9m0-2a4 4 0 00-4 4v2h8V9a4 4 0 00-4-4z"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Manage Attendance Seamlessly
                </h3>
                <p class="text-gray-500">
                  Keep track of participants and event logistics.
                </p>
              </div>
            </div>

            <!-- Feature 4 -->
            <div class="flex items-center">
              <div class="flex-shrink-0 bg-red-500 text-white rounded-full p-3">
                <!-- Icon for analytics -->
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 19V6a1 1 0 011-1h1a1 1 0 011 1v13a1 1 0 01-1 1h-1a1 1 0 01-1-1zm-4 0v-7a1 1 0 011-1h1a1 1 0 011 1v7a1 1 0 01-1 1H8a1 1 0 01-1-1zm8 0v-9a1 1 0 011-1h1a1 1 0 011 1v9a1 1 0 01-1 1h-1a1 1 0 01-1-1z"
                  ></path>
                </svg>
              </div>
              <div class="ml-4">
                <h3 class="font-semibold text-lg text-gray-700">
                  Track Event Analytics
                </h3>
                <p class="text-gray-500">
                  Gain insights with real-time data and reports.
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="lg:w-1/3 flex justify-center lg:justify-end">
          <NuxtLink to="/user/events/create-event"
            ><button
              class="bg-blue-600 text-white px-6 py-3 rounded-full font-semibold hover:bg-blue-700 transition-colors duration-200 mt-6 lg:mt-0"
            >
              Create Your Event Here
            </button></NuxtLink
          >
        </div>
      </div>
    </div>
    <CustomFooter />
  </div>
</template>
<script setup>
import CustomEventCard from "@/components/CustomEventCard.vue";
import { ref, watch, computed } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";

const events = ref([]);
const visibleEvents = ref([]);
const { result, loading, error } = useQuery(getEvents);

const itemsPerPage = 3;
const currentPage = ref(1);

const categories = ref([
  {
    description:
      "Experience the world of art through gallery exhibitions, theater performances, film screenings, and cultural festivals.",
    id: "c7ad0fc6-5572-4274-88ff-e3fbd7a20509",
    name: "Art & Culture",
  },
  {
    description:
      "Dive into the world of literature with book signings, author talks, poetry readings, and literary festivals.",
    id: "171f72bc-3318-4f5a-aa93-e8c02027cad3",
    name: "Books & Litrature",
  },
  {
    description:
      "Attend conferences, workshops, and networking events aimed at entrepreneurs, business professionals, and startups.",
    id: "0cd0f9b7-2f5a-4913-97e0-ed4d71da5d2c",
    name: "Bussiness & Networking",
  },
  {
    description:
      "Support various causes by participating in charity runs, fundraising events, and community outreach programs.",
    id: "c83c4305-046a-463d-ad56-99cd767007c0",
    name: "Charity & Causes",
  },
  {
    description:
      "Enroll in educational workshops, training programs, seminars, and courses that enhance your skills and knowledge.",
    id: "a5f376d2-da8a-4d53-8191-2579cdb3ecbc",
    name: "Education & Training",
  },
  {
    description:
      "Explore the latest trends in fashion with runway shows, beauty workshops, and fashion design exhibitions.",
    id: "36b4bf60-0528-4515-bd49-efc20a744fbd",
    name: "Fasion & Beauty",
  },
  {
    description:
      "Celebrate life with cultural festivals, holiday events, and themed celebrations that bring people together.",
    id: "0fa6b097-e484-4d8b-a8a6-c9cf825847c8",
    name: "Festival & Celebrations",
  },
  {
    description:
      " Discover a variety of culinary events including food festivals, tasting sessions, and cooking classes that explore global cuisines and local flavors.",
    id: "bc07826a-1bba-4234-8ab7-64ec827e0d1d",
    name: "Food & Beverage",
  },
  {
    description:
      "Events promoting a healthy lifestyle, such as yoga classes, fitness bootcamps, meditation workshops, and wellness retreats.",
    id: "a808b5dd-fc73-4d40-896a-9cc678ddf153",
    name: "Health & Wellness",
  },
  {
    description:
      "Enjoy family-friendly events like carnivals, educational workshops, and kid-focused recreational activities.",
    id: "e8671982-944b-402c-8d8c-589853ab5c6c",
    name: "Kids & Family",
  },
  {
    description:
      "From intimate acoustic sessions to large-scale music festivals, find concerts and musical performances across diverse genres.",
    id: "9f2e6230-beea-4ac6-89ff-a083bacd6bba",
    name: "Music & Concerts",
  },
  {
    description:
      "Enjoy a vibrant nightlife scene with club events, themed parties, and late-night gatherings.",
    id: "afabcfdb-adf1-42a5-8bfd-fac3bf9f05d0",
    name: "Nightlife & Party",
  },
  {
    description:
      "Explore events related to the natural world and science, including ecological tours, astronomy nights, and science fairs.",
    id: "b27df5bd-e175-403d-b76e-b11f3c3d2141",
    name: "Science & Nature",
  },
  {
    description:
      "Participate in or watch events like marathons, outdoor adventures, sports competitions, and recreational activities.",
    id: "eebe4dc3-72ff-43ea-b7bc-8d3dcfd0f5d8",
    name: "Sports & Outdoors",
  },
  {
    description:
      "Keep up with the latest in tech with hackathons, coding bootcamps, AI conferences, and seminars on emerging trends.",
    id: "82ba38a7-f3ad-40ed-8dee-1dea5fad28b0",
    name: "Technology & Innovation",
  },
  {
    description:
      "Discover travel experiences, from guided tours and adventure trips to travel expos showcasing exciting destinations.",
    id: "06c9194b-3faa-4ab2-bb26-4439d1c6c0e0",
    name: "Travel & Venture",
  },
  {
    description:
      "Explore hands-on workshops and classes focused on DIY crafts, creative hobbies, and skill-building activities.",
    id: "1f90bd11-5447-4354-8a2b-4db165a0bbb8",
    name: "Workshops & Hobbies",
  },
]);
const tags = ref([
  {
    description:
      "Outdoor adventure events, such as hiking, camping, and extreme sports.",
    id: "1053c440-3a22-490c-b308-ffde5834d59f",
    name: "Advanture",
  },
  {
    description:
      "Events discussing artificial intelligence and machine learning.",
    id: "89802cbd-70f7-4577-b88e-e4e16dad3035",
    name: "AI",
  },
  {
    description:
      "Conferences and workshops focused on blockchain technology and cryptocurrency.",
    id: "5336362e-a130-4f69-85e8-ddb56d23ac37",
    name: "Blockchain",
  },
  {
    description:
      "Courses offering professional certifications in various fields.",
    id: "e85c8bb5-ce45-4670-8a32-ddac4acc7f16",
    name: "Certefication",
  },
  {
    description: "Fundraising events and charity drives for various causes.",
    id: "5744796d-2499-4895-ab67-2ef379c5ec3c",
    name: "Charity",
  },
  {
    description:
      "Events focused on Christmas celebrations, markets, and parties.",
    id: "7476d9a0-120c-4460-9a28-8a5aa61f312c",
    name: "Christmas",
  },
  {
    description: "Hackathons, coding bootcamps, and developer meetups.",
    id: "48eaae4e-dcf0-4b14-bd09-594d875ef7d2",
    name: "Coding",
  },
  {
    description: " Events with DJ performances, parties, and music festivals.",
    id: "301dcd5b-3452-41eb-a99b-2da1a123245f",
    name: "DJ",
  },
  {
    description: "Art galleries, exhibitions, and visual arts events.",
    id: "62a9e698-4222-439f-a7e4-10ed6ff5c7a4",
    name: "Exhibition",
  },
  {
    description:
      "Large-scale events like music festivals and cultural festivals.",
    id: "148f9fa2-ec55-4ff7-94c6-784726c4be7c",
    name: "Festival",
  },
  {
    description: "Movie screenings, film festivals, and cinematic events.",
    id: "3d748c36-acc4-4204-a26d-6b3df410d98c",
    name: "Film",
  },
  {
    description: "Large food and drink events showcasing different cuisines.",
    id: "36e9db80-c57b-4c4c-abda-4b2d06317aa7",
    name: "Food festival",
  },
  {
    description:
      "Seminars and workshops focused on leadership skills and personal development.",
    id: "e573c64a-4eb9-427c-bee3-f5a80ed5c921",
    name: "Leadership",
  },
  {
    description:
      "Live Music: Events featuring live band performances and concerts.",
    id: "d15f957f-94c1-4caa-bc61-18b2887cca7c",
    name: "Live Music",
  },
  {
    description:
      "Conferences and sessions on digital marketing, branding, and advertising.",
    id: "f305dfaf-bc85-4b01-b803-252f0978fd2f",
    name: "Marketing",
  },
  {
    description:
      "Parties, countdowns, and celebrations ringing in the new year.",
    id: "acd8b525-1e68-4bfe-9213-95eca4c2f239",
    name: "New Year’s Eve",
  },
  {
    description:
      "Events featuring photography exhibitions, walks, and workshops.",
    id: "4bbbf383-178b-414d-8f65-b817b3a945bb",
    name: "Photography",
  },
  {
    description:
      "Events tailored for startups, including pitch nights and entrepreneur networking.",
    id: "8b2400b3-666e-4688-9d53-b4e91ccfb446",
    name: "Startup",
  },
  {
    description: "Travel expos, tours, and destination-focused events.",
    id: "86ca238d-64e7-4a1b-adf2-8f34fa854ee5",
    name: "Travel",
  },
  {
    description: "Events centered around giving back and community service.",
    id: "b8c3dcac-e244-4175-8736-263267e93eae",
    name: "Volunteering",
  },
  {
    description:
      "High-energy fitness sessions, boot camps, and group exercises.",
    id: "39b14fe1-f80f-4a93-9016-1f1c68bdfd8f",
    name: "Workout",
  },
  {
    description: "Interactive training sessions and skill-building workshops.",
    id: "4ff6b69a-d686-4d5d-a8c7-5932c2d1a693",
    name: "Workshop",
  },
  {
    description: "Events and classes focused on yoga practices.",
    id: "00fe46d4-e118-41ca-a2e8-b4907ad0b250",
    name: "Yoga",
  },
]);

watch(result, (newResult) => {
  if (newResult?.events) {
    events.value = newResult.events;
    updateVisibleEvents();
  }
});

watch(loading, (newLoading) => {
  console.log("Loading:", newLoading);
});

watch(error, (newError) => {
  console.error("Error:", newError);
});

const updateVisibleEvents = () => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  visibleEvents.value = events.value.slice(
    startIndex,
    startIndex + itemsPerPage
  );
};

const loadMoreEvents = () => {
  if (hasMoreEvents) {
    currentPage.value += 1;
    updateVisibleEvents();
  }
};

const hasMoreEvents = computed(
  () => events.value.length > visibleEvents.value.length
);

definePageMeta({ layout: "authenticated" });
</script>
