<script>
    import { onMount } from "svelte";
    import { Timer } from "lucide-svelte";
    import { token } from "$lib/stores/auth";
    import { getNoticias } from "$lib/api";
    import Cargando from "$lib/components/cargando/Cargando.svelte";
  
    let tokenId = "";
    let noticias = [];
    let cargando = false;
    let startDate = "";
    let endDate = "";
  
    async function handleNoticias(tokenId) {
      try {
        cargando = true;
        const { data } = await getNoticias(tokenId, startDate, endDate);
        noticias = data;
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
  
      await handleNoticias(tokenId);
    });
  </script>
  
  <div class="container_noticias">
    <h1 class="noticias_titulo">NOTICIAS</h1>
    {#if cargando}
      <Cargando />
    {:else}
      <div class="noticias_grid">
        {#each noticias as noticia}
          <div class="noticia_card">
            <p class="noticia_fecha"><Timer /> {new Date(noticia.PublishedAt).toLocaleString()}</p>
            <h2 class="noticia_titulo">{noticia.Title}</h2>
            <p class="noticia_contenido">{noticia.Content}</p>
          </div>
        {/each}
      </div>
    {/if}
  </div>
  
  <style>
    :root {
    --title-color_1: #B17D62;
  }
    .container_noticias {
      padding: 0 2rem;
      min-height: 100vh;
      overflow-y: auto;
      .noticias_titulo {
        color: var(--title-color_1);
      }
    }
  
    .container_noticias h1 {
      text-align: center;
      font-size: 2.5rem;
      margin-bottom: 2rem;
      color: #222; /* Color más sobrio */
      font-weight: 700;
    }
  
    .noticias_grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 1.5rem;
    }
  
    .noticia_card {
      background-color: white;
      border-radius: 20px; /* Bordes más redondeados */
      box-shadow: 0 8px 20px rgba(0, 0, 0, 0.08); /* Sombra más suave */
      padding: 1.5rem;
      transition: transform 0.3s ease, box-shadow 0.3s ease;
      display: flex;
      flex-direction: column;
      gap: 0.75rem;
    }
  
    .noticia_card:hover {
      transform: translateY(-6px);
      box-shadow: 0 12px 24px rgba(0, 0, 0, 0.12); /* Hover más sutil */
    }
  
    .noticia_fecha {
      font-size: 0.85rem;
      color: #888;
      text-align: right;
      font-style: italic;
    }
  
    .noticia_titulo {
      font-size: 1.3rem;
      font-weight: 600;
      color: var(--title-color_1); /* Azul más moderno (puedes cambiarlo por --primary-color) */
      margin: 0;
      line-height: 1.4;
    }
  
    .noticia_contenido {
      font-size: 1rem;
      color: #555;
      line-height: 1.5;
    }
  
    @media (max-width: 600px) {
      .container_noticias {
        padding: 1rem;
      }
  
      .noticia_card {
        padding: 1rem;
      }
  
      .noticia_titulo {
        font-size: 1.1rem;
      }
    }
  </style>
  