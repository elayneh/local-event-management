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
</style>

<template>
  <div>
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

    <div
      v-if="showRegisterForm"
      class="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50 transition-opacity duration-500 ease-in-out"
    >
      <div class="animate__animated animate__zoomIn">
        <CustomCard
          title="Sign Up"
          customClass="bg-white p-6 rounded-lg shadow-lg w-full max-w-md transform transition-transform duration-300 hover:scale-105 hover:shadow-2xl"
        >
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
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transform transition-transform duration-300 hover:scale-110"
                  >
                    Register
                  </button>
                  <button
                    type="button"
                    @click="navigateTo('/')"
                    class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transform transition-transform duration-300 hover:scale-110"
                  >
                    Cancel
                  </button>
                </div>
                <div>
                  <p class="mt-5 text-gray-700 text-center">
                    Already have an account?
                    <NuxtLink
                      class="text-blue-500 hover:text-blue-700 font-semibold ml-1 transition-colors duration-300"
                      to="/users/login"
                    >
                      <span>Login</span>
                    </NuxtLink>
                  </p>
                </div>
              </div>
            </DynamicForm>
          </template>
        </CustomCard>
      </div>
    </div>
  </div>
</template>
