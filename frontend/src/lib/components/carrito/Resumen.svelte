<script>
    import { cart } from "$lib/stores/cart";
    import { createEventDispatcher } from "svelte";
    import { onMount } from "svelte";
  
    const dispatch = createEventDispatcher();
  
    let nombre = "";
    let direccion = "";
    let tarjeta = "";
  
    $: total = $cart.reduce((sum, item) => sum + item.Price * item.cantidad, 0);
  
    function confirmarPago() {
      if (!nombre || !direccion || !tarjeta) {
        alert("Por favor completa todos los campos");
        return;
      }
  
      // Simular éxito
      dispatch("pedidoRealizado");
      cart.set([]); // Vaciar el carrito
    }
  </script>
  
  <div class="resumen_container">
    <h2>Resumen del Pedido</h2>
  
    <ul>
      {#each $cart as item}
        <li>{item.cantidad} x {item.Name} - S/. {item.Price * item.cantidad}</li>
      {/each}
    </ul>
  
    <p><strong>Total: S/. {total.toFixed(2)}</strong></p>
  
    <div class="formulario_pago">
      <label>Nombre:</label>
      <input type="text" bind:value={nombre} />
  
      <label>Dirección de entrega:</label>
      <input type="text" bind:value={direccion} />
  
      <label>Número de tarjeta:</label>
      <input type="text" bind:value={tarjeta} maxlength="16" />
  
      <button on:click={confirmarPago}>Pagar</button>
    </div>
  </div>
  
  <style>
    .resumen_container {
      padding: 20px;
    }
  
    .resumen_container ul {
      margin: 10px 0;
    }
  
    .formulario_pago {
      margin-top: 20px;
      display: flex;
      flex-direction: column;
      gap: 10px;
    }
  
    .formulario_pago input {
      padding: 5px;
      font-size: 1rem;
      border-radius: 5px;
      border: 1px solid #ccc;
    }
  
    .formulario_pago button {
      padding: 10px;
      background-color: #007bff;
      color: white;
      border: none;
      font-weight: bold;
      border-radius: 6px;
      cursor: pointer;
    }
  
    .formulario_pago button:hover {
      background-color: #0056b3;
    }
  </style>
  