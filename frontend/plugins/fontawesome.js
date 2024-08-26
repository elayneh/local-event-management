// plugins/fontawesome.js

import { defineNuxtPlugin } from "#app";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faCaretDown,
  faFilter,
  faSearch,
  faBookmark,
} from "@fortawesome/free-solid-svg-icons";

export default defineNuxtPlugin((nuxtApp) => {
  library.add(faSearch, faFilter, faCaretDown, faBookmark);

  nuxtApp.vueApp.component("font-awesome-icon", FontAwesomeIcon);
});
