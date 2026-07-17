<script lang="ts">
	import CardThumb from '$lib/components/CardThumb.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	const filters = [
		{ label: 'All', href: '/cards', match: () => !data.arcana && !data.suit },
		{ label: 'Major', href: '/cards?arcana=major', match: () => data.arcana === 'major' },
		{ label: 'Wands', href: '/cards?suit=wands', match: () => data.suit === 'wands' },
		{ label: 'Cups', href: '/cards?suit=cups', match: () => data.suit === 'cups' },
		{ label: 'Swords', href: '/cards?suit=swords', match: () => data.suit === 'swords' },
		{ label: 'Pentacles', href: '/cards?suit=pentacles', match: () => data.suit === 'pentacles' }
	];
</script>

<svelte:head><title>Cassandra — The Deck</title></svelte:head>

<section>
	<header class="panel head">
		<h1>The Deck</h1>
		<p>All seventy-eight cards. Choose one to learn its secrets.</p>
	</header>

	<nav class="filters">
		{#each filters as f (f.label)}
			<a href={f.href} class:active={f.match()}>{f.label}</a>
		{/each}
	</nav>

	<div class="grid">
		{#each data.cards as card (card.id)}
			<CardThumb {card} href={`/cards/${card.nameShort}`} />
		{/each}
	</div>
</section>

<style>
	.head {
		text-align: center;
		padding: 1.25rem;
		margin: 1.5rem 0 1.5rem;
	}
	.head h1 {
		margin: 0 0 0.4rem;
	}
	.head p {
		margin: 0;
		opacity: 0.8;
	}
	.filters {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		gap: 0.75rem;
		margin-bottom: 2rem;
	}
	.filters a {
		color: rgb(197, 176, 176);
		text-decoration: none;
		padding: 0.3rem 0.9rem;
		border: 1px solid rgba(135, 22, 22, 0.5);
		border-radius: 1rem;
		opacity: 0.8;
	}
	.filters a.active,
	.filters a:hover {
		opacity: 1;
		border-color: #cf3105;
		box-shadow: 0 0 0.6em 1px rgba(207, 49, 5, 0.4);
	}
	.grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(7rem, 1fr));
		gap: 1.25rem;
	}
</style>
