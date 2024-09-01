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
  "/admin/tags",
  "/admin/tags/add",
  "/admin/settings",
  /^\/admin\/tags\/delete\/[a-zA-Z0-9-]+$/,
  /^\/admin\/tags\/update\/[a-zA-Z0-9-]+$/,
  /^\/admin\/categories\/update\/[a-zA-Z0-9-]+$/,
];

export const userRoutes = [
  "/user",
  "/user/events",
  "/user/events/bookmarked",
  "/user/events/following",
  "/user/events/tickets",
  "/user/events/create-event",
  "/user/profile",
  /^\/user\/events\/[a-zA-Z0-9-]+$/,
  /^\/events\/[a-zA-Z0-9-]+$/,
  "/events/update",
];
