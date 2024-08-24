<template>
  <div class="p-6 bg-white shadow-md rounded-lg">
    <h2 class="text-2xl font-bold mb-4">User Management</h2>
    <table class="min-w-full bg-gray-100 border">
      <thead>
        <tr>
          <th class="py-2 px-4 border">#</th>
          <th class="py-2 px-4 border">ID</th>
          <th class="py-2 px-4 border">Username</th>
          <th class="py-2 px-4 border">Email</th>
          <th class="py-2 px-4 border">Role</th>
          <th class="py-2 px-4 border">Created At</th>
          <th class="py-2 px-4 border">Updated At</th>
          <th class="py-2 px-4 border">Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(user, index) in users" :key="user.id" class="border-t">
          <td class="py-2 px-4 border">{{ index + 1 }}</td>
          <td class="py-2 px-4 border">{{ user.id }}</td>
          <td class="py-2 px-4 border">{{ user.username }}</td>
          <td class="py-2 px-4 border">{{ user.email }}</td>
          <td class="py-2 px-4 border">{{ user.role }}</td>
          <td class="py-2 px-4 border">{{ formatDate(user.updated_at) }}</td>
          <td class="py-2 px-4 border">{{ formatDate(user.created_at) }}</td>
          <td class="py-2 px-4 border text-center">
            <button
              @click="deleteUser(user.id)"
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
import getUsers from "~/graphql/queries/users/getUsers.gql";

const users = ref([]);
const { result, loading, error } = useQuery(getUsers);

definePageMeta({ layout: "admin-dashboard" });

function formatDate(dateString) {
  const options = { year: "numeric", month: "short", day: "numeric" };
  return new Date(dateString).toLocaleDateString(undefined, options);
}

watch(result, (newResult) => {
  if (newResult?.users) {
    users.value = newResult.users;
  }
});

watch(loading, (newLoading) => {
  console.log("Loading:", newLoading);
});

watch(error, (newError) => {
  console.error("Error:", newError);
});
</script>
