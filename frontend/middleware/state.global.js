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
  const isAdminRoute = adminRoutes.some((route) =>
    route instanceof RegExp ? route.test(to.path) : route === to.path
  );
  const isUserRoute = userRoutes.some((route) =>
    route instanceof RegExp ? route.test(to.path) : route === to.path
  );
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
      return navigateTo("/admin");
    } else if (userRole === "user" && !isUserRoute) {
      return navigateTo("/user");
    } else if (!isAdminRoute && !isUserRoute) {
      if (to.path !== "/users/unauthorized") {
        return navigateTo("/users/unauthorized");
      }
    }
  }
});
