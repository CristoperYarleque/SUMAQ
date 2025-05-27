<script>
  import { cart } from "$lib/stores/cart";
  import { createEventDispatcher } from "svelte";

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
    cart.set([]);
  }
</script>

<div class="resumen_container">
  <h2 class="color_title">Resumen del Pedido</h2>

  <ul class="resumen_list">
    {#each $cart as item}
      <li>{item.cantidad} x {item.Name} - S/. {item.Price * item.cantidad}</li>
    {/each}
  </ul>

  <p class="color_title"><strong>Total: S/. {total.toFixed(2)}</strong></p>

  <div class="formulario_pago">
    <input type="text" bind:value={nombre} placeholder="Nombre" />
    <input type="text" bind:value={direccion} placeholder="Dirección" />
    <input
      type="text"
      bind:value={tarjeta}
      maxlength="16"
      placeholder="Número de tarjeta"
    />
    <button on:click={confirmarPago}>Pagar</button>
  </div>
</div>

<style>
  :root {
    --primary-color: #ebb2bd;
    --secondary-color: #ede9e4;
    --text-color: #333;
    --hover-color: #f29dae;
    --background-light: #f9f9f9;
    --title-color_1: #b17d62;
    --title-color_2: #a0bea5;
  }
  .resumen_container {
    padding: 20px;
    color: var(--text-color);
  }

  .formulario_pago {
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .formulario_pago input {
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    font-size: 1rem;
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    color: black;
    margin: 0 0.75rem;
  }

  .formulario_pago button {
    padding: 0.5rem 1rem;
    background-color: #ff85a2;
    color: white;
    font-weight: bold;
    border: none;
    width: 50%;
    margin: auto;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  .formulario_pago button:hover {
    background-color: #f06292;
  }

  .color_title {
    color: var(--title-color_1);
  }

  .resumen_list {
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin-top: 2rem;
    background-color: var(--background-light);
    padding: 2rem;
    border-radius: 0.5rem;
  }
</style>
