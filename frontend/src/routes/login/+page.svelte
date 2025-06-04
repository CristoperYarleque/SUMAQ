<script>
  import { loginUser } from "$lib/api";
  import { login } from "$lib/stores/auth";
  import Cargando from "$lib/components/cargando/Cargando.svelte";
  import { goto } from "$app/navigation";

  let cargando = false;
  let email = "",
    password = "",
    role = "entrepreneur";

  const handleLogin = async () => {
    try {
      cargando = true;
      const { data } = await loginUser({ email, password, role });
      login(data);
      cargando = false;
      goto(
        data.Role === "entrepreneur" || data.Role === "admin"
          ? "/emprendedor"
          : "/cliente"
      );
    } catch (err) {
      cargando = false;
      email = "";
      password = "";
      role = "entrepreneur";
      alert(err.message);
    }
  };

  const handleRegister = () => {
    goto("/register");
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
      <img src="/logo_inicio.png" alt="Logo Llama" class="llama-logo" />
      <h1 class="Bienvenida">Bienvenidos</h1>
      <form on:submit|preventDefault={handleLogin}>
        <input
          bind:value={email}
          type="email"
          placeholder="Correo Electrónico"
          required
          class="inputlogin"
        />
        <input
          bind:value={password}
          type="password"
          placeholder="Contraseña"
          required
          class="inputlogin"
        />

        <div class="role-toggle">
          <label class:active={role === "entrepreneur"}>
            <input type="radio" bind:group={role} value="entrepreneur" />
            <span class="iconsseleccionlogin">
              <img src="/emprendedor.png" alt="Emprende" class="emprende" /> Emprendedor</span
            >
          </label>
          <label class:active={role === "client"}>
            <input type="radio" bind:group={role} value="client" />
            <span class="iconsseleccionlogin">
              <img src="/cliente.png" alt="Cliente" class="cliente" /> Cliente</span
            >
          </label>
        </div>

        <button type="submit" class="login-button">Ingresar</button>
      </form>

      <p class="register-text">
        ¿No tienes cuenta? <span
          class="register-text-span"
          on:click={handleRegister}>Regístrate aquí</span
        >
      </p>
    {/if}
  </div>
</div>

<style>
  :root {
    --primary-color: #ebb2bd;
    --secondary-color: #ede9e4;
    --login-color: #a0bea5;
    --title-color_1: #b17d62;
    --text-color: #333;
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
    width: 400px;
    height: 400px;
    margin: auto;
  }

  .iconsseleccionlogin {
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

  .inputlogin {
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

  .register-text {
    margin-top: 1rem;
    font-size: 0.9rem;
    text-align: center;
  }

  .register-text-span {
    color: #d88c6d;
    text-decoration: none;
    font-weight: bold;
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
</style>
