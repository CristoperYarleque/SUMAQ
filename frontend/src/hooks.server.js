import { redirect } from "@sveltejs/kit";

const publicRoutes = ["/login", "/register"];

export async function handle({ event, resolve }) {
  const token = event.cookies.get("token") || null;
  const role = event.cookies.get("role") || null;
  const path = event.url.pathname;

  const publicRoutes = ["/login", "/register"];

  if (!token && !publicRoutes.includes(path)) {
    throw redirect(303, "/login");
  }

  if (token && publicRoutes.includes(path)) {
    const home =
      role === "entrepreneur" || role === "admin" ? "/emprendedor" : "/cliente";
    throw redirect(303, home);
  }

  if (
    token &&
    path.startsWith("/emprendedor") &&
    role !== "entrepreneur" &&
    role !== "admin"
  ) {
    throw redirect(303, "/cliente");
  }
  
  if (token && path.startsWith("/cliente") && role !== "client") {
    throw redirect(303, "/emprendedor");
  }

  return resolve(event);
}
