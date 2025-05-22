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

<div>
  <h1>PROMOCIONES</h1>
  {#if cargando}
    <Cargando />
  {:else}
      {#each promociones as promocion}
        <p>{promocion.Description}</p>
        <p>{promocion.StartDate}</p>
        <p>{promocion.EndDate}</p>
    {/each}
  {/if}
</div>

<style>
</style>
