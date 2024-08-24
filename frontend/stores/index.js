import { defineStore } from "pinia";
import Cookies from "js-cookie";
import jwt_decode from "jwt-decode";
import getUser from "@/graphql/queries/users/getSingleUser.gql";
import authQuery from "@/composables/getUser";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: Cookies.get("token") || null,
    id: null,
    role: null,
    user: null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    getUserRole: (state) => state.role,
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
      const { onResult, onError } = await authQuery(id, getUser);

      onResult((result) => {
        if (result.data && result.data.users_by_pk) {
          this.user = { ...result.data.users_by_pk };
        }
      });

      onError((error) => {
        console.error("GraphQL Error: ", error);
      });
    },

    autoLogin() {
      if (this.token) {
        const decoded = jwt_decode(this.token);
        if (decoded.exp * 1000 > Date.now()) {
          const id =
            decoded["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
          const role = decoded["https://hasura.io/jwt/claims"]["x-hasura-role"];
          this.setId(id);
          this.setUser(id);
          this.setRole(role);
        } else {
          this.logout();
        }
      }
    },
  },
});
