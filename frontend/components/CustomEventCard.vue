<template>
  <div
    :class="[
      'bg-white shadow-lg rounded-lg overflow-hidden flex flex-col',
      'transition-transform duration-300 ease-in-out transform hover:scale-105',
      customClass,
    ]"
    class="w-full max-w-md mx-auto p-4 hover:bg-indigo-50"
  >
    <NuxtLink :to="`/user/events/${event.id}`">
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
            'cursor-pointer transition-colors duration-200',
            {
              'text-yellow-500': bookmark,
              'text-gray-400': !bookmark,
            },
          ]"
          @click="toggleBookmark"
        />
      </div>

      <div
        @click="toggleBuyTicket"
        v-if="!isFree"
        :class="{
          'bg-gray-200': isBought,
          'bg-green-500': !isBought,
          'bg-red-300': !isticketavailable,
        }"
        class="px-4 py-2 rounded text-white cursor-pointer"
      >
        <span>{{
          isTicketAvailable ? (isBought ? "Bought" : "Buy") : "Not available"
        }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import defaultImage from "~/assets/images/home.png";
import { useAuthStore } from "~/stores";
import { useRouter } from "vue-router";
import { InsertFollowers } from "~/graphql/mutations/event_followers/InsertFollowers.gql";
import { DeleteFollowers } from "~/graphql/mutations/event_followers/DeleteFollowers.gql";
import { InsertBookmarks } from "~/graphql/mutations/event_bookmarks/InsertBookmarks.gql";
import { DeleteBookmarks } from "~/graphql/mutations/event_bookmarks/DeleteBookmarks.gql";
import { InsertTickets } from "~/graphql/mutations/event_tickets/InsertTickets.gql";
import { DeleteTickets } from "~/graphql/mutations/event_tickets/DeleteTickets.gql";

const authStore = useAuthStore();
const isAuthenticated = authStore.isAuthenticated;
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

const isFree = ref(props.event?.is_free);

const isTicketAvailable = ref(
  props.event?.events_tickets?.quantity &&
    props.event?.events_tickets?.quantity < props.event?.capacity
);

const isticketavailable = isTicketAvailable.value;
const isBought = ref(
  props.event?.events_tickets?.some((ticket) => {
    ticket?.user_id === user_id;
  })
);

const following = ref(
  props.event?.events_followers?.some((follow) => follow?.user_id === user_id)
);
const bookmark = ref(
  props.event?.events_bookmarks?.some(
    (bookmark) => bookmark?.user_id === user_id
  )
);

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

// Buy Ticket
const {
  mutate: InsertBuyTicketMutation,
  onDone: onInsertBuyTicketDone,
  loading: loadingBuyTicket,
  onError: onInsertBuyTicketError,
} = useAuthenticatedMutation(InsertTickets);

const {
  mutate: DeleteBuyTicketMutation,
  onDone: onDeleteBuyTicketDone,
  loading: loadingDeleteBuyTicket,
  onError: onDeleteBuyTicketError,
} = useAuthenticatedMutation(DeleteTickets);

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
        event_id: props.event?.id,
      };
      await InsertBookmarkMutation(bookmarkPayload);
      bookmark.value = true;
    } else {
      const unBookmarkPayload = {
        user_id: user_id,
        event_id: props.event?.id,
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

// toggle buy

const toggleBuyTicket = async (quantity) => {
  if (!user_id) {
    return router.push("/users/login");
  }
  const isbought = isBought.value;

  try {
    if (!isbought || !isticketavailable) {
      const ticketPayload = {
        user_id: user_id,
        event_id: props.event?.id,
        quantity: quantity,
      };
      await InsertBuyTicketMutation(ticketPayload);
      isBought.value = true;
    } else if (isbought) {
      const cancelTicket = {
        user_id: user_id,
        event_id: props.event?.id,
      };
      await DeleteBuyTicketMutation(cancelTicket);
      isBought.value = false;
    } else return;
  } catch (error) {
    console.log(
      `Error occurred while ${isbought ? "cancelTicket" : "boy"} the event:`,
      error
    );
  }
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

definePageMeta({ layout: isAuthenticated ? "authenticated" : "" });
</script>
