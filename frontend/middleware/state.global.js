import { useAuthStore } from "~/stores";
import {
  adminRoutes,
  unAuthenticatedRoutes,
  userRoutes,
} from "~/constants/routes";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStore();
  authStore.autoLogin();
  if (!authStore.isAuthenticated && !unAuthenticatedRoutes.includes(to.path)) {
    return navigateTo("/users/login");
  }

  const userRole = authStore.role?.toLowerCase();
  console.log("USER: ", userRole);
  if (userRole) {
    if (!adminRoutes.includes(to.path) && userRole == "admin") {
      return navigateTo("/admin");
    }
    if (!userRoutes.includes(to.path) && userRole == "user") {
      return navigateTo("/user");
    }
  }
});
