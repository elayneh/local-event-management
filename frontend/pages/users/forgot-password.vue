<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
      <h2 class="text-2xl font-bold mb-6 text-center">Reset Password</h2>
      <form @submit.prevent="requestPasswordReset" class="space-y-4">
        <input
          v-model="email"
          type="email"
          placeholder="Enter your email"
          class="w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          required
        />
        <button
          type="submit"
          class="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          Reset Password
        </button>
      </form>
    </div>
  </div>
</template>


<script setup>
import { ref } from 'vue';
import { useMutation } from '@vue/apollo-composable';
import RequestPasswordReset from '~/graphql/mutations/users/requestPasswordReset.gql';

const email = ref('');

const { mutate: requestPasswordResetMutation, onDone, onError } = useMutation(RequestPasswordReset);

const requestPasswordReset = async () => {
  try {
    const response = await requestPasswordResetMutation({ email: email.value });
  } catch (error) {
    console.error('Password reset request failed:', error);
  }
};

onDone((response) => {
  console.log('Mutation completed:', response.data);
});

// Optional: handle mutation error
onError((error) => {
  console.error('Mutation error:', error);
}); 
</script>
