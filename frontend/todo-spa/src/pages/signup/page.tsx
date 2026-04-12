import { useForm } from "react-hook-form";
import { signup } from "../../features/auth/api/signup";

type SignupFormValues = {
  email: string;
  password: string;
  confirm_password: string;
};

export function SignupPage() {
  const { register, handleSubmit } = useForm<SignupFormValues>();

  const onSubmit = async (data: SignupFormValues) => {
    const result = await signup(
      data.email,
      data.password,
      data.confirm_password,
    );
    console.log(result);
  };
  return (
    <div>
      <h1>Signup Page</h1>
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
        <input
          {...register("confirm_password")}
          type="password"
          placeholder="パスワード"
        />
        <button type="submit">登録</button>
      </form>
    </div>
  );
}
