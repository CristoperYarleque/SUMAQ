import { writable } from "svelte/store";

// Leer el carrito desde localStorage (si existe)
const storedCart = typeof localStorage !== "undefined"
  ? JSON.parse(localStorage.getItem("cart") || "[]")
  : [];

export const cart = writable(storedCart);

// Guardar en localStorage cada vez que el carrito cambia
cart.subscribe((value) => {
  if (typeof localStorage !== "undefined") {
    localStorage.setItem("cart", JSON.stringify(value));
  }
});

// Agregar producto (aumenta cantidad si ya existe)
export function addToCart(producto) {
  cart.update((items) => {
    const index = items.findIndex(item => item.ProductId === producto.ProductId);
    if (index !== -1) {
      // Ya existe, aumentar cantidad
      items[index].cantidad += 1;
      return [...items];
    } else {
      // Nuevo producto con cantidad 1
      return [...items, { ...producto, cantidad: 1 }];
    }
  });
}

