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
  faTachometerAlt,
  faUser,
  faCalendarAlt,
  faTag,
  faCog,
  faSignOutAlt,
} from "@fortawesome/free-solid-svg-icons";

export default defineNuxtPlugin((nuxtApp) => {
  library.add(
    faSearch,
    faFilter,
    faCaretDown,
    faBookmark,
    faUserCircle,
    faTimes,
    faTachometerAlt,
    faUser,
    faCalendarAlt,
    faTag,
    faCog,
    faSignOutAlt
  );

  nuxtApp.vueApp.component("font-awesome-icon", FontAwesomeIcon);
});
