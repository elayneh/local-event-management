import vue from "eslint-plugin-vue";

export default [
  {
    parser: "vue-eslint-parser",
    parserOptions: {
      ecmaVersion: 2020,
      sourceType: "module",
    },
    extends: [
      "eslint:recommended",
      "plugin:vue/vue3-recommended",
      "@vue/eslint-config-airbnb",
    ],
    plugins: ["vue"],
    rules: {
      "no-unused-vars": "warn",
      "no-undef": "warn",
    },
    ignores: ["node_modules", "dist"],
  },
];
