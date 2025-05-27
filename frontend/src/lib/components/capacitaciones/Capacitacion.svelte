<script>
  import { onMount } from "svelte";
  import { token, users } from "$lib/stores/auth";
  import { getCapacitaciones, createCapacitacion } from "$lib/api";
  import { getEmbedUrl } from "$lib/helpers/utils";
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

<div class="container_capacitaciones">
  {#if role === "admin"}
    <div class="container_capacitaciones_form">
      <form on:submit|preventDefault={handleSubmit}>
        <input
          class="input_capacitaciones_video"
          type="text"
          name="url"
          bind:value={url}
          placeholder="Url Youtube"
          required
        />
        <textarea
          class="input_capacitaciones_description"
          type="text"
          name="description"
          bind:value={description}
          placeholder="Descripción"
        ></textarea>
        <button class="button_capacitaciones" type="submit">Guardar</button>
      </form>
    </div>
  {/if}
  <div class="container_capacitaciones_info">
    <h1 class="capacitaciones_titulo">CAPACITACIONES</h1>
    {#if cargando}
      <Cargando />
    {:else}
      {#each capacitacion as capacitacion}
        <div class="capacitaciones_banner">
          <div class="capacitaciones_video_wrapper">
            <iframe
              src={getEmbedUrl(capacitacion.Url)}
              allowfullscreen
              loading="lazy"
            ></iframe>
          </div>
          <p class="capacitaciones_description">{capacitacion.Description}</p>
        </div>
      {/each}
    {/if}
  </div>
</div>

<style>
  :root {
    --title-color_1: #b17d62;
    --title-color_2: #a0bea5;
  }
  .container_capacitaciones {
    height: 100vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }
  .container_capacitaciones_form form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin: 2rem;
  }
  .input_capacitaciones_video {
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    color: black;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin: 0 0.75rem;
  }

  .input_capacitaciones_description {
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    color: black;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin: 0 0.75rem;
    resize: none;
  }
  .button_capacitaciones {
    background-color: #c08060;
    color: white;
    border: none;
    border-radius: 12px;
    padding: 0.5rem 1rem;
    font-weight: bold;
    font-size: 1rem;
    width: 50%;
    margin: auto;
    cursor: pointer;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
    transition: background-color 0.3s ease;
  }
  .container_capacitaciones_info {
    border: 1px solid #ccc;
    flex: 1;
    height: 100%;
    overflow-y: auto;
    padding: 2rem 1rem;
    background-color: #fdf9f7;
    border-left: 1px solid #e5e5e5;
    max-height: 100vh;
    scroll-behavior: smooth;
    .capacitaciones_titulo {
      text-align: center;
      font-size: 2.5rem;
      margin-bottom: 2rem;
      font-weight: bold;
      color: var(--title-color_1);
      font-size: 1.8rem;
    }
  }

  .capacitaciones_banner {
    border-radius: 16px;
    box-shadow: 0px 2px 10px var(--title-color_2);
    margin-bottom: 2rem;
    padding: 1rem;
    max-width: 700px;
    width: 100%;
    margin-left: auto;
    margin-right: auto;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .capacitaciones_video_wrapper {
    position: relative;
    width: 100%;
    padding-top: 56.25%;
    border-radius: 12px;
    overflow: hidden;
  }

  .capacitaciones_video_wrapper iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border: none;
  }

  .capacitaciones_description {
    font-size: 1rem;
    color: #333;
    background-color: #fafafa;
    padding: 1rem;
    border-radius: 10px;
    max-height: 150px;
    overflow-y: auto;
    white-space: pre-wrap;
    word-wrap: break-word;
    border: 1px solid #eee;
  }
</style>
