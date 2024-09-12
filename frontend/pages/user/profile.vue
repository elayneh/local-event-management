<script setup>
import { ref } from "vue";
import useUserFetchData from "~/composables/useUserFetchData";
import { useAuthStore } from "~/stores";
import { UpdateUser } from "~/graphql/mutations/users/update.gql";
import { ChangePassword } from "~/graphql/mutations/users/changePassword.gql";
import { useMutation } from "@vue/apollo-composable";
import { toast } from "vue3-toastify";
import { Form, Field } from "vee-validate";
import { object, string } from "yup";

const authStore = useAuthStore();
const user_id = authStore.id;
const { user } = useUserFetchData(user_id);

const isChangePassword = ref(false);
const isEditing = ref(false);
const editableUser = ref({ ...user.value });

const schema = object({
  username: string()
    .required("Username is required")
    .min(3, "Username must be at least 3 characters"),
  email: string()
    .required("Email is required")
    .email("Email must be a valid email address"),
});

const passwordSchema = object({
  newPassword: string().required("Password is required"),
});

const { mutate: updateUserMutation } = useMutation(UpdateUser);
const { mutate: changePasswordMutation } = useMutation(changePassword);

const editProfile = () => {
  isEditing.value = true;
  editableUser.value = { ...user.value };
};
const passwords = ref({ currentPassword: "", newPassword: "" });

const saveProfile = async (values) => {
  try {
    const { data } = await updateUserMutation({
      id: user.value.id,
      email: values.email,
      username: values.username,
    });

    user.value = { ...editableUser.value };

    toast.success("Your Profile Updated Successfully!");
    isEditing.value = false;
  } catch (err) {
    toast.error("Error Updating Your Profile, try again");
    console.error("Error updating user: ", err);
  }
};

const changePassword = async (values) => {
  try {
    const { data } = await changePasswordMutation({
      newPassword: values.newPassword,
      currentPassword: values.currentPassword,
    });

    user.value = { ...editableUser.value };

    toast.success("Your password has been changed successfully!");
    isChangePassword.value = false;
  } catch (err) {
    toast.error("Error changing password, try again");
    console.error("Error changing password: ", err);
  }
};

const cancelEdit = () => {
  isEditing.value = false;
};

const openChangePassword = () => {
  isChangePassword = !isChangePassword;
};

definePageMeta({
  layout: "authenticated",
});
</script>

<template>
  <div
    class="bg-gradient-to-r from-gray-100 via-red-300 to-gray-500 w-full min-h-screen flex justify-center items-start"
  >
    <div
      class="bg-white shadow-xl rounded-lg p-8 w-full max-w-4xl mt-24 mx-8 transform transition-transform duration-300 ease-in-out hover:scale-105"
    >
      <div class="flex justify-end">
        <button
          @click="navigateTo('/user')"
          class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 pr-4 pt-2"
        >
          <i class="fas fa-times text-2xl"></i>
        </button>
      </div>
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
          <Form
            v-if="isEditing"
            :validation-schema="schema"
            @submit="saveProfile"
            v-slot="{ errors }"
          >
            <div class="pl-8">
              <label class="text-gray-800 font-semibold">Username</label>
              <Field
                name="username"
                v-model="editableUser.username"
                class="border border-gray-300 rounded p-2 w-full mt-2"
                type="text"
              />
              <span class="text-red-500">{{ errors.username }}</span>
            </div>
            <div class="pl-8 mt-6">
              <label class="text-gray-800 font-semibold">Email</label>
              <Field
                name="email"
                v-model="editableUser.email"
                class="border border-gray-300 rounded p-2 w-full mt-2"
                type="email"
              />
              <span class="text-red-500">{{ errors.email }}</span>
            </div>
            <div class="pl-8 mt-6">
              <label class="text-gray-800 font-semibold">Role</label>
              <input
                v-model="editableUser.role"
                class="border border-gray-300 rounded p-2 w-full mt-2"
                type="text"
                disabled
              />
            </div>
            <div v-if="isChangePassword">
              <Form
                :validation-schema="passwordSchema"
                @submit="changePassword"
                v-slot="{ errors }"
              >
                <div class="pl-8 mt-6">
                  <label class="text-gray-800 font-semibold"
                    >Current Password</label
                  >
                  <Field
                    name="currentPassword"
                    v-model="passwords.currentPassword"
                    class="border border-gray-300 rounded p-2 w-full mt-2"
                    type="password"
                  />

                  <span class="text-red-500">{{ errors.currentPassword }}</span>
                </div>
                <div class="pl-8 mt-6">
                  <label class="text-gray-800 font-semibold"
                    >New Password</label
                  >
                  <Field
                    name="newPassword"
                    v-model="passwords.newPassword"
                    class="border border-gray-300 rounded p-2 w-full mt-2"
                    type="password"
                  />
                  <span class="text-red-500">{{ errors.newPassword }}</span>
                </div>
              </Form>
            </div>
            <div v-else class="mt-24 pl-8">
              <button @click="openChangePassword">Change Password</button>
            </div>
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
          <div v-else>
            <div class="pl-8">
              <label class="text-gray-800 font-semibold">Username</label>
              <p class="text-gray-600">{{ user.username }}</p>
            </div>
            <div class="pl-8 mt-6">
              <label class="text-gray-800 font-semibold">Email</label>
              <p class="text-gray-600">{{ user.email }}</p>
            </div>
            <div class="pl-8 mt-6">
              <label class="text-gray-800 font-semibold">Role</label>
              <p class="text-gray-600">{{ user.role }}</p>
            </div>
          </div>
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
      <div class="mt-12 flex justify-end">
        <button
          v-if="!isEditing"
          @click="editProfile"
          class="bg-blue-600 text-white px-6 py-2 rounded-lg shadow hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          Edit Profile
        </button>
      </div>
    </div>
  </div>
</template>
