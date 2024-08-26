// constants/routes.js
export const unAuthenticatedRoutes = [
  "/",
  "/users",
  "/about",
  "/events",
  "/users/login",
  "/users/register",
  "/users/profile/:id",
  "/users/",
  "/users/unauthorized",
];

export const adminRoutes = [
  "/admin",
  "/events/admin-event",
  "/admin/users",
  "/admin/events",
  "/categories/update/:id",
  "/categories/delete/:id",
  "/categories",
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
  "/events/user-event",
  "/user/events/create-event",
];
