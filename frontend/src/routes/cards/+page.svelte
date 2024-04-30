<script>
  import { onMount } from "svelte";
  import Typewriter from "svelte-typewriter";

  let userName = "";
  let userStory = "";
  let threeCards = [];
  let requestID = null;
  let isLoading = true;
  let error = null;
  let interpretation = null;
  let showInterpretation = false;
  let loadCard1 = false;
  let loadCard2 = false;
  let loadCard3 = false;
  let showButton = false;

  function handleButtonClick() {
    window.location.href = "/";
  }

  if (typeof window !== "undefined") {
    const urlParams = new URLSearchParams(window.location.search);
    userName = urlParams.get("name") || "";
    userStory = urlParams.get("userstory") || "";
  }
  console.log(userName);
  console.log(userStory);

  async function getFate() {
    try {
      let data;
      do {
        const res = await fetch(
          `http://localhost:8082/cards/interpret/${requestID}`,
        );
        data = await res.json();
        if (data.interpretation) {
          interpretation = data.interpretation;
        } else {
          // Wait for 1 second before trying again
          await new Promise((resolve) => setTimeout(resolve, 1000));
        }
      } while (!data.interpretation);
    } catch (err) {
      error = err.message;
      console.error("Error:", err);
    } finally {
      showInterpretation = true;
    }
  }

  onMount(async () => {
    setTimeout(() => {
      loadCard1 = true;
    }, 3000);
    setTimeout(() => {
      loadCard2 = true;
    }, 5500);
    setTimeout(() => {
      loadCard3 = true;
    }, 8000);
    setTimeout(() => {
      showButton = true;
    }, 13000);

    try {
      const res = await fetch(
        `http://localhost:8082/cards?name=${userName}&userstory=${userStory}`,
      );
      const data = await res.text();
      threeCards = JSON.parse(data).cards;
      requestID = JSON.parse(data).requestID;
    } catch (err) {
      error = err.message;
      console.error("Error:", err);
    } finally {
      isLoading = false;
    }
  });

  let isHovering = [false, false, false];

  function hoverCard(index) {
    isHovering[index] = true;
  }

  function unhoverCard(index) {
    isHovering[index] = false;
  }
</script>

<link
  href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css"
  rel="stylesheet"
