<script>
  import { getContext } from "svelte";
  
  const Context = Symbol('interpretation');
  const interpretation = getContext(Context);
</script>

<p>{interpretation}</p>