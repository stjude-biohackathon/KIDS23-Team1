<script lang="ts">
    import ManageOrderContainer from "./ManageOrderContainer.svelte";
    //TODO - unstub and replace with actual handleSubmit
    const handleSubmit:
    //EventHandler<Event & { readonly submitter: HTMLElement; }, HTMLFormElement>
        any
        = (event, form) => {
        event.preventDefault();
        console.log("form submitted")
    }
  // handle form submission
</script>

<!-- TODO - unstub and replace with actual component -->
<ManageOrderContainer title="Place Order">
  <form on:submit={handleSubmit}>
    <label for="sample-select">Select Sample(s):</label>
    <select id="sample-select" name="sample-select">
      <option value="">--Please choose an option--</option>
      <option value="sample1">Sample 1</option>
      <option value="sample2">Sample 2</option>
      <option value="sample3">Sample 3</option>
    </select>
    <br><br>
    <label for="sample-select">Select Sample File(s):</label>
    <select id="sample-select" name="sample-file-select">
      <option value="">--Please choose an option--</option>
      <option value="sample1">Sample File 1</option>
      <option value="sample2">Sample File 2</option>
      <option value="sample3">Sample File 3</option>
    </select>
    <br><br>
    <label for="analysis-select">Select Analysis:</label>
    <select id="analysis-select" name="analysis-select">
      <option value="">--Please choose an option--</option>
      <option value="analysis1">Analysis 1</option>
      <option value="analysis2">Analysis 2</option>
      <option value="analysis3">Analysis 3</option>
    </select>
    <br><br>
    <button type="submit">Submit</button>
  </form>
</ManageOrderContainer>