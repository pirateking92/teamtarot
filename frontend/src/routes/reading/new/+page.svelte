<script lang="ts">
	import { enhance } from '$app/forms';
	import Typewriter from '$lib/components/Typewriter.svelte';
	import OrnateButton from '$lib/components/OrnateButton.svelte';
	import type { PageData, ActionData } from './$types';

	let { data, form }: { data: PageData; form: ActionData } = $props();

	let ready = $state(false);
	let spreadType = $state(form?.spreadType ?? 'past-present-future');
	let submitting = $state(false);
</script>

<svelte:head>
	<title>Cassandra — A New Reading</title>
</svelte:head>

<section class="panel new-reading">
	<div class="intro">
		<Typewriter
			lines={['Sit with me a while.', 'Tell me what weighs on you, and choose how the cards shall fall.']}
			ondone={() => (ready = true)}
		/>
	</div>

	{#if ready}
		<form method="POST" use:enhance={() => {
			submitting = true;
			return async ({ update }) => {
				await update();
				submitting = false;
			};
		}}>
			<label class="field">
				<span>Your name</span>
				<input
					class="input-field"
					name="querentName"
					placeholder="Enter your name"
					value={form?.querentName ?? data.name}
				/>
			</label>

			<fieldset class="spreads">
				<legend>Choose your spread</legend>
				{#each data.spreads as s (s.type)}
					<label class="spread-option" class:selected={spreadType === s.type}>
						<input type="radio" name="spreadType" value={s.type} bind:group={spreadType} />
						<span class="spread-name">{s.name}</span>
						<span class="spread-count">{s.positions.length} card{s.positions.length > 1 ? 's' : ''}</span>
						<span class="spread-desc">{s.description}</span>
					</label>
				{/each}
			</fieldset>

			<label class="field">
				<span>Your story <em>(optional)</em></span>
				<textarea
					class="input-field story"
					name="story"
					rows="4"
					placeholder="What's on your mind, seeker?">{form?.story ?? ''}</textarea>
			</label>

			{#if form?.error}
				<p class="error">{form.error}</p>
			{/if}

			<div class="submit-row">
				<OrnateButton type="submit" disabled={submitting}>
					{submitting ? 'Consulting the cards…' : 'Draw the cards'}
				</OrnateButton>
			</div>
		</form>
	{/if}
</section>

<style>
	.new-reading {
		margin: 2.5rem auto;
		padding: 2.5rem;
		max-width: 44rem;
	}
	.intro {
		text-align: center;
		font-size: 1.3rem;
		margin-bottom: 1.5rem;
	}
	.field {
		display: block;
		margin: 1.25rem 0;
	}
	.field > span {
		display: block;
		margin-bottom: 0.4rem;
		font-size: 1.05rem;
	}
	.field .input-field {
		width: 100%;
		text-align: left;
	}
	.story {
		resize: vertical;
	}
	.spreads {
		border: none;
		padding: 0;
		margin: 1.5rem 0;
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(12rem, 1fr));
		gap: 0.9rem;
	}
	.spreads legend {
		font-size: 1.05rem;
		margin-bottom: 0.6rem;
	}
	.spread-option {
		display: grid;
		gap: 0.2rem;
		padding: 0.9rem 1rem;
		border: 2px solid rgba(135, 22, 22, 0.5);
		border-radius: 0.9rem;
		cursor: pointer;
		background: rgba(19, 5, 0, 0.5);
		transition:
			border-color 0.2s,
			box-shadow 0.2s;
	}
	.spread-option:hover {
		border-color: #871616;
	}
	.spread-option.selected {
		border-color: #cf3105;
		box-shadow: 0 0 0.8em 2px rgba(207, 49, 5, 0.5);
	}
	.spread-option input {
		position: absolute;
		opacity: 0;
		pointer-events: none;
	}
	.spread-name {
		font-size: 1.1rem;
	}
	.spread-count {
		font-size: 0.85rem;
		opacity: 0.7;
	}
	.spread-desc {
		font-size: 0.82rem;
		opacity: 0.75;
		line-height: 1.3;
	}
	.error {
		color: #ff9b7a;
		text-align: center;
	}
	.submit-row {
		text-align: center;
		margin-top: 1.75rem;
	}
</style>
