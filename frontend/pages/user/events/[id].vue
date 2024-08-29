<template>
  <div
    class="relative bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 min-h-screen w-full"
  >
    <div class="relative w-full h-96 overflow-hidden">
      <img
        :src="event_images[0]"
        class="w-full h-full object-contain rounded-b-lg shadow-lg transition-transform duration-500 ease-in-out transform hover:scale-105"
        alt="Main Event Image"
      />
      <div
        class="absolute inset-0 bg-gradient-to-b from-transparent to-black opacity-60"
      ></div>
      <div class="absolute bottom-0 left-0 p-8">
        <h1 class="text-5xl text-white font-bold">{{ event?.title || "" }}</h1>
        <p class="text-xl text-gray-200 mt-4">{{ event?.description || "" }}</p>
      </div>
    </div>

    <div class="w-full flex flex-col lg:flex-row gap-8 mt-12 px-8">
      <div
        class="lg:w-1/2 bg-white p-6 rounded-lg shadow-lg flex flex-col justify-between"
      >
        <div class="grid grid-cols-2 gap-2 flex-grow">
          <div
            v-for="(image, index) in event_images.slice(1)"
            :key="index"
            class="flex items-center"
          >
            <img
              :src="image"
              alt="Additional Event Image"
              class="w-full h-full object-cover rounded-lg shadow-lg transition-transform duration-500 ease-in-out transform hover:scale-105"
            />
          </div>
        </div>
      </div>

      <div
        class="lg:w-2/3 bg-white p-6 rounded-lg shadow-lg flex flex-col justify-between"
      >
        <div class="flex-grow">
          <div class="flex justify-between items-center mb-6">
            <h2 class="text-3xl font-bold text-gray-800">Event Information</h2>
            <button
              v-if="ownEvent"
              @click="editEvent"
              class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-all duration-300"
            >
              Edit
            </button>
          </div>
          <p class="text-lg mb-6">
            <strong>Date:</strong> {{ event?.start_date || "" }} to
            {{ event?.end_date || "" }}
          </p>
          <p class="text-lg mb-6">
            <strong>Address:</strong> {{ event?.address || "" }}
          </p>
          <p class="text-lg mb-6">
            <strong>Venue:</strong> {{ event?.venue || "" }}
          </p>
          <p class="text-lg mb-6">
            <strong>Capacity:</strong> {{ event?.capacity || "N/A" }} available
            spaces
          </p>
          <p class="text-lg mb-6" v-if="event?.isFree === false">
            <strong>Price:</strong> ${{ event?.price || "N/A" }}
          </p>
          <p :class="[freeClass, 'inline-block']">{{ isFreeText }}</p>
        </div>

        <!-- Action Buttons -->
        <div class="mt-8 flex justify-between items-center">
          <button
            @click="toggleFollow"
            :class="{
              'bg-blue-500': following,
              'bg-gray-300': !following,
            }"
            class="px-4 py-2 rounded text-white font-semibold transition-all duration-300 hover:bg-blue-600"
          >
            {{ following ? "Following" : "Follow" }}
          </button>

          <font-awesome-icon
            :icon="['fas', 'bookmark']"
            :class="[
              'cursor-pointer text-2xl',
              {
                'text-yellow-500': bookmark,
                'text-gray-400': !bookmark,
              },
            ]"
            @click="toggleBookmark"
          />

          <button
            v-if="!event?.isFree"
            @click="toggleBuy"
            :class="{
              'bg-green-600': !isBought,
              'bg-gray-500': isBought,
            }"
            class="px-4 py-2 rounded text-white font-semibold transition-all duration-300 hover:bg-green-700"
          >
            {{ isBought ? "Bought" : "Buy" }}
          </button>
        </div>
      </div>
    </div>

    <div class="mt-12 bg-white mx-8 p-6 rounded-lg shadow-lg">
      <div class="flex justify-between items-center">
        <p class="text-xl font-semibold">
          <strong>Event Location</strong>
        </p>
        <button
          @click="redirectToGoogleMap"
          class="text-blue-500 hover:text-blue-700 transition-all duration-300"
        >
          Show on Google Map
        </button>
      </div>
      <DisplayLeafletMap
        :latitude="location[0]"
        :longitude="location[1]"
        :address="event?.address"
        class="mt-6"
      />
    </div>
    <div class="right-20 m-12"><button>Edit</button></div>

    <CustomFooter class="mt-12" />
  </div>
</template>

<script setup>
import { ref, computed, watch } from "vue";
import { useRoute } from "vue-router";
import { useQuery } from "@vue/apollo-composable";
import defaultImage from "@/assets/images/home.png";
import getEvent from "~/graphql/queries/events/getSingleEvent.gql";
import { useAuthStore } from "~/stores";
import { InsertFollowers } from "~/graphql/mutations/event_followers/InsertFollowers.gql";
import { DeleteFollowers } from "~/graphql/mutations/event_followers/DeleteFollowers.gql";
import { InsertBookmarks } from "~/graphql/mutations/event_bookmarks/InsertBookmarks.gql";
import { DeleteBookmarks } from "~/graphql/mutations/event_bookmarks/DeleteBookmarks.gql";

