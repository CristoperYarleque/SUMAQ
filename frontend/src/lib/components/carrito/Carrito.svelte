<script>
  import { cart } from "$lib/stores/cart";
  import { ShoppingCart, Trash2 } from "lucide-svelte";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

   const regaloId = "REGALO_BOLSA";

  // Eliminar un producto completamente del carrito
  function removeFromCart(id) {
    cart.update((items) => items.filter((item) => item.ProductId !== id));
  }

  // Aumentar cantidad
  function increaseQty(id) {
    cart.update((items) =>
      items.map((item) =>
        item.ProductId === id ? { ...item, cantidad: item.cantidad + 1 } : item
      )
    );
  }

  // Disminuir cantidad (mínimo 1)
  function decreaseQty(id) {
    cart.update((items) =>
      items.map((item) =>
        item.ProductId === id && item.cantidad > 1
          ? { ...item, cantidad: item.cantidad - 1 }
          : item
      )
    );
  }

  // Alternar bolsa de regalo
  function toggleRegalo() {
    const yaExiste = $cart.some(item => item.ProductId === regaloId);
    cart.update(items => {
      if (yaExiste) {
        return items.filter(item => item.ProductId !== regaloId);
      } else {
        return [
          ...items,
          {
            ProductId: regaloId,
            Name: "Bolsa de regalo",
            Price: 2.99,
            cantidad: 1,
            Url: "/bolsa.png"
          }
        ];
      }
    });
  }

  $: tieneRegalo = $cart.some(item => item.ProductId === regaloId);

  $: total = $cart.reduce((sum, item) => sum + item.cantidad * item.Price, 0);

  function confirmarPedido() {
    dispatch("confirmar");
  }
</script>

<div class="carrito_container">
  <h2 class="color_title"><ShoppingCart /> Tu carrito</h2>

  {#if $cart.length === 0}
    <p class="color_title">El carrito está vacío.</p>
  {:else}
    <ul class="carrito_list">
      {#each $cart as item}
        <li class="carrito_item">
          <div class="carrito_item_left">
            <img class="carrito_item_img" src={item.Url} alt={item.Name} />
            <div>
              <p><strong>{item.Name}</strong></p>
              <p>S/. {item.Price.toFixed(2)} x {item.cantidad} = <strong>S/. {(item.Price * item.cantidad).toFixed(2)}</strong></p>
            </div>
          </div>
          <div class="carrito_controls">
            <button on:click={() => decreaseQty(item.ProductId)}>-</button>
            <span>{item.cantidad}</span>
            <button on:click={() => increaseQty(item.ProductId)}>+</button>
            <button
              class="remove"
              on:click={() => removeFromCart(item.ProductId)}><Trash2 /></button
            >
          </div>
        </li>
      {/each}
    </ul>
    <div class="regalo_toggle">
      {#if !tieneRegalo}
        <button on:click={toggleRegalo}>
        {tieneRegalo ? "Quitar bolsa de regalo" : "Agregar bolsa de regalo (+ S/ 2.99)"}
      </button>
      {/if}
    </div>
    <div class="carrito_total_container">
      <p class="carrito_total color_title">
        Total: <strong>S/. {total.toFixed(2)}</strong>
      </p>
      <button class="confirmar_button" on:click={confirmarPedido}
        >Confirmar pedido</button
      >
    </div>
  {/if}
</div>

<style>
  :root {
    --primary-color: #ebb2bd;
    --secondary-color: #ede9e4;
    --text-color: #333333d7;
    --hover-color: #f29dae;
    --background-light: #f9f9f9;
    --title-color_1: #b17d62;
    --title-color_2: #a0bea5;
  }
  .carrito_container {
    padding: 20px;
    overflow-y: auto;
    height: 100%;
    color: var(--text-color);
  }

  .carrito_item_left {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .carrito_item_img {
    width: 60px;
    height: 60px;
    object-fit: cover;
    border-radius: 5px;
  }

  .carrito_item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #ccc;
    padding: 10px 0;
  }

  .carrito_controls {
    display: flex;
    align-items: center;
  }

  .carrito_controls button {
    margin: 0px 5px;
    padding: 8px 10px;
    font-size: 1rem;
    border-radius: 4px;
    border: 2px solid var(--title-color_2);
    cursor: pointer;
  }

  .carrito_controls .remove {
    background-color: #ff4d50be;
    color: white;
    border: none;
    border-radius: 4px;
    padding: 5px;
    cursor: pointer;
  }

  .carrito_total {
    margin-top: 20px;
    font-size: 1.2rem;
    font-weight: bold;
  }
  .confirmar_button {
    margin-top: 20px;
    padding: 10px 15px;
    background-color: #ff85a2;
    color: white;
    font-weight: bold;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  .confirmar_button:hover {
    background-color: #f06292;
  }

  .carrito_total_container {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .color_title {
    color: var(--title-color_1);
  }

  .carrito_list {
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin-top: 2rem;
    background-color: var(--background-light);
    padding: 2rem;
    border-radius: 0.5rem;
  }

  .regalo_toggle {
  margin-top: 1.5rem;
  text-align: end;
}

.regalo_toggle button {
  padding: 10px 20px;
  background-color: var(--secondary-color);
  border: 2px dashed var(--title-color_2);
  color: var(--text-color);
  font-weight: bold;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.regalo_toggle button:hover {
  background-color: #ff85a2;
  color: white;
  border-style: none;
}
</style>
