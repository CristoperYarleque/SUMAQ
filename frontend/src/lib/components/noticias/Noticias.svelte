<script>
    import { onMount } from "svelte";
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
    <h1>NOTICIAS</h1>
    {#if cargando}
      <Cargando />
    {:else}
        {#each noticias as noticia}
          <p>{noticia.Title}</p>
          <p>{noticia.Content}</p>
          <p>{noticia.PublishedAt}</p>
      {/each}
    {/if}
  </div>

  <style>
    .container_noticias {
      height: 100vh;
      overflow: hidden;
      overflow-y: auto;
    }
  </style>
  