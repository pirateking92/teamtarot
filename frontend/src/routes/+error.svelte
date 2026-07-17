<script lang="ts">
	import { page } from '$app/state';
</script>

<svelte:head><title>Cassandra — {page.status}</title></svelte:head>

<section class="panel err">
	<h1>{page.status}</h1>
	<p>{page.error?.message ?? 'The cards have clouded over.'}</p>
	<a href="/">Return to Cassandra</a>
</section>

<style>
	.err {
		max-width: 30rem;
		margin: 4rem auto;
		padding: 2.5rem;
		text-align: center;
	}
	.err h1 {
		font-size: 3rem;
		margin: 0;
		text-shadow: 0 0 0.5em rgba(207, 49, 5, 0.6);
	}
	.err a {
		color: rgb(197, 176, 176);
		display: inline-block;
		margin-top: 1.5rem;
	}
</style>
