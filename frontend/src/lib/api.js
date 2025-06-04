// src/lib/api.js
import { logout } from "$lib/stores/auth";
import { goto } from "$app/navigation";

const API_BASE = import.meta.env.VITE_BACKEND_SUMAQ_URL;
const IMGBB_API_KEY = import.meta.env.VITE_IMGBB_API_KEY;
const IMGBB_URL = import.meta.env.VITE_IMGBB_URL;

async function authFetch(url, options = {}) {
  const res = await fetch(url, options);

  if (res.status === 401) {
    logout();
    goto("/login");
    throw new Error("Sesión expirada. Por favor inicia sesión nuevamente.");
  }

  return res;
}

export async function registerUser(data) {
  const res = await fetch(`${API_BASE}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Error en el registro o email ya existe");
  return res.json();
}

export async function loginUser(data) {
  const res = await fetch(`${API_BASE}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Credenciales inválidas");
  return res.json();
}

export async function getCategories(token) {
  const res = await authFetch(`${API_BASE}/categories`, {
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!res.ok) throw new Error("Error al obtener categorías");
  return res.json();
}

export async function createProduct(token, data) {
  const res = await authFetch(`${API_BASE}/products`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Error al crear producto");
  return res.json();
}

export async function getProducts(token, entrepreneurId, categoryId) {
  const res = await authFetch(
    `${API_BASE}/products?entrepreneurId=${entrepreneurId}&categoryId=${categoryId}`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
  if (!res.ok) throw new Error("Error al obtener productos");
  return res.json();
}

export async function deleteProduct(token, productId) {
  const res = await authFetch(`${API_BASE}/products/${productId}`, {
    method: "DELETE",
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!res.ok) throw new Error("Error al eliminar producto");
  return res.json();
}

export async function getCapacitaciones(token, entrepreneurId, type) {
  const res = await authFetch(
    `${API_BASE}/info-entrepreneurs?entrepreneurId=${entrepreneurId}&type=${type}`,
    {
      headers: { Authorization: `Bearer ${token}` },
      "Content-Type": "application/json",
    }
  );
  if (!res.ok) throw new Error("Error al obtener capacitaciones");
  return res.json();
}

export async function createCapacitacion(token, data) {
  const res = await authFetch(`${API_BASE}/info-entrepreneurs`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Error al crear capacitación");
  return res.json();
}

export async function getBienestar(token, entrepreneurId, type) {
  const res = await authFetch(
    `${API_BASE}/info-entrepreneurs?entrepreneurId=${entrepreneurId}&type=${type}`,
    {
      headers: { Authorization: `Bearer ${token}` },
      "Content-Type": "application/json",
    }
  );
  if (!res.ok) throw new Error("Error al obtener bienestar");
  return res.json();
}

export async function createBienestar(token, data) {
  const res = await authFetch(`${API_BASE}/info-entrepreneurs`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Error al crear bienestar");
  return res.json();
}

export async function getEmprendedores(token, entrepreneurId, type) {
  const res = await authFetch(
    `${API_BASE}/info-entrepreneurs?entrepreneurId=${entrepreneurId}&type=${type}`,
    {
      headers: { Authorization: `Bearer ${token}` },
      "Content-Type": "application/json",
    }
  );
  if (!res.ok) throw new Error("Error al obtener emprendedores");
  return res.json();
}

export async function getEmprendedor(token, entrepreneurId) {
  const res = await authFetch(`${API_BASE}/entrepreneurs/${entrepreneurId}`, {
    headers: { Authorization: `Bearer ${token}` },
    "Content-Type": "application/json",
  });
  if (!res.ok) throw new Error("Error al obtener emprendedor");
  return res.json();
}

export async function createEmprendedor(token, data) {
  const res = await authFetch(`${API_BASE}/info-entrepreneurs`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Error al crear emprendedor");
  return res.json();
}

export async function uploadImageToImgBB(file) {
  const formData = new FormData();
  formData.append("image", file);
  const res = await authFetch(`${IMGBB_URL}/upload?key=${IMGBB_API_KEY}`, {
    method: "POST",
    body: formData,
  });
  if (!res.ok) throw new Error("Error al subir imagen a ImgBB");
  return res.json();
}

export async function updateUrlEmprendedor(token, entrepreneurId, data) {
  const res = await authFetch(`${API_BASE}/entrepreneurs/${entrepreneurId}`, {
    method: "PUT",
    headers: { Authorization: `Bearer ${token}` },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Error al actualizar url emprendedor");
  return res.json();
}

export async function getPromociones(token, startDate, endDate) {
  const res = await authFetch(
    `${API_BASE}/promotions?startDate=${startDate}&endDate=${endDate}`,
    {
      headers: { Authorization: `Bearer ${token}` },
      "Content-Type": "application/json",
    }
  );
  if (!res.ok) throw new Error("Error al obtener promociones");
  return res.json();
}

export async function getNoticias(token, startDate, endDate) {
  const res = await authFetch(
    `${API_BASE}/news?startDate=${startDate}&endDate=${endDate}`,
    {
      headers: { Authorization: `Bearer ${token}` },
      "Content-Type": "application/json",
    }
  );
  if (!res.ok) throw new Error("Error al obtener noticias");
  return res.json();
}

export async function getEmprendedorByCategory(token, categoryId) {
  const res = await authFetch(
    `${API_BASE}/entrepreneurs?categoryId=${categoryId}`,
    {
      headers: { Authorization: `Bearer ${token}` },
      "Content-Type": "application/json",
    }
  );
  if (!res.ok) throw new Error("Error al obtener emprendedor por categoría");
  return res.json();
}

export async function getChatbot(token) {
  const res = await authFetch(`${API_BASE}/chatbot`, {
    headers: { Authorization: `Bearer ${token}` },
    "Content-Type": "application/json",
  });
  if (!res.ok) throw new Error("Error al obtener chatbot");
  return res.json();
}

export async function getReactionInfo(token, userId, newsId) {
  const res = await authFetch(`${API_BASE}/news-reactions?userId=${userId}&newsId=${newsId}`, {
    headers: { Authorization: `Bearer ${token}` },
     "Content-Type": "application/json",
  });
  if (!res.ok) throw new Error('Error al obtener información de reacciones');
  return res.json();
}

export async function upsertReaction(token, data) {
  const res = await authFetch(`${API_BASE}/news-reactions`, {
    method: 'POST',
    headers: { 
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}` 
    },
    body: JSON.stringify(data)
  });
  if (!res.ok) throw new Error('Error al registrar reacción');
  return res.json();
}

export async function deleteReaction(token, userId, newsId) {
  const res = await authFetch(`${API_BASE}/news-reactions?userId=${userId}&newsId=${newsId}`, {
    method: 'DELETE',
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!res.ok) throw new Error('Error al eliminar reacción');
  return res.json();
}