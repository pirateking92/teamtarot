<script lang="ts">
	import OrnateButton from '$lib/components/OrnateButton.svelte';
	import type { PageData, ActionData } from './$types';

	let { data, form }: { data: PageData; form: ActionData } = $props();
</script>

<svelte:head><title>Cassandra — Sign In</title></svelte:head>

<section class="panel auth">
	<h1>Welcome back, seeker</h1>
	<form method="POST">
		<input type="hidden" name="redirectTo" value={data.redirectTo} />
		<label>
			<span>Email</span>
			<input class="input-field" type="email" name="email" value={form?.email ?? ''} required />
		</label>
		<label>
			<span>Password</span>
			<input class="input-field" type="password" name="password" required />
		</label>
		{#if form?.error}<p class="error">{form.error}</p>{/if}
		<div class="row"><OrnateButton type="submit">Sign in</OrnateButton></div>
	</form>
	<p class="alt">No account yet? <a href="/register">Join Cassandra</a></p>
</section>

<style>
	.auth {
		max-width: 26rem;
		margin: 3rem auto;
		padding: 2.25rem;
		text-align: center;
	}
	.auth h1 {
		margin-top: 0;
		font-size: 1.5rem;
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
