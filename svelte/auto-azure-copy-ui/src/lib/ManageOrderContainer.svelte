<script lang="ts">
    export let title: string;
</script>

<div class ="manage-order-container">
  <div class ="inner-container">
      <h1>{title}</h1>
  <slot></slot>
  </div>

</div>

  <style>
  .manage-order-container {
    height: 50vh;
    width: 50vw;
    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    align-content: center;
  }
  </style>