/>
<div class="container mx-auto mt-8">
  <div class="flex justify-between items-center mb-8">
    <Typewriter delay={0} interval={75}>
      <h3 class="text-2xl font-bold card-display-title">Cassandra's Parlor</h3>
    </Typewriter>
    <button on:click={handleButtonClick} style="--clr:#871616"
      ><span>Leave the Parlor</span><i></i></button
    >
  </div>

  {#if isLoading}
    <div>
      <p>Loading...</p>
    </div>
  {:else if error}
    <p class="text-red-500">Error: {error}</p>
  {:else if !showInterpretation}
    <div class="three-cards-container mx-auto px-20">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 ml-12 mr-12">
        {#each threeCards as card, index}
          {#if index === 0 && loadCard1}
            <div
              class="bg-card-container"
              role="link"
              tabindex="0"
              on:mouseenter={() => hoverCard(index)}
              on:mouseleave={() => unhoverCard(index)}
            >
              <h2 class="text-xl font-semibold mb-2">
                <span>Past</span>
              </h2>
              <h2 class="text-xl font-semibold mb-2">{card.name}</h2>
              <p class="text-#fed7aa-600 mb-2">Type: {card.type}</p>
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

              {#if card.reversed && isHovering[index]}
                <p class="text-#fed7aa-600 mb-2">
                  Meaning Reversed:
                  {#if card.meaning_rev.length > 90}
                    {card.meaning_rev.slice(0, 90) + "..."}
                  {:else}
                    {card.meaning_rev}
                  {/if}
                </p>
              {:else if isHovering[index]}
                <p class="text-#fed7aa-600 mb-2">
                  Meaning Upright:
                  {#if card.meaning_up.length > 90}
                    {card.meaning_up.slice(0, 90) + "..."}
                  {:else}
                    {card.meaning_up}
                  {/if}
                </p>
              {/if}
            </div>
          {:else if index === threeCards.length - 1 && loadCard3}
            <div
              class="bg-card-container"
              role="link"
              tabindex="0"
              on:mouseenter={() => hoverCard(index)}
              on:mouseleave={() => unhoverCard(index)}
            >
              <h2 class="text-xl font-semibold mb-2">
                <span>Future</span>
              </h2>
              <h2 class="text-xl font-semibold mb-2">{card.name}</h2>
              <p class="text-#fed7aa-600 mb-2">Type: {card.type}</p>
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

              {#if card.reversed && isHovering[index]}
                <p class="text-#fed7aa-600 mb-2">
                  Meaning Reversed:
                  {#if card.meaning_rev.length > 90}
                    {card.meaning_rev.slice(0, 90) + "..."}
                  {:else}
                    {card.meaning_rev}
                  {/if}
                </p>
              {:else if isHovering[index]}
                <p class="text-#fed7aa-600 mb-2">
                  Meaning Upright:
                  {#if card.meaning_up.length > 90}
                    {card.meaning_up.slice(0, 90) + "..."}
                  {:else}
                    {card.meaning_up}
                  {/if}
                </p>
              {/if}
            </div>
          {:else if index === threeCards.length - 2 && loadCard2}
            <div
              class="bg-card-container"
              role="link"
              tabindex="0"
              on:mouseenter={() => hoverCard(index)}
              on:mouseleave={() => unhoverCard(index)}
            >
              <h2 class="text-xl font-semibold mb-2">
                <span>Present</span>
              </h2>
              <h2 class="text-xl font-semibold mb-2">{card.name}</h2>
              <p class="text-#fed7aa-600 mb-2">Type: {card.type}</p>
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

              {#if card.reversed && isHovering[index]}
                <p class="text-#fed7aa-600 mb-2">
                  Meaning Reversed:
                  {#if card.meaning_rev.length > 90}
                    {card.meaning_rev.slice(0, 90) + "..."}
                  {:else}
                    {card.meaning_rev}
                  {/if}
                </p>
              {:else if isHovering[index]}
                <p class="text-#fed7aa-600 mb-2">
                  Meaning Upright:
                  {#if card.meaning_up.length > 90}
                    {card.meaning_up.slice(0, 90) + "..."}
                  {:else}
                    {card.meaning_up}
                  {/if}
                </p>
              {/if}
            </div>
          {/if}
        {/each}
      </div>
      {#if showButton}
        <button
          on:click={getFate}
          class="mx-auto mt-8 block"
          id="get-fate-btn"
          style="--clr:#200505"><span>Ask for reading</span><i></i></button
        >
      {/if}
    </div>
  {:else}
    <div class="mt-4 bg-card-container shadow-md p-4 rounded-lg">
      <p class="interpretation-text">{interpretation}</p>
    </div>
  {/if}
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");

  :global(html) {
    margin: 0;
    padding: 0;
    background-image: url("../../lib/assets/2-parlor.png");
    background-size: cover; /* Adjust as needed */
    background-position: center; /* Adjust as needed */
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    color: rgb(197, 176, 176);
  }

  .card-display-title {
    color: whitesmoke;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    opacity: 70%;
  }

  .card-image {
    max-width: 200px; /* Set the desired maximum width */
    max-height: 300px; /* Set the desired maximum height */
    display: block;
    margin: 1em auto;
    opacity: 80%;
    padding: 0.25em;
    box-shadow: 0 0 20px 10px rgba(197, 126, 34, 0.5);
  }

  .three-cards-container {
    min-height: 45em;
  }

  .bg-card-container {
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.9),
      rgba(80, 13, 2, 0.8)
    );
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    padding: 2em;
    margin: 1em;
    text-align: center;
    border: 3px solid black;
    border-radius: 50px;
    min-height: 37em;
    align-content: center;
  }

  .bg-card-container:hover {
    box-shadow: 0 0 20px 10px rgba(197, 126, 34, 0.5);
  }

  #get-fate-btn {
    letter-spacing: 0.4rem;
    font-size: 1.6rem;
    width: 15em;
    padding: 0.5em 0;
  }

  #get-fate-btn:hover {
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.4),
      rgba(12, 3, 2, 0.4)
    );
  }

  .interpretation-text {
    font-size: 1.5em;
    line-height: 1.6em;
  }

  button {
    position: relative;
    background: #200505;
    color: rgb(197, 176, 176);
    text-decoration: none;
    text-transform: uppercase;
    border: none;
    letter-spacing: 0.1rem;
    font-size: 0.8rem;
    padding: 0.6rem 2rem;
    transition: 0.2s;
    animation: box 2s infinite;
    border-radius: 20px;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
  }

  button:hover {
    letter-spacing: 0.1rem;
    padding: 0.6rem 2rem;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.7),
      rgba(80, 13, 2, 0.5)
    );
    /* box-shadow: 0 0 35px var(--clr); */
    animation: box 1.5s infinite;
  }

  button::before {
    content: "";
    position: absolute;
    inset: 2px;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.7),
      rgba(80, 13, 2, 0.5)
    );
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
