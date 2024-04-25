<script>
  import "../../app.css";
</script>

<div class="container">
  <br /><br />
  <div class="intro-text">
    <p>
      I can sense hope and fear within yourself, seeker.
      Let me guide you through the land where fate, fortune and destiny
      intertwine.
    </p>
    <br />
    <p>
      But be warned - the truths we may uncover could shake the very foundations
      of what you think you know.
    </p>
    <br />
    <p>
      After this, there is no turning back. If you decide to leave - the story
      ends, you return to your world and believe whatever you want to believe.
      If you enter - I will show you how deep the rabbit hole goes.
    </p>
    <br />
    <br />
    <button
      on:click={() => (window.location.href = "/card_display")}
      style="--clr:#c377d4"><span>Enter</span><i></i></button
    >
  </div>
  <br />
</div>

<style>
  @import url("https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap");

  :global(html, body) {
    height: 100%;
    margin: 0;
    padding: 0;
    background: linear-gradient(to right, #20051c, #46204e);
  }

  .container {
    margin: auto;
    color: whitesmoke;
    padding: 0.5em;
    font-weight: 600;
    font-family: MedievalSharp, Arial, Helvetica, sans-serif;
    font-size: 1.7rem;
    opacity: 70%;
  }

  .intro-text {
    max-width: 80%;
    text-align: center;
    margin: auto;
    padding: 0.5em;
    border-radius: 50px;
  }

  button {
    position: relative;
    background: #1d1242;
    color: #fff;
    text-decoration: none;
    text-transform: uppercase;
    border: none;
    letter-spacing: 0.2rem;
    font-size: 1.5rem;
    padding: 1.3rem 3rem;
    transition: 0.2s;
    animation: box 2s infinite;
    border-radius: 20px;
  }

  button:hover {
    letter-spacing: 0.3rem;
    padding: 1.3rem 3rem;
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
