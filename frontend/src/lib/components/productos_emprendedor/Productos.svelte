<script>
  import { onMount} from "svelte";
  import { getProducts, getCategories, createProduct } from "$lib/api";
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
    price = "";

  let cargando = false;
  

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
    }
  }

  async function handleSubmit() {
    try {
      const { code } = await createProduct(tokenId, {
        name,
        description,
        price,
        categoryId,
        entrepreneurId,
      });

      if (code === 201) {
        name = "";
        description = "";
        price = "";
        file = "";
        previewUrl = "";
        if (categorias.length > 0) {
          categoryId = categorias[0].CategoryId;
        }
      }
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
  <div>
    <form on:submit|preventDefault={handleSubmit}>
      <input
        type="text"
        bind:value={name}
        name="name"
        placeholder="Nombre"
        required
      />
      <input
        type="text"
        bind:value={description}
        name="description"
        placeholder="Descripción"
      />
      <input
        type="number"
        step="0.01"
        bind:value={price}
        placeholder="Precio"
        required
      />
      <input
        type="file"
        name="url"
        on:change={handleFileChange}
        accept="image/*"
      />
      <select name="categoryId" id="categoryId" bind:value={categoryId} required>
        {#each categorias as categoria}
          <option value={categoria.CategoryId}>{categoria.Name}</option>
        {/each}
      </select>
      <button type="submit">Agregar</button>
    </form>
    {#if previewUrl}
      <img class="preview-image" src={previewUrl} alt="Vista previa" />
    {/if}
  </div>
  <div class="container_productos">
    <h1>PRODUCTOS</h1>
    {#each products as product}
      <Tarjetas {product} on:delete={handleDelete} />
    {/each}
  </div>
</div>


<style>
  .container_productos_emprendedor {
    height: 100vh;
    overflow: hidden;
  }
  .container_productos {
    overflow-y: auto;
    height: 100%;
  }
  .preview-image {
    max-width: 200px;
    margin-top: 10px;
  }
</style>
