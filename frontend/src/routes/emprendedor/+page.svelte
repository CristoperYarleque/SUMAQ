<script>
  import { ArrowRight, ClipboardList, Presentation, FolderHeart } from "lucide-svelte";
  import { onMount } from "svelte";
  import Capacitacion from "$lib/components/capacitaciones/Capacitacion.svelte";
  import Bienestar from "$lib/components/bienestar_emocional/Bienestar.svelte";
  import Productos from "$lib/components/productos_emprendedor/Productos.svelte";
  import Emprendedor from "$lib/components/emprendedor/Emprendedor.svelte";
  import Chatbot from "$lib/components/chatbot/Chatbot.svelte";
  import { getEmprendedor } from "$lib/api";
  import { logout, users, token } from "$lib/stores/auth";
  import { goto } from "$app/navigation";

  let productos = false;
  let capacitaciones = false;
  let bienestarEmocional = false;
  let emprendedor = false;
  let logo = true;
  let entrepreneurId = "";
  let tokenId = "";
  let name = "";
  let logoEntrepreneur = "";

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


  async function handleEmprendedores(tokenId, entrepreneurId) {
    try {
      const { data } = await getEmprendedor(tokenId, entrepreneurId);
      logoEntrepreneur = data.Url;
    } catch (error) {
      console.error(error);
    }
  }

  onMount(async () => {
    users.subscribe((user) => {
      if (user) {
        entrepreneurId = user.userId;
        name = user.name;
      }
    });

    token.subscribe(async (token) => {
      if (token) {
        tokenId = token;
      }
    });

    await handleEmprendedores(tokenId, entrepreneurId);
  });
</script>

<div class="container_emprendedor">
  <div class="container_emprendedor_left">
    <button class="button_emprendedor" on:click={handleEmprendedor}>
      <img src={logoEntrepreneur} alt="logo" class="logo_emprendedor" /><br>
      {name.toUpperCase()}
    </button>
    <button class="button_emprendedor" on:click={handleProductos}>
      <p><ClipboardList size="40" /></p>
      INVENTARIO
    </button>
    <button class="button_emprendedor" on:click={handleCapacitaciones}>
      <p><Presentation size="40" /></p>
      CAPACITACIONES
    </button>
    <button class="button_emprendedor" on:click={handleBienestarEmocional}>
      <p><FolderHeart size="40" /></p>
      BIENESTAR EMOCIONAL
    </button>
    <button class="button_logout" on:click={handleLogout}>
      CERRAR SESIÓN <ArrowRight size="20" />
    </button>
  </div>
  <div class="container_emprendedor_right">
    <Chatbot />
    {#if emprendedor}
      <Emprendedor on:logoUpdated={async () => await handleEmprendedores(tokenId, entrepreneurId)} />
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
    --title-color_2: #a0bea5;
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
    /* color: white; */
    color: var(--text-color);
    background: none;
    /* border: 2px solid white; */
    border: none;
    /* border: 2px solid black; */
    padding: 0.5rem 1rem;
    border-radius: 20px;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .button_emprendedor:hover {
    /* background-color: white; */
    /* color: var(--primary-color); */
    color: white;
  }

  .logo_emprendedor {
    width: 50px;
    height: 50px;
    border-radius: 0.5rem;
    margin-bottom: 1rem;
    object-fit: cover;
  }

  .button_logout {
    background-color: white;
    color: var(--primary-color);
    padding: 0.5rem 1rem;
    border-radius: 20px;
    cursor: pointer;
    transition: all 0.3s ease;
    border: 1px solid var(--title-color_2);
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
