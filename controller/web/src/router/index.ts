import { createRouter, createWebHashHistory, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import AuthView from "../views/AuthView.vue";

import Session from "supertokens-web-js/recipe/session";
import { EmailVerificationClaim } from "supertokens-web-js/recipe/emailverification";

async function shouldLoadRoute(): Promise<boolean> {
  if (await Session.doesSessionExist()) {
    const validationErrors = await Session.validateClaims();
    if (validationErrors.length === 0) {
      // user has verified their email address
      return true;
    } else {
      for (const err of validationErrors) {
        if (err.id === EmailVerificationClaim.id) {
          // email is not verified
        }
      }
    }
  }
  // a session does not exist, or email is not verified
  return false;
}

const router = createRouter({
  //history:  createWebHashHistory(), //createWebHistory(import.meta.env.BASE_URL),
  history:  createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/auth/:pathMatch(.*)*",
      name: "auth",
      component: AuthView,
    },
  ],
});

export default router;
