// // plugins/apollo.ts

// import {
//   ApolloClient,
//   InMemoryCache,
//   createHttpLink,
// } from "@apollo/client/core";
// import { defineNuxtPlugin } from "#app";

// export default defineNuxtPlugin((nuxtApp) => {
//   // Access runtime configuration
//   const { hasuraAdminSecret } = useRuntimeConfig().private;

//   const httpLink = createHttpLink({
//     uri: "http://localhost:8080/v1/graphql", // Replace with your actual Hasura endpoint
//     headers: {
//       "x-hasura-admin-secret": hasuraAdminSecret || "",
//     },
//   });

//   const apolloClient = new ApolloClient({
//     link: httpLink,
//     cache: new InMemoryCache(),
//   });

//   nuxtApp.provide("apollo", apolloClient);
// });
