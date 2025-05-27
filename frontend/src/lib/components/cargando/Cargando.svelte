<script>
  export let size = "80px";
</script>

<div class="loading-container">
  <div
    class="spinner-border"
    style="width: {size}; height: {size};"
    role="status"
  ></div>
</div>

<style>
  .loading-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }

  .spinner-border {
    border: 8px solid #f3f3f3;
    border-top: 8px solid #eb8ab2;
    border-radius: 50%;
    width: 100px;
    height: 100px;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
