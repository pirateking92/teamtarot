<script>
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import Typewriter from "svelte-typewriter";
  import Flashcard from "./Flashcard.svelte";
  import "../../tailwind.css";

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

  function detectDevice() {
    const isDesktop =
      /^((?!chrome|android).)*safari/i.test(navigator.userAgent) ||
      /.(up.browser|midp|silk|xoom|tablet|playbook|kindle|kindle fire|blazer|mobile|fennec|bolt|dalvik|armada|operamini|operamobi|openwave|nexusone|sam|series.?[0-9]+|x11).*/i.test(
        navigator.userAgent
      );
    return isDesktop ? "Desktop" : "Mobile";
  }
  let deviceType = detectDevice();
  console.log(deviceType);

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

<!-- <link
  href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css"
  rel="stylesheet"
/> -->
<div class="font-medieval text-green-500">
  <h1 class="text-xl">Hello, world!</h1>
</div>

<div
  class="bg-cover h-screen w-screen"
  style="background-image: url(/assets/2-parlor.png);"
>
  <div class="container mx-auto mt-8 hidden">
    <div class="flex justify-between items-center mb-8">
      <h3 class="text-2xl font-bold">Cassandra's Parlor</h3>
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
      <div class="mx-auto px-20">
        <div
          class="grid grid-cols-1 md:grid-cols-3"
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
                class="bg-card-container grid-area-card-past"
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
                    <h2 class="text-xl mt-2 font-semibold mb-2">
                      {card.name}
                    </h2>
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
                      <div
                        class="flip-card-inner"
                        class:flip-it={flipped[index]}
                      >
                        <div
                          on:click={() => handleCardClick(index)}
                          on:keydown={(event) => {}}
                          tabindex="0"
                          role="button"
                        >
                          <Flashcard
                            cardBack={`/assets/tarot_back.png`}
                            cardFront={`/assets/${card.image_file_name}`}
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
                    <h2 class="text-xl mt-2 font-semibold mb-2">
                      {card.name}
                    </h2>
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
                      <div
                        class="flip-card-inner"
                        class:flip-it={flipped[index]}
                      >
                        <div
                          on:click={() => handleCardClick(index)}
                          on:keydown={(event) => {}}
                          tabindex="0"
                          role="button"
                        >
                          <Flashcard
                            cardBack={`/assets/tarot_back.png`}
                            cardFront={`/assets/${card.image_file_name}`}
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
                    <h2 class="text-xl mt-2 font-semibold mb-2">
                      {card.name}
                    </h2>
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
                      <div
                        class="flip-card-inner"
                        class:flip-it={flipped[index]}
                      >
                        <div
                          on:click={() => handleCardClick(index)}
                          on:keydown={(event) => {}}
                          tabindex="0"
                          role="button"
                        >
                          <Flashcard
                            cardBack={`/assets/tarot_back.png`}
                            cardFront={`/assets/${card.image_file_name}`}
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
            <button
              on:click={drawCard}
              on:keydown={(event) => {
                if (event.key === "Enter" || event.key === " ") {
                  drawCard();
                }
              }}
              class="mb-2 rounded-lg deck-image"
              style="grid-area: deck;"
            >
              <img src={`/assets/tarot_deck.png`} alt="deck" />
            </button>
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
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");
</style>
