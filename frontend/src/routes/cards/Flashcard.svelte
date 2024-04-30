<script>
    export let cardBack;
    export let cardFront;
    export let flipped;
</script>

<div class="flip-card-front">
    <div id="image-cont">
        <img src={cardBack} alt={"Back of tarot card"} />
    </div>
</div>

<div class="flip-card-back" class:conceal-answer={flipped}>
    <div id="image-cont">
        <img src={cardFront} alt={"Front of tarot card"} />
    </div>
</div>

<!-- Script
 let flashcardIndex = 0;
 $: clue = vocab[flashcardIndex].image;
 $: answer = vocab[flashcardIndex].word;
let showCardBack = false; 
 const toggleShowBack = () => showCardBack = !showCardBack;

 HTML
 <div class="flip-box-inner" class:flip-it={showCardBack}>
   <Flashcard   {clue} 
                {answer} 
                {showCardBack} 
      />
    </div>
↓  Applies to the button, in Svelte docs example, but could be applied to image? ↓ 
<button on:click={toggleShowBack}>
   {showCardBack ? "Hide Answer" : "Show Answer"}
  </button>

  CSS
.flip-it {
  transform: rotateY(180deg);
 } -->

<style>
    /* Position the front and back side */
    .flip-card-front,
    .flip-card-back {
        position: absolute;
        width: 100%;
        height: 100%;
        -webkit-backface-visibility: hidden; /* Safari */
        backface-visibility: hidden;
        box-shadow: -1px 1px 3px black;
    }

    /* Style the front side */
    .flip-card-front {
        background-color: #ddd;
        color: black;
        display: flex;
        justify-content: center;
    }

    @keyframes revealTextSlowly {
        to {
            color: white;
        }
    }

    .conceal-answer {
        animation: revealTextSlowly 0.3s forwards;
    }

    /* Style the back side */
    .flip-card-back {
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: #61592f;
        color: #61592f;
        /* 		color: white; */
        font-size: 1.3rem;
        transform: rotateY(180deg);
    }

    #image-cont {
        background-color: hsl(50, 65%, 75%);
        max-width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    img {
        max-height: 80%;
    }
</style>
