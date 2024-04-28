<!-- <script>
  import { onMount } from "svelte";
  import { useParams } from "svelte-routing";

  let finalInterpretation = [];
  let isLoading = true;
  let error = null;

  const params = useParams();
  const { id } = params;
    
  // onMount(async () => {
  //   async function load() {
  //       // Fetch interpretation based on id
  //       const res = await fetch(`http://localhost:8082/cards/interpret/05ded632-c07d-42d5-b651-c098139b7921`);
  //       const data = await res.text();

  //       return{}
  //   }
  // });


  onMount(async () => {
    try {
      console.log([id])
      const res = await fetch(`http://localhost:8082/cards/interpret/${id}`);
      finalInterpretation = await res.text();
    } catch (err) {
      error = err.message;
      console.error("Error:", err);
    } finally {
      isLoading = false;
    }
  });
  console.log([id]) -->
<!-- </script>

{#if isLoading}
  <p>Loading...</p>
{:else if error}
  <p class="text-red-500">Error: {error}</p>
{:else}
  <h1>Interpretation</h1>
  <p>#{finalInterpretation}</p>
{/if} -->

<!-- <script context="module">
  export async function load({ params }) {
    const { id } = params;

    if (!id) {
      return {
        props: {
          finalInterpretation: "No ID provided.",
        },
      };
    }

    // Fetch interpretation based on id
    const res = await fetch(`http://localhost:8082/cards/interpret/${id}`);
    const data = await res.text();

    console.log(data);

    // Return the fetched interpretation
    return {
      props: {
        finalInterpretation: data,
      },
    };
  }
</script>

<script>
  export let finalInterpretation;
</script>

<p>{finalInterpretation}</p> -->
<script context="module">
  export async function load({ params }) {
    const { uuid } = params;

    console.log('uuid:', uuid);

    const url = `http://localhost:8082/cards/interpret/${uuid}`;
    console.log('fetch URL:', url);

    try {
      const res = await fetch(url);
      const data = await res.text();

      console.log('server response:', data);

      if (!res.ok) {
        console.error('Error in server response:', data);
        return {
          props: {
            finalInterpretation: '', // Return an empty finalInterpretation prop
          },
        };
      }

      console.log('fetched data:', data);

      return {
        props: {
          finalInterpretation: data,
        },
      };
    } catch (err) {
      console.error('Error in fetch request:', err);
      return {
        props: {
          finalInterpretation: '', // Return an empty finalInterpretation prop
        },
      };
    }
  }
</script>

<script>
  export let finalInterpretation;
</script>

{#if finalInterpretation}
  <p>{finalInterpretation}</p>
{:else}
  <p>Loading...</p>
{/if}