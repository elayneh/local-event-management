// constants/routes.js
export const unAuthenticatedRoutes = [
  "/",
  "/users",
  "/about",
  "/events",
  "/users/login",
  "/users/register",
  "/users/profile/:id",
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
  "/events",
  "/user/events/:id",
  "/user/events",
  "/user/events/bookmarked",
  "/user/events/following",
  "/user/events/tickets",
  "/events/:id",
  "/events/update",
  "/user/events/create-event",
  "/user",
];
