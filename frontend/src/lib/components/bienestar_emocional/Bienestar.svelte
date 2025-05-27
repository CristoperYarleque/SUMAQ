<script>
  import { onMount } from "svelte";
  import { token, users } from "$lib/stores/auth";
  import { getBienestar, createBienestar } from "$lib/api";
  import { getEmbedUrl } from "$lib/helpers/utils";
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
    <div class="container_bienestar_emocional_form">
      <form on:submit|preventDefault={handleSubmit}>
        <input
          class="input_bienestar_emocional_video"
          type="text"
          name="url"
          bind:value={url}
          placeholder="Url Youtube"
          required
        />
        <textarea
          class="input_bienestar_emocional_description"
          type="text"
          name="description"
          bind:value={description}
          placeholder="Descripción"
        ></textarea>
        <button class="button_bienestar_emocional" type="submit">Guardar</button
        >
      </form>
    </div>
  {/if}
  <div class="container_bienestar_emocional_info">
    <h1 class="bienestar_emocional_titulo">BIENESTAR EMOCIONAL</h1>
    {#if cargando}
      <Cargando />
    {:else}
      {#each bienestar as bienestar}
        <div class="bienestar_emocional_banner">
          <div class="bienestar_emocional_video_wrapper">
            <iframe
              src={getEmbedUrl(bienestar.Url)}
              allowfullscreen
              loading="lazy"
            ></iframe>
          </div>
          <p class="bienestar_emocional_description">{bienestar.Description}</p>
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
  .container_bienestar_emocional {
    height: 100vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }
  .container_bienestar_emocional_form form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin: 2rem;
  }
  .input_bienestar_emocional_video {
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    color: black;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin: 0 0.75rem;
  }

  .input_bienestar_emocional_description {
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
  .button_bienestar_emocional {
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
  .container_bienestar_emocional_info {
    border: 1px solid #ccc;
    flex: 1;
    height: 100%;
    overflow-y: auto;
    padding: 2rem 1rem;
    background-color: #fdf9f7;
    border-left: 1px solid #e5e5e5;
    max-height: 100vh;
    scroll-behavior: smooth;
    .bienestar_emocional_titulo {
      text-align: center;
      font-size: 2.5rem;
      margin-bottom: 2rem;
      font-weight: bold;
      color: var(--title-color_1);
      font-size: 1.8rem;
    }
  }

  .bienestar_emocional_banner {
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

  .bienestar_emocional_video_wrapper {
    position: relative;
    width: 100%;
    padding-top: 56.25%; /* 16:9 ratio */
    border-radius: 12px;
    overflow: hidden;
  }

  .bienestar_emocional_video_wrapper iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border: none;
  }

  .bienestar_emocional_description {
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
