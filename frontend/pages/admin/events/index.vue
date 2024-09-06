<script setup>
import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import deleteEvent from "~/graphql/mutations/events/delete.gql";
import useAuthenticatedMutation from "~/composables/useAuthenticatedMutation";
import { toast } from "vue3-toastify";

const events = ref([]);

function formatDate(dateString) {
  const options = {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  };
  return new Date(dateString).toLocaleDateString(undefined, options);
}
const limit = ref(10);
const offset = ref(0);
const where = ref(null);

const { onResult: onEventsResult } = useQuery(getEvents, {
  limit: limit.value,
  offset: offset.value,
  where: where.value,
});

const { mutate: deleteEventMutation } = useAuthenticatedMutation(deleteEvent);

onEventsResult((response) => {
  const newResult = response.data?.events;
  if (newResult) {
    events.value = newResult;
  }
});

const handleDeleteEvent = async (event_id) => {
  try {
    const { data } = await deleteEventMutation(
      { id: event_id },
      {
        onDone: () => {
          events.value = events.value.filter((event) => event.id !== event_id);
        },
      }
    );

    if (data) {
      toast.success("Event deleted successfully", {
        position: "top-right",
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
      });
    }
  } catch (error) {
    console.error("Error deleting event:", error);
    toast.error("Failed to delete event", {
      position: "top-right",
      autoClose: 3000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
    });
  }
};

const freezeEvent = (event_id) => {
  console.log(`Event with ID ${event_id} freezed successfully`);
};

definePageMeta({ layout: "admin-dashboard" });
</script>

<template>
  <div class="p-6 bg-white shadow-md rounded-lg">
    <h2 class="text-2xl font-bold mb-4 text-gray-800">
      Admin Event Management
    </h2>
    <table
      class="min-w-full bg-gray-100 border border-gray-300 rounded-lg overflow-hidden"
    >
      <thead>
        <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
          <th class="py-3 px-4 border-b">#</th>
          <th class="py-3 px-4 border-b">Title</th>
          <th class="py-3 px-4 border-b">Description</th>
          <th class="py-3 px-4 border-b">Created At</th>
          <th class="py-3 px-4 border-b text-center">Freeze</th>
          <th class="py-3 px-4 border-b text-center">Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(event, index) in events"
          :key="event.id"
          class="border-t transition-transform duration-300 ease-in-out transform hover:scale-105 hover:bg-gray-50"
        >
          <td class="py-2 px-4 border">{{ index + 1 }}</td>
          <td class="py-2 px-4 border">{{ event.title }}</td>
          <td class="py-2 px-4 border">{{ event.description }}</td>
          <td class="py-2 px-4 border">{{ formatDate(event?.created_at) }}</td>
          <td class="py-2 px-4 border text-center">
            <button
              @click="freezeEvent(event.id)"
              class="text-blue-500 hover:text-blue-700 transition-colors duration-200 transform hover:scale-110"
            >
              <i class="fa fa-snowflake" aria-hidden="true"></i>
            </button>
          </td>
          <td class="py-2 px-4 border text-center">
            <button
              @click="handleDeleteEvent(event.id)"
              class="text-red-500 hover:text-red-700 transition-colors duration-200 transform hover:scale-110"
            >
              <i class="fa fa-trash" aria-hidden="true"></i>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
