<script>
  import { onMount } from "svelte";
  import { token } from "$lib/stores/auth";
  import { getCategories, getEmprendedorByCategory, getProducts  } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";
  import { addToCart } from "$lib/stores/cart";
 

  let tokenId = "";
  let categorias = [];
  let emprendedoresPorCategoria = [];
  let productos = [];
  let cargando = false;
  let selectedCategoryId = null;

  async function handleCategorias() {
    try {
      cargando = true;
      const { data } = await getCategories(tokenId);
      categorias = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  async function handleEmprendedorByCategory(categoryId) {
    try {
      cargando = true;
      selectedCategoryId = categoryId;
      const { data } = await getEmprendedorByCategory(tokenId, categoryId);
      productos = [];
      emprendedoresPorCategoria = data;
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  async function handleProducts(categoryId, emprendedorId) {
    try {
      cargando = true;
      const { data } = await getProducts(tokenId, emprendedorId, categoryId);
      productos = data;
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
    await handleCategorias();
  });
</script>

<div class="container_productos_layout">
  <h1>PRODUCTOS</h1>
  <div class="container_categorias">
    {#if cargando}
      <Cargando />
    {:else}
      {#each categorias as categoria}
        <span on:click={() => handleEmprendedorByCategory(categoria.CategoryId)}
          >{categoria.Name}</span
        >
      {/each}
    {/if}
  </div>
  <div class="container_scrollable_sections">
    <div class="container_emprendedores">
        {#if cargando}
        <Cargando />
        {:else}
        {#each emprendedoresPorCategoria as emprendedor}
            <div on:click={() => handleProducts(selectedCategoryId, emprendedor.Id)} style="cursor: pointer;">
            <p>{emprendedor.Id}</p>
            <p>{emprendedor.Name}</p>
            <p>{emprendedor.Email}</p>
            </div>
        {/each}
        {/if}
    </div>
    <div class="container_productos">
        {#if cargando}
        <Cargando />
        {:else}
        {#each productos as producto}
            <div>
            <p>{producto.Name}</p>
            <p>{producto.Description}</p>
            <p>{producto.Price}</p>
            <p>{producto.Url}</p>
            <button on:click={() => addToCart(producto)}>Agregar al carrito</button>
            </div>
        {/each}
        {/if}
    </div>
  </div>
</div>

<style>
  .container_productos_layout {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: hidden;
  }
  .container_productos_layout h1 {
    text-align: center;
  }
  .container_categorias {
    display: flex;
    flex-direction: row;
    gap: 10px;
    justify-content: space-around;
    align-items: center;
    padding: 10px;
    flex: 0 0 auto;
  }
  .container_scrollable_sections {
    display: flex;
    flex-direction: row;
    flex: 1 1 auto;
    overflow: hidden;
  }
  .container_emprendedores,
  .container_productos {
    flex: 1 1 auto;
    overflow-y: auto;
    padding: 10px;
    border: 1px solid #ccc;
  }
  .container_categorias span {
    text-align: center;
    cursor: pointer;
    font-size: 1.1rem;
    transition: font-size 0.3s ease;
    padding: 10px;
    &:hover {
      background-color: #000000;
      color: #ffffff;
      font-size: 1.2rem;
      border-radius: 10px;
    }
  }
</style>
