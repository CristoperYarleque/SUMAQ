<script>
  import { onMount } from "svelte";
  import { token, users } from "$lib/stores/auth";
  import { getEmprendedores, createEmprendedor } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";
  import { getEmbedUrl } from "$lib/helpers/utils";

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
  <div class="container_emprendedor_form">
    <form on:submit|preventDefault={handleSubmit}>
      <input class="input_emprendedor_video"
        type="text"
        name="url"
        bind:value={url}
        placeholder="Url"
        required
        />
      <!-- <input
        type="text"
        name="description"
        bind:value={description}
        placeholder="Descripción"
      /> -->
      <textarea class="input_emprendedor_description"
      name="description"
      bind:value={description}
      placeholder="Descripción"
    ></textarea>

      <button class="button_emprendedor" type="submit">Guardar</button>
    </form>
  </div>

  <div class="container_emprendedor_info">
    <h1>EMPRENDEDOR</h1>
    {#if cargando}
      <Cargando />
    {:else}
      <!-- {#each emprendedores as emprendedor}
        <p>{emprendedor.Url}</p>
        <p>{emprendedor.Description}</p>
      {/each} -->
      {#each emprendedores as emprendedor}
      <div class="emprendedor_banner">
        <div class="emprendedor_video_wrapper">
          <iframe
            src={getEmbedUrl(emprendedor.Url)}
            allowfullscreen
            loading="lazy"
          ></iframe>
        </div>
        <p class="emprendedor_description">{emprendedor.Description}</p>
      </div>
    {/each}    
    

    {/if}
  </div>
</div>

<style>

  .container_emprendedor {
    height: 100vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }
  .container_emprendedor_info {
    border: 1px solid #ccc;
    flex: 1;
    height: 100%;
    overflow-y: auto;
    padding: 2rem 1rem;
  }

  .input_emprendedor {
  background-color: white;
  border: none;
  border-radius: 10px;
  padding: 0.5rem 1rem;
  font-size: 1rem;
  color: black;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  width: 100%;
  /* max-width: 500px; */
  height: 100px;
  resize: none;
  overflow-y: auto;
}

form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  /* max-width: 500px; */
  margin: 2rem;
}

.input_emprendedor_video{
  background-color: white;
  border: none;
  border-radius: 10px;
  padding: 0.5rem 1rem;
  font-size: 1rem;
  color: black;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  margin: 0 0.75rem;
  
}

.input_emprendedor_description{
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

.button_emprendedor{
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



.container_emprendedor_info {
  background-color: #fdf9f7;
  border-left: 1px solid #e5e5e5;
  padding: 2rem 1rem;
  overflow-y: auto;
  max-height: 100vh;
  scroll-behavior: smooth;
}

.container_emprendedor_info h1 {
  font-size: 1.8rem;
  font-weight: bold;
  color: #b17d62;
  margin-bottom: 2rem;
  text-align: center;
}

.emprendedor_banner {
  border-radius: 16px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.08);
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

.emprendedor_video_wrapper {
  position: relative;
  width: 100%;
  padding-top: 56.25%; /* 16:9 ratio */
  border-radius: 12px;
  overflow: hidden;
}

.emprendedor_video_wrapper iframe {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border: none;
}

.emprendedor_description {
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
