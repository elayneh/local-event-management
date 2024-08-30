import { defineNuxtPlugin } from "#app";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faCaretDown,
  faFilter,
  faSearch,
  faBookmark,
  faUserCircle,
  faTimes,
} from "@fortawesome/free-solid-svg-icons";

export default defineNuxtPlugin((nuxtApp) => {
  library.add(
    faSearch,
    faFilter,
    faCaretDown,
    faBookmark,
    faUserCircle,
    faTimes
  );

  nuxtApp.vueApp.component("font-awesome-icon", FontAwesomeIcon);
});
