import { useAuthStore } from "@/stores";
export default defineNuxtRouteMiddleware((to, from, next) => {
  const authStore = useAuthStore();
  console.log("setup.global.js this is the global middleware");
  if (!authStore.isAuthenticated) {
    authStore.autoLogin();
  }
  return;
});
