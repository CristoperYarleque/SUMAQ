// src/lib/stores/auth.js
import { writable } from "svelte/store";

export const users = writable({});
export const token = writable(null);

export function login({ UserId, Token, Name, Role }) {
  const userData = { name: Name, role: Role, userId: UserId };
  users.set(userData);
  token.set(Token);
  sessionStorage.setItem("user", JSON.stringify(userData));
  sessionStorage.setItem("token", Token);
  document.cookie = `token=${Token}; Path=/;`;
  document.cookie = `role=${Role}; Path=/;`;
}

export function logout() {
  users.set({});
  token.set(null);
  sessionStorage.removeItem("token");
  sessionStorage.removeItem("user");
  document.cookie = "token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
  document.cookie = "role=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
}

if (typeof sessionStorage !== "undefined") {
  const savedToken = sessionStorage.getItem("token");
  const savedUser = sessionStorage.getItem("user");
  if (savedToken) token.set(savedToken);
  if (savedUser) users.set(JSON.parse(savedUser));
}
