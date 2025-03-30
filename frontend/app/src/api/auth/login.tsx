import { Context, Hono } from "hono";
import { OidcClient } from "../../oidc/oidcClient";
import { getSignedCookie, setSignedCookie } from "hono/cookie";
import axios from "redaxios";

export const authApi = new Hono();

authApi.use(async (ctx: Context, next) => {
  const oidcClient = new OidcClient({
    clientId: import.meta.env.VITE_CLIENT_ID,
    authEndpoint: import.meta.env.VITE_AUTH_ENDPOINT,
    redirectUrl: import.meta.env.VITE_REDIRECT_URL,
  });
  ctx.set("oidcClient", oidcClient);
  await next();
});

authApi.get("/login", async (ctx: Context) => {
  const state = crypto.randomUUID();
  await setSignedCookie(
    ctx,
    "state",
    state,
    import.meta.env.VITE_COOKIE_SECRET
  );

  const oidcClient: OidcClient = ctx.var.oidcClient;

  const authUrl = oidcClient.getAuthUrl({
    respType: "code",
    scopes: ["openid", "email", "profile"],
    state,
  });

  return ctx.redirect(authUrl, 302);
});

authApi.get("/google/callback", async (ctx: Context) => {
  const state = await getSignedCookie(
    ctx,
    import.meta.env.VITE_COOKIE_SECRET,
    "state"
  );
  const queryState = ctx.req.query("state");

  if (state !== queryState) {
    return ctx.render(<h1>Try Again</h1>);
  }

  const code = ctx.req.query("code");
  if (!code) {
    return ctx.render(<h1>Try Again</h1>);
  }

  try {
    const token = await axios.get(
      `${import.meta.env.VITE_SERVER_ENDPOINT}/v1/token_request?code=${code}`
    );

    const response = await axios.post(
      `${import.meta.env.VITE_SERVER_ENDPOINT}/v1/auth/login`,
      {},
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token.data.id_token}`,
        },
      }
    );

    return ctx.render(<h1>Login</h1>);
  } catch (err) {
    console.log(err);
  }
});
