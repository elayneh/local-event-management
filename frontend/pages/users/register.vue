<script setup>
import { ref } from "vue";
import CustomCard from "~/components/CustomCard.vue";
import DynamicForm from "~/components/DynamicForm.vue";
import register from "~/graphql/mutations/users/register.gql";
import { toast } from "vue3-toastify";
import * as yup from "yup";

const emit = defineEmits(["close"]);

const { mutate: userRegistration, onDone, onError } = authentication(register);

const formSchema = {
  fields: [
    {
      name: "username",
      as: "input",
      placeholder: "Username",
      label: "Username",
      rules: yup.string().required("Username is required"),
    },
    {
      name: "email",
      as: "input",
      placeholder: "Email",
      rules: yup
        .string()
        .email("Invalid email address")
        .required("Email is required"),
    },
    {
      name: "password",
      as: "input",
      type: "password",
      placeholder: "Password",
      rules: yup
        .string()
        .min(6, "Password must be at least 6 characters")
        .required("Password is required"),
    },
    {
      placeholder: "Confirm Password",
      name: "cpassword",
      as: "input",
      type: "password",
      rules: yup
        .string()
        .oneOf([yup.ref("password")], "Passwords must match")
        .required("Confirm Password is required"),
    },
  ],
};

const showModal = ref(false);
const showRegisterForm = ref(true);

function submitHandler(values) {
  console.log("User form data: ", values);
  const formData = {
    username: values.username,
    email: values.email,
    password: values.password,
    role: "user",
  };
  userRegistration(formData);
}

onDone((result) => {
  const { id, role, token } = result.data.register;
  toast.success("User registered successfully", {
    autoClose: 2000,
  });

  closeRegistrationForm();
  openModal();
});

onError((error) => {
  console.log("Error: ", error.message);
  toast.error("Something went wrong while registering the user", {
    autoClose: 10000,
  });
});

function closeRegistrationForm() {
  showRegisterForm.value = false;
}

function openModal() {
  showModal.value = true;
}

function closeModal() {
  showModal.value = false;
}

const getStarted = () => {
  console.log("Get Started clicked");
  closeModal();
  window.open("https://mailtrap.io/inboxes/2756361/messages", "_blank");
};
</script>

<style scoped>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-title {
  margin: 0;
}

.close-button {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 18px;
}

.modal-body {
  margin-top: 10px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 10px;
}

.primary-button {
  background-color: #007bff;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.door-wrapper {
  perspective: 800px;
}

.door {
  transform-origin: right center;
  transform: rotateY(90deg);
  transition: transform 1s ease-in-out;
}

.door-animation {
  animation: doorOpen 1.5s forwards ease-in-out;
}

@keyframes doorOpen {
  0% {
    transform: rotateY(90deg);
  }
  100% {
    transform: rotateY(0deg);
  }
}

/* Animated Arrow */
.arrow-right {
  width: 0;
  height: 0;
  border-top: 10px solid transparent;
  border-bottom: 10px solid transparent;
  border-left: 20px solid #fff;
}

@keyframes arrowBounce {
  0% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
  100% {
    transform: translateY(0);
  }
}

.animate-arrow-bounce {
  animation: arrowBounce 1s infinite;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

.animate-fade-in {
  animation: fadeIn 1s forwards;
}

.delay-500 {
  animation-delay: 0.5s;
}

.delay-1000 {
  animation-delay: 1s;
}
</style>

<template>
  <div class="flex min-h-screen relative overflow-hidden">
    <div class="absolute inset-0 z-0">
      <img
        src="~/assets/images/logo.png"
        class="w-full h-full object-cover"
        alt="Background Image"
      />
      <div class="absolute inset-0 bg-gray-800 bg-opacity-80"></div>
    </div>

    <div
      class="relative z-10 flex text-white items-center justify-center lg:mr-8 px-4 lg:px-8"
    >
      <div class="absolute right-0 z-30 animate-arrow-bounce">
        <div class="arrow-right"></div>
      </div>
      <div class="text-center">
        <h1 class="text-4xl font-bold mb-4 opacity-0 animate-fade-in">
          Create Your Account
        </h1>
        <p class="text-lg mb-6 opacity-0 animate-fade-in delay-500">
          Join our platform and explore all the features. Sign up today!
        </p>
        <div class="opacity-0 animate-fade-in delay-1000">
          <button
            @click="navigateTo('/about')"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-lg transition transform hover:scale-105 shadow-lg"
          >
            Learn More
          </button>
        </div>
      </div>
    </div>

    <div class="door-wrapper flex items-center justify-center w-1/2">
      <div
        class="door relative z-20 transform-gpu door-animation p-8 lg:p-16 w-full max-w-lg shadow-lg rounded-lg -translate-x-5"
      >
        <CustomCard title="Sign Up" class="bg-white p-6 rounded-lg shadow-lg">
          <template #body>
            <DynamicForm
              ref="DynamicForm"
              :schema="formSchema"
              :submitHandler="submitHandler"
            >
              <div class="mt-24">
                <div class="flex items-center justify-center gap-4 mt-4">
                  <button
                    type="submit"
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                  >
                    Register
                  </button>
                  <button
                    type="button"
                    @click="navigateTo('/')"
                    class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                  >
                    Cancel
                  </button>
                </div>
                <div class="mt-10 text-gray-700 text-center">
                  <p class="mb-2">
                    Already have an account?
                    <NuxtLink
                      class="text-blue-500 hover:text-blue-700 font-semibold ml-1"
                      to="/users/login"
                    >
                      Login
                    </NuxtLink>
                  </p>
                </div>
              </div>
            </DynamicForm>
          </template>
        </CustomCard>
      </div>
    </div>

    <div class="modal" v-if="showModal">
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">Confirm Email</h3>
          <button class="close-button" @click="closeModal">X</button>
        </div>
        <div class="modal-body">
          <p>An email has been sent to confirm your registration.</p>
        </div>
        <div class="modal-footer">
          <button class="primary-button" @click="getStarted">
            Go to Gmail
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
