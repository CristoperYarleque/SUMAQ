<script>
  import { onMount } from "svelte";
  import { Image, X } from "lucide-svelte";
  import {
    getProducts,
    getCategories,
    createProduct,
    uploadImageToImgBB,
  } from "$lib/api";
  import { token, users } from "$lib/stores/auth";
  import Tarjetas from "$lib/components/tarjetas/Tarjetas.svelte";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let tokenId = "";
  let entrepreneurId = 0;
  let categoryId = 0;

  let categorias = [];
  let products = [];
  let previewUrl = "";

  let file = "",
    name = "",
    description = "",
    price = "",
    url = "";

  let cargando = false;

  let fileLabelText = "Subir Imagen";

  async function handleCategorias(token) {
    try {
      cargando = true;
      const { data } = await getCategories(token);
      categorias = data;
      if (data.length > 0) {
        categoryId = data[0].CategoryId;
      }
      cargando = false;
    } catch (err) {
      cargando = false;
      console.log("err", err);
    }
  }

  async function handleProducts(tokenId) {
    try {
      cargando = true;
      const { data } = await getProducts(tokenId, entrepreneurId, 0);
      products = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.log("error", error);
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
      fileLabelText = "Subir Imagen";
    }
  }

  function clearPreview() {
    previewUrl = "";
    file = "";
    url = "";
    fileLabelText = "Subir Imagen";
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
          url = data.url;
        } else {
          alert("Error al subir la imagen");
          return;
        }
      }

      const { code } = await createProduct(tokenId, {
        name,
        description,
        price,
        categoryId,
        entrepreneurId,
        url,
      });

      if (code === 201) {
        name = "";
        description = "";
        price = "";
        if (categorias.length > 0) {
          categoryId = categorias[0].CategoryId;
        }
      }
      clearPreview();
      await handleProducts(tokenId);
    } catch (error) {
      console.log("error", error);
    }
  }

  async function handleDelete() {
    await handleProducts(tokenId);
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
    await handleCategorias(tokenId);
    await handleProducts(tokenId);
  });
</script>

<div class="container_productos_emprendedor">
  <div class="container_productos_forms">
    <form on:submit|preventDefault={handleSubmit}>
      <input
        class="input_productos"
        type="text"
        bind:value={name}
        name="name"
        placeholder="Nombre"
        required
      />
      <input
        class="input_productos"
        type="text"
        bind:value={description}
        name="description"
        placeholder="Descripción"
      />
      <input
        class="input_productos"
        type="number"
        step="0.01"
        bind:value={price}
        placeholder="Precio"
        required
      />
      <select
        class="select_productos"
        name="categoryId"
        id="categoryId"
        bind:value={categoryId}
        required
      >
        {#each categorias as categoria}
          <option value={categoria.CategoryId}>{categoria.Name}</option>
        {/each}
      </select>

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

      <button class="agregar_productos" type="submit">Agregar</button>
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
  <div class="container_productos">
    <h1 class="title_productos">PRODUCTOS</h1>
    <div class="container_productos_emprendedor_right">
      {#if cargando}
        <Cargando />
      {:else}
        {#each products as product}
          <Tarjetas {product} on:delete={handleDelete} />
        {/each}
      {/if}
    </div>
  </div>
</div>

<style>
  :root {
    --primary-color: #ebb2bd;
    --secondary-color: #ede9e4;
    --text-color: #333;
    --hover-color: #f29dae;
    --background-light: #f9f9f9;
    --title-color_1: #b17d62;
    --title-color_2: #a0bea5;
  }
  .input_hidden {
    display: none;
  }
  .custom_file_button {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.5rem 1rem;
    background-color: var(--background-light);
    color: rgba(0, 0, 0, 0.589);
    border-radius: 10px;
    cursor: pointer;
    font-weight: bold;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
    transition: background-color 0.3s ease;
    margin: 0 0.75rem;
    width: 200px;
  }

  .custom_file_button:hover {
    background-color: #ebb2bd;
  }
  .container_productos_emprendedor {
    height: 100vh;
    overflow: hidden;
    display: flex;
    flex-direction: column;
  }
  .title_productos {
    text-align: center;
    font-size: 2.5rem;
    margin-bottom: 2rem;
    font-weight: bold;
    color: var(--title-color_1);
    font-size: 1.8rem;
  }
  .container_productos {
    border: 1px solid #ccc;
    flex: 1;
    height: 100%;
    overflow-y: auto;
    padding: 1rem;
    background-color: white;
  }
  .container_productos_emprendedor_right {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    justify-content: center;
  }

  .container_productos_forms form {
    display: flex;
    flex-wrap: wrap;
    gap: 0.75rem;
    justify-content: center;
    align-items: center;
    margin: 2rem 0;
  }

  .input_productos,
  .select_productos {
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    color: black;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.2s ease;
  }

  .input_productos:focus {
    outline: none;
    box-shadow: 0 0 0 2px #c08060;
  }

  .container_productos_forms button {
    background-color: #c08060;
    color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1.2rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  .container_productos_forms button:hover {
    background-color: #a8664d;
  }

  .select_productos {
    background-color: white;
    border: none;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    color: black;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    margin: 0 0.75rem;
    width: 220px;
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

  .agregar_productos {
    background-color: #c08060;
    color: white;
    border: none;
    border-radius: 12px;
    padding: 0.5rem 1rem;
    font-weight: bold;
    font-size: 1rem;
    cursor: pointer;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
    transition: background-color 0.3s ease;
  }

  .agregar_productos:hover {
    background-color: #a8664d;
  }
</style>
