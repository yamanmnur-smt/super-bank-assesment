
const base_url = process.env.NEXT_PUBLIC_CLIENT_URL || "http://localhost:3000";
const routes = {
  logout: `${base_url}/api/logout`,
}


export const LogoutRequest = async () => {
  const res = await fetch(`${routes.logout}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!res.ok) {
    throw new Error("Invalid credentials");
  }
};
