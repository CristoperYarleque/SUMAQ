<script>
  import { onMount } from "svelte";
  import { LogOut, ShoppingCart } from "lucide-svelte";
  import { logout, users } from "$lib/stores/auth";
  import { goto } from "$app/navigation";
  import Promociones from "$lib/components/promociones/Promociones.svelte";
  import Noticias from "$lib/components/noticias/Noticias.svelte";
  import Productos from "$lib/components/productos_cliente/Productos.svelte";
  import Carrito from "$lib/components/carrito/Carrito.svelte";
  import Resumen from "$lib/components/carrito/Resumen.svelte";
  import ModalExito from "$lib/components/carrito/ModalExito.svelte";
  import { cart } from "$lib/stores/cart";

  let productos = false;
  let noticias = true;
  let promociones = false;
  let cartComponent = false;
  let resumenComponent = false;
  let exito = false;
  let name = "";

  function handleProductos() {
    productos = true;
    noticias = false;
    promociones = false;
    cartComponent = false;
    resumenComponent = false;
    exito = false;
  }

  function handleNoticias() {
    productos = false;
    noticias = true;
    promociones = false;
    cartComponent = false;
    resumenComponent = false;
    exito = false;
  }

  function handlePromociones() {
    productos = false;
    noticias = false;
    promociones = true;
    cartComponent = false;
    resumenComponent = false;
    exito = false;
  }

  function handleCart() {
    cartComponent = true;
    resumenComponent = false;
    productos = false;
    noticias = false;
    promociones = false;
    exito = false;
  }

  function handleConfirmar() {
    cartComponent = false;
    resumenComponent = true;
    exito = false;
    productos = false;
    noticias = false;
    promociones = false;
  }

  function handleExito() {
    resumenComponent = false;
    exito = true;
    productos = false;
    noticias = false;
    promociones = false;
    cartComponent = false;
  }

  function handleLogout() {
    logout();
    cart.set([]);
    goto("/login");
  }

  onMount(async () => {
    users.subscribe((user) => {
      if (user) {
        name = user.name;
      }
    });
  });
</script>

<div class="container_cliente">
  <div class="container_cliente_left">
    <button class="button_cliente" on:click={handleProductos}>PRODUCTOS</button>
    <button class="button_cliente" on:click={handleNoticias}>NOTICIAS</button>
    <button class="button_cliente" on:click={handlePromociones}
      >PROMOCIONES</button
    >
    <button class="button_logout" on:click={handleLogout}>
      CERRAR SESIÓN <LogOut size="20" />
    </button>
  </div>
  <div class="container_cliente_right">
    <div class="container_cliente_right_header">
      <p class="nombre_usuario">{name.toUpperCase()}</p>
      <div class="floating_cart" on:click={handleCart}>
        <ShoppingCart />
        {$cart.reduce((total, item) => total + item.cantidad, 0)}
      </div>
    </div>
    {#if productos}
      <Productos />
    {:else if cartComponent}
      <Carrito on:confirmar={handleConfirmar} />
    {:else if resumenComponent}
      <Resumen on:pedidoRealizado={handleExito} />
    {:else if exito}
      <ModalExito />
    {:else if noticias}
      <Noticias />
    {:else if promociones}
      <Promociones on:changeToProducts={handleProductos} />
    {/if}
  </div>
</div>

<style>
  :root {
    --primary-color: #ebb2bd;
    --secondary-color: #ede9e4;
    --text-color: #333;
    --hover-color: #f29dae;
    --background-light: #f9f9f9;
    --title-color_2: #a0bea5;
  }

  .container_cliente {
    display: grid;
    grid-template-columns: 1fr 3fr;
    height: 100vh;
    background-color: var(--background-light);
  }

  .container_cliente_left {
    background-color: var(--primary-color);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    color: white;
  }

  .nombre_usuario {
    font-weight: 800;
    color: var(--title-color_1);
  }

  .button_cliente {
    font-size: 1.1rem;
    font-weight: 600;
    color: white;
    background: none;
    border: 2px solid white;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .button_cliente:hover {
    background-color: white;
    color: var(--primary-color);
  }

  .button_logout {
    background-color: white;
    color: var(--primary-color);
    padding: 0.5rem 1rem;
    border-radius: 20px;
    cursor: pointer;
    transition: all 0.3s ease;
    border: 1px solid var(--title-color_2);
  }

  .button_logout:hover {
    background-color: var(--hover-color);
    color: white;
  }

  .container_cliente_right {
    background-color: var(--secondary-color);
    position: relative;
    overflow-y: auto;
  }

  .container_cliente_right_header {
    position: fixed;
    top: 1rem;
    right: 2rem;
    display: flex;
    align-items: center;
    gap: 1rem;
    background: white;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    z-index: 100;
  }

  .floating_cart {
    background-color: var(--primary-color);
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .floating_cart:hover {
    background-color: var(--hover-color);
  }
</style>
