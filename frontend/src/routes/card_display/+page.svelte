<script>
  import { onMount } from "svelte";

  function handleButtonClick() {
    window.location.href = "/";
  }

  async function getFate() {
    try {
      const res = await fetch("http://localhost:8082/interpret-");
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
    window.location.href = "/cassandra";
  }

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

<!-- 
    this script text is the js that parses the JSON data.
Tried to change routes and get a response, but didnt get very far.
Perhaps some changes to the routes? Tried getting responses from the backend but couldnt.
Code has been changed in routes package, some added to main
-->
<link
  href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css"
  rel="stylesheet"
/>
<div class="container mx-auto mt-8">
  <div class="flex justify-between items-center mb-8">
    <h3 class="text-2xl font-bold card-display-title">
      Welcome to Cassandra's Parlor
    </h3>
    <button on:click={handleButtonClick} style="--clr:#c377d4"
      ><span>Leave the Parlor</span><i></i></button
    >
  </div>

  {#if isLoading}
    <p>Loading...</p>
  {:else if error}
    <p class="text-red-500">Error: {error}</p>
  {:else}
    <div class="mx-auto px-20">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 ml-12 mr-12">
        {#each threeCards as card, index}
          <div class="bg-card-container shadow-md p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-2">
              {#if index === 0}
                <span>Past</span>
              {:else if index === threeCards.length - 1}
                <span>Future</span>
              {:else}
                <span>Present</span>
              {/if}
            </h2>

            <h2 class="text-xl font-semibold mb-2">{card.name}</h2>
            {#if card.reversed}
              <img
                src={`src/lib/assets/${card.image_file_name}`}
                alt={card.name}
                class="mb-2 rounded-lg card-image"
                style="transform: rotate(180deg);"
              />
            {:else}
              <img
                src={`src/lib/assets/${card.image_file_name}`}
                alt={card.name}
                class="mb-2 rounded-lg card-image"
              />
            {/if}
            <p class="text-gray-600 mb-2">Type: {card.type}</p>
            {#if card.reversed}
              <p class="text-gray-600 mb-2">
                Meaning Reversed: {card.meaning_rev}
              </p>
            {:else}
              <p class="text-gray-600 mb-2">
                Meaning Upright: {card.meaning_up}
              </p>
            {/if}
          </div>
        {/each}
      </div>
    </div>
  {/if}
  <button on:click={getFate} class="mx-auto mt-8 block" style="--clr:#c377d4"
  ><span>Get Fate</span><i></i></button
>
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");

  :global(html, body) {
    height: 100%;
    margin: 0;
    padding: 0;
    background: linear-gradient(to right, #20051c, #46204e);
  }

  .card-display-title {
    color: whitesmoke;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
  }

  .card-image {
    max-width: 200px; /* Set the desired maximum width */
    max-height: 300px; /* Set the desired maximum height */
    display: block;
    margin: 1em auto;
  }

  .bg-card-container {
    background: linear-gradient(to right, #b9b7ce, #4f86c8);
  }

  button {
    position: relative;
    background: #1d1242;
    color: #fff;
    text-decoration: none;
    text-transform: uppercase;
    border: none;
    letter-spacing: 0.1rem;
    font-size: 0.8rem;
    padding: 0.5rem 2rem;
    transition: 0.2s;
    animation: box 2s infinite;
    border-radius: 20px;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
  }

  button:hover {
    letter-spacing: 0.1rem;
    padding: 0.6rem 2rem;
    background: var(--clr);
    color: var(--clr);
    /* box-shadow: 0 0 35px var(--clr); */
    animation: box 1.5s infinite;
  }

  button::before {
    content: "";
    position: absolute;
    inset: 2px;
    background: linear-gradient(to right, #20051c, #46204e);
    border-radius: 20px;
  }

  button span {
    position: relative;
    z-index: 1;
  }
  button i {
    position: absolute;
    inset: 0;
    display: block;
  }

  button i::before {
    content: "";
    position: absolute;
    width: 10px;
    height: 2px;
    left: 80%;
    top: -2px;
    border: 2px solid var(--clr);
    background: #272822;
    transition: 0.2s;
  }

  button:hover i::before {
    width: 15px;
    left: 20%;
    animation: move 3s infinite;
  }

  button i::after {
    content: "";
    position: absolute;
    width: 10px;
    height: 2px;
    left: 20%;
    bottom: -2px;
    border: 2px solid var(--clr);
    background: #272822;
    transition: 0.2s;
  }

  button:hover i::after {
    width: 15px;
    left: 80%;
    animation: move 3s infinite;
  }

  @keyframes move {
    0% {
      transform: translateX(0);
    }
    50% {
      transform: translateX(5px);
    }
    100% {
      transform: translateX(0);
    }
  }

  @keyframes box {
    0% {
      box-shadow: #27272c;
    }
    50% {
      box-shadow: 0 0 25px var(--clr);
    }
    100% {
      box-shadow: #27272c;
    }
  }
</style>
