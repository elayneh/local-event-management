<template>
  <div v-if="ticket" class="bg-white shadow-md rounded-lg overflow-hidden hover:shadow-lg transition-shadow duration-200">
    <div class="p-4">
      <h3 class="text-lg font-bold text-gray-800 mb-2">
        Event ID: {{ ticket.event_id }}
      </h3>
      <p class="text-gray-600 mb-1">
        <span class="font-semibold">Price:</span> ${{ ticket.price }}
      </p>
      <p class="text-gray-600 mb-1">
        <span class="font-semibold">Quantity:</span> {{ ticket.quantity }}
      </p>
      <p class="text-gray-600 mb-1">
        <span class="font-semibold">Purchase Date:</span> {{ formatDate(ticket.purchase_date) }}
      </p>
      <p :class="ticket.status === 'confirmed' ? 'bg-green-200 text-green-800' : 'bg-red-200 text-red-800'">
        Status: {{ ticket.status === "confirmed" ? "Confirmed" : "Pending" }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { format } from "date-fns";
import { defineProps } from "vue";

const props = defineProps({
  ticket: {
    type: Object,
    required: true,
  },
});

// Log the ticket to the console
console.log("Ticket data:", props.ticket);

// Utility function to format the purchase date
const formatDate = (dateString) => format(new Date(dateString), "MMMM dd, yyyy");
</script>
