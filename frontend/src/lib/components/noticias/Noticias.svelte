<script>
  import { onMount } from "svelte";
  import { Timer, ThumbsUp, Heart } from "lucide-svelte";
  import { token, users } from "$lib/stores/auth";
  import { getNoticias, getReactionInfo, upsertReaction, deleteReaction } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let tokenId = "";
  let userId = "";
  let noticias = [];
  let cargando = false;
  let startDate = "";
  let endDate = "";

  async function handleNoticias(tokenId) {
    try {
      cargando = true;
      const { data } = await getNoticias(tokenId, startDate, endDate);
      noticias = data;
      await loadReactions();
      cargando = false;
    } catch (error) {
      cargando = false;
      console.error(error);
    }
  }

  async function handleReaction(noticia, reaction) {
    try {
      const prevReaction = noticia.userReaction;

      if (prevReaction === reaction) {
        await deleteReaction(tokenId, userId, noticia.NewsId);

        if (reaction === 'like') {
          noticia.likeCount = Math.max(0, (noticia.likeCount || 1) - 1);
        } else if (reaction === 'love') {
          noticia.loveCount = Math.max(0, (noticia.loveCount || 1) - 1);
        }

        noticia.userReaction = null;
        noticias = [...noticias];
        return;
      };

      await upsertReaction(tokenId, { userId, newsId: noticia.NewsId, reaction });

      // Ajustar conteos localmente
      if (reaction === 'like') {
        noticia.likeCount = (noticia.likeCount || 0) + 1;
        if (prevReaction === 'love') noticia.loveCount = (noticia.loveCount || 1) - 1;
      } else if (reaction === 'love') {
        noticia.loveCount = (noticia.loveCount || 0) + 1;
        if (prevReaction === 'like') noticia.likeCount = (noticia.likeCount || 1) - 1;
      }

      // Actualizar la reacción del usuario
      noticia.userReaction = reaction;

      noticias = [...noticias];
    } catch (error) {
      console.error(error);
    }
  }



  async function loadReactions() {
    for (const noticia of noticias) {
      if (userId && noticia.NewsId) {
        try {
          const { data } = await getReactionInfo(tokenId, userId, noticia.NewsId);
          noticia.userReaction = data.UserReaction;
          noticia.likeCount = data.LikeCount;
          noticia.loveCount = data.LoveCount;
        } catch (e) {
          console.error(e);
        }
      }
    }

    console.log("Reactions loaded:", noticias);
    
  }

  onMount(async () => {
    token.subscribe(async (token) => {
      if (token) {
        tokenId = token;
      }
    });

    users.subscribe(async (user) => {
      if (user) {
        userId = user.userId;
      }
    });

    const sixMonthsAgo = new Date();
    sixMonthsAgo.setMonth(sixMonthsAgo.getMonth() - 6);
    const sixMonthsLater = new Date();
    sixMonthsLater.setMonth(sixMonthsLater.getMonth() + 6);

    startDate = sixMonthsAgo.toISOString().slice(0, 10);
    endDate = sixMonthsLater.toISOString().slice(0, 10);

    await handleNoticias(tokenId);
  });
</script>

<div class="container_noticias">
  <h1 class="noticias_titulo">NOTICIAS</h1>
  {#if cargando}
    <Cargando />
  {:else}
    <div class="noticias_grid">
      {#each noticias as noticia}
        <div class="noticia_card">
          <p class="noticia_fecha">
            <Timer />
            {new Date(noticia.PublishedAt).toLocaleString()}
          </p>
          <h2 class="noticia_titulo">{noticia.Title}</h2>
          <p class="noticia_contenido">{noticia.Content}</p>
          <div>
            <img class="noticia_img" src={noticia.Url} alt={noticia.Title} />
          </div>
          <div class="reactions">
            <button
              class="reaction_button"
              class:active={noticia.userReaction === 'like'}
              on:click={() => handleReaction(noticia, 'like')}
            >
              <ThumbsUp /> {noticia.likeCount || 0}
            </button>

            <button
              class="reaction_button"
              class:active={noticia.userReaction === 'love'}
              on:click={() => handleReaction(noticia, 'love')}
            >
              <Heart /> {noticia.loveCount || 0}
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  :root {
    --title-color_1: #b17d62;
  }
  .container_noticias {
    padding: 0 2rem;
    min-height: 100vh;
    overflow-y: auto;
    .noticias_titulo {
      text-align: center;
      font-size: 2.5rem;
      margin-bottom: 2rem;
      font-weight: 700;
      color: var(--title-color_1);
    }
  }

  .noticias_grid {
    /* display: flex; */
    /* flex-direction: column; */
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(500px, 1fr));
    gap: 1.5rem;
  }

  .noticia_card {
    /* height: 300px; */
    background-color: white;
    border-radius: 20px;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.08);
    padding: 1.5rem;
    transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .noticia_card:hover {
    background-color: #f8bbd053;
    transform: translateY(-6px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.12);
  }

  .noticia_fecha {
    font-size: 0.85rem;
    color: #888;
    text-align: right;
    font-style: italic;
    display: flex;
    gap: 0.5rem;
    justify-content: flex-end;
    align-items: center;
  }

  .noticia_titulo {
    font-size: 1.3rem;
    font-weight: 600;
    color: var(--title-color_1);
    margin: 0;
    line-height: 1.4;
  }

  .noticia_contenido {
    font-size: 1rem;
    color: #555;
    line-height: 1.5;
  }

  .noticia_img {
    width: 100%;
    height: 200px;
    border-radius: 10px;
    margin-top: 1rem;
    object-fit: cover;
  }

  .reactions {
    display: flex;
    justify-content: flex-end;
    margin-top: 1rem;
  }

  .reaction_button {
    background: transparent;
    border: none;
    cursor: pointer;
    font-size: 1.2rem;
    margin-right: 0.5rem;
  }
  .active {
    color: var(--title-color_1);
    font-weight: bold;
  }

  @media (max-width: 600px) {
    .container_noticias {
      padding: 1rem;
    }

    .noticia_card {
      padding: 1rem;
    }

    .noticia_titulo {
      font-size: 1.1rem;
    }
  }
</style>
