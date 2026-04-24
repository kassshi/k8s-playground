import { BrowserRouter, Route, Routes } from "react-router";
import { LoginPage } from "../pages/login/login";
import { SignupPage } from "../pages/signup/page";

export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/todos" element={<div>Todo Page</div>} />
      </Routes>
    </BrowserRouter>
  );
}
