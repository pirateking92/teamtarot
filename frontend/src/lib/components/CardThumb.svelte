<script lang="ts">
	import { cardImage, type Card } from '$lib/types';

	interface Props {
		card: Card;
		reversed?: boolean;
		href?: string;
	}

	let { card, reversed = false, href }: Props = $props();
</script>

{#snippet body()}
	<img src={cardImage(card)} alt={card.name} class:reversed />
	<span class="name">{card.name}</span>
{/snippet}

{#if href}
	<a class="thumb" {href}>{@render body()}</a>
{:else}
	<div class="thumb">{@render body()}</div>
{/if}

<style>
	.thumb {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.4rem;
		text-decoration: none;
		color: rgb(197, 176, 176);
	}
	.thumb img {
		width: 100%;
		aspect-ratio: 5 / 8.6;
		object-fit: cover;
		border-radius: 0.6rem;
		box-shadow: 0 0 0.6em 1px rgba(197, 126, 34, 0.4);
		transition:
			transform 0.2s,
			box-shadow 0.2s;
	}
	a.thumb:hover img {
		transform: translateY(-4px);
		box-shadow: 0 0 1em 3px rgba(207, 49, 5, 0.5);
	}
	.reversed {
		transform: rotate(180deg);
	}
	a.thumb:hover .reversed {
		transform: translateY(-4px) rotate(180deg);
	}
	.name {
		font-size: 0.85rem;
		text-align: center;
		line-height: 1.2;
	}
</style>
