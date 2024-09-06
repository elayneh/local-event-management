import { defineNuxtPlugin } from "#app";
import vClickOutside from "vue-click-outside";

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(vClickOutside);
});
