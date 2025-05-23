<script>
    import { cart } from "$lib/stores/cart";
    import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  
    // Eliminar un producto completamente del carrito
    function removeFromCart(id) {
      cart.update(items => items.filter(item => item.ProductId !== id));
    }
  
    // Aumentar cantidad
    function increaseQty(id) {
      cart.update(items =>
        items.map(item =>
          item.ProductId === id ? { ...item, cantidad: item.cantidad + 1 } : item
        )
      );
    }
  
    // Disminuir cantidad (mínimo 1)
    function decreaseQty(id) {
      cart.update(items =>
        items.map(item =>
          item.ProductId === id && item.cantidad > 1
            ? { ...item, cantidad: item.cantidad - 1 }
            : item
        )
      );
    }
  
    $: total = $cart.reduce((sum, item) => sum + item.cantidad * item.Price, 0);

    function confirmarPedido() {
    dispatch("confirmar");
  }
  </script>
  
  <div class="carrito_container">
    <h2>🛒 Tu carrito</h2>
  
    {#if $cart.length === 0}
      <p>El carrito está vacío.</p>
    {:else}
      <ul>
        {#each $cart as item}
          <li class="carrito_item">
            <div>
              <strong>{item.Name}</strong><br />
              S/. {item.Price.toFixed(2)} x {item.cantidad} = <strong>S/. {(item.Price * item.cantidad).toFixed(2)}</strong>
            </div>
            <div class="carrito_controls">
              <button on:click={() => decreaseQty(item.ProductId)}>-</button>
              <span>{item.cantidad}</span>
              <button on:click={() => increaseQty(item.ProductId)}>+</button>
              <button class="remove" on:click={() => removeFromCart(item.ProductId)}>🗑</button>
            </div>
          </li>
        {/each}
      </ul>
      <p class="carrito_total">Total: <strong>S/. {total.toFixed(2)}</strong></p>
      <button class="confirmar_button" on:click={confirmarPedido}>Confirmar pedido</button>
    {/if}
  </div>
  
  <style>
    .carrito_container {
      padding: 20px;
      overflow-y: auto;
      height: 100%;
    }
  
    .carrito_item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      border-bottom: 1px solid #ccc;
      padding: 10px 0;
    }
  
    .carrito_controls button {
      margin: 0 5px;
      padding: 5px 8px;
      font-size: 1rem;
    }
  
    .carrito_controls .remove {
      background-color: #ff4d4f;
      color: white;
      border: none;
      border-radius: 4px;
      padding: 5px 10px;
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
    background-color: #28a745;
    color: white;
    font-weight: bold;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  .confirmar_button:hover {
    background-color: #218838;
  }
  </style>
  