import { createClient } from "@connectrpc/connect";
import { AuthService } from "../../../gen/auth/v1/auth_pb";
import { transport } from "../../../shared/lib/connect";
const client = createClient(AuthService, transport);

export async function login(email: string, password: string) {
  return await client.login({
    email,
    password,
  });
}
