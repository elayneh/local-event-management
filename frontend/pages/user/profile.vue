<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 w-full min-h-screen flex justify-center items-start"
  >
    <div
      class="bg-white shadow-xl rounded-lg p-8 w-full max-w-4xl mt-24 mx-8 transform transition-transform duration-300 ease-in-out hover:scale-105"
    >
      <div class="flex flex-col justify-center items-center">
        <div
          class="w-16 h-16 bg-gray-700 rounded-full flex items-center justify-center focus:outline-none focus:ring-2 focus:ring-gray-500"
        >
          <font-awesome-icon
            icon="user-circle"
            class="text-gray-500 w-full h-full"
          />
        </div>
        <div class="text-center mt-4">
          <h1 class="text-2xl font-bold">Profile Details</h1>
        </div>
        <div class="mt-12 w-full">
          <div class="pl-8">
            <label class="text-gray-800 font-semibold">Username</label>
            <input
              v-if="isEditing"
              v-model="editableUser.username"
              class="border border-gray-300 rounded p-2 w-full mt-2"
              type="text"
            />
            <p v-else class="text-gray-600">{{ user.username }}</p>
          </div>
          <div class="pl-8 mt-6">
            <label class="text-gray-800 font-semibold">Email</label>
            <input
              v-if="isEditing"
              v-model="editableUser.email"
              class="border border-gray-300 rounded p-2 w-full mt-2"
              type="email"
            />
            <p v-else class="text-gray-600">{{ user.email }}</p>
          </div>
          <div class="pl-8 mt-6">
            <label class="text-gray-800 font-semibold">Role</label>
            <input
              v-if="isEditing"
              v-model="editableUser.role"
              class="border border-gray-300 rounded p-2 w-full mt-2"
              type="text"
              disabled
            />
            <p v-else class="text-gray-600">{{ user.role }}</p>
          </div>
        </div>
        <div v-if="!isEditing && user.role === 'user'" class="mt-24">
          <div
            class="pl-8 mt-6 bg-gray-50 p-6 rounded-lg shadow-md border border-gray-200"
          >
            <p class="text-gray-800 font-semibold text-lg mb-4">Permissions</p>
            <p class="text-gray-600 mb-2">
              You are permitted to post events as an event creator or you can
              browse, reserve, buy, follow, and bookmark any event you prefer.
            </p>
            <p class="text-gray-600">You can edit or update your profile.</p>
          </div>
        </div>
      </div>
      <div class="mt-12">
        <div v-if="isEditing" class="flex flex-row justify-center items-center">
          <button
            @click="saveProfile"
            class="bg-green-600 text-white px-6 py-2 rounded-lg shadow hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 mr-64"
          >
            Save
          </button>
          <button
            @click="cancelEdit"
            class="bg-red-600 text-white px-6 py-2 rounded-lg shadow hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
          >
            Cancel
          </button>
        </div>
        <div v-else class="flex justify-end">
          <button
            @click="editProfile"
            class="bg-blue-600 text-white px-6 py-2 rounded-lg shadow hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            Edit Profile
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import useUserFetchData from "~/composables/useUserFetchData";
import { useAuthStore } from "~/stores";
import { UpdateUser } from "~/graphql/mutations/users/update.gql";
import { useMutation } from "@vue/apollo-composable";
import { toast } from "vue3-toastify";

const authStore = useAuthStore();
const user_id = authStore.id;

const { user } = useUserFetchData(user_id);

const isEditing = ref(false);
const editableUser = ref({ ...user.value });

const {
  mutate: updateUserMutation,
  loading,
  error,
  onResult,
  onDone,
} = useMutation(UpdateUser);

// Methods
const editProfile = () => {
  isEditing.value = true;
  editableUser.value = { ...user.value };
};

const saveProfile = async () => {
  try {
    const { data } = await updateUserMutation({
      id: user.value.id,
      email: editableUser.value.email,
      username: editableUser.value.username,
    });

    console.log("Updated user: ", data);
    user.value = { ...editableUser.value };

    toast.success("Your Profile Updated Successfully!");
    isEditing.value = false;
  } catch (err) {
    toast.error("Error Updatting Your Profile, try again", {
      transition: toast.TRANSITIONS.ZOOM,
    });
    console.error("Error updating user: ", err);
  }
};

const cancelEdit = () => {
  isEditing.value = false;
};

definePageMeta({
  layout: "authenticated",
});
</script>
