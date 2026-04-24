import { createClient } from "@connectrpc/connect";
import { AuthService } from "../../../gen/auth/v1/auth_pb";
import { transport } from "../../../shared/lib/connect";
const client = createClient(AuthService, transport);

export async function signup(
  email: string,
  password: string,
  confirmPassword: string,
) {
  return await client.signup({
    email,
    password,
    confirmPassword,
  });
}
