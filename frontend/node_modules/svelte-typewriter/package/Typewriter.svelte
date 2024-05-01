<script>
    import '@formatjs/intl-segmenter/polyfill'

    export let mode = "concurrent"

    // general-purpose props
	export let interval = 30
	export let cursor = true
    export let keepCursorOnFinish = false
	export let delay = 0
    export let showCursorOnDelay = false
    export let disabled = false
    export let element = "div"

    // mode-specific props
    export let scrambleDuration = 3000
    export let scrambleSlowdown = true
	export let unwriteInterval = 30
    export let wordInterval = 1500

    $: isLoopMode = /^loop(Once|Random)?$/.test(mode)

    // these modes stop once all given elements have finished their animations
    // and support the cursor
    $: isFiniteCursorMode = ["concurrent", "cascade", "loopOnce"].includes(mode)

    $: invalidCursorOnFinish = !isFiniteCursorMode && keepCursorOnFinish
    $: invalidCursorOnDelay = delay < 1 && showCursorOnDelay
    $: invalidLoopProps = !isLoopMode && ($$props.unwriteInterval || $$props.wordInterval)
    $: invalidScrambleProps = mode !== "scramble" && ($$props.scrambleDuration || $$props.scrambleSlowdown)
    $: unnecessaryCursorOnFinish = typeof keepCursorOnFinish === 'number' && keepCursorOnFinish < 1

    const modes = {
        concurrent: () => import("./modes/concurrent"),
        cascade: () => import("./modes/cascade"),
        loop: () => import("./modes/loop"),
        loopOnce: () => import("./modes/loopOnce"),
        loopRandom: () => import("./modes/loopRandom"),
        scramble: () => import("./modes/scramble")
    }

    $: invalidCursorOnFinish && console.warn("[svelte-typewriter] The prop 'keepCursorOnFinish' is compatible only with finite modes")
    $: invalidCursorOnDelay && console.warn("[svelte-typewriter] The prop 'showCursorOnDelay' has no effect if the delay is 0")
    $: invalidLoopProps && console.warn("[svelte-typewriter] The props 'unwriteInterval' and 'wordInterval' are only compatible with loop modes")
    $: invalidScrambleProps && console.warn("[svelte-typewriter] The props 'scrambleDuration' and 'scrambleSlowdown' are only compatible with scramble mode")
    $: unnecessaryCursorOnFinish && console.warn("[svelte-typewriter] The prop 'keepCursorOnFinish' has no effect with values lower than 1")

    $: delayPromise = () => new Promise(resolve => setTimeout(() => resolve(delay), delay))

    $: props = {
        interval,
        cursor,
        keepCursorOnFinish,
        delay,
        showCursorOnDelay,
        disabled,
        element,
        scrambleDuration,
        scrambleSlowdown,
        unwriteInterval,
        wordInterval
    }
</script>

<style>
	@keyframes cursorFade {
		0%,
		100% {
			opacity: 1;
		}

		50% {
			opacity: 0;
		}
	}

	.typewriter-container :global(*:not(.typing):not(.finished-typing):not([data-static])) {
		display: none;
	}

	.typewriter-container :global(.finished-typing::after) {
		content: none;
	}

    .cursor :global(.typing::after) {
        content: '';
        width: var(--cursor-width, 1ch);
        height: 2ch;
        display: inline-block;
        vertical-align: text-top;
        background-color: var(--cursor-color, #000000);
        animation: cursorFade 1.25s infinite;
    }
</style>

<noscript>
    <slot />
</noscript>

{#key $$props}
    {#if disabled}
        <div class="typewriter-container">
            <slot />
        </div>
    {:else}
        {#await delayPromise()}
            {#if showCursorOnDelay}
                <div class="typewriter-container cursor">
                    <p class="typing"></p>
                </div>
            {/if}
        {:then}
            {#await modes[mode]() then selectedMode}
                <svelte:element this={element} use:selectedMode.default={props} class:cursor class="typewriter-container">
                    <slot />
                </svelte:element>
            {/await}
        {/await}
    {/if}
{/key}
