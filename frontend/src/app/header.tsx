import { Button } from "@/components/ui/button";

function Logout() {
  localStorage.removeItem("token");
  window.location.href = "/login";
}

export function Header() {
  return (
    <header className="p-6 items-right flex justify-end border-b">
      <Button onClick={Logout} variant="destructive" className="ml-4">
        Logout
      </Button>
    </header>
  );
}
