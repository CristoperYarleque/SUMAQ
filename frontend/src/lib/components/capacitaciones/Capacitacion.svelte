<script>
  import { onMount } from "svelte";
  import { token, users } from "$lib/stores/auth";
  import { getCapacitaciones, createCapacitacion } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let capacitacion = [];
  let tokenId = "";
  let entrepreneurId = "";
  let role = "";
  let type = "training";
  let cargando = false;

  let url = "",
    description = "";

  async function handleCapacitaciones(tokenId, entrepreneurId, type) {
    try {
      cargando = true;
      const { data } = await getCapacitaciones(tokenId, entrepreneurId, type);
      capacitacion = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  async function handleSubmit() {
    try {
      const { code } = await createCapacitacion(tokenId, {
        url,
        description,
        entrepreneurId,
        type,
      });
      if (code === 201) {
        url = "";
        description = "";
        await handleCapacitaciones(tokenId, entrepreneurId, type);
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
    await handleCapacitaciones(tokenId, entrepreneurId, type);
  });
</script>

<div>
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
  <div>
    <h1>CAPACITACIONES</h1>
    {#if cargando}
      <Cargando />
    {:else}
      {#each capacitacion as capacitacion}
        <p>{capacitacion.Url}</p>
        <p>{capacitacion.Description}</p>
      {/each}
    {/if}
  </div>
</div>

<style>
</style>
