<script>
  import { loginUser } from "$lib/api";
  import { login } from "$lib/stores/auth";
  import { goto } from "$app/navigation";

  let email = "",
    password = "",
    role = "entrepreneur",
    error = "";

  const handleLogin = async () => {
    try {
      const { data } = await loginUser({ email, password, role });
      login(data);
      goto(data.Role === "entrepreneur" ? "/emprendedor" : "/cliente");
    } catch (err) {
      email = "";
      password = "";
      role = "entrepreneur";
      error = err.message;
    }
  };
</script>

<div>
  <h1>Iniciar Sesión</h1>
  {#if error}<p style="color:red">{error}</p>{/if}
  <form on:submit|preventDefault={handleLogin}>
    <input bind:value={email} type="email" placeholder="Email" required />
    <input
      bind:value={password}
      type="password"
      placeholder="Contraseña"
      required
    />
    <label
      ><input type="radio" bind:group={role} value="entrepreneur" checked /> Emprendedor
    </label>
    <label
      ><input type="radio" bind:group={role} value="client" /> Cliente
    </label>
    <button type="submit">Entrar</button>
  </form>
</div>

<style>
  h1
</style>
