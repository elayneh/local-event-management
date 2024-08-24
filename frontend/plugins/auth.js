import { useAuthStore } from "~/stores";

export default defineNuxtPlugin((nuxtApp) => {
  const authStore = useAuthStore();
  authStore.autoLogin();
});
