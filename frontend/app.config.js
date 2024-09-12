// app.config.js
export default {
  router: {
    scrollBehavior(to, from, savedPosition) {
      // Always scroll to the top of the page on route change
      return { top: 0 };
    },
  },
};
