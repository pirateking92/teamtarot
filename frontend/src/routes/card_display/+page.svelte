<!-- 
    this script text is the js that parses the JSON data.
Tried to change routes and get a response, but didnt get very far.
Perhaps some changes to the routes? Tried getting responses from the backend but couldnt.
Code has been changed in routes package, some added to main
-->
<script>
  import { onMount } from "svelte";

  let threeCards = [];
  let isLoading = true;
  let error = null;

  onMount(async () => {
    try {
      const res = await fetch("/card_display");
      console.log("Response Status:", res.status); // test 1
      console.log("Response Headers:", res.headers); // test 2
      const data = await res.text(); //changed from res.json()
      console.log("Response Data:", data);
      threeCards = JSON.parse(data).cards;
    } catch (err) {
      error = err.message;
      console.error("Error:", err); // test 3
    } finally {
      isLoading = false;
    }
  });
</script>

<h1>Arcana</h1>
<!-- this subsequent code is what shows the data once it is properly parsed -->
{#if isLoading}
  <p>Loading...</p>
{:else if error}
  <p>Error: {error}</p>
{:else}
  <ul>
    {#each threeCards as card}
      <li>
        <h2>{card.cardName}</h2>
        <img src="/images/{card.imageName}" alt={card.cardName} />
        <p>Type: {card.type}</p>
        <p>Meaning Upright: {card.meaningUp}</p>
        <p>Meaning Reversed: {card.meaningReverse}</p>
        <p>{card.description}</p>
      </li>
    {/each}
  </ul>
{/if}

<h2>Lol just kiddin hun, tarot costs extra xxx</h2>
