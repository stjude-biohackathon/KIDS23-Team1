<script lang="ts">
  import OrderProcessingInfo from "./lib/OrderProcessingInfo.svelte";
  import PlaceOrder from "./lib/PlaceOrder.svelte";
</script>

<main>
  <div class="order-components-container">
    <PlaceOrder/>
    <OrderProcessingInfo/>
  </div>
</main>

<style>
  .order-components-container {
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    align-content: center;
  }
  @media only screen and (max-width: 768px) {
    .order-components-container {
      display: flex;;
      flex-direction: row;
    }
  }
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
