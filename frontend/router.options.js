// app/router.options.js
export default function (to, from, savedPosition) {
  // if savedPosition is available, it's a popstate navigation, so scroll to the saved position
  if (savedPosition) {
    return savedPosition;
  } else {
    // otherwise, scroll to the top of the page
    return { left: 0, top: 0 };
  }
}
