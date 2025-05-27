import { writable } from "svelte/store";

const storedCart =
  typeof localStorage !== "undefined"
    ? JSON.parse(localStorage.getItem("cart") || "[]")
    : [];

export const cart = writable(storedCart);

cart.subscribe((value) => {
  if (typeof localStorage !== "undefined") {
    localStorage.setItem("cart", JSON.stringify(value));
  }
});

export function addToCart(producto) {
  cart.update((items) => {
    const index = items.findIndex(
      (item) => item.ProductId === producto.ProductId
    );
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
