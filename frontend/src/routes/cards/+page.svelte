<script>
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import Typewriter from "svelte-typewriter";
  import Flashcard from "./Flashcard.svelte";

  let userName = "";
  let userStory = "";
  let threeCards = [];
  let requestID = null;
  let isLoading = true;
  let error = null;
  let interpretation = null;
  let showInterpretation = false;
  let postShowInterpretation = false;
  let showButton = false;
  let flipped = [false, false, false];
  let postFlip = [false, false, false];
  let isCardVisible = [false, false, false];
  let showDeck = true;
  let showGuide = true;
  let hovered = [false, false, false];

  // function showPopup() {
  //   document.getElementById("popup").classList.remove("hidden");
  // }

  // function hidePopup() {
  //   document.getElementById("popup").classList.add("hidden");
  // }

  function handleButtonClick() {
    window.location.href = "/";
  }

  function handleCardClick(index) {
    if (isCardVisible.every((value) => value === true)) {
      if (flipped[index] == false) {
        flipped[index] = !flipped[index];
      }
      if (flipped.every((value) => value === true)) {
        showButton = true;
        showGuide = false;
      }
    }
  }

  function drawCard() {
    for (let i = 0; i < isCardVisible.length; i++) {
      if (isCardVisible[i] === false) {
        isCardVisible[i] = true;
        if (isCardVisible.every((value) => value === true)) {
          showDeck = false;
        }
        break;
      }
    }
  }

  function closeInterpretationModal() {
    showInterpretation = false;
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
          `https://teamtarot.onrender.com/cards/interpret/${requestID}`
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
      postShowInterpretation = true;
    }
  }

  onMount(async () => {
    document.querySelector(".container").classList.remove("hidden");
    try {
      const res = await fetch(
        `https://teamtarot.onrender.com/cards?name=${userName}&userstory=${userStory}`
      );
      const data = await res.text();
      threeCards = JSON.parse(data).cards;
      requestID = JSON.parse(data).requestID;
      postFlip = [false, false, false];
    } catch (err) {
      error = err.message;
      console.error("Error:", err);
    } finally {
      isLoading = false;
    }
  });

  // let isHovering = [false, false, false];

  // function hoverCard(index) {
  //   isHovering[index] = true;
  // }

  // function unhoverCard(index) {
  //   isHovering[index] = false;
  // }

  // <!-- on:mouseenter={() => hoverCard(index)}
  //             on:mouseleave={() => unhoverCard(index)} -->
</script>

<link
  href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css"
  rel="stylesheet"
