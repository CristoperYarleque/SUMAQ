<script>
  import { LogOut } from "lucide-svelte";
  import { onMount } from "svelte";
  import { logout, users } from "$lib/stores/auth";
  import { goto } from "$app/navigation";
  import Promociones from "$lib/components/promociones/Promociones.svelte";
  import Noticias from "$lib/components/noticias/Noticias.svelte";

  let name = "";
  let categorias = false;
  let noticias = false;
  let promociones = false;

  function handleCategorias() {
    categorias = true;
    noticias = false;
    promociones = false;
  }

  function handleNoticias() {
    categorias = false;
    noticias = true;
    promociones = false;
  }

  function handlePromociones() {
    categorias = false;
    noticias = false;
    promociones = true;
  }

  function handleLogout() {
    logout();
    goto("/login");
  }

  onMount(() => {
    users.subscribe((user) => {
      name = user.name;
    });
  });
</script>

<div class="container_cliente">
  <div class="container_cliente_left">
    <!-- <button class="button_cliente">{name.toUpperCase()}</button> -->
    <button class="button_cliente" on:click={handleCategorias}>CATEGORIAS</button>
    <button class="button_cliente" on:click={handleNoticias}>NOTICIAS</button>
    <button class="button_cliente" on:click={handlePromociones}>PROMOCIONES</button>
    <button class="button_logout" on:click={handleLogout}
      >CERRAR SESIÓN <LogOut class="w-1 h-1" /></button
    >
  </div>
  <div class="container_cliente_right">
    {#if categorias}
      <h1>CATEGORIAS</h1>
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
    border-right: 3px solid #00000074;
  }
  .container_cliente_right {
    height: 100vh;
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
</style>
