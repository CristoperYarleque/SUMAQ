<script>
  import { logout, token, users } from "$lib/stores/auth";
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";

  $: {
    if ($token) {
      const home = users.role === "entrepreneur" || users.role === "admin" ? "/emprendedor" : "/cliente";
      goto(home);
    }
  }
</script>

<slot />
