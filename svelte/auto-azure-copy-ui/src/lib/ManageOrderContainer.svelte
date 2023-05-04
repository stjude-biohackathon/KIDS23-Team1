
<script lang="ts">
    export let title: string;
  </script>

  <div class ="manage-order-container">
    <h1>{title}</h1>
    <slot></slot>
  </div>

  <style>
    div {
      /* Your styles here */
    }
  </style>
