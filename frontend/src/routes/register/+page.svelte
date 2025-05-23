<script>
  import { registerUser } from "$lib/api";
  import { goto } from "$app/navigation";
  import { LogOut } from "lucide-svelte";

  let name = "",
    email = "",
    password = "",
    role = "entrepreneur";

  const handleRegister = async () => {
    try {
      const { code } = await registerUser({ name, email, password, role });
      code === 201 && goto("/login");
    } catch (err) {
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
      <img src="/logo2.png" alt="Mujer tejiendo" />
    </div>
  </div>

  <div class="login-form">
    <button class="button_logout" on:click={handleLogin}
      ><LogOut class="w-1 h-1 " /></button
    >
    <img src="/logo.png" alt="Logo Llama" class="llama-logo" />
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
        <label class:active={role === 'entrepreneur'}>
          <input type="radio" bind:group={role} value="entrepreneur" />
          <span class="iconsseleccionregister">
            <img src="/emprende.png" alt="Emprende" class="emprende" /> Emprendedor
          </span>
        </label>
        <label class:active={role === 'client'}>
          <input type="radio" bind:group={role} value="client" />
          <span class="iconsseleccionregister">
            <img src="/cliente.png" alt="Cliente" class="cliente" /> Cliente
          </span>
        </label>
      </div>

      <button type="submit" class="login-button">Registrarse</button>
    </form>
  </div>
</div>

<style>
  .container {
    display: flex;
    height: 100vh;
    font-family: 'Georgia', serif;
  }

  .illustration {
    flex: 1;
    background-color: #DF9A7D;
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
    margin: auto
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
  }

  .login-form h1 {
    font-size: 2.5rem;
    margin-bottom: 1.5rem;
    color: #3e4a3c;
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
    background: #f1e8e1;
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
    transition: background 0.3s, color 0.3s;
  }

  .role-toggle label.active {
    background: #DF9A7D;
    color: white;
  }

  .role-toggle input {
    display: none;
  }

  .login-button {
    width: 100%;
    padding: 1rem;
    font-size: 1.2rem;
    background-color: #83927e;
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


