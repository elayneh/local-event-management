import { defineStore } from "pinia";
import Cookies from "js-cookie";
import jwt_decode from "jwt-decode";
import get_single_user from "@/graphql/queries/user/getSingleUser.gql";
import authQuery from "@/composables/auth";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: Cookies.get("token") || null,
    id: null,
    role: null,
    user: null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },

  actions: {
    setToken(token) {
      this.token = token;
    },

    setRole(role) {
      this.role = role;
    },

    setId(id) {
      this.id = id;
    },

    logout() {
      this.token = null;
      this.user = null;
      this.id = null;
      Cookies.remove("token");
    },

    async setUser(id) {
      const { onResult, onError } = authQuery(get_single_user, {
        id,
      });

      onResult((result) => {
        this.user = { ...result.data.users_by_pk };
      });

      onError((error) => {
        console.log(error);
      });
    },

    autoLogin() {
      if (this.token) {
        const decoded = jwt_decode(this.token);
        if (decoded.exp * 1000 > Date.now()) {
          const id =
            decoded["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
          this.setId(id);
          this.setUser(id);
        } else {
          this.logout();
        }
      }
    },
  },
});
