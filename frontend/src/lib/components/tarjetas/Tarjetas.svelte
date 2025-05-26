<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { deleteProduct } from "$lib/api";
  import { token } from "$lib/stores/auth";

  let tokenId = "";
  const dispatch = createEventDispatcher();
  export let product = {};

  async function handleDelete(productId) {
    try {
      await deleteProduct(tokenId, productId);
      dispatch("delete");
    } catch (error) {
      console.error(error);
    }
  }

  onMount(() => {
    token.subscribe(async (token) => {
      if (token) {
        tokenId = token;
      }
    });
  });
</script>

<!-- <div>
  <p>{product.Name}</p>
  <p>{product.Price}</p>
  <img src={product.Url} alt={product.Name} />
  <p>{product.Description}</p>
  <p>{product.CategoryName}</p>
  <button on:click={() => handleDelete(product.ProductId)}>Eliminar</button>
</div> -->

<div class="producto-card">
  <p class="producto-nombre">{product.Name}</p>
  <p class="producto-precio">${product.Price}</p>
  <img class="producto-imagen" src={product.Url} alt={product.Name} />
  <p class="producto-descripcion">{product.Description}</p>
  <p class="producto-categoria">{product.CategoryName}</p>
  <button class="producto-boton" on:click={() => handleDelete(product.ProductId)}>Eliminar</button>
</div>


<style>
  .producto-card {
  border: 2px solid #b17d62;
  border-radius: 16px;
  padding: 16px;
  width: 260px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: transform 0.3s ease;
}

.producto-card:hover {
  transform: scale(1.02);
}

.producto-imagen {
  width: 100%;
  height: 180px;
  object-fit: cover;
  border-radius: 12px;
  margin-bottom: 12px;
}

.producto-nombre {
  font-size: 1.2rem;
  font-weight: bold;
  color: #b17d62;
  margin-bottom: 6px;
  text-align: center;
}

.producto-precio {
  font-size: 1rem;
  color: #a0be95;
  font-weight: bold;
  margin-bottom: 8px;
}

.producto-descripcion {
  font-size: 0.9rem;
  color: #6d4f3f;
  text-align: center;
  margin-bottom: 8px;
}

.producto-categoria {
  font-size: 0.85rem;
  color: #b17d62;
  font-style: italic;
  margin-bottom: 12px;
}

.producto-boton {
  background-color: #ebb2bd;
  color: white;
  padding: 10px 16px;
  border: none;
  border-radius: 12px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.producto-boton:hover {
  background-color: #d59da9;
}

</style>


