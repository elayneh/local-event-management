// nuxt.config.js
export default {
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
    autoImports: true,
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
    inject: true,
    quality: 80,
  },

  compatibilityDate: "2024-08-01",

  router: {
    options: {
      scrollBehavior(to, from, savedPosition) {
        // if (savedPosition) {
        return savedPosition;
        // } else {
        //   return { left: 0, top: 0 };
        // }
      },
    },
  },
};
