<template>
  <div
    :class="[
      'bg-white shadow-lg rounded-lg overflow-hidden flex flex-col',
      'transition-transform duration-300 ease-in-out transform hover:scale-105',
      customClass,
    ]"
    class="w-full max-w-md mx-auto p-4 hover:bg-indigo-300"
  >
    <NuxtLink :to="`/events/${event.id}`">
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
    </NuxtLink>
    <div
      class="p-4 rounded-lg border border-gray-200 flex justify-between items-center"
    >
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
              'text-gray-400': isBookmarked.value,
              'text-gray-700': !isBookmarked.value,
            },
          ]"
          @click="toggleBookmark"
        />
      </div>

      <div
        @click="toggleBuy"
        v-if="!isFree.value"
        :class="{
          'bg-gray-200': isBought.value,
          'bg-green-500': !isBought.value,
        }"
        class="px-4 py-2 rounded text-white cursor-pointer"
      >
        <span>{{ isBought.value ? "Bought" : "Buy" }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import defaultImage from "~/assets/images/home.png";
import { useAuthStore } from "~/stores";
import { useRouter } from "vue-router";
import { InsertFollowers } from "~/graphql/mutations/event_followers/insert.gql";
import { DeleteFollowers } from "~/graphql/mutations/event_followers/delete.gql";

const authStore = useAuthStore();
const user_id = authStore.id;
const router = useRouter();

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
const isBought = ref(false);

const following = ref(
  props.event?.events_followers?.some((follow) => follow?.user_id === user_id)
);

console.log("Event Object: ", props.event);

console.log("EVents: ", props.event?.event_followers?.user_id, "\n", user_id);
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

const toggleFollow = async () => {
  if (!user_id) {
    return router.push("/users/login");
  }

  const isCurrentlyFollowing = following.value;
  try {
    if (!isCurrentlyFollowing) {
      const followingPayload = {
        user_id: user_id,
        event_id: props.event?.id,
      };
      await InsertFollowersMutation(followingPayload);
      following.value = true;
    } else {
      const unFollowingPayload = {
        user_id: user_id,
        event_id: props.event?.id,
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
      .split(",")
      .map((image) => image.trim());
    return imagesArray.length ? imagesArray : [defaultImage];
  }
  return [defaultImage];
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
