export default {
  devtools: { enabled: false },
  ssr: false,

  runtimeConfig: {
    public: {
      hasuraAdminSecret: process.env.HASURA_GRAPHQL_ADMIN_SECRET || "",
      CALLBACK_URL: process.env.NUXT_ENV_CALLBACK_URL,
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
      title: "local event management at minab",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: "description",
          content:
            "A comprehensive platform designed to streamline the planning and management of events. It offers features for organizing, scheduling, and promoting events, managing ticket sales, and tracking attendee engagement. Users can easily create and customize events, handle registrations, and manage event-related communications. The system provides a user-friendly interface for both event organizers and attendees, enhancing the overall experience with real-time updates and interactive tools.",
        },
      ],
      link: [
        { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
        {
          rel: "apple-touch-icon",
          sizes: "180x180",
          href: "/apple-touch-icon.png",
        },
        {
          rel: "icon",
          type: "image/png",
          sizes: "32x32",
          href: "/favicon-32x32.png",
        },
        {
          rel: "icon",
          type: "image/png",
          sizes: "16x16",
          href: "/favicon-16x16.png",
        },
      ],
      link: [
        {
          rel: "stylesheet",
          href: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css",
        },
      ],
    },
  },

  plugins: ["~/plugins/fontawesome.js", "~/plugins/click-outside.js"],
  router: {
    middleware: ["scroll-to-top"],
  },
  image: {
    inject: true,
    quality: 80,
  },

  compatibilityDate: "2024-08-01",
};
