import SuperTokens from "supertokens-web-js";
import Session from "supertokens-web-js/recipe/session";

export function initSuperTokensUI() {
  (window as any).supertokensUIInit("supertokensui", {
    appInfo: {
      websiteDomain: "http://127.0.0.1:3000",
      apiDomain: "http://127.0.0.1:3001",
      appName: "XTuples Web Front ",
    },
    recipeList: [
      (window as any).supertokensUIEmailPassword.init(),
      (window as any).supertokensUIEmailVerification.init({ mode: "REQUIRED" }),
      (window as any).supertokensUISession.init(),
    ],
  });
}

export function initSuperTokensWebJS() {
  SuperTokens.init({
    appInfo: {
      appName: "XTuples Web Front",
      apiDomain: "http://127.0.0.1:3001",
    },
    recipeList: [Session.init()],
  });
}