const authStore = useAuthStore();
const isAuthenticated = computed(() => authStore.isAuthenticated);

const user_id = useAuthStore().id;

const router = useRoute();
const event = ref(null);
const isBought = ref(false);

const route = useRoute();
const { id } = route.params;

const { result, loading, error } = useQuery(getEvent, { id });

const following = ref(false);
const bookmark = ref(false);

watch(
  result,
  (newValue) => {
    if (newValue && newValue.events_by_pk) {
      event.value = newValue.events_by_pk;

      following.value = event.value.events_followers?.some(
        (follow) => follow.user_id === user_id
      );

      bookmark.value = event.value.events_bookmarks?.some(
        (bookmark) => bookmark.user_id === user_id
      );
    }
  },
  { immediate: true }
);

const event_images = computed(() => {
  if (event.value?.event_images) {
    const imagesArray = event?.value.event_images
      .replace(/{|}/g, "")
      .split(",");
    return imagesArray.length > 0 ? imagesArray : [defaultImage];
  }
  return [defaultImage];
});

const location = computed(() => {
  if (event.value?.location) {
    return event.value.location
      .replace(/{|}/g, "")
      .split(",")
      .map((coord) => parseFloat(coord));
  }
  return [0, 0];
});

const ownEvent = event?.value?.uid == user_id


const {
  mutate: InsertFollowersMutation,
  onDone: onInsertFollowersDone,
  loading: loadingFollowers,
  onError: onInsertFollowersError,
} = useAuthenticatedMutation(InsertFollowers);

const {
  mutate: DeleteFollowersMutation,
  onDone: onDeleteFollowersDone,
  loading: loadingDeleteFollowers,
  onError: onDeleteFollowersError,
} = useAuthenticatedMutation(DeleteFollowers);

// bookmark
const {
  mutate: InsertBookmarkMutation,
  onDone: onInsertBookmarkDone,
  loading: loadingBookmark,
  onError: onInsertBookmarkError,
} = useAuthenticatedMutation(InsertBookmarks);

const {
  mutate: DeleteBookmarkMutation,
  onDone: onDeleteBookmarkDone,
  loading: loadingDeleteBookmark,
  onError: onDeleteBookmarkError,
} = useAuthenticatedMutation(DeleteBookmarks);

// Toggle following
const toggleFollow = async () => {
  if (!user_id) {
    return router.push("/users/login");
  }

  const isCurrentlyFollowing = following.value;
  try {
    if (!isCurrentlyFollowing) {
      const followingPayload = {
        user_id: user_id,
        event_id: event?.value.id,
      };
      await InsertFollowersMutation(followingPayload);
      following.value = true;
    } else {
      const unFollowingPayload = {
        user_id: user_id,
        event_id: event?.value.id,
      };
      await DeleteFollowersMutation(unFollowingPayload);
      following.value = false;
    }
  } catch (error) {
    console.log(
      `Error occurred while ${
        isCurrentlyFollowing ? "unfollowing" : "following"
      } the event:`,
      error
    );
  }
};

// toggle bookmark
const toggleBookmark = async () => {
  if (!user_id) {
    return router.push("/users/login");
  }
  const isCurrentlyBookmarked = bookmark.value;
  try {
    if (!isCurrentlyBookmarked) {
      const bookmarkPayload = {
        user_id: user_id,
        event_id: event?.value.id,
      };
      await InsertBookmarkMutation(bookmarkPayload);
      bookmark.value = true;
    } else {
      const unBookmarkPayload = {
        user_id: user_id,
        event_id: event?.value.id,
      };
      await DeleteBookmarkMutation(unBookmarkPayload);
      bookmark.value = false;
    }
  } catch (error) {
    console.log(
      `Error occurred while ${
        isCurrentlyBookmarked ? "unbookmarked" : "bookmark"
      } the event:`,
      error
    );
  }
};

const editEvent = () => {
  // Handle edit event logic here
  console.log("Edit event clicked");
};

const toggleBuy = () => {
  isBought.value = !isBought.value;
};

const isFreeText = computed(() => (event.value?.isFree ? "Free" : "Paid"));

const freeClass = computed(() =>
  event.value?.isFree
    ? "bg-green-300 text-white px-2 py-1 rounded"
    : "bg-red-800 text-white px-2 py-1 rounded"
);

const redirectToGoogleMap = () => {
  const [latitude, longitude] = location.value;
  const url = `https://www.google.com/maps?q=${latitude},${longitude}`;
  window.open(url, "_blank");
};

definePageMeta({ layout: "authenticated" });
</script>
