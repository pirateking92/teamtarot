<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
	import type { Snippet } from 'svelte';
	import type { User } from '$lib/types';

	let { data, children }: { data: { user: User | null }; children: Snippet } = $props();

	// Staged background scene by route, reinforcing the narrative progression.
	const scene = $derived.by(() => {
		const p = page.url.pathname;
		if (p === '/reading/new') return 'scene-alley';
		if (p.startsWith('/reading/')) return 'scene-parlor';
		return 'scene-town';
	});
</script>

<div class="scene {scene}" aria-hidden="true"></div>

<header class="site-header">
	<a class="brand" href="/">Cassandra</a>
	<nav>
		<a href="/cards">The Deck</a>
		<a href="/daily">Daily Card</a>
		{#if data.user}
			<a href="/history">My Readings</a>
			<span class="who">{data.user.displayName}</span>
			<form method="POST" action="/logout" class="inline-form">
				<button type="submit" class="linkish">Sign out</button>
			</form>
		{:else}
			<a href="/login">Sign in</a>
			<a href="/register">Join</a>
		{/if}
	</nav>
</header>

<main>
	{@render children()}
</main>

<style>
	.site-header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
		padding: 0.9rem 1.5rem;
		flex-wrap: wrap;
	}
	.brand {
		font-size: 1.6rem;
		letter-spacing: 0.15em;
		text-transform: uppercase;
		text-decoration: none;
		color: rgb(197, 176, 176);
		text-shadow: 0 0 0.5em rgba(207, 49, 5, 0.6);
	}
	nav {
		display: flex;
		align-items: center;
		gap: 1.1rem;
		flex-wrap: wrap;
	}
	nav a,
	.linkish {
		color: rgb(197, 176, 176);
		text-decoration: none;
		font-size: 1rem;
		background: none;
		border: none;
		cursor: pointer;
		font-family: inherit;
		opacity: 0.85;
		transition: opacity 0.2s;
	}
	nav a:hover,
	.linkish:hover {
		opacity: 1;
		text-shadow: 0 0 0.5em rgba(207, 49, 5, 0.6);
	}
	.who {
		opacity: 0.65;
		font-style: italic;
	}
	.inline-form {
		display: inline;
	}
	main {
		max-width: 60rem;
		margin: 0 auto;
		padding: 1rem 1.25rem 4rem;
	}
</style>
