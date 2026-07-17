<script lang="ts">
	import CardThumb from '$lib/components/CardThumb.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const { card, reversed } = data.daily;
	const meaning = reversed ? card.meaningReverse : card.meaningUp;
	const today = new Date(data.daily.date).toLocaleDateString(undefined, {
		weekday: 'long',
		year: 'numeric',
		month: 'long',
		day: 'numeric'
	});
</script>

<svelte:head><title>Cassandra — Card of the Day</title></svelte:head>

<section class="daily">
	<header class="panel head">
		<h1>The Card of the Day</h1>
		<p class="date">{today}</p>
	</header>

	<div class="body">
		<div class="art">
			<CardThumb {card} {reversed} href={`/cards/${card.nameShort}`} />
		</div>
		<div class="panel meaning">
			<h2>{card.name}{reversed ? ' (Reversed)' : ''}</h2>
			<p class="type">
				{card.arcana === 'major' ? 'Major Arcana' : `${card.suit} · Minor Arcana`}
			</p>
			<p>{meaning}</p>
			<p class="desc">{card.description}</p>
		</div>
	</div>
</section>

<style>
	.head {
		text-align: center;
		padding: 1.25rem;
		margin: 1.5rem 0 2rem;
	}
	.head h1 {
		margin: 0;
	}
	.date {
		opacity: 0.75;
		margin: 0.4rem 0 0;
	}
	.body {
		display: grid;
		grid-template-columns: minmax(9rem, 14rem) 1fr;
		gap: 2rem;
		align-items: start;
	}
	.meaning {
		padding: 1.75rem;
	}
	.meaning h2 {
		margin-top: 0;
	}
	.type {
		opacity: 0.65;
		margin-top: -0.5rem;
	}
	.desc {
		opacity: 0.8;
		font-size: 0.95rem;
	}
	@media (max-width: 40rem) {
		.body {
			grid-template-columns: 1fr;
			justify-items: center;
		}
		.art {
			max-width: 12rem;
		}
	}
</style>
