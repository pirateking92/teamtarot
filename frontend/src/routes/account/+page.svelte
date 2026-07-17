<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const user = data.user!;
	const joined = new Date(user.createdAt).toLocaleDateString();
</script>

<svelte:head><title>Cassandra — Your Account</title></svelte:head>

<section class="panel account">
	<h1>{user.displayName}</h1>
	<dl>
		<dt>Email</dt>
		<dd>{user.email}</dd>
		<dt>Known to Cassandra since</dt>
		<dd>{joined}</dd>
	</dl>
	<div class="actions">
		<a class="link" href="/history">View my readings</a>
		<form method="POST" action="/logout">
			<button type="submit" class="ornate-button">Sign out</button>
		</form>
	</div>
</section>

<style>
	.account {
		max-width: 30rem;
		margin: 3rem auto;
		padding: 2.25rem;
		text-align: center;
	}
	.account h1 {
		margin-top: 0;
	}
	dl {
		text-align: left;
		margin: 1.5rem auto;
		max-width: 22rem;
	}
	dt {
		opacity: 0.7;
		font-size: 0.85rem;
		margin-top: 0.75rem;
	}
	dd {
		margin: 0.2rem 0 0;
		font-size: 1.1rem;
	}
	.actions {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;
		margin-top: 1.5rem;
	}
	.link {
		color: rgb(197, 176, 176);
		opacity: 0.85;
	}
</style>
