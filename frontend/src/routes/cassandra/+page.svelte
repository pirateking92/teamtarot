<script>
  import "../../app.css";
  import Typewriter from "svelte-typewriter";
  import { onMount } from "svelte";

  let userName = "";
  let userStory = "";
  let typewriterFinished = false;
  let inputFieldAppeared = false;

  onMount(() => {
    document.querySelector('.container').classList.remove('hidden');
    // Simulating typewriter effect completion after 5 seconds
    setTimeout(() => {
      typewriterFinished = true;
    }, 23000);
    setTimeout(() => {
      inputFieldAppeared = true;
    }, 27000);
  });

  if (typeof window !== "undefined") {
    const urlParams = new URLSearchParams(window.location.search);
    userName = urlParams.get("name") || "";
  }

  function handleButtonClick() {
    const url = `/cards?name=${userName}&userstory=${encodeURIComponent(userStory)}`;
    window.location.href = url;
  }
</script>

<div class="container hidden">
  <br /><br />
  <div class="intro-text">
    <p>
      Welcome, {#if userName}{userName}{:else}seeker{/if}.
    </p>
    <Typewriter delay={500} interval={60}>
      <p>I can sense hope and fear within yourself.</p>
    </Typewriter>
    <br />
    <Typewriter delay={4500} interval={60}>
      <p>
        Let me guide you through the land where fate, fortune and destiny
        intertwine.
      </p>
    </Typewriter>
    <Typewriter delay={12000} interval={60}>
      <p>
        In order to do this, we will draw three cards, representing Past,
        Present and Future.
      </p>
    </Typewriter>
    <br />
    <Typewriter delay={20000} interval={60}>
      <p>But now tell me, what has been tormenting you?</p>
    </Typewriter>
    <br />
    {#if typewriterFinished}
      <textarea
        class="input-field"
        placeholder="Cassandra, I'd like to know about..."
        bind:value={userStory}
        style="resize: none;"
      />
      <br />
      <br />
    {/if}
    {#if inputFieldAppeared}
      <button on:click={handleButtonClick} style="--clr:#871616"
        ><span>Enter the Parlor</span><i></i></button
      >
    {/if}
  </div>
  <br />
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");

  :global(html) {
    background-image: url("../../lib/assets/1-alley.png");
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

  .container {
    display: flex;
    margin: 5em;
    min-height: 25em;
    box-shadow: 0 0 1em 4px rgba(207, 49, 5, 0.5);
    padding: 0.5em 5em;
    font-weight: 600;
    font-size: 1.7rem;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.9),
      rgba(19, 5, 0, 0.9)
    );
    border-radius: 50px;
  }

  .intro-text {
    text-align: center;
    margin: auto;
    padding: 2em 1em;
    border-radius: 50px;
  }

  .input-field {
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.7),
      rgba(19, 5, 0, 0.7)
    );
    box-shadow: 0 0 1em 4px rgba(149, 35, 4, 0.5);
    border-radius: 20px;
    padding: 0.4em;
    margin: 0 auto 2em auto;
    border: #200505 solid 2px;
    outline: none;
    min-width: 20em;
    text-align: center;
    font-size: 1.2rem;
  }

  button {
    position: relative;
    background: #200505;
    color: rgb(197, 176, 176);
    text-decoration: none;
    text-transform: uppercase;
    border: none;
    letter-spacing: 0.2rem;
    font-size: 1.2rem;
    padding: 1.3rem 3rem;
    transition: 0.2s;
    animation: box 2s infinite;
    border-radius: 20px;
  }

  button:hover {
    letter-spacing: 0.3rem;
    padding: 1.3rem 3rem;
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
</style>
