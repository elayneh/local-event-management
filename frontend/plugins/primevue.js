import PrimeVue from "primevue/config";

export default defineNuxtPlugin((nuxtApp) => {
  const app = nuxtApp.vueApp;
  app.use(PrimeVue);
});
