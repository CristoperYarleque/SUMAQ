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
  let noticias = false;
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
    
  }

  function handleNoticias() {
    productos = false;
    noticias = true;
    promociones = false;
    cartComponent = false;
  }

  function handlePromociones() {
    productos = false;
    noticias = false;
    promociones = true;
    cartComponent = false;
  }

  function handleCart() {
    cartComponent = true;
    resumenComponent = false;
    productos = false;
    noticias = false;
    promociones = false;
  }

  function handleConfirmar() {
    cartComponent = false;
    resumenComponent = true;
  }

  function handleExito() {
    resumenComponent = false;
    exito = true;
  }

  function handleLogout() {
    logout();
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
    <!-- <button class="button_cliente">{name.toUpperCase()}</button> -->
    <button class="button_cliente" on:click={handleProductos}>PRODUCTOS</button>
    <button class="button_cliente" on:click={handleNoticias}>NOTICIAS</button>
    <button class="button_cliente" on:click={handlePromociones}>PROMOCIONES</button>
    <button class="button_logout" on:click={handleLogout}
      >CERRAR SESIÓN <LogOut class="w-1 h-1" /></button
    >
  </div>
  <div class="container_cliente_right">

      
    <div class="container_cliente_right_header">
      <p>{name.toUpperCase()}</p>
      <div class="floating_cart" on:click={handleCart}>
        <ShoppingCart /> {
          $cart.reduce((total, item) => total + item.cantidad, 0)
        }
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
      <Promociones />
    {/if}
  </div>
</div>

<style>
  .container_cliente {
    display: grid;
    grid-template-columns: 1fr 3fr;
  }
  .container_cliente_left {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    height: 100vh;
    border-right: 0.5px solid #ccc;
  }
  .container_cliente_right {
    height: 100vh;
    border-left: 0.5px solid #ccc;
  }

  .container_cliente_left .button_cliente {
    font-size: 1rem;
    font-weight: 400;
    color: #000000;
    text-align: center;
    cursor: pointer;
    transition: font-size 0.3s ease;
    background: none;
    border: none;
    padding: 0;
    &:hover {
      font-size: 1.1rem;
    }
  }

  .container_cliente_left .button_logout {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 1rem;
    font-weight: 400;
    color: #000000;
    text-align: center;
    cursor: pointer;
    width: 60%;
    border-radius: 10px;
    border: none;
    margin: 0 auto;
    padding: 10px;
    transition: all 0.3s ease;
    &:hover {
      background-color: #000000;
      color: #ffffff;
      width: 65%;
    }
  }
  .floating_cart {
    position: fixed;
    right: 50px;
    top: 30px;
    background-color: #000;
    color: white;
    padding: 7px 10px;
    border-radius: 50px;
    font-size: 1rem;
    font-weight: bold;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
    cursor: pointer;
    transition: all 0.3s ease;
    &:hover {
      background-color: #000000;
      color: #ffffff;
      padding: 8px 11px;
    }
  }

  .container_cliente_right_header {
    position: fixed;
    top: -12px;
    right: 10px;
    z-index: 9999;
  }
</style>
