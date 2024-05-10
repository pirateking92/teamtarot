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
          `http://localhost:8082/cards/interpret/${requestID}`
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
        `http://localhost:8082/cards?name=${userName}&userstory=${userStory}`
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
      /^((?!mozilla|chrome|android).)*safari/i.test(navigator.userAgent) ||
      /.(up.browser|midp|silk|xoom|tablet|playbook|kindle|kindle fire|blazer|mobile|fennec|bolt|dalvik|armada|operamini|operamobi|openwave|nexusone|sam|series.?[0-9]+|x11).*/i.test(
        navigator.userAgent
      );
    return isDesktop ? "Desktop" : "Mobile";
  }
  let deviceType = detectDevice();
  console.log(deviceType);
</script>

<div
  class="bg-cover min-h-screen w-screen bg-center relative"
  style="background-image: url(/assets/2-parlor.png);"
>
  <header class="text-white top-0 left-0 right-0 z-10">
    <section class="mac-w-4xl mx-auto p-4 flex justify-between items-center">
      <h1 class="text-3xl font-medieval">Cassandra</h1>
      <button
        on:click={handleButtonClick}
        class="bg-green-700 rounded-md text-2xl font-medieval p-4 shadow-lg hover:bg-green-800 transition-colors duration-200"
        >Leave</button
      >
    </section>
  </header>
  <main class="flex items-center justify-center p-8">
    <div class="grid grid-cols-3 gap-7 justify-items-center items-center">
      <div><img src="/assets/tarot_back.png" alt="Tarot back" /></div>
      <div><img src="/assets/tarot_back.png" alt="Tarot back" /></div>
      <div><img src="/assets/tarot_back.png" alt="Tarot back" /></div>
    </div>
  </main>
  <div class="flex items-center justify-center">
    <button
      class="bg-green-700 text-white flex rounded-md text-2xl font-medieval p-4 shadow-lg hover:bg-green-800 transition-colors duration-200"
      >Draw Cards</button
    >
  </div>
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");
</style>
