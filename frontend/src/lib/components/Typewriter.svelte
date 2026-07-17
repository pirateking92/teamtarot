<script lang="ts">
	import { onMount } from 'svelte';

	interface Props {
		/** Lines to type in sequence. */
		lines: string[];
		/** Milliseconds per character. */
		speed?: number;
		/** Delay before typing begins. */
		startDelay?: number;
		/** Called once all lines have finished (or the user skips). */
		ondone?: () => void;
	}

	let { lines, speed = 42, startDelay = 300, ondone }: Props = $props();

	let displayed = $state<string[]>(lines.map(() => ''));
	let done = $state(false);
	let cancelled = false;

	const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

	function finish() {
		if (done) return;
		displayed = [...lines];
		done = true;
		ondone?.();
	}

	function skip() {
		cancelled = true;
		finish();
	}

	onMount(() => {
		(async () => {
			await sleep(startDelay);
			for (let i = 0; i < lines.length; i++) {
				for (let c = 0; c <= lines[i].length; c++) {
					if (cancelled) return;
					displayed[i] = lines[i].slice(0, c);
					displayed = [...displayed];
					await sleep(speed);
				}
				await sleep(280);
			}
			finish();
		})();
		return () => {
			cancelled = true;
		};
	});
</script>

<!-- Click anywhere in the block to skip the animation. -->
<div
	class="typewriter"
	role="button"
	tabindex="0"
	onclick={skip}
	onkeydown={(e) => (e.key === 'Enter' || e.key === ' ') && skip()}
>
	{#each displayed as line, i (i)}
		<p class:pending={!done && line.length < lines[i].length}>{line}</p>
	{/each}
</div>

<style>
	.typewriter {
		cursor: pointer;
	}
	.typewriter p {
		margin: 0.35em 0;
		min-height: 1.4em;
	}
	.pending::after {
		content: '▌';
		animation: blink 1s steps(1) infinite;
		opacity: 0.8;
	}
	@keyframes blink {
		50% {
			opacity: 0;
		}
	}
</style>
