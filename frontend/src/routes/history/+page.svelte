<script lang="ts">
	import { cardImage } from '$lib/types';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	function when(iso: string): string {
		return new Date(iso).toLocaleString(undefined, {
			dateStyle: 'medium',
			timeStyle: 'short'
		});
	}
</script>

<svelte:head><title>Cassandra — My Readings</title></svelte:head>

<section>
	<header class="panel head">
		<h1>My Readings</h1>
	</header>

	{#if data.readings.length === 0}
		<p class="empty">
			You have no readings yet. <a href="/reading/new">Draw your first.</a>
		</p>
	{:else}
		<ul class="list">
			{#each data.readings as r (r.id)}
				<li>
					<a class="panel entry" href={`/reading/${r.id}`}>
						<div class="mini">
							{#each r.cards as rc (rc.positionIndex)}
								<img src={cardImage(rc.card)} alt={rc.card.name} class:reversed={rc.reversed} />
							{/each}
						</div>
						<div class="meta">
							<span class="names">{r.cards.map((c) => c.card.name).join(' · ')}</span>
							<span class="sub">{when(r.createdAt)}</span>
							{#if r.status !== 'complete'}
								<span class="badge">{r.status}</span>
							{/if}
						</div>
					</a>
				</li>
			{/each}
		</ul>
	{/if}
</section>

<style>
	.head {
		text-align: center;
		padding: 1.25rem;
		margin: 1.5rem 0;
	}
	.head h1 {
		margin: 0;
	}
	.empty {
		text-align: center;
		opacity: 0.85;
	}
	.empty a {
		color: rgb(197, 176, 176);
	}
	.list {
		list-style: none;
		padding: 0;
		margin: 0;
		display: grid;
		gap: 1rem;
	}
	.entry {
		display: flex;
		gap: 1rem;
		align-items: center;
		padding: 0.9rem 1.1rem;
		text-decoration: none;
		color: rgb(197, 176, 176);
	}
	.mini {
		display: flex;
		gap: 0.3rem;
		flex-shrink: 0;
	}
	.mini img {
		width: 2.2rem;
		aspect-ratio: 5 / 8.6;
		object-fit: cover;
		border-radius: 0.25rem;
	}
	.mini .reversed {
		transform: rotate(180deg);
	}
	.meta {
		display: flex;
		flex-direction: column;
		gap: 0.2rem;
	}
	.names {
		font-size: 1rem;
	}
	.sub {
		opacity: 0.65;
		font-size: 0.85rem;
	}
	.badge {
		align-self: flex-start;
		font-size: 0.72rem;
		text-transform: uppercase;
		opacity: 0.8;
		border: 1px solid rgba(207, 49, 5, 0.5);
		border-radius: 0.6rem;
		padding: 0.05rem 0.4rem;
	}
</style>
