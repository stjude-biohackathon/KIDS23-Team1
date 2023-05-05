
<script lang="ts">
    export let title: string;
  </script>

  <div class ="manage-order-container">
    <h1>{title}</h1>
    <slot></slot>
  </div>

  <style>

    /* Your styles here */
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
  @media only screen and (max-width: 768px) {
    .manage-order-container {
      flex-direction: column;
    }
  }

  </style>
