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
      >
        <div
          class="flex flex-col absolute bottom-0 justify-center items-center p-8 justify-center"
        >
          <h1 class="text-5xl text-white font-bold">
            {{ event?.title || "" }}
          </h1>
          <p class="text-xl text-gray-200 mt-4">
            {{ event?.description || "" }}
          </p>
        </div>
      </div>
    </div>
    <div v-if="!isEditing">
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
              <h2 class="text-3xl font-bold text-gray-800">
                Event Information
              </h2>
              <button
                v-if="ownEvent && !isEditing"
                @click="editHandler"
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
              <strong>Capacity:</strong>
              {{ event?.capacity || "N/A" }} available spaces
            </p>
            <p class="text-lg mb-6" v-if="event?.is_free === false">
              <strong>Price:</strong> ${{ event?.price || "N/A" }}
            </p>
            <p v-if="isClosed"><strong>Closed</strong></p>
            <p :class="[freeClass, 'inline-block']">{{ isFreeText }}</p>
          </div>

          <div v-if="!ownEvent" class="mt-8 flex justify-between items-center">
            <button
              @click="toggleFollow"
              :class="{ 'bg-gray-500': following, 'bg-gray-300': !following }"
              class="px-4 py-2 rounded text-white font-semibold transition-all duration-300 hover:bg-blue-600"
            >
              {{ following ? "Following" : "Follow" }}
            </button>

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

            <button
              v-if="!event?.is_free"
              @click="toggleBuy"
              :class="{
                'bg-green-600': !isBought,
                'bg-gray-500': isBought,
              }"
              class="px-4 py-2 rounded text-white font-semibold transition-all duration-300 hover:bg-green-700"
            >
              {{ isBought ? "Bought" : "Buy" }}
            </button>
            <button
              v-else
              @click="toggleBuy"
              :class="{
                'bg-green-600': !isReserved,
                'bg-gray-500': isReserved,
              }"
              class="px-4 py-2 rounded text-white font-semibold transition-all duration-300 hover:bg-green-700"
            >
              {{ isReserved ? "Reserved" : "Reserve" }}
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
    </div>

    <div
      v-if="isEditing"
      class="flex justify-center w-full mb-64 overflow-x-auto px-32"
    >
      <div class="min-w-[600px] bg-white p-6 rounded-lg shadow-lg px-12">
        <Form
          v-if="isEditing"
          :validation-schema="schema"
          @submit="saveEventChanges"
          class="px-8 py-6 bg-white rounded-lg shadow-lg mt-12"
          v-slot="{ errors }"
        >
          <h2
            class="flex text-3xl font-extrabold mb-6 relative text-transparent bg-clip-text bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 justify-center"
          >
            <span
              class="absolute -top-1 left-0 w-full h-1 bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500"
            ></span>
            <svg
              class="absolute -top-4 left-0 w-6 h-6 text-blue-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 12l2 2 4-4 4 4 8-8"
              ></path>
            </svg>
            Edit Event Informations
          </h2>

          <div class="mb-4">
            <label for="title" class="block text-lg font-medium mb-2"
              >Title</label
            >
            <Field
              v-model="editableEvent.title"
              id="title"
              name="title"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Title"
            />
            <span class="text-red-500 text-sm mt-1">{{ errors.title }}</span>
          </div>

          <div class="mb-4">
            <label for="description" class="block text-lg font-medium mb-2"
              >Description</label
            >
            <Field
              id="description"
              v-model="editableEvent.description"
              name="description"
              as="textarea"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Description"
            />
            <span class="text-red-500 text-sm mt-1">{{
              errors.description
            }}</span>
          </div>

          <div class="mb-4">
            <label for="start_date" class="block text-lg font-medium mb-2"
              >Start Date</label
            >
            <Field
              v-model="editableEvent.start_date"
              id="start_date"
              name="start_date"
              type="datetime-local"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
            />
            <span class="text-red-500 text-sm mt-1">{{
              errors.start_date
            }}</span>
          </div>

          <div class="mb-4">
            <label for="end_date" class="block text-lg font-medium mb-2"
              >End Date</label
            >
            <Field
              v-model="editableEvent.end_date"
              id="end_date"
              name="end_date"
              type="datetime-local"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
            />
            <span class="text-red-500 text-sm mt-1">{{ errors.end_date }}</span>
          </div>

          <!-- <div class="mb-4">
            <label for="address" class="block text-lg font-medium mb-2"
              >Address</label
            >
            <Field
              v-model="editableEvent.address"
              id="address"
              name="address"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Address"
            />
            <ErrorMessage name="address" class="text-red-500 text-sm mt-1" />
          </div> -->

          <div class="mb-4">
            <label for="venue" class="block text-lg font-medium mb-2"
              >Venue</label
            >
            <Field
              v-model="editableEvent.venue"
              id="venue"
              name="venue"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Venue"
            />
            <span class="text-red-500 text-sm mt-1">{{ errors.venue }}</span>
          </div>

          <div class="mb-4">
            <label for="capacity" class="block text-lg font-medium mb-2"
              >Capacity</label
            >
            <Field
              v-model="editableEvent.capacity"
              id="capacity"
              name="capacity"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Capacity"
            />
            <span class="text-red-500 text-sm mt-1">{{ errors.capacity }}</span>
          </div>

          <div class="mb-4">
            <label for="is_free" class="block text-lg font-medium mb-2"
              >Is free</label
            >
            <div class="flex items-center space-x-4">
              <label class="inline-flex items-center">
                <Field
                  v-model="editableEvent.is_free"
                  id="is_freeTrue"
                  name="is_free"
                  type="radio"
                  :value="true"
                  class="form-radio text-blue-600"
                />
                <span class="ml-2">Yes</span>
              </label>
              <label class="inline-flex items-center">
                <Field
                  v-model="editableEvent.is_free"
                  id="is_freeFalse"
                  name="is_free"
                  type="radio"
                  :value="false"
                  class="form-radio text-blue-600"
                />
                <span class="ml-2">No</span>
              </label>
            </div>
            <span class="text-red-500 text-sm mt-1">{{ errors.is_free }}</span>
          </div>

          <div class="mb-4" v-if="!editableEvent.is_free">
            <label for="price" class="block text-lg font-medium mb-2"
              >Price</label
            >
            <Field
              v-model="editableEvent.price"
              id="price"
              name="price"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Price"
            />
            <span class="text-red-500 text-sm mt-1">{{ errors.price }}</span>
          </div>

          <!-- <div class="mb-4" v-if="!values.is_free">
            <label for="price" class="block text-lg font-medium mb-2"
              >Price</label
            >
            <Field
              v-model="editableEvent.price"
              id="price"
              name="price"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              placeholder="Event Price"
            />
            <ErrorMessage name="price" class="text-red-500 text-sm mt-1" />
          </div> -->
          <div class="mt-12 flex justify-center items-center">
            <button
              type="submit"
              class="bg-green-600 text-white px-6 py-2 rounded-lg shadow hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 mr-32"
            >
              Save
            </button>
            <button
              type="button"
              @click="cancelEdit"
              class="bg-red-600 text-white px-6 py-2 rounded-lg shadow hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
            >
              Cancel
            </button>
          </div>
        </Form>
      </div>
    </div>

    <CustomFooter class="mt-auto" />
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useRoute } from "vue-router";
import { useQuery } from "@vue/apollo-composable";
import { useForm, Field, Form, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import defaultImage from "@/assets/images/home.png";
import getEvent from "~/graphql/queries/events/getSingleEvent.gql";
import { useAuthStore } from "~/stores";
import { InsertFollowers } from "~/graphql/mutations/event_followers/InsertFollowers.gql";
import { DeleteFollowers } from "~/graphql/mutations/event_followers/DeleteFollowers.gql";
import { InsertBookmarks } from "~/graphql/mutations/event_bookmarks/InsertBookmarks.gql";
import { DeleteBookmarks } from "~/graphql/mutations/event_bookmarks/DeleteBookmarks.gql";
import { UpdateEvent } from "~/graphql/mutations/events/update.gql";
import { toast } from "vue3-toastify";

const authStore = useAuthStore();
const isAuthenticated = computed(() => authStore.isAuthenticated);
const user_id = authStore.id;

const router = useRoute();
const event = ref({});
const isBought = ref(false);
const isReserved = ref(false);
const ownEvent = ref(false);
const isClosed = ref(false);
const following = ref(false);
const bookmark = ref(false);
const isEditing = ref(false);
const editableEvent = ref({});

const { result, loading, error, onResult } = useQuery(getEvent, {
  id: router.params.id,
});
onResult((response) => {
  const newResult = response.data?.events_by_pk;
  if (newResult) {
    event.value = newResult;
    editableEvent.value = { ...event.value };
    ownEvent.value = event?.value?.uid == user_id;
    isClosed.value = new Date(event?.value?.end_date) < Date.now();

    following.value = event.value.events_followers?.some(
      (follow) => follow?.user_id === user_id
    );

    bookmark.value = event.value.events_bookmarks?.some(
      (bookmark) => bookmark?.user_id === user_id
    );
  } else {
    console.error.call("Error: events_by_pk not found");
  }
  console.log("Events: ", event);
});

const { mutate: updateEventMutation } = useMutation(UpdateEvent);

const editHandler = () => {
  isEditing.value = true;
  editableEvent.value = { ...event.value };
};

const cancelEdit = () => {
  isEditing.value = false;
};

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

const toggleBuy = () => {
  isBought.value = !isBought.value;
  isReserved.value = !isReserved.value;
};

const isFreeText = computed(() => (event?.value?.is_free ? "Free" : "Paid"));

const freeClass = computed(() =>
  event.value?.is_free
    ? "bg-green-300 text-white px-2 py-1 rounded"
    : "bg-red-800 text-white px-2 py-1 rounded"
);

const redirectToGoogleMap = () => {
  const [latitude, longitude] = location.value;
  const url = `https://www.google.com/maps?q=${latitude},${longitude}`;
  window.open(url, "_blank");
};

// Form Setup
const schema = yup.object({
  title: yup.string().required("Title is required"),
  description: yup.string().required("Description is required"),
  start_date: yup.date().required("Start date is required"),
  end_date: yup.date().required("End date is required"),
  // address: yup.string().required("Address is required"),
  venue: yup.string().required("Venue is required"),
  capacity: yup
    .number()
    .required("Capacity is required")
    .positive("Capacity must be positive"),
  price: yup.number().nullable(),
});

// const { resetForm, handleSubmit, errors, isSubmitting, setFieldValue, values } =
//   useForm({
//     validationSchema: schema,
//     initialValues: {
//       title: event.value?.title || "",
//       description: event.value?.description || "",
//       start_date: event.value?.start_date || "",
//       end_date: event.value?.end_date || "",
//       // address: event.value?.address || "",
//       venue: event.value?.venue || "",
//       capacity: event.value?.capacity || "",
//       price: event.value?.price || 0,
//       is_free: event.value?.is_free || false,
//     },
//   });

const saveEventChanges = async (values) => {
  console.log("Values: ", values);
  try {
    const { data } = await updateEventMutation({
      id: event.value.id,
      price: values.price || 0,
      description: values.description,
      title: values.title,
      is_free: values.is_free,
      venue: values.venue,
      capacity: values.capacity,
      end_date: values.end_date,
      start_date: values.start_date,
    });
    console.log("Updatting");
    event.value = { ...editableEvent.value };

    toast.success("The Event Updated Successfully!");
    isEditing.value = false;
  } catch (err) {
    toast.error("Error Updating the event, try again");
    console.error("Error updating the event: ", err);
  }
};

definePageMeta({ layout: "authenticated" });
</script>
