<template>
  <div class="pt-4 md:pt-24">
    <div class="flex flex-col gap-8">
      <div class="flex flex-col lg:flex-row gap-8 mx-12">
        <div class="lg:w-1/2 bg-gray-100 p-6 rounded-lg shadow-lg">
          <img
            :src="event?.image?.length > 0 ? event.image[0] : defaultImage"
            alt="Event Image"
            class="w-full h-64 object-cover rounded-lg shadow-lg"
          />

          <div class="grid grid-cols-2 gap-2 mt-4">
            <div
              v-for="(image, index) in event?.image?.slice(1) || []"
              :key="index"
              class="h-32"
            >
              <img
                :src="image"
                alt="Additional Event Image"
                class="w-full h-full object-cover rounded-lg shadow-md"
              />
            </div>
          </div>
        </div>

        <div
          class="lg:w-1/2 bg-gray-100 p-6 rounded-lg shadow-lg flex flex-col"
        >
          <div class="flex-grow">
            <h1 class="text-3xl font-bold mb-4">{{ event?.title || "" }}</h1>
            <p class="text-lg mb-2">
              <strong>Description:</strong> {{ event?.description || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>Address:</strong> {{ event?.address || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>Start Date:</strong> {{ event?.start_date || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>End Date:</strong> {{ event?.end_date || "" }}
            </p>
            <p :class="[freeClass, 'inline-block']">{{ isFreeText }}</p>
          </div>
          <div class="p-4 flex justify-between items-center">
            <div
              @click="toggleFollow"
              :class="{
                'bg-gray-300': isFollowing,
                'bg-gray-500': !isFollowing,
              }"
              class="px-4 py-2 rounded text-white cursor-pointer"
            >
              <span>{{ isFollowing ? "Following" : "Follow" }}</span>
            </div>

            <div class="flex items-center justify-center h-full">
              <font-awesome-icon
                :icon="['fas', 'bookmark']"
                :class="[
                  'cursor-pointer',
                  {
                    'text-gray-400': isBookmarked,
                    'text-gray-700': !isBookmarked,
                  },
                ]"
                @click="toggleBookmark"
              />
            </div>

            <div
              v-if="!event?.isFree"
              @click="toggleBuy"
              :class="{ 'bg-green-500': isBought, 'bg-red-300': !isBought }"
              class="px-4 py-2 rounded text-white cursor-pointer"
            >
              <span>{{ isBought ? "Bought" : "Buy" }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-gray-100 mx-12">
        <p class="text-lg text-center pb-4"><strong>Event Location</strong></p>
        <DisplayLeafletMap
          :latitude="event?.location?.latitude"
          :longitude="event?.location?.longitude"
          :address="event?.address"
        />
      </div>
    </div>
    <CustomFooter />
  </div>
</template>
<script setup>
import { ref, computed, watch } from "vue";
import { useRoute } from "vue-router";
import { useQuery } from "@vue/apollo-composable";
import defaultImage from "@/assets/images/home.png";
import DisplayLeafletMap from "~/components/DisplayLeafletMap.vue";
import getEvent from "~/graphql/queries/events/getSingleEvent.gql";

const event = ref(null);
const isBookmarked = ref(false);
const isFollowing = ref(false);
const isBought = ref(false);

const route = useRoute();
const { id } = route.params;

// Use Apollo's useQuery to fetch the event data
const { result, loading, error } = useQuery(getEvent, { id });

// Watch for changes in the result to assign the event
watch(
  () => result.value,
  (newValue) => {
    if (newValue && newValue.events_by_pk) {
      event.value = newValue.events_by_pk;
    }
  }
);

const toggleFollow = () => {
  isFollowing.value = !isFollowing.value;
};

const toggleBookmark = () => {
  isBookmarked.value = !isBookmarked.value;
};

const toggleBuy = () => {
  isBought.value = !isBought.value;
};

// Computed properties for display logic
const isFreeText = computed(() => (event.value?.isFree ? "Free" : "Paid"));

const freeClass = computed(() =>
  event.value?.isFree
    ? "bg-green-300 text-white px-2 py-1 rounded"
    : "bg-red-800 text-white px-2 py-1 rounded"
);
</script>
