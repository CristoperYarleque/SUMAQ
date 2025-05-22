<script>
  import { registerUser } from "$lib/api";
  import { goto } from "$app/navigation";

  let name = "",
    email = "",
    password = "",
    role = "entrepreneur",
    error = "";

  const handleRegister = async () => {
    try {
      const { code } = await registerUser({ name, email, password, role });
      code === 201 && goto("/login");
    } catch (err) {
      name = "";
      email = "";
      password = "";
      role = "entrepreneur";
      error = err.message;
    }
  };
</script>

<div>
  <h1>Registro</h1>
  {#if error}<p style="color:red">{error}</p>{/if}
  <form on:submit|preventDefault={handleRegister}>
    <input bind:value={name} placeholder="Nombre" required />
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
    <button type="submit">Registrarse</button>
  </form>
</div>

<style>
</style>
