<script lang="ts">
	import OrnateButton from '$lib/components/OrnateButton.svelte';
	import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
</script>

<svelte:head><title>Cassandra — Join</title></svelte:head>

<section class="panel auth">
	<h1>Bind your fate to Cassandra</h1>
	<p class="sub">An account lets you keep your readings. You can always read as a guest.</p>
	<form method="POST">
		<label>
			<span>Name Cassandra should call you</span>
			<input class="input-field" name="displayName" value={form?.displayName ?? ''} />
		</label>
		<label>
			<span>Email</span>
			<input class="input-field" type="email" name="email" value={form?.email ?? ''} required />
		</label>
		<label>
			<span>Password <em>(8+ characters)</em></span>
			<input class="input-field" type="password" name="password" required minlength="8" />
		</label>
		{#if form?.error}<p class="error">{form.error}</p>{/if}
		<div class="row"><OrnateButton type="submit">Create account</OrnateButton></div>
	</form>
	<p class="alt">Already known to Cassandra? <a href="/login">Sign in</a></p>
</section>

<style>
	.auth {
		max-width: 28rem;
		margin: 3rem auto;
		padding: 2.25rem;
		text-align: center;
	}
	.auth h1 {
		margin-top: 0;
		font-size: 1.5rem;
	}
	.sub {
		opacity: 0.8;
		margin-top: 0;
	}
	label {
		display: block;
		text-align: left;
		margin: 1rem 0;
	}
	label span {
		display: block;
		margin-bottom: 0.35rem;
	}
	label em {
		opacity: 0.65;
	}
	label .input-field {
		width: 100%;
		text-align: left;
	}
	.row {
		margin-top: 1.5rem;
	}
	.error {
		color: #ff9b7a;
	}
	.alt {
		margin-top: 1.5rem;
		opacity: 0.8;
	}
	.alt a {
		color: rgb(197, 176, 176);
	}
</style>
