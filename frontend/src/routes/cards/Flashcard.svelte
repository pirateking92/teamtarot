<script>
    export let cardBack;
    export let cardFront;
    export let flipped;
    export let reversed;
</script>

<div class="flip-card-inner" class:flip-it={flipped}>
    <div class="flip-card-front">
        <div id="image-cont" class:reveal-delay={flipped}>
            <img
                src={flipped ? cardFront : cardBack}
                alt={"Front of tarot card"}
                style={reversed ? "transform: rotate(180deg)" : ""}
            />
        </div>
    </div>

    <div class="flip-card-back">
        <div id="image-cont">
            <img src={cardBack} alt={"Front of tarot card"} />
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
        transform: scaleX(-1);
    }

    @keyframes delayReveal {
        0% {
            opacity: 0;
        }
        70% {
            opacity: 0.3;
        }
        100% {
            opacity: 1;
        }
    }
    .reveal-delay {
        animation: delayReveal 0.4s linear;
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
        max-height: 88%;
        padding: 0.25em;
        box-shadow: 0 0 20px 10px rgba(197, 126, 34, 0.5);
        border-radius: 10px;
    }
</style>
