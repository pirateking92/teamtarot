<script>
  import "../app.css";
  import Typewriter from "svelte-typewriter";
  import { onMount } from "svelte";

  let userName = "";
  let typewriterFinished = false;
  let inputFieldAppeared = false;

  onMount(() => {
    // Simulating typewriter effect completion after 5 seconds
    setTimeout(() => {
      typewriterFinished = true;
    }, 9000);
    setTimeout(() => {
      inputFieldAppeared = true;
    }, 13000);
  });

  function handleButtonClick() {
    const url = `/cassandra?name=${encodeURIComponent(userName)}`;
    window.location.href = url;
  }
</script>

<div class="container">
  <br /><br />
  <div class="intro-text">
    <Typewriter delay={500} interval={75}>
      <h1>Ah, I've been expecting you, seeker.</h1>
    </Typewriter>
    <Typewriter delay={4000} interval={75}>
      <h1>My name is Cassandra.</h1>
    </Typewriter>
    <br />
    <Typewriter delay={6000} interval={75}>
      <h1>What's yours?</h1>
    </Typewriter>
    <br /><br />
    {#if typewriterFinished}
      <input
        class="input-field"
        placeholder="Enter your name"
        bind:value={userName}
      />
    {/if}
    <br /><br /><br />
    {#if inputFieldAppeared}
      <button on:click={handleButtonClick} style="--clr:#871616"
        ><span>Follow Cassandra</span><i></i></button
      >
    {/if}
  </div>
  <br />
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");

  :global(html, body) {
    margin: 0;
    padding: 0;
    background-image: url("../lib/assets/0-town.jpeg");
    background-size: cover; /* Adjust as needed */
    background-position: center; /* Adjust as needed */
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
