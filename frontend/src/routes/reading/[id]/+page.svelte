<script lang="ts">
	import { onMount } from 'svelte';
	import Flashcard from '$lib/components/Flashcard.svelte';
	import OrnateButton from '$lib/components/OrnateButton.svelte';
	import type { Reading, ReadingStatus } from '$lib/types';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	const reading: Reading = data.reading;
	const cards = reading.cards;

	let flipped = $state(cards.map(() => false));
	let status = $state<ReadingStatus>(reading.status);
	let interpretation = $state(reading.interpretation);
	let revealRequested = $state(false);

	const allRevealed = $derived(flipped.every(Boolean));
	const anyRevealed = $derived(flipped.some(Boolean));

	function reveal(i: number) {
		if (flipped[i]) return;
		flipped[i] = true;
		flipped = [...flipped];
	}

	function meaningFor(i: number): string {
		const rc = cards[i];
		return rc.reversed ? rc.card.meaningReverse : rc.card.meaningUp;
	}

	// Poll the interpretation until it's ready (bounded).
	onMount(() => {
		if (status === 'complete' || status === 'error') return;
		let tries = 0;
		let stopped = false;
		const tick = async () => {
			if (stopped || tries >= 40) return;
			tries++;
			try {
				const res = await fetch(`/api/readings/${reading.id}`);
				if (res.ok) {
					const { reading: latest } = (await res.json()) as { reading: Reading };
					status = latest.status;
					interpretation = latest.interpretation;
					if (status === 'complete' || status === 'error') return;
				}
			} catch {
				// transient; keep polling
			}
			setTimeout(tick, 1500);
		};
		setTimeout(tick, 1200);
		return () => {
			stopped = true;
		};
	});
</script>

<svelte:head>
	<title>Cassandra — Your Reading</title>
</svelte:head>

<section class="reading">
	<header class="reading-head panel">
		<h1>{reading.querentName ? `${reading.querentName}’s Reading` : 'Your Reading'}</h1>
		{#if !allRevealed}
			<p class="hint">Touch each card to turn it over.</p>
		{/if}
	</header>

	<div class="cards" style="--n: {cards.length}">
		{#each cards as rc, i (i)}
			<div class="card-slot">
				<span class="position">{rc.positionLabel}</span>
				<Flashcard
					card={rc.card}
					reversed={rc.reversed}
					flipped={flipped[i]}
					onclick={() => reveal(i)}
					label={flipped[i] ? rc.card.name : `Reveal the ${rc.positionLabel} card`}
				/>
				{#if flipped[i]}
					<div class="card-detail">
						<strong>{rc.card.name}{rc.reversed ? ' (Reversed)' : ''}</strong>
						<em>{rc.card.arcana === 'major' ? 'Major Arcana' : `${rc.card.suit} · Minor Arcana`}</em>
						<p>{meaningFor(i)}</p>
					</div>
				{/if}
			</div>
		{/each}
	</div>

	{#if allRevealed && !revealRequested}
		<div class="reveal-row">
			<OrnateButton onclick={() => (revealRequested = true)}>Reveal Cassandra’s reading</OrnateButton>
		</div>
	{/if}

	{#if revealRequested}
		<div class="interpretation panel">
			{#if status === 'complete'}
				{#each interpretation.split('\n').filter((p) => p.trim()) as para}
					<p>{para}</p>
				{/each}
			{:else if status === 'error'}
				<p class="error">{interpretation || 'The reading could not be completed.'}</p>
			{:else}
				<p class="waiting">Cassandra gazes into the cards…</p>
			{/if}
		</div>
	{/if}

	<div class="again">
		<a href="/reading/new">Draw again</a>
	</div>
</section>

<style>
	.reading {
		margin: 1.5rem auto;
	}
	.reading-head {
		text-align: center;
		padding: 1.25rem 1.5rem;
		margin-bottom: 2rem;
	}
	.reading-head h1 {
		margin: 0;
		font-size: 1.8rem;
	}
	.hint {
		opacity: 0.75;
		margin: 0.5rem 0 0;
	}
	.cards {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(9rem, 1fr));
		gap: 1.5rem;
		align-items: start;
	}
	.card-slot {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.6rem;
	}
	.position {
		font-size: 1.05rem;
		letter-spacing: 0.05em;
		opacity: 0.85;
	}
	.card-detail {
		text-align: center;
		font-size: 0.85rem;
		line-height: 1.4;
	}
	.card-detail strong {
		display: block;
		font-size: 1rem;
	}
	.card-detail em {
		opacity: 0.65;
		font-size: 0.78rem;
	}
	.card-detail p {
		margin: 0.4rem 0 0;
		opacity: 0.9;
	}
	.reveal-row {
		text-align: center;
		margin: 2.5rem 0 1rem;
	}
	.interpretation {
		margin: 1.5rem auto 0;
		padding: 2rem 2.25rem;
		max-width: 46rem;
		font-size: 1.1rem;
		line-height: 1.7;
	}
	.interpretation p {
		margin: 0 0 1rem;
	}
	.waiting {
		text-align: center;
		font-style: italic;
		opacity: 0.8;
	}
	.error {
		color: #ff9b7a;
	}
	.again {
		text-align: center;
		margin-top: 2.5rem;
	}
	.again a {
		color: rgb(197, 176, 176);
		opacity: 0.8;
	}
</style>
