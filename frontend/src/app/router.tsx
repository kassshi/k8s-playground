import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import { LoginPage } from "../pages/login/login";
import { SignupPage } from "../pages/signup/page";
import { TodosPage } from "../pages/todos/page";
import { Header } from "./header";

function PrivateRoute({ children }: { children: React.ReactNode }) {
  const isAuthenticated = !!localStorage.getItem("token");
  return isAuthenticated ? (
    <>
      <Header />
      {children}
    </>
  ) : (
    <Navigate to="/login" />
  );
}

export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route
          path="/todos"
          element={
            <PrivateRoute>
              <TodosPage />
            </PrivateRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}
