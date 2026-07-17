<script lang="ts">
	import { cardImage, type Card } from '$lib/types';

	interface Props {
		card: Card;
		reversed?: boolean;
		flipped?: boolean;
		onclick?: () => void;
		label?: string;
	}

	let { card, reversed = false, flipped = false, onclick, label }: Props = $props();
</script>

<button
	class="flip-card"
	class:is-flipped={flipped}
	{onclick}
	aria-label={label ?? (flipped ? card.name : 'Reveal card')}
>
	<div class="flip-card-inner">
		<div class="flip-face flip-face-front">
			<img src="/cards/tarot_back.png" alt="A face-down tarot card" />
		</div>
		<div class="flip-face flip-face-back">
			<img src={cardImage(card)} alt={card.name} class:reversed />
		</div>
	</div>
</button>

<style>
	.flip-card {
		display: block;
		width: 100%;
		aspect-ratio: 5 / 8.6;
		padding: 0;
		border: none;
		background: transparent;
		cursor: pointer;
	}
	.flip-card img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		display: block;
	}
	.flip-face {
		box-shadow: 0 0 0.8em 2px rgba(197, 126, 34, 0.5);
	}
	/* Reversed cards are shown upside-down, per tarot tradition. */
	.flip-face-back img.reversed {
		transform: rotate(180deg);
	}
</style>
