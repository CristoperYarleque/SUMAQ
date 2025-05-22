<script>
  import { LogOut } from "lucide-svelte";
  import { onMount } from "svelte";
  import Capacitacion from "$lib/components/capacitaciones/Capacitacion.svelte";
  import Bienestar from "$lib/components/bienestar_emocional/Bienestar.svelte";
  import Productos from "$lib/components/productos_emprendedor/Productos.svelte";
  import Emprendedor from "$lib/components/emprendedor/Emprendedor.svelte";
  import { logout, users } from "$lib/stores/auth";
  import { goto } from "$app/navigation";

  let categorias = false;
  let capacitaciones = false;
  let bienestarEmocional = false;
  let emprendedor = false;
  let name = "";

  function handleCategorias() {
    categorias = true;
    capacitaciones = false;
    bienestarEmocional = false;
    emprendedor = false;
  }

  function handleCapacitaciones() {
    categorias = false;
    capacitaciones = true;
    bienestarEmocional = false;
    emprendedor = false;
  }

  function handleBienestarEmocional() {
    categorias = false;
    capacitaciones = false;
    bienestarEmocional = true;
    emprendedor = false;
  }

  function handleEmprendedor() {
    categorias = false;
    capacitaciones = false;
    bienestarEmocional = false;
    emprendedor = true;
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
    <button class="button_emprendedor" on:click={handleCategorias}
      >INVENTARIO</button
    >
    <button class="button_emprendedor" on:click={handleCapacitaciones}
      >CAPACITACIONES</button
    >
    <button class="button_emprendedor" on:click={handleBienestarEmocional}
      >BIENESTAR EMOCIONAL</button
    >
    <button class="button_logout" on:click={handleLogout}
      >CERRAR SESIÓN <LogOut class="w-1 h-1 " /></button
    >
  </div>
  <div class="container_emprendedor_right">
    {#if emprendedor}
      <Emprendedor />
    {:else if categorias}
      <Productos />
    {:else if capacitaciones}
      <Capacitacion />
    {:else if bienestarEmocional}
      <Bienestar />
    {/if}
  </div>
</div>

<style>
  .container_emprendedor {
    display: grid;
    grid-template-columns: 1fr 3fr;
  }
  .container_emprendedor_left {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    height: 100vh;
    border-right: 3px solid #00000074;
  }
  .container_emprendedor_right {
    height: 100vh;
  }

  .container_emprendedor_left .button_emprendedor {
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

  .container_emprendedor_left .button_logout {
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
