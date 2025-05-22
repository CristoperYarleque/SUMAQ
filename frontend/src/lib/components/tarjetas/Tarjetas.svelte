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

<div>
  <p>{product.Name}</p>
  <p>{product.Price}</p>
  <img src={product.Url} alt={product.Name} />
  <p>{product.Description}</p>
  <p>{product.CategoryName}</p>
  <button on:click={() => handleDelete(product.ProductId)}>Eliminar</button>
</div>

<style>
</style>
