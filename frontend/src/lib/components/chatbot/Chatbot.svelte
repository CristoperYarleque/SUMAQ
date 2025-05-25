<script>
  import { onMount } from "svelte";
  import { MessageCircleMore } from "lucide-svelte";
  import { users, token } from "$lib/stores/auth";
  import { getChatbot } from "$lib/api";
  import Cargando from "$lib/components/cargando/Cargando.svelte";

  let open = false;
  let messages = [];
  let name = "";
  let tokenId = "";
  let options = [];
  let cargando = false;

  function toggleChat() {
    open = !open;
    if (open && messages.length === 0) {
      messages.push({
        from: "bot",
        text: `¡Hola ${name || "emprendedor"}! ¿En qué puedo ayudarte?`,
      });
    }
  }

  function addMessages(newMessages) {
    messages = [...messages, ...newMessages];
  }

  function sendMessage(item) {
    addMessages([
      { from: "user", text: item.Question },
      { from: "bot", text: item.Answer },
    ]);
  }

  async function handleChatbot(tokenId) {
    try {
      cargando = true;
      const { data } = await getChatbot(tokenId);
      options = data;
      cargando = false;
    } catch (err) {
      console.error("Error al cargar chatbot:", err);
      cargando = false;
    }
  }

  onMount(() => {
    token.subscribe((token) => {
      if (token) tokenId = token;
    });
    users.subscribe((user) => {
      if (user) name = user.name;
    });

    handleChatbot(tokenId);
  });
</script>

<div class="chat-icon" on:click={toggleChat}><MessageCircleMore /></div>

{#if open}
  <div class="chat-window">
    {#if cargando}
      <Cargando />
    {:else}
      {#each messages as msg}
        <div class="message {msg.from}">{msg.text}</div>
      {/each}
    {/if}
    <div class="options">
      {#if cargando}
        <Cargando />
      {:else}
        {#each options as item}
          <button on:click={() => sendMessage(item)}>{item.Question}</button>
        {/each}
      {/if}
    </div>
  </div>
{/if}

<style>
  .chat-icon {
    position: fixed;
    bottom: 30px;
    right: 30px;
    background-color: #000;
    color: white;
    padding: 15px;
    border-radius: 50%;
    cursor: pointer;
    z-index: 10000;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  }

  .chat-window {
    position: fixed;
    bottom: 90px;
    right: 30px;
    width: 300px;
    max-height: 400px;
    background-color: white;
    border: 1px solid #ccc;
    border-radius: 15px;
    overflow: auto;
    padding: 10px;
    z-index: 10000;
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .message.bot {
    background: #f0f0f0;
    border-radius: 10px;
    padding: 6px 10px;
    align-self: flex-start;
  }

  .message.user {
    background: #d1e7dd;
    border-radius: 10px;
    padding: 6px 10px;
    align-self: flex-end;
  }

  .options button {
    margin-top: 5px;
    font-size: 0.9rem;
    background: #000;
    color: white;
    padding: 5px 10px;
    border: none;
    border-radius: 8px;
    cursor: pointer;
  }
</style>
