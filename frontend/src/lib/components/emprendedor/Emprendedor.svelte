<script>
  import { onMount } from "svelte";
  import { token, users } from "$lib/stores/auth";
  import { getEmprendedores, createEmprendedor } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let emprendedores = [];
  let tokenId = "";
  let entrepreneurId = "";
  let type = "info";
  let cargando = false;

  let url = "",
    description = "";

  async function handleEmprendedores(tokenId, entrepreneurId, type) {
    try {
      cargando = true;
      const { data } = await getEmprendedores(tokenId, entrepreneurId, type);
      emprendedores = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  async function handleSubmit() {
    try {
      const { code } = await createEmprendedor(tokenId, {
        url,
        description,
        entrepreneurId,
        type,
      });
      if (code === 201) {
        url = "";
        description = "";
        await handleEmprendedores(tokenId, entrepreneurId, type);
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
      }
    });
    await handleEmprendedores(tokenId, entrepreneurId, type);
  });
</script>

<div class="container_emprendedor">
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

  <div class="container_emprendedor_info">
    <h1>EMPRENDEDOR</h1>
    {#if cargando}
      <Cargando />
    {:else}
      {#each emprendedores as emprendedor}
        <p>{emprendedor.Url}</p>
        <p>{emprendedor.Description}</p>
      {/each}
    {/if}
  </div>
</div>

<style>
  .container_emprendedor {
    height: 100vh;
    overflow: hidden;
  }
  .container_emprendedor_info {
    border: 1px solid #ccc;
    height: 100%;
    overflow-y: auto;
  }
</style>
