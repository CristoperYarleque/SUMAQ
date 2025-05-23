<script>
  import { onMount } from "svelte";
  import { token, users } from "$lib/stores/auth";
  import { getBienestar, createBienestar } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let bienestar = [];
  let tokenId = "";
  let entrepreneurId = "";
  let role = "";
  let type = "welfare";
  let cargando = false;

  let url = "",
    description = "";

  async function handleBienestar(tokenId, entrepreneurId, type) {
    try {
      cargando = true;
      const { data } = await getBienestar(tokenId, entrepreneurId, type);
      bienestar = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  async function handleSubmit() {
    try {
      const { code } = await createBienestar(tokenId, {
        url,
        description,
        entrepreneurId,
        type,
      });
      if (code === 201) {
        url = "";
        description = "";
        await handleBienestar(tokenId, entrepreneurId, type);
      }
    } catch (error) {
      console.error(error);
    }
  }

  onMount(async () => {
    token.subscribe(async (token) => {
      if (token) {
        tokenId = token;
      }
    });
    users.subscribe(async (user) => {
      if (user) {
        entrepreneurId = user.userId;
        role = user.role;
      }
    });
    await handleBienestar(tokenId, entrepreneurId, type);
  });
</script>

<div class="container_bienestar_emocional">
  {#if role === "admin"}
    <div>
      <form on:submit|preventDefault={handleSubmit}>
        <input
          type="text"
          name="url"
          bind:value={url}
          placeholder="Url"
          required
        />
        <input
          type="text"
          name="description"
          bind:value={description}
          placeholder="Descripción"
        />
        <button type="submit">Guardar</button>
      </form>
    </div>
  {/if}
  <div class="container_bienestar_emocional_info">
    <h1>BIENESTAR EMOCIONAL</h1>
    {#if cargando}
      <Cargando />
    {:else}
      {#each bienestar as bienestar}
        <p>{bienestar.Url}</p>
        <p>{bienestar.Description}</p>
      {/each}
    {/if}
  </div>
</div>

<style>
  .container_bienestar_emocional {
    height: 100vh;
    overflow: hidden;
  }
  .container_bienestar_emocional_info {
    border: 1px solid #ccc;
    height: 100%;
    overflow-y: auto;
  }
</style>
