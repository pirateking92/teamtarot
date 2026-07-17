<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		children: Snippet;
		type?: 'button' | 'submit';
		disabled?: boolean;
		onclick?: (event: MouseEvent) => void;
		class?: string;
	}

	let { children, type = 'button', disabled = false, onclick, class: klass = '' }: Props = $props();
</script>

<button class="ornate-button {klass}" {type} {disabled} {onclick}>
	{@render children()}
</button>
