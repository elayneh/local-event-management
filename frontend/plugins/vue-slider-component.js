// import VueApexCharts from "vue3-apexcharts";
import VueSlider from "vue-slider-component/dist-css/vue-slider-component.umd.min.js";
import "vue-slider-component/dist-css/vue-slider-component.css";
import "vue-slider-component/theme/default.css";
// import 'vue-slider-component/theme/material.css'
export default defineNuxtPlugin((nuxtApp) => {
	nuxtApp.vueApp.component("VueSlider", VueSlider);
});
