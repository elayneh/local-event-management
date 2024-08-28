import { useAuthStore } from "~/stores";
import {
  adminRoutes,
  unAuthenticatedRoutes,
  userRoutes,
} from "~/constants/routes";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStore();
  await authStore.autoLogin();

  const isUnAuthenticatedRoute = unAuthenticatedRoutes.includes(to.path);
  const isAdminRoute = adminRoutes.includes(to.path);
  const isUserRoute = userRoutes.includes(to.path);
  const userRole = authStore.role?.toLowerCase();

  if (!authStore.isAuthenticated) {
    if (isUnAuthenticatedRoute) {
      return;
    }
    if (isAdminRoute || isUserRoute) {
      return navigateTo("/users/unauthorized");
    }
  } else {
    if (userRole === "admin" && !isAdminRoute) {
      return navigateTo("/admin"); // Redirect to the admin dashboard if not on an admin route
    } else if (userRole === "user" && !isUserRoute) {
      return navigateTo("/user"); // Redirect to the user dashboard if not on a user route
    } else if (!isAdminRoute && !isUserRoute) {
      // Prevent redirect loops by ensuring the current path is not an unauthorized page
      if (to.path !== "/users/unauthorized") {
        return navigateTo("/users/unauthorized");
      }
    }
  }
});
