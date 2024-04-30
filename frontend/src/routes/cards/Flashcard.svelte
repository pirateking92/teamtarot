<script>
    export let cardBack;
    export let cardFront;
    export let flipped;
</script>

<div class="flip-card-inner" class:flip-it={flipped}>
    <div class="flip-card-front">
        <div id="image-cont">
            <img src={cardBack} alt={"Back of tarot card"} />
        </div>
    </div>

    <div class="flip-card-back">
        <div id="image-cont">
            <img src={cardFront} alt={"Front of tarot card"} />
        </div>
    </div>
</div>

<style>
    /* Position the front and back side */
    .flip-card-front,
    .flip-card-back {
        position: absolute;
        width: 100%;
        height: 100%;
        -webkit-backface-visibility: hidden; /* Safari */
        backface-visibility: hidden;
    }

    /* Style the front side */
    .flip-card-front {
        color: black;
        display: flex;
        justify-content: center;
    }

    @keyframes revealTextSlowly {
        to {
            color: white;
        }
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
        max-width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    img {
        max-height: 80%;
        box-shadow: 0 0 25px 12px rgba(197, 126, 34, 0.5);
        border-radius: 10px;
        border: 3px rgb(97, 35, 9) solid;
    }
</style>
