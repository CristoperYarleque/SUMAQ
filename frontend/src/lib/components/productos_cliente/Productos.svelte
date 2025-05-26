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
  let cargandoCategorias = false;
  let cargandoEmprendedores = false;
  let cargandoProductos = false;
  let selectedCategoryId = null;

  async function handleCategorias() {
    try {
      cargandoCategorias = true;
      const { data } = await getCategories(tokenId);
      categorias = data;
      cargandoCategorias = false;
    } catch (error) {
      cargandoCategorias = false;
      console.error(error);
    }
  }

  async function handleEmprendedorByCategory(categoryId) {
    try {
      cargandoEmprendedores = true;
      selectedCategoryId = categoryId;
      const { data } = await getEmprendedorByCategory(tokenId, categoryId);
      productos = [];
      emprendedoresPorCategoria = data;
      cargandoEmprendedores = false;
    } catch (error) {
      cargandoEmprendedores = false;
      console.error(error);
    }
  }

  async function handleProducts(categoryId, emprendedorId) {
    try {
      cargandoProductos = true;
      const { data } = await getProducts(tokenId, emprendedorId, categoryId);
      productos = data;
      cargandoProductos = false;
    } catch (error) {
      cargandoProductos = false;
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
  <h1 class="productos_titulo">PRODUCTOS</h1>

  <div class="categorias">
    {#if cargandoCategorias}
      <Cargando />
    {:else}
      {#each categorias as categoria}
        <button class:selected={selectedCategoryId === categoria.CategoryId} 
                on:click={() => handleEmprendedorByCategory(categoria.CategoryId)}>
          {categoria.Name}
        </button>
      {/each}
    {/if}
  </div>

  <div class="scrollable_sections">
    <div class="emprendedores">
      {#if cargandoEmprendedores}
        <Cargando />
      {:else}
        <div class="emprendedores_grid">
          {#each emprendedoresPorCategoria as emprendedor}
            <div class="emprendedor_card" on:click={() => handleProducts(selectedCategoryId, emprendedor.Id)}>
              <img src={emprendedor.Url} alt={emprendedor.Name} class="emprendedor_img" />
              <h3>{emprendedor.Name}</h3>
              <p>{emprendedor.Email}</p>
            </div>
          {/each}
        </div>
      {/if}
    </div>
    

    <div class="productos">
      {#if cargandoProductos}
        <Cargando />
      {:else}
        {#each productos as producto}
          <div class="producto_card">
            <img src={producto.Url} alt={producto.Name} />
            <h4>{producto.Name}</h4>
            <p>{producto.Description}</p>
            <p class="precio">S/ {producto.Price}</p>
            <button on:click={() => addToCart(producto)}>Agregar al carrito</button>
          </div>
        {/each}
      {/if}
    </div>
  </div>
</div>

<style>
  :root {
    --primary-color: #EBB2BD;
    --secondary-color: #EDE9E4;
    --text-color: #333;
    --hover-color: #f29dae;
    --background-light: #f9f9f9;
    --title-color_1: #B17D62;
    --title-color_2: #A0BEA5;
  }
  .container_productos_layout {
    display: flex;
    flex-direction: column;
    padding: 2rem;
    background-color: var(--secondary-color);
    height: 100vh;
    box-sizing: border-box;
    .productos_titulo {
      text-align: center;
      font-size: 2.5rem;
      margin-bottom: 2rem;
      color: var(--title-color_1);
      font-weight: 700;
    }
  }

  .categorias {
    display: flex;
    justify-content: space-around;
    margin-bottom: 1rem;
  }

  .categorias button {
    padding: 0.5rem 1rem;
    background-color: #f8e1e1;
    border: 2px solid #ff85a2;
    border-radius: 20px;
    cursor: pointer;
    font-weight: bold;
    color: #333;
    transition: background-color 0.3s, transform 0.2s;
  }

  .categorias button.selected, .categorias button:hover {
    background-color: #ff85a2;
    color: white;
    transform: scale(1.05);
  }

  .scrollable_sections {
    display: flex;
    flex: 1;
    gap: 1rem;
    overflow: hidden;
  }

  .emprendedores, .productos {
    flex: 1;
    overflow-y: auto;
    background-color: #fff;
    border-radius: 12px;
    padding: 1rem;
    box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  }

  /* .emprendedor_card {
    padding: 0.5rem;
    margin-bottom: 0.5rem;
    background-color: #fce4ec;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.3s;
  }

  .emprendedor_card:hover {
    background-color: #f8bbd0;
  } */

.emprendedores_grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem;
}

.emprendedor_card {
  background-color: #fce4ec;
  border-radius: 12px;
  padding: 1rem;
  text-align: center;
  cursor: pointer;
  transition: background-color 0.3s;
  box-shadow: 0 4px 12px rgba(0,0,0,0.05);
}

.emprendedor_card:hover {
  background-color: #f8bbd0;
}

.emprendedor_img {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 0.5rem;
}

  .producto_card {
    padding: 1rem;
    border-radius: 12px;
    background-color: #f9f9f9;
    border: 1px solid var(--title-color_2);
    margin-bottom: 1rem;
    text-align: center;
    box-shadow: 0 4px 12px rgba(0,0,0,0.05);
  }

  .producto_card img {
    max-width: 100%;
    height: 150px;
    object-fit: cover;
    border-radius: 8px;
    margin-bottom: 0.5rem;
  }

  .producto_card h4 {
    margin: 0.5rem 0 0.3rem 0;
  }

  .producto_card .precio {
    font-weight: bold;
    color: #e91e63;
    margin-bottom: 0.5rem;
  }

  .producto_card button {
    background-color: #ff85a2;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    color: white;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s;
  }

  .producto_card button:hover {
    background-color: #f06292;
  }
</style>