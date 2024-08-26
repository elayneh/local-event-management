// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  ssr: false,

  runtimeConfig: {
    public: {
      hasuraAdminSecret: process.env.HASURA_GRAPHQL_ADMIN_SECRET || "",
    },
    private: {
      hasuraAdminSecret: process.env.HASURA_GRAPHQL_ADMIN_SECRET || "",
    },
  },
  modules: [
    "@nuxtjs/tailwindcss",
    "@nuxtjs/apollo",
    "@pinia/nuxt",
    "nuxt-rating",
    "nuxt-icon",
    "@vee-validate/nuxt",
    "@vueuse/nuxt",
    "@nuxtjs/color-mode",
    "@nuxt/image",
  ],
  // pinia: { storesDirs: ["./stores/**"] },
  colorMode: {
    classSuffix: "",
  },

  apollo: {
    autoImports: true,
    authType: "Bearer",
    authHeader: "Authorization",
    tokenStorage: "cookie",
    proxyCookies: true,
    clients: {
      default: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
      },
      authClient: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
        tokenName: "token",
      },
    },
  },

  veeValidate: {
    // disable or enable auto imports
    autoImports: true,
    // Use different names for components
    componentNames: {
      Form: "VeeForm",
      Field: "VeeField",
      FieldArray: "VeeFieldArray",
      ErrorMessage: "VeeErrorMessage",
    },
  },

  app: {
    head: {
      link: [
        {
          rel: "stylesheet",
          href: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css",
        },
      ],
    },
  },
  plugins: ["~/plugins/fontawesome.js", "~/plugins/click-outside.js"],

  image: {
    // Options
    inject: true,
    quality: 80,
  },
  compatibilityDate: "2024-08-01",
});
