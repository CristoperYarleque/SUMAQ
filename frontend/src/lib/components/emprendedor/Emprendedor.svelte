<script>
  import { onMount } from "svelte";
  import { Image, X } from "lucide-svelte";
  import { token, users } from "$lib/stores/auth";
  import {
    getEmprendedores,
    createEmprendedor,
    updateUrlEmprendedor,
    uploadImageToImgBB,
  } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";
  import { getEmbedUrl } from "$lib/helpers/utils";

  let emprendedores = [];
  let tokenId = "";
  let entrepreneurId = "";
  let type = "info";
  let cargando = false;
  let previewUrl = "";
  let file = "";

  let url = "",
    description = "",
    imageUrl = "";

  let fileLabelText = "Subir Fotografía";

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

  function handleFileChange(event) {
    file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = () => {
        previewUrl = reader.result;
      };
      reader.readAsDataURL(file);
      fileLabelText = file.name;
    } else {
      fileLabelText = "Subir Fotografía";
    }
  }

  function clearPreview() {
    previewUrl = "";
    file = "";
    imageUrl = "";
    fileLabelText = "Subir Fotografía";
    const fileInput = document.getElementById("fileInput");
    if (fileInput) {
      fileInput.value = "";
    }
  }

  async function handleSubmit() {
    try {
      if (file) {
        const { data, status } = await uploadImageToImgBB(file);
        if (status === 200) {
          imageUrl = data.url;
        } else {
          alert("Error al subir la imagen");
          return;
        }
      }

      const { code } = await createEmprendedor(tokenId, {
        url,
        description,
        entrepreneurId,
        type,
      });
      if (code === 201) {
        url = "";
        description = "";
        const { code } = await updateUrlEmprendedor(tokenId, entrepreneurId, {
          url: imageUrl,
        });
        if (code === 200) {
          await handleEmprendedores(tokenId, entrepreneurId, type);
          clearPreview();
        }
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
      <input
        class="input_emprendedor_video"
        type="text"
        name="url"
        bind:value={url}
        placeholder="Url Youtube"
        required
      />

      <textarea
        class="input_emprendedor_description"
        name="description"
        bind:value={description}
        placeholder="Descripción"
      ></textarea>

      <input
        id="fileInput"
        class="input_hidden"
        type="file"
        name="url"
        on:change={handleFileChange}
        accept="image/*"
      />

      <label for="fileInput" class="custom_file_button">
        <Image />
        {fileLabelText}
      </label>

      <button class="button_emprendedor" type="submit">Guardar</button>
    </form>
    {#if previewUrl}
      <div class="preview-container">
        <img class="preview-image" src={previewUrl} alt="Vista previa" />
        <span type="button" class="delete_preview" on:click={clearPreview}>
          <X />
        </span>
      </div>
    {/if}
  </div>

  <div class="container_emprendedor_info">
    <h1 class="emprendedor_titulo">EMPRENDEDOR</h1>
    {#if cargando}
      <Cargando />
    {:else}
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
  :root {
    --title-color_1: #b17d62;
    --title-color_2: #a0bea5;
  }
  .input_hidden {
    display: none;
  }
  .custom_file_button {
    display: flex;
    align-items: center;
    margin: auto;
    gap: 0.5rem;
    padding: 0.5rem;
    background-color: var(--background-light);
    color: rgba(0, 0, 0, 0.589);
    border-radius: 10px;
    cursor: pointer;
    font-weight: bold;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
    transition: background-color 0.3s ease;
    width: 49%;
  }

  .custom_file_button:hover {
    background-color: #ebb2bd;
  }

  .preview-image {
    border-radius: 10px;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin: auto;
    width: 200px;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .preview-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    max-width: 200px;
    position: relative;
    margin: auto;
    margin-bottom: 1rem;
  }

  .delete_preview {
    position: absolute;
    right: 0;
    border: none;
    border-radius: 8px;
    padding: 0.3rem 0.7rem;
    cursor: pointer;
    color: var(--text-color);
  }
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
    background-color: #fdf9f7;
    border-left: 1px solid #e5e5e5;
    max-height: 100vh;
    scroll-behavior: smooth;
    .emprendedor_titulo {
      text-align: center;
      font-size: 2.5rem;
      margin-bottom: 2rem;
      font-weight: bold;
      color: var(--title-color_1);
      font-size: 1.8rem;
    }
  }

  .container_emprendedor_form form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin: 2rem;
  }

  .input_emprendedor_video {
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    color: black;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin: 0 0.75rem;
  }

  .input_emprendedor_description {
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

  .button_emprendedor {
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

  .emprendedor_banner {
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

  .emprendedor_video_wrapper {
    position: relative;
    width: 100%;
    padding-top: 56.25%;
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
