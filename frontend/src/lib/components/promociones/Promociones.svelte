<script>
  import { onMount } from "svelte";
  import { token } from "$lib/stores/auth";
  import { getPromociones } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let tokenId = "";
  let promociones = [];
  let cargando = false;
  let startDate = "";
  let endDate = "";

  async function handlePromociones(tokenId) {
    try {
      cargando = true;
      const { data } = await getPromociones(tokenId, startDate, endDate);
      promociones = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  onMount(async () => {
    token.subscribe(async (token) => {
      if (token) {
        tokenId = token;
      }
    });

    const sixMonthsAgo = new Date();
    sixMonthsAgo.setMonth(sixMonthsAgo.getMonth() - 6);
    const sixMonthsLater = new Date();
    sixMonthsLater.setMonth(sixMonthsLater.getMonth() + 6);

    // Formato YYYY-MM-DD
    startDate = sixMonthsAgo.toISOString().slice(0, 10);
    endDate = sixMonthsLater.toISOString().slice(0, 10);

    await handlePromociones(tokenId);
  });
</script>

<div class="container_promociones">
  <h1 class="promociones_titulo">PROMOCIONES</h1>
  {#if cargando}
    <Cargando />
  {:else}
    <div class="banners">
      {#each promociones as promocion}
        <div class="banner">
          <img src={promocion.Url} alt={promocion.Description} />
          <div class="banner_texto">
            <h2>{promocion.Description}</h2>
            <p>Desde {new Date(promocion.StartDate).toLocaleDateString()} hasta {new Date(promocion.EndDate).toLocaleDateString()}</p>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  :root {
    --title-color_1: #B17D62;
  }
  .container_promociones {
    padding: 0 2rem;
    min-height: 100vh;
    overflow-y: auto;
    .promociones_titulo {
      text-align: center;
      margin-bottom: 2rem;
      font-size: 2.5rem;
      font-weight: 700;
      color: var(--title-color_1);
    }
  }

  .banners {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .banner {
    position: relative;
    width: 100%;
    height: 300px;
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s, box-shadow 0.3s;
  }

  .banner:hover {
    transform: translateY(-5px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
  }

  .banner img {
    width: 100%;
    height: 100%;
    object-fit: fill;
    filter: brightness(0.7);
  }

  .banner_texto {
    position: absolute;
    bottom: 1rem;
    left: 1rem;
    color: white;
    text-shadow: 1.5px 1.5px 4px rgb(0, 0, 0);
  }

  .banner_texto h2 {
    font-size: 1.8rem;
    margin: 0 0 0.5rem 0;
  }

  .banner_texto p {
    font-size: 1rem;
    margin: 0;
  }

  @media (max-width: 600px) {
    .banner {
      height: 200px;
    }

    .banner_texto h2 {
      font-size: 1.2rem;
    }

    .banner_texto p {
      font-size: 0.9rem;
    }
  }
</style>