<script lang="ts">
  import ManageOrderContainer from "./lib/ManageOrderContainer.svelte";
</script>

<main>
  <div class="order-components-container">
    <ManageOrderContainer title="Place Order">
      <!-- TODO - place order section -->
      <div>hey from place order inside app.svelte</div>
    </ManageOrderContainer>
    <ManageOrderContainer title ="Order Info">
      <!-- TODO order info section -->
      <div>hey from order info inside app.svelte</div>
    </ManageOrderContainer>
  </div>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
