// constants/routes.js
export const unAuthenticatedRoutes = [
  "/",
  "/users",
  "/about",
  "/events",
  "/users/login",
  "/users/register",
  "/users/unauthorized",
];

export const adminRoutes = [
  "/admin",
  "/events/admin-event",
  "/admin/users",
  "/admin/events",
  "/admin/categories/update/:id",
  "/admin/categories/delete/:id",
  "/admin/categories",
];

export const userRoutes = [
  /^\/user\/events\/[a-zA-Z0-9-]+$/,
  "/user/events",
  "/user/events/bookmarked",
  "/user/events/following",
  "/user/events/tickets",
  /^\/events\/[a-zA-Z0-9-]+$/,
  "/events/update",
  "/user/events/create-event",
  "/user",
];
