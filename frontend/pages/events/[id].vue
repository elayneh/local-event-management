<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 h-64 w-full"
  >
    <div class="w-full flex justify-center items-center py-24">
      <div>
        <img :src="event_images[0]" class="rounded-lg shadow-lg" />
      </div>
    </div>
    <div
      class="w-full flex-col gap-8 flex justify-center pt-24 pb-12 space-x-6"
    >
      <div class="flex flex-col lg:flex-row gap-8 mx-12">
        <div class="lg:w-1/2 bg-gray-100 p-6 rounded-lg shadow-lg">
          <img
            :src="event_images[0]"
            alt="Event Image"
            class="w-full h-64 object-cover animate-bounce rounded-lg shadow-lg transition-transform duration-500 ease-in-out transform hover:scale-105"
          />

          <div class="grid grid-cols-2 gap-2 mt-4">
            <div
              v-for="(image, index) in event_images.slice(1)"
              :key="index"
              class="h-32"
            >
              <img
                :src="image"
                alt="Additional Event Image"
                class="w-full animate-pulse h-full object-cover rounded shadow-lg transition-transform duration-500 ease-in-out transform hover:scale-105"
              />
            </div>
          </div>
        </div>

        <div
          class="lg:w-1/2 bg-gray-100 p-6 rounded-lg shadow-lg flex flex-col"
        >
          <div class="flex-grow">
            <h1 class="text-3xl font-bold mb-4 text-center mb-12">
              {{ event?.title || "" }}
            </h1>
            <p class="text-lg mb-12">
              {{ event?.description || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>Posted at </strong> {{ event?.start_date || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>The event will be closed at </strong>
              {{ event?.end_date || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>Address:</strong> {{ event?.address || "" }}
            </p>
            <p class="text-lg mb-2">
              <strong>Venue:</strong> {{ event?.venue || "" }}
            </p>
            <p class="text-lg mb-2">
              There are {{ event?.capacity }} available spaces
            </p>
            <p class="text-lg mb-2" v-if="event?.isFree === false">
              ${{ event?.price }}
            </p>
            <p :class="[freeClass, 'inline-block']">{{ isFreeText }}</p>
          </div>
          <div class="pt-24 flex justify-between items-center">
            <div
              @click="toggleFollow"
              :class="{ 'bg-gray-500': following, 'bg-gray-300': !following }"
              class="px-4 py-2 rounded text-white cursor-pointer"
            >
              <span>{{ following ? "Following" : "Follow" }}</span>
            </div>

            <div class="flex items-center justify-center h-full">
              <font-awesome-icon
                :icon="['fas', 'bookmark']"
                :class="[
                  'cursor-pointer',
                  {
                    'text-yellow-500': bookmark,
                    'text-gray-400': !bookmark,
                  },
                ]"
                @click="toggleBookmark"
              />
            </div>

            <div
              v-if="!event?.isFree"
              @click="toggleBuy"
              :class="{ 'bg-gray-500': isBought, 'bg-green-700': !isBought }"
              class="px-4 py-2 rounded text-white cursor-pointer"
            >
              <span>{{ isBought ? "Bought" : "Buy" }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-gray-100 mx-12">
        <p class="text-lg text-center pb-4">
          <strong>Event Location</strong>
          <button @click="redirectToGoogleMap" class="ml-2 text-blue-500">
            Show on Google Map
          </button>
        </p>
        <DisplayLeafletMap
          :latitude="location[0]"
          :longitude="location[1]"
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
import getEvent from "~/graphql/queries/events/getSingleEvent.gql";
import { useAuthStore } from "~/stores";
import { InsertFollowers } from "~/graphql/mutations/event_followers/insert.gql";
import { DeleteFollowers } from "~/graphql/mutations/event_followers/delete.gql";
import { InsertBookmarks } from "~/graphql/mutations/event_bookmarks/insert.gql";
import { DeleteBookmarks } from "~/graphql/mutations/event_bookmarks/delete.gql";

const isAuthenticated = useAuthStore().isAuthenticated;
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

      // Update the following state based on the loaded event data
      following.value = event.value.events_followers?.some(
        (follow) => follow.user_id === user_id
      );

      // Update the bookmark state based on the loaded event data
      bookmark.value = event.value.events_bookmarks?.some(
        (bookmark) => bookmark.user_id === user_id
      );
    }
  },
  { immediate: true }
);
console.log("EVENTS: ", event);

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
  return [0, 0]; // Default to [0, 0] or another valid default location
});

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

///////////////////
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

// definePageMeta(() => ({
//   layout: isAuthenticated ? "authenticated" : "",
// }));
</script>