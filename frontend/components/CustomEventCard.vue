<template>
  <div
    :class="[
      'bg-white shadow-lg rounded-lg overflow-hidden flex flex-col',
      'transition-transform duration-300 ease-in-out transform hover:scale-105',
      customClass,
    ]"
    class="w-full max-w-md mx-auto p-4 hover:bg-indigo-300"
  >
    <img
      :src="event_images[0]"
      alt="Event Image"
      class="w-full h-48 object-cover rounded-lg"
    />

    <div class="text-black text-center py-4">
      <h2 class="text-xl font-semibold">{{ event.title }}</h2>
    </div>

    <div class="flex-1 p-4">
      <p class="text-gray-700 mb-2">
        <strong>Description:</strong> {{ truncatedDescription }}
      </p>
      <p class="text-gray-700 mb-2">
        <strong>Address:</strong> {{ event.address || "N/A" }}
      </p>
      <p class="text-gray-700 mb-2 my-4">
        <strong>End Date:</strong> {{ formatDate(event.end_date) || "N/A" }}
      </p>
      <p class="text-gray-700 mt-12">
        <span :class="freeClass">{{ isFreeText }}</span>
      </p>
    </div>
    <div
      class="p-4 rounded-lg border border-gray-200 flex justify-between items-center"
    >
      <div
        @click="toggleFollow"
        :class="{ 'bg-gray-500': isFollowing, 'bg-gray-300': !isFollowing }"
        class="px-4 py-2 rounded text-white cursor-pointer"
      >
        <span>{{ isFollowing ? "Following" : "Follow" }}</span>
      </div>

      <div class="flex items-center justify-center h-full">
        <font-awesome-icon
          :icon="['fas', 'bookmark']"
          :class="[
            'cursor-pointer',
            { 'text-gray-400': isBookmarked, 'text-gray-700': !isBookmarked },
          ]"
          @click="toggleBookmark"
        />
      </div>

      <div
        @click="toggleBuy"
        v-if="!isFree"
        :class="{ 'bg-gray-200': isBought, 'bg-green-500': !isBought }"
        class="px-4 py-2 rounded text-white cursor-pointer"
      >
        <span>{{ isBought ? "Bought" : "Buy" }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import defaultImage from "~/assets/images/home.png";

const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
  customClass: {
    type: String,
    default: "",
  },
});

const isBookmarked = ref(false);
const isFollowing = ref(false);
const isBought = ref(false);


const toggleFollow = () => {
  i()
  isFollowing.value = !isFollowing.value;
};
const toggleBookmark = () => {
  isBookmarked.value = !isBookmarked.value;
};
const toggleBuy = () => {
  isBought.value = !isBought.value;
};

const truncatedDescription = computed(() => {
  const maxLength = 100;
  return props.event?.description?.length > maxLength
    ? props.event.description.slice(0, maxLength) + "..."
    : props.event?.description || "N/A";
});

const event_images = computed(() => {
  if (props.event?.event_images) {
    const imagesArray = props.event.event_images
      .replace(/{|}/g, "")
      .split(",");
    return imagesArray || defaultImage;
  }
  return defaultImage;
});

const isFree = computed(() => props.event?.is_free);

const isFreeText = computed(() => (isFree.value ? "Free" : "Paid"));

const freeClass = computed(() =>
  isFree.value
    ? "bg-green-300 text-white px-2 py-1 rounded"
    : "bg-red-800 text-white px-2 py-1 rounded"
);

const formatDate = (dateString) => {
  if (!dateString) return "N/A";
  const date = new Date(dateString);
  return date.toLocaleDateString();
};
</script>
