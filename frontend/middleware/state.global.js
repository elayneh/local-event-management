import { useAuthStore } from "~/stores";
import {
  adminRoutes,
  unAuthenticatedRoutes,
  userRoutes,
} from "~/constants/routes";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStore();
  await authStore.autoLogin();

  const isUnAuthenticatedRoute = unAuthenticatedRoutes.some((route) =>
    to.path.startsWith(route)
  );

  if (!authStore.isAuthenticated) {
    if (isUnAuthenticatedRoute) {
      return;
    } else {
      return navigateTo("/users/unauthorized");
    }
  }

  const userRole = authStore.role?.toLowerCase();

  if (userRole === "admin") {
    const isAdminRoute = adminRoutes.some((route) => to.path.startsWith(route));
    if (!isAdminRoute && to.path !== "/admin") {
      return navigateTo("/admin"); // Redirect to the admin dashboard if not on an admin route
    }
  }

  if (userRole === "user") {
    const isUserRoute = userRoutes.some((route) => to.path.startsWith(route));
    if (!isUserRoute && to.path !== "/user") {
      return navigateTo("/user"); // Redirect to the user dashboard if not on a user route
    }
  }
});
