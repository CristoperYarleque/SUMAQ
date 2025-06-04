<script>
  import { registerUser } from "$lib/api";
  import { goto } from "$app/navigation";
  import Cargando from "$lib/components/cargando/Cargando.svelte";
  import { LogOut } from "lucide-svelte";

  let cargando = false;
  let name = "",
    email = "",
    password = "",
    role = "entrepreneur";
    

  const handleRegister = async () => {
    try {
      cargando = true;
      const { code } = await registerUser({ name, email, password, role });
      cargando = false;
      code === 201 && goto("/login");
    } catch (err) {
      cargando = false;
      name = "";
      email = "";
      password = "";
      role = "entrepreneur";
      alert(err.message);
    }
  };

  const handleLogin = () => {
    goto("/login");
  };
</script>

<div class="container">
  <div class="illustration">
    <div class="weaver">
      <img src="/logo_hero.png" alt="Mujer tejiendo" />
    </div>
  </div>

  <div class="login-form">
    {#if cargando}
      <Cargando />
    {:else}
      <button class="button_logout" on:click={handleLogin}
        ><LogOut class="w-1 h-1 " /></button
      >
      <img src="/logo_inicio.png" alt="Logo Llama" class="llama-logo" />
      <h1 class="Bienvenida">Registro</h1>
      <form on:submit|preventDefault={handleRegister}>
        <input
          bind:value={name}
          type="text"
          placeholder="Nombre completo"
          required
          class="inputregister"
        />
        <input
          bind:value={email}
          type="email"
          placeholder="Correo Electrónico"
          required
          class="inputregister"
        />
        <input
          bind:value={password}
          type="password"
          placeholder="Contraseña"
          required
          class="inputregister"
        />

        <div class="role-toggle">
          <label class:active={role === "entrepreneur"}>
            <input type="radio" bind:group={role} value="entrepreneur" />
            <span class="iconsseleccionregister">
              <img src="/emprendedor.png" alt="Emprende" class="emprende" /> Emprendedor
            </span>
          </label>
          <label class:active={role === "client"}>
            <input type="radio" bind:group={role} value="client" />
            <span class="iconsseleccionregister">
              <img src="/cliente.png" alt="Cliente" class="cliente" /> Cliente
            </span>
          </label>
        </div>

        <button type="submit" class="login-button">Registrarse</button>
      </form>
    {/if}
  </div>
</div>

<style>
  :root {
    --primary-color: #ebb2bd;
    --secondary-color: #ede9e4;
    --login-color: #a0bea5;
    --text-color: #333;
    --title-color_1: #b17d62;
  }
  .container {
    display: flex;
    height: 100vh;
    font-family: "Georgia", serif;
  }

  .illustration {
    flex: 1;
    background-color: var(--primary-color);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .weaver img {
    width: 300px;
    max-width: 80%;
  }

  .llama-logo {
    width: 300px;
    height: 300px;
    margin: auto;
  }

  .iconsseleccionregister {
    display: flex;
    align-items: center;
  }

  .login-form {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 3rem;
    background-color: var(--secondary-color);
  }

  .login-form h1 {
    font-size: 2.5rem;
    margin-bottom: 1.5rem;
    color: var(--title-color_1);
  }

  .inputregister {
    width: 100%;
    padding: 1rem;
    margin-bottom: 1rem;
    border: 1px solid #ddd;
    border-radius: 10px;
    font-size: 1rem;
  }

  .role-toggle {
    display: flex;
    justify-content: space-between;
    margin-bottom: 1.5rem;
    background: #ffffff;
    border-radius: 10px;
    overflow: hidden;
  }

  .role-toggle label {
    flex: 1;
    text-align: center;
    padding: 1rem;
    cursor: pointer;
    font-weight: bold;
    color: #555;
    transition:
      background 0.3s,
      color 0.3s;
  }

  .role-toggle label.active {
    background: var(--primary-color);
    color: var(--text-color);
    /* color: white; */
  }

  .role-toggle input {
    display: none;
  }

  .login-button {
    width: 100%;
    padding: 1rem;
    font-size: 1.2rem;
    background-color: var(--login-color);
    color: white;
    border: none;
    border-radius: 10px;
    cursor: pointer;
  }

  .emprende {
    width: 35px;
    height: 40px;
  }

  .cliente {
    width: 35px;
    height: 40px;
  }

  .button_logout {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    cursor: pointer;
    border: none;
    background: none;
  }
</style>
