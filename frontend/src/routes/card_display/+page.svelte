<!-- 
    this script text is the js that parses the JSON data.
Tried to change routes and get a response, but didnt get very far.
Perhaps some changes to the routes? Tried getting responses from the backend but couldnt.
Code has been changed in routes package, some added to main
-->
<link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">

<script>
  import { onMount } from "svelte";

  let threeCards = [];
  let isLoading = true;
  let error = null;

  onMount(async () => {
    try {
      const res = await fetch("http://localhost:8082/cards");
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
<div class="container mx-auto mt-8">
  <h1 class="text-4xl font-bold mb-4">Arcana Tarot Card Reading</h1>

  {#if isLoading}
    <p>Loading...</p>
  {:else if error}
    <p class="text-red-500">Error: {error}</p>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      {#each threeCards as card, index}
        <div class="bg-white shadow-md p-4 rounded-lg">
          <h2 class="text-xl font-semibold mb-2">
            {#if index === 0}
              <span>Past</span>
            {:else if index === threeCards.length - 1}
              <span>Future</span>
            {:else}
              <span>Present</span>
            {/if}
            <!-- {card.name} -->
          </h2>
          
          <h2 class="text-xl font-semibold mb-2">{card.name}</h2>
          <img src={`src/lib/assets/${card.image_file_name}`} alt="{card.name}" class="mb-2 rounded-lg">
          <p class="text-gray-600 mb-2">Type: {card.type}</p>
          <p class="text-gray-600 mb-2">Meaning Upright: {card.meaning_up}</p>
          <p class="text-gray-600 mb-2">Meaning Reversed: {card.meaning_rev}</p>
          <!-- <p class="text-gray-600">{card.desc}</p> -->
        </div>
      {/each}
    </div>
  {/if}

  <p class="mt-8 text-gray-500">Lol just kiddin hun, tarot costs extra xxx</p>
</div>