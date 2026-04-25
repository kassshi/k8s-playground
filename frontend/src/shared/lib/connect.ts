import { Code, ConnectError, type Interceptor } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const jwtInterceptor: Interceptor = (next) => async (req) => {
  const token = localStorage.getItem("token");
  if (token) {
    req.header.set("Authorization", `Bearer ${token}`);
  }
  return await next(req);
};

const authErrorInterceptor: Interceptor = (next) => async (req) => {
  try {
    const res = await next(req);
    return res;
  } catch (error) {
    if (error instanceof ConnectError && error.code == Code.Unauthenticated) {
      window.location.href = "/login";
    }
  }
};

export const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  interceptors: [jwtInterceptor, authErrorInterceptor],
});
