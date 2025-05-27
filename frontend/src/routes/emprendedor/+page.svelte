<script>
  import { LogOut } from "lucide-svelte";
  import { onMount } from "svelte";
  import Capacitacion from "$lib/components/capacitaciones/Capacitacion.svelte";
  import Bienestar from "$lib/components/bienestar_emocional/Bienestar.svelte";
  import Productos from "$lib/components/productos_emprendedor/Productos.svelte";
  import Emprendedor from "$lib/components/emprendedor/Emprendedor.svelte";
  import Chatbot from "$lib/components/chatbot/Chatbot.svelte";
  import { logout, users } from "$lib/stores/auth";
  import { goto } from "$app/navigation";

  let productos = false;
  let capacitaciones = false;
  let bienestarEmocional = false;
  let emprendedor = false;
  let logo = true;
  let name = "";

  function handleProductos() {
    productos = true;
    capacitaciones = false;
    bienestarEmocional = false;
    emprendedor = false;
    logo = false;
  }

  function handleCapacitaciones() {
    productos = false;
    capacitaciones = true;
    bienestarEmocional = false;
    emprendedor = false;
    logo = false;
  }

  function handleBienestarEmocional() {
    productos = false;
    capacitaciones = false;
    bienestarEmocional = true;
    emprendedor = false;
    logo = false;
  }

  function handleEmprendedor() {
    productos = false;
    capacitaciones = false;
    bienestarEmocional = false;
    emprendedor = true;
    logo = false;
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

<div class="container_emprendedor">
  <div class="container_emprendedor_left">
    <button class="button_emprendedor" on:click={handleEmprendedor}
      >{name.toUpperCase()}</button
    >
    <button class="button_emprendedor" on:click={handleProductos}
      >INVENTARIO</button
    >
    <button class="button_emprendedor" on:click={handleCapacitaciones}
      >CAPACITACIONES</button
    >
    <button class="button_emprendedor" on:click={handleBienestarEmocional}
      >BIENESTAR EMOCIONAL</button
    >
    <button class="button_logout" on:click={handleLogout}>
      CERRAR SESIÓN <LogOut size="20" />
    </button>
  </div>
  <div class="container_emprendedor_right">
    <Chatbot />
    {#if emprendedor}
      <Emprendedor />
    {:else if productos}
      <Productos />
    {:else if logo}
      <div class="container_logo">
        <img src="/logo_hero.png" alt="logo" class="logo_hero" />
      </div>
    {:else if capacitaciones}
      <Capacitacion />
    {:else if bienestarEmocional}
      <Bienestar />
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
  }

  .container_emprendedor {
    display: grid;
    grid-template-columns: 1fr 3fr;
    height: 100vh;
    background-color: var(--background-light);
  }

  .container_emprendedor_left {
    background-color: var(--primary-color);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    color: white;
  }

  .button_emprendedor {
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

  .button_emprendedor:hover {
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
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .button_logout:hover {
    background-color: var(--hover-color);
    color: white;
  }

  .container_emprendedor_right {
    background-color: var(--secondary-color);
    overflow-y: auto;
    position: relative;
  }

  .container_logo {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }
</style>
