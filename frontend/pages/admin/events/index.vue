<template>
  <div class="p-6 bg-white shadow-md rounded-lg">
    <h2 class="text-2xl font-bold mb-4 text-gray-800">
      Admin Event Management
    </h2>
    <table class="min-w-full bg-gray-100 border">
      <thead>
        <tr>
          <th class="py-2 px-4 border">#</th>
          <th class="py-2 px-4 border">Title</th>
          <th class="py-2 px-4 border">Description</th>
          <th class="py-2 px-4 border">Created At</th>
          <th class="py-2 px-4 border">Created By</th>
          <th class="py-2 px-4 border text-center">Freeze</th>
          <th class="py-2 px-4 border text-center">Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(event, index) in events" :key="event.id" class="border-t">
          <td class="py-2 px-4 border">{{ index + 1 }}</td>
          <td class="py-2 px-4 border">{{ event.title }}</td>
          <td class="py-2 px-4 border">{{ event.description }}</td>
          <td class="py-2 px-4 border">{{ formatDate(event.created_at) }}</td>
          <td class="py-2 px-4 border">{{ event.user_id }}</td>
          <td class="py-2 px-4 border text-center">
            <button
              @click="freezeEvent(event.id)"
              class="text-blue-500 hover:text-blue-700 transition-colors duration-200"
            >
              <i class="fa fa-snowflake" aria-hidden="true"></i>
            </button>
          </td>
          <td class="py-2 px-4 border text-center">
            <button
              @click="deleteEvent(event.id)"
              class="text-red-500 hover:text-red-700 transition-colors duration-200"
            >
              <i class="fa fa-trash" aria-hidden="true"></i>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import deleteEventMutation from "~/graphql/mutations/events/delete.gql";
import useAuthenticatedMutation from "~/composables/useAuthenticatedMutation";
import { toast } from "vue3-toastify"; // Importing the toast function

// State to store events
const events = ref([]);

// Utility function to format date
function formatDate(dateString) {
  const options = { year: "numeric", month: "short", day: "numeric" };
  return new Date(dateString).toLocaleDateString(undefined, options);
}

// Query to fetch events
const {
  result: eventResult,
  loading: eventLoading,
  error: eventError,
  onResult: onEventsResult,
} = useQuery(getEvents);

// Mutation to delete event
const { mutate: deleteEvent, loading: deleteLoading, error: deleteError } = useAuthenticatedMutation(deleteEventMutation);

// Handling the result from the query
onEventsResult(({ data }) => {
  if (data) {
    events.value = data.events;
  }
});

// Function to handle event deletion
const handleDeleteEvent = (event_id) => {
  deleteEvent(
    { id: event_id },
    {
      onDone: () => {
        events.value = events.value.filter((event) => event.id !== event_id);
        toast.success("Event deleted successfully");
      },
      onError: (error) => {
        console.error("Error deleting event:", error);
        toast.error("Failed to delete event");
      },
    }
  );
};

// Example function for freezing an event
const freezeEvent = (event_id) => {
  console.log(`Event with ID ${event_id} freezed successfully`);
};

// Define the page meta layout
definePageMeta({ layout: "admin-dashboard" });
</script>
