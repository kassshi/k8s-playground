import { Button } from "@/components/ui/button";
import { ConnectError } from "@connectrpc/connect";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router";
import { login } from "../../features/auth/api/login";

type LoginFormValues = {
  email: string;
  password: string;
};

export function LoginPage() {
  const { register, handleSubmit } = useForm<LoginFormValues>();
  const navigator = useNavigate();
  const onSubmit = async (data: LoginFormValues) => {
    try {
      const result = await login(data.email, data.password);
      localStorage.setItem("token", result.accessToken);
      navigator("/todos");
    } catch (error) {
      if (error instanceof ConnectError) {
        console.log("signin failed:", error.message);
      }
    }
  };
  return (
    <div>
      <h1>Login Page</h1>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          {...register("email")}
          type="email"
          placeholder="メールアドレス"
        />
        <input
          {...register("password")}
          type="password"
          placeholder="パスワード"
        />
        <Button type="submit">ログイン</Button>
      </form>
      <div className=" grid-cols-3 gap-4 p-4">
        <Button>
          <Link to={"/signup"}>アカウント作成</Link>
        </Button>
      </div>
    </div>
  );
}
