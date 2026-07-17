<script lang="ts">
	import { cardImage } from '$lib/types';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const card = data.card;
</script>

<svelte:head><title>Cassandra — {card.name}</title></svelte:head>

<section class="detail">
	<div class="art">
		<img src={cardImage(card)} alt={card.name} />
	</div>
	<div class="panel info">
		<a class="back" href="/cards">← The Deck</a>
		<h1>{card.name}</h1>
		<p class="type">{card.arcana === 'major' ? 'Major Arcana' : `${card.suit} · Minor Arcana`}</p>

		<h2>Upright</h2>
		<p>{card.meaningUp}</p>
		<h2>Reversed</h2>
		<p>{card.meaningReverse}</p>
		<h2>Description</h2>
		<p class="desc">{card.description}</p>
	</div>
</section>

<style>
	.detail {
		display: grid;
		grid-template-columns: minmax(10rem, 16rem) 1fr;
		gap: 2rem;
		align-items: start;
		margin: 2rem 0;
	}
	.art img {
		width: 100%;
		border-radius: 0.75rem;
		box-shadow: 0 0 1em 3px rgba(197, 126, 34, 0.5);
	}
	.info {
		padding: 1.75rem 2rem;
	}
	.back {
		color: rgb(197, 176, 176);
		opacity: 0.75;
		text-decoration: none;
		font-size: 0.9rem;
	}
	.info h1 {
		margin: 0.5rem 0 0;
	}
	.type {
		opacity: 0.65;
		margin-top: 0.25rem;
	}
	.info h2 {
		font-size: 1.05rem;
		margin: 1.25rem 0 0.3rem;
		color: #d98a5a;
	}
	.info p {
		margin: 0;
		line-height: 1.6;
	}
	.desc {
		opacity: 0.85;
		font-size: 0.95rem;
	}
	@media (max-width: 40rem) {
		.detail {
			grid-template-columns: 1fr;
			justify-items: center;
		}
		.art {
			max-width: 14rem;
		}
	}
</style>