/>
<div class="container mx-auto mt-8 hidden">
  <div class="flex justify-between items-center mb-8">
    <h3 class="text-2xl font-bold card-display-title">Cassandra's Parlor</h3>
    {#if showGuide && isCardVisible.every((value) => value === true)}
      <div
        class="mx-auto mt-8 min-h-50"
        id="guide-box"
        style="--clr:#200505; grid-area: guide;"
        transition:fade
      >
        <span>Now, flip the cards</span><i></i>
      </div>
    {:else if showGuide}
      <div
        class="mx-auto mt-8 min-h-50"
        id="guide-box"
        style="--clr:#200505; grid-area: guide;"
      >
        <span>Draw three cards from the deck, {userName}</span><i></i>
      </div>
    {/if}
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
      <div
        class="grid grid-cols-3 md:grid-cols-3 gap-4"
        style="grid-template-columns: 1fr 1fr 1fr;
        grid-template-rows: 0.1fr 0.3fr 0.3fr 0.1fr 0.3fr;
        grid-template-areas: 
        'button card-present-1 guide'
        'card-past-0 card-present-1 card-future-2'
        'card-past-0 card-present-1 card-future-2'
        'card-past-0 deck card-future-2'
        '........... deck .............'"
      >
        {#each threeCards as card, index}
          {#if index === 0}
            <div
              class="bg-card-container"
              on:mouseenter={() => {
                if (postFlip[index]) hovered[index] = true;
                console.log("Postflip below");
                console.log(postFlip[index]);
                console.log("Hovered below");
                console.log(hovered[index]);
              }}
              on:mouseleave={() => {
                if (isCardVisible[index] && flipped[index])
                  (hovered[index] = false), (postFlip[index] = true);
                console.log(postFlip[index]);
              }}
              role="link"
              tabindex="0"
              style="grid-area: card-past-{index};"
              transition:fade
            >
              <h2 class="text-xl font-semibold mb-2">
                <span>Past</span>
              </h2>
              {#if isCardVisible[index]}
                {#if flipped[index]}
                  <h2 class="text-xl mt-2 font-semibold mb-2">{card.name}</h2>
                  <!-- <p class="text-#fed7aa-600 mb-2">Type: {card.type}</p> -->
                {/if}
                {#if hovered[index] && flipped[index]}
                  {#if card.reversed}
                    <p class="meaning-text text-#fed7aa-600 mb-2">
                      Meaning: <br /><br />
                      {card.meaning_rev}
                    </p>
                  {:else}
                    <p class="meaning-text text-#fed7aa-600 mb-2">
                      Meaning: <br /><br />
                      {card.meaning_up}
                    </p>
                  {/if}
                {:else if isCardVisible[index]}
                  <div class="flip-card">
                    <div class="flip-card-inner" class:flip-it={flipped[index]}>
                      <div
                        on:click={() => handleCardClick(index)}
                        on:keydown={(event) => {}}
                        tabindex="0"
                        role="button"
                      >
                        <Flashcard
                          cardBack={`$app/assets/assets/tarot_back.png`}
                          cardFront={`$app/assets/assets/${card.image_file_name}`}
                          flipped={flipped[index]}
                          reversed={card.reversed}
                        />
                      </div>
                    </div>
                  </div>
                {/if}
              {/if}
            </div>
          {:else if index === threeCards.length - 1}
            <div
              class="bg-card-container"
              on:mouseenter={() => {
                if (postFlip[index]) hovered[index] = true;
                console.log(postFlip[index]);
              }}
              on:mouseleave={() => {
                if (isCardVisible[index] && flipped[index])
                  (hovered[index] = false), (postFlip[index] = true);
                console.log(postFlip[index]);
              }}
              role="link"
              tabindex="0"
              style="grid-area: card-future-{index};"
            >
              <h2 class="text-xl font-semibold mb-2">
                <span>Future</span>
              </h2>
              {#if isCardVisible[index]}
                {#if flipped[index]}
                  <h2 class="text-xl mt-2 font-semibold mb-2">{card.name}</h2>
                  <!-- <p class="text-#fed7aa-600 mb-2">Type: {card.type}</p> -->
                {/if}
                {#if hovered[index] && flipped[index]}
                  {#if card.reversed}
                    <p class="meaning-text text-#fed7aa-600 mb-2">
                      Meaning: <br /><br />
                      {card.meaning_rev}
                    </p>
                  {:else}
                    <p class="meaning-text text-#fed7aa-600 mb-2">
                      Meaning: <br /><br />
                      {card.meaning_up}
                    </p>
                  {/if}
                {:else if isCardVisible[index]}
                  <div class="flip-card">
                    <div class="flip-card-inner" class:flip-it={flipped[index]}>
                      <div
                        on:click={() => handleCardClick(index)}
                        on:keydown={(event) => {}}
                        tabindex="0"
                        role="button"
                      >
                        <Flashcard
                          cardBack={`src/lib/assets/tarot_back.png`}
                          cardFront={`src/lib/assets/${card.image_file_name}`}
                          flipped={flipped[index]}
                          reversed={card.reversed}
                        />
                      </div>
                    </div>
                  </div>
                {/if}
              {/if}
            </div>
          {:else if index === threeCards.length - 2}
            <div
              class="bg-card-container"
              on:mouseenter={() => {
                if (postFlip[index]) hovered[index] = true;
                console.log(postFlip[index]);
              }}
              on:mouseleave={() => {
                if (isCardVisible[index] && flipped[index])
                  (hovered[index] = false), (postFlip[index] = true);
                console.log(postFlip[index]);
              }}
              role="link"
              tabindex="0"
              style="grid-area: card-present-{index};"
            >
              <h2 class="text-xl font-semibold mb-2">
                <span>Present</span>
              </h2>
              {#if isCardVisible[index]}
                {#if flipped[index]}
                  <h2 class="text-xl mt-2 font-semibold mb-2">{card.name}</h2>
                  <!-- <p class="text-#fed7aa-600 mb-2">Type: {card.type}</p> -->
                {/if}
                {#if hovered[index] && flipped[index]}
                  {#if card.reversed}
                    <p class="meaning-text text-#fed7aa-600 mb-2">
                      Meaning: <br /><br />
                      {card.meaning_rev}
                    </p>
                  {:else}
                    <p class="meaning-text text-#fed7aa-600 mb-2">
                      Meaning: <br /><br />
                      {card.meaning_up}
                    </p>
                  {/if}
                {:else if isCardVisible[index]}
                  <div class="flip-card">
                    <div class="flip-card-inner" class:flip-it={flipped[index]}>
                      <div
                        on:click={() => handleCardClick(index)}
                        on:keydown={(event) => {}}
                        tabindex="0"
                        role="button"
                      >
                        <Flashcard
                          cardBack={`src/lib/assets/tarot_back.png`}
                          cardFront={`src/lib/assets/${card.image_file_name}`}
                          flipped={flipped[index]}
                          reversed={card.reversed}
                        />
                      </div>
                    </div>
                  </div>
                {/if}
              {/if}
            </div>
          {/if}
        {/each}
        {#if showDeck}
          <img
            on:click={drawCard}
            src={`src/lib/assets/tarot_deck.png`}
            alt="deck"
            class="mb-2 rounded-lg deck-image"
            style="grid-area: deck;"
          />
        {/if}
        {#if showButton && !postShowInterpretation}
          <button
            transition:fade
            on:click={getFate}
            class="mx-auto mt-8"
            id="get-fate-btn"
            style="--clr:#200505; grid-area: button;"
            ><span>Ask for reading</span><i></i></button
          >
        {:else if showButton}
          <button
            transition:fade
            on:click={getFate}
            class="mx-auto mt-8"
            id="get-fate-btn"
            style="--clr:#200505; grid-area: button;"
            ><span>Show reading</span><i></i></button
          >
        {/if}
      </div>
    </div>
  {:else}
    <div class="modal-backdrop">
      <div class="interpretation-modal">
        <div
          class="interpretation-container shadow-md p-4 rounded-lg"
          transition:fade
        >
          <button class="close-button" on:click={closeInterpretationModal}
            >back</button
          >
          <p class="interpretation-text" style="white-space: pre-line;">
            {interpretation}
          </p>
        </div>
      </div>
    </div>
  {/if}
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");

  :global(html) {
    background-image: url("../../lib/assets/2-parlor.png");
    background-size: cover; /* Adjust as needed */
    background-position: center; /* Adjust as needed */
    background-repeat: no-repeat;
    height: 100%;
  }

  :global(body) {
    margin: 0;
    padding: 0;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    color: rgb(197, 176, 176);
  }

  .card-display-title {
    color: whitesmoke;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    opacity: 70%;
  }

  .deck-image {
    max-width: 160px; /* Set the desired maximum width */
    max-height: 270px; /* Set the desired maximum height */
    border: 0.25em rgb(50, 17, 3) solid;
    box-shadow: 0 0 20px 10px rgba(197, 126, 34, 0.5);
    transition: 0.2s;
    position: absolute;
    align-self: flex-end;
    justify-self: center;
    margin: 5em 0 0 0;
    animation: box 2s infinite;
  }

  .deck-image::before {
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

  .three-cards-container {
    min-height: 45em;
    display: grid;
  }

  .bg-card-container {
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.9),
      rgba(80, 13, 2, 0.8)
    );
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    padding: 2em auto;
    margin: 1em auto;
    text-align: center;
    border: 3px solid black;
    border-radius: 50px;
    min-height: 35em;
    min-width: 336px;
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
    margin: 5em auto auto auto;
    max-height: 3em;
    position: absolute;
    justify-self: center;
    align-self: flex-end;
  }

  #guide-box {
    letter-spacing: 0.16rem;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.9),
      rgba(80, 13, 2, 1)
    );
    border-radius: 10px;
    font-size: 1.2rem;

    padding: 0.5em 1.2em;
    margin: auto auto -1.7em auto;
    max-height: 5em;
    border: 3px solid black;
    justify-self: center;
    text-align: center;
    align-self: center;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
  }

  #get-fate-btn:hover {
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.4),
      rgba(12, 3, 2, 0.4)
    );
  }

  .meaning-text {
    max-width: 18em;
    display: block;
    margin: auto;
    transition: 1s ease;
    line-height: 2em;
  }

  .interpretation-text {
    font-size: 1.1em;
    line-height: 1.6em;
  }

  .modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0); /* Semi-transparent black */
    z-index: 999;
  }

  .interpretation-modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 1000; /* Ensure it's above the backdrop */
  }

  .close-button {
    position: absolute;
    top: 15px;
    right: 30px;
    background: none;
    border: none;
    color: #fff;
    font-size: 15px;
    cursor: pointer;
  }

  .interpretation-container {
    display: block;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.9),
      rgba(80, 13, 2, 0.8)
    );
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    padding: 2em;
    margin: 13em auto;
    text-align: center;
    border: 3px solid black;
    border-radius: 50px;
    min-height: 35em;
    align-content: center;
    align-self: center;
    position: relative;
  }

  /* BUTTON STYLING BELOW */

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

  .container.hidden {
    display: none;
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

  .flip-card {
    background-color: transparent;
    width: 330px;
    height: 340px; /* the max this can be without distorting the grid */
    perspective: 1000px; /* Remove this if you don't want the 3D effect */
  }

  /* This container is needed to position the front and back side */
  .flip-card-inner {
    position: relative;
    width: 100%;
    height: 100%;
    text-align: center;
    transition: transform 0.4s;
    transform-style: preserve-3d;
  }

  .flip-it {
    transform: rotateY(180deg);
  }
</style>
