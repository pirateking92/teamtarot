<script lang="ts">
	import { goto } from '$app/navigation';
	import Typewriter from '$lib/components/Typewriter.svelte';
	import OrnateButton from '$lib/components/OrnateButton.svelte';

	let name = $state('');
	let ready = $state(false);

	function begin() {
		const trimmed = name.trim();
		const q = trimmed ? `?name=${encodeURIComponent(trimmed)}` : '';
		goto(`/reading/new${q}`);
	}
</script>

<svelte:head>
	<title>Cassandra — Tarot Readings</title>
</svelte:head>

<section class="panel intro">
	<div class="intro-text">
		<Typewriter
			lines={[
				"Ah, I've been expecting you, seeker.",
				'My name is Cassandra.',
				"What's yours?"
			]}
			ondone={() => (ready = true)}
		/>

		{#if ready}
			<form
				class="entry"
				onsubmit={(e) => {
					e.preventDefault();
					begin();
				}}
			>
				<input class="input-field" placeholder="Enter your name" bind:value={name} />
				<OrnateButton type="submit">Follow Cassandra</OrnateButton>
			</form>
		{/if}
	</div>
</section>

<style>
	.intro {
		margin: 3rem auto;
		padding: 2.5rem 3rem;
		max-width: 42rem;
		text-align: center;
		font-size: 1.5rem;
		font-weight: 600;
	}
	.entry {
		margin-top: 2rem;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.5rem;
	}
	.entry .input-field {
		min-width: 18rem;
	}
</style>
