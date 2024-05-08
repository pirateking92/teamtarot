import { c as create_ssr_component, i as is_promise, n as noop } from "./ssr.js";
import "@formatjs/intl-segmenter/polyfill.js";
/* empty css                                         */
const void_element_names = /^(?:area|base|br|col|command|embed|hr|img|input|keygen|link|meta|param|source|track|wbr)$/;
function is_void(name) {
  return void_element_names.test(name) || name.toLowerCase() === "!doctype";
}
const css = {
  code: "@keyframes svelte-1gv2i7t-cursorFade{0%,100%{opacity:1}50%{opacity:0}}.typewriter-container.svelte-1gv2i7t *:not(.typing):not(.finished-typing):not([data-static]){display:none}.typewriter-container.svelte-1gv2i7t .finished-typing::after{content:none}.cursor.svelte-1gv2i7t .typing::after{content:'';width:var(--cursor-width, 1ch);height:2ch;display:inline-block;vertical-align:text-top;background-color:var(--cursor-color, #000000);animation:svelte-1gv2i7t-cursorFade 1.25s infinite}",
  map: `{"version":3,"file":"Typewriter.svelte","sources":["Typewriter.svelte"],"sourcesContent":["<script>\\n    import '@formatjs/intl-segmenter/polyfill'\\n\\n    export let mode = \\"concurrent\\"\\n\\n    // general-purpose props\\n\\texport let interval = 30\\n\\texport let cursor = true\\n    export let keepCursorOnFinish = false\\n\\texport let delay = 0\\n    export let showCursorOnDelay = false\\n    export let disabled = false\\n    export let element = \\"div\\"\\n\\n    // mode-specific props\\n    export let scrambleDuration = 3000\\n    export let scrambleSlowdown = true\\n\\texport let unwriteInterval = 30\\n    export let wordInterval = 1500\\n\\n    $: isLoopMode = /^loop(Once|Random)?$/.test(mode)\\n\\n    // these modes stop once all given elements have finished their animations\\n    // and support the cursor\\n    $: isFiniteCursorMode = [\\"concurrent\\", \\"cascade\\", \\"loopOnce\\"].includes(mode)\\n\\n    $: invalidCursorOnFinish = !isFiniteCursorMode && keepCursorOnFinish\\n    $: invalidCursorOnDelay = delay < 1 && showCursorOnDelay\\n    $: invalidLoopProps = !isLoopMode && ($$props.unwriteInterval || $$props.wordInterval)\\n    $: invalidScrambleProps = mode !== \\"scramble\\" && ($$props.scrambleDuration || $$props.scrambleSlowdown)\\n    $: unnecessaryCursorOnFinish = typeof keepCursorOnFinish === 'number' && keepCursorOnFinish < 1\\n\\n    const modes = {\\n        concurrent: () => import(\\"./modes/concurrent\\"),\\n        cascade: () => import(\\"./modes/cascade\\"),\\n        loop: () => import(\\"./modes/loop\\"),\\n        loopOnce: () => import(\\"./modes/loopOnce\\"),\\n        loopRandom: () => import(\\"./modes/loopRandom\\"),\\n        scramble: () => import(\\"./modes/scramble\\")\\n    }\\n\\n    $: invalidCursorOnFinish && console.warn(\\"[svelte-typewriter] The prop 'keepCursorOnFinish' is compatible only with finite modes\\")\\n    $: invalidCursorOnDelay && console.warn(\\"[svelte-typewriter] The prop 'showCursorOnDelay' has no effect if the delay is 0\\")\\n    $: invalidLoopProps && console.warn(\\"[svelte-typewriter] The props 'unwriteInterval' and 'wordInterval' are only compatible with loop modes\\")\\n    $: invalidScrambleProps && console.warn(\\"[svelte-typewriter] The props 'scrambleDuration' and 'scrambleSlowdown' are only compatible with scramble mode\\")\\n    $: unnecessaryCursorOnFinish && console.warn(\\"[svelte-typewriter] The prop 'keepCursorOnFinish' has no effect with values lower than 1\\")\\n\\n    $: delayPromise = () => new Promise(resolve => setTimeout(() => resolve(delay), delay))\\n\\n    $: props = {\\n        interval,\\n        cursor,\\n        keepCursorOnFinish,\\n        delay,\\n        showCursorOnDelay,\\n        disabled,\\n        element,\\n        scrambleDuration,\\n        scrambleSlowdown,\\n        unwriteInterval,\\n        wordInterval\\n    }\\n<\/script>\\n\\n<style>\\n\\t@keyframes cursorFade {\\n\\t\\t0%,\\n\\t\\t100% {\\n\\t\\t\\topacity: 1;\\n\\t\\t}\\n\\n\\t\\t50% {\\n\\t\\t\\topacity: 0;\\n\\t\\t}\\n\\t}\\n\\n\\t.typewriter-container :global(*:not(.typing):not(.finished-typing):not([data-static])) {\\n\\t\\tdisplay: none;\\n\\t}\\n\\n\\t.typewriter-container :global(.finished-typing::after) {\\n\\t\\tcontent: none;\\n\\t}\\n\\n    .cursor :global(.typing::after) {\\n        content: '';\\n        width: var(--cursor-width, 1ch);\\n        height: 2ch;\\n        display: inline-block;\\n        vertical-align: text-top;\\n        background-color: var(--cursor-color, #000000);\\n        animation: cursorFade 1.25s infinite;\\n    }\\n</style>\\n\\n<noscript>\\n    <slot />\\n</noscript>\\n\\n{#key $$props}\\n    {#if disabled}\\n        <div class=\\"typewriter-container\\">\\n            <slot />\\n        </div>\\n    {:else}\\n        {#await delayPromise()}\\n            {#if showCursorOnDelay}\\n                <div class=\\"typewriter-container cursor\\">\\n                    <p class=\\"typing\\"></p>\\n                </div>\\n            {/if}\\n        {:then}\\n            {#await modes[mode]() then selectedMode}\\n                <svelte:element this={element} use:selectedMode.default={props} class:cursor class=\\"typewriter-container\\">\\n                    <slot />\\n                </svelte:element>\\n            {/await}\\n        {/await}\\n    {/if}\\n{/key}\\n"],"names":[],"mappings":"AAiEC,WAAW,yBAAW,CACrB,EAAE,CACF,IAAK,CACJ,OAAO,CAAE,CACV,CAEA,GAAI,CACH,OAAO,CAAE,CACV,CACD,CAEA,oCAAqB,CAAS,uDAAyD,CACtF,OAAO,CAAE,IACV,CAEA,oCAAqB,CAAS,uBAAyB,CACtD,OAAO,CAAE,IACV,CAEG,sBAAO,CAAS,cAAgB,CAC5B,OAAO,CAAE,EAAE,CACX,KAAK,CAAE,IAAI,cAAc,CAAC,IAAI,CAAC,CAC/B,MAAM,CAAE,GAAG,CACX,OAAO,CAAE,YAAY,CACrB,cAAc,CAAE,QAAQ,CACxB,gBAAgB,CAAE,IAAI,cAAc,CAAC,QAAQ,CAAC,CAC9C,SAAS,CAAE,yBAAU,CAAC,KAAK,CAAC,QAChC"}`
};
const Typewriter = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let isLoopMode;
  let isFiniteCursorMode;
  let invalidCursorOnFinish;
  let invalidCursorOnDelay;
  let invalidLoopProps;
  let invalidScrambleProps;
  let unnecessaryCursorOnFinish;
  let delayPromise;
  let { mode = "concurrent" } = $$props;
  let { interval = 30 } = $$props;
  let { cursor = true } = $$props;
  let { keepCursorOnFinish = false } = $$props;
  let { delay = 0 } = $$props;
  let { showCursorOnDelay = false } = $$props;
  let { disabled = false } = $$props;
  let { element = "div" } = $$props;
  let { scrambleDuration = 3e3 } = $$props;
  let { scrambleSlowdown = true } = $$props;
  let { unwriteInterval = 30 } = $$props;
  let { wordInterval = 1500 } = $$props;
  const modes = {
    concurrent: () => import("./concurrent.js"),
    cascade: () => import("./cascade.js"),
    loop: () => import("./loop.js"),
    loopOnce: () => import("./loopOnce.js"),
    loopRandom: () => import("./loopRandom.js"),
    scramble: () => import("./scramble.js")
  };
  if ($$props.mode === void 0 && $$bindings.mode && mode !== void 0)
    $$bindings.mode(mode);
  if ($$props.interval === void 0 && $$bindings.interval && interval !== void 0)
    $$bindings.interval(interval);
  if ($$props.cursor === void 0 && $$bindings.cursor && cursor !== void 0)
    $$bindings.cursor(cursor);
  if ($$props.keepCursorOnFinish === void 0 && $$bindings.keepCursorOnFinish && keepCursorOnFinish !== void 0)
    $$bindings.keepCursorOnFinish(keepCursorOnFinish);
  if ($$props.delay === void 0 && $$bindings.delay && delay !== void 0)
    $$bindings.delay(delay);
  if ($$props.showCursorOnDelay === void 0 && $$bindings.showCursorOnDelay && showCursorOnDelay !== void 0)
    $$bindings.showCursorOnDelay(showCursorOnDelay);
  if ($$props.disabled === void 0 && $$bindings.disabled && disabled !== void 0)
    $$bindings.disabled(disabled);
  if ($$props.element === void 0 && $$bindings.element && element !== void 0)
    $$bindings.element(element);
  if ($$props.scrambleDuration === void 0 && $$bindings.scrambleDuration && scrambleDuration !== void 0)
    $$bindings.scrambleDuration(scrambleDuration);
  if ($$props.scrambleSlowdown === void 0 && $$bindings.scrambleSlowdown && scrambleSlowdown !== void 0)
    $$bindings.scrambleSlowdown(scrambleSlowdown);
  if ($$props.unwriteInterval === void 0 && $$bindings.unwriteInterval && unwriteInterval !== void 0)
    $$bindings.unwriteInterval(unwriteInterval);
  if ($$props.wordInterval === void 0 && $$bindings.wordInterval && wordInterval !== void 0)
    $$bindings.wordInterval(wordInterval);
  $$result.css.add(css);
  isLoopMode = /^loop(Once|Random)?$/.test(mode);
  isFiniteCursorMode = ["concurrent", "cascade", "loopOnce"].includes(mode);
  invalidCursorOnFinish = !isFiniteCursorMode && keepCursorOnFinish;
  invalidCursorOnDelay = delay < 1 && showCursorOnDelay;
  invalidLoopProps = !isLoopMode && ($$props.unwriteInterval || $$props.wordInterval);
  invalidScrambleProps = mode !== "scramble" && ($$props.scrambleDuration || $$props.scrambleSlowdown);
  unnecessaryCursorOnFinish = typeof keepCursorOnFinish === "number" && keepCursorOnFinish < 1;
  invalidCursorOnFinish && console.warn("[svelte-typewriter] The prop 'keepCursorOnFinish' is compatible only with finite modes");
  invalidCursorOnDelay && console.warn("[svelte-typewriter] The prop 'showCursorOnDelay' has no effect if the delay is 0");
  invalidLoopProps && console.warn("[svelte-typewriter] The props 'unwriteInterval' and 'wordInterval' are only compatible with loop modes");
  invalidScrambleProps && console.warn("[svelte-typewriter] The props 'scrambleDuration' and 'scrambleSlowdown' are only compatible with scramble mode");
  unnecessaryCursorOnFinish && console.warn("[svelte-typewriter] The prop 'keepCursorOnFinish' has no effect with values lower than 1");
  delayPromise = () => new Promise((resolve) => setTimeout(() => resolve(delay), delay));
  return `<noscript>${slots.default ? slots.default({}) : ``}</noscript> ${disabled ? `<div class="typewriter-container svelte-1gv2i7t">${slots.default ? slots.default({}) : ``}</div>` : `${function(__value) {
    if (is_promise(__value)) {
      __value.then(null, noop);
      return ` ${showCursorOnDelay ? `<div class="typewriter-container cursor svelte-1gv2i7t" data-svelte-h="svelte-1ssf3hl"><p class="typing"></p></div>` : ``} `;
    }
    return function() {
      return ` ${function(__value2) {
        if (is_promise(__value2)) {
          __value2.then(null, noop);
          return ``;
        }
        return function(selectedMode) {
          return ` ${((tag) => {
            return tag ? `<${element} class="${["typewriter-container svelte-1gv2i7t", cursor ? "cursor" : ""].join(" ").trim()}">${is_void(tag) ? "" : `${slots.default ? slots.default({}) : ``}`}${is_void(tag) ? "" : `</${tag}>`}` : "";
          })(element)} `;
        }();
      }(modes[mode]())} `;
    }();
  }(delayPromise())}`}`;
});
export {
  Typewriter as T
};
