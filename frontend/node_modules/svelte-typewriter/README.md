# svelte-typewriter

> A simple and reusable typewriter effect for your Svelte applications

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![npm Version](https://img.shields.io/npm/v/svelte-typewriter)](https://www.npmjs.com/package/svelte-typewriter)
[![npm Downloads](https://img.shields.io/npm/dw/svelte-typewriter)](https://www.npmjs.com/package/svelte-typewriter)
[![MadeWithSvelte.com shield](https://madewithsvelte.com/storage/repo-shields/2074-shield.svg)](https://madewithsvelte.com/p/svelte-typewriter/shield-link)

[DEMO](https://henriquehbr.srht.site/svelte-typewriter)

## Summary

- [Installation](#installation)
- [Usage](#usage)
  - [Component-based approach](#component-based-approach)
  - [Directive-based approach](#directive-based-approach)
- [Props](#props)
  - [Settings](#settings)
    - [Modes](#modes)
  - [Event listeners](#event-listeners)
  - [Child attributes](#child-attributes)
  - [CSS variables](#css-variables)
- [Used by](#used-by)
- [FAQs](#faqs)
  - [UMD and IIFE output formats are not supported for code-splitting builds](#umd-and-iife-output-formats-are-not-supported-for-code-splitting-builds)
  - [Test suite failed to run: SyntaxError: Unexpected token '<'](#test-suite-failed-to-run-syntaxerror-unexpected-token-)
- [Contributing](#contributing)

## Installation

```bash
pnpm i -D svelte-typewriter
```

> `pnpm` is used here just as an example, you can use your package manager of choice

## Usage

### Component-based approach

In order to use this method, you need to import the Svelte component, and wrap
your elements with the `<Typewriter>` component

```svelte
<script>
    import Typewriter from 'svelte-typewriter'
</script>

<Typewriter>
    <h1>Testing the typewriter effect</h1>
</Typewriter>
```

### Directive-based approach

This method relies on [Svelte actions](https://svelte.dev/docs#template-syntax-element-directives-use-action)
(more specifically, the `use:action` directive), in order to animate your
components with this approach, you must import the directive of the animation
mode you want to use and include it as a attribute on your element

```svelte
<script>
    import { concurrent } from 'svelte-typewriter'
</script>

<p use:concurrent={{ interval: 30 }}>Testing the typewriter effect</p>
```

Each mode has it's own directive, which accepts a single object parameter that
can be used to set the animation props (just like on the component-based approach)

There are just a few limitations of this approach to keep in mind:

- For now, there's no way to have the cursor caret on the text being animated
- Event listeners (like `on:done`) won't be triggered
- Depending on the animation mode you're using, some essential animation
properties must be explicitly specified, because they don't have default values
when used on a directive, otherwise the animation won't work properly, those
include `interval`, `wordInterval`, `writeInterval` and `scrambleDuration`

## Props

The `<Typewriter>` component can receive props that allows to manipulate the
behavior of the resulting animation, these props are divided into the following
groups

- **Settings:** general animation properties
- **Event listeners:** functions executed based on the condition of a trigger
- **Child attributes:** child elements animation properties
- **CSS variables:** styling-related properties

### Settings

| Prop | Type | Description | Default |
| --- | --- | --- | --- |
| `mode` | `concurrent`, `cascade`, `loop`, `loopOnce`, `loopRandom` or `scramble` | The animation mode to be used | `concurrent` |
| `interval` | `number` or `array` | The interval (in milliseconds) between each letter, you can also pass a array of distinct intervals to mimic human typing | `30` |
| `cursor` | `boolean` | Enables/disables the cursor on the Typewriter animation | `true` |
| `keepCursorOnFinish` | `number` or `boolean` | Keep the cursor visible (indefinitely, or for a given amount of time, in milliseconds) after the animation has finished | `false` |
| `delay` | `number` | The interval (in milliseconds) before the animation starts | `0` |
| `showCursorOnDelay` | `boolean` | (only usable when `delay` is not 0) Shows the cursor during delay period | `false` |
| `disabled` | `boolean` | Enables/disables the typewriter animation | `false` |
| `element` | `string` | Sets the tag that will be used for the container element | `div` |
| `wordInterval` | `number` | **(`loop`, `loopOnce` and `loopRandom` modes only)** Sets the interval (in milliseconds) between each word | `1500` |
| `unwriteInterval` | `number` | **(`loop`, `loopOnce` and `loopRandom` modes only)** The interval (in milliseconds) between each letter unwrite, if not defined it uses `interval` | `30` |
| `scrambleDuration` | `number` | **(`scramble` mode only)** Sets the duration (in milliseconds) of the scramble animation | `3000` |
| `scrambleSlowdown` | `boolean` | **(`scramble` mode only)** Enables/disables the slowdown effect right before the scramble animation ends (only works in scramble mode) | `true` |

#### Modes

You can control the behavior of the typewriter effect by passing specific props to the `<Typewriter>` component, the table below contains information about all modes:

| Mode | Description |
| --- | --- |
| `concurrent` | Apply animation simultaneously on all elements, as opposed to the sequential animation of `cascade` mode |
| `cascade` | Apply animation on all elements sequentially instead of simultaneously |
| `loop` | Cycles the animation between the children elements of the parent `Typewriter` component |
| `loopOnce` | It's very similar to the `loop` mode, but the animation stops once it reaches the last element |
| `loopRandom` | It's very similar to `loop` mode, but instead of cycling the animation in a linear way, it picks a random child element to animate each time |
| `scramble` | Slowly reveals the a word by continuously randomizing all of it's letters for a specific amount of time |

### Event listeners

| Event | Trigger | Detail |
| --- | --- | --- |
| `on:done` | Is executed at the end of the animation, if used with one of the loop modes, this event will be fired after each word gets written and erased | `"write" \| "unwrite" \| null` |

### Child attributes

| Attribute | Description |
| --- | --- |
| `data-static` | Marks an element as static, excluding it from receiving animations from the `<Typewriter>` component |

### CSS variables

| Variable | Description |
| --- | --- |
| `--cursor-color` | Sets the cursor color (accepts any valid color name, hex code and rgb/rgba values) |
| `--cursor-width` | Sets the cursor width |

## Used by

<div align="center">
    <a href="https://github.com/GeopJr/SveltePress">
        <img src="https://git.sr.ht/~henriquehbr/svelte-typewriter/blob/main/sveltepress-logo.png" width="128" alt="SveltePress logo">
    </a>
    <a href="https://github.com/ThatConference/that.us">
        <img src="https://git.sr.ht/~henriquehbr/svelte-typewriter/blob/main/that-us-logo.png" width="256" alt="That.us logo">
    </a>
    <a href="https://github.com/Fronvo/site">
        <img src="https://git.sr.ht/~henriquehbr/svelte-typewriter/blob/main/fronvo-logo.png" width="256" alt="Fronvo logo">
    </a>
    <a href="https://github.com/PlebeianTech/plebeian-market">
        <img src="https://git.sr.ht/~henriquehbr/svelte-typewriter/blob/main/plebeian-market-logo.jpg" height="128" alt="plebeian.market logo" />
    </a>
    <a href="https://github.com/kishansripada/pada">
        <img src="https://git.sr.ht/~henriquehbr/svelte-typewriter/blob/main/pada-logo.png" height="128" alt="pada logo" />
    </a>
</div>

## FAQs

#### UMD and IIFE output formats are not supported for code-splitting builds.

From version 2.1.17 onwards, this library makes use of dynamic imports, if your
Rollup configuration `output.format` is set to `iife` or `umd`, consider
setting `inlineDynamicImports` to `true`, otherwise, change `output.format` to
something else, like `esm` (for more details, consider checking [#21](https://github.com/henriquehbr/svelte-typewriter/issues/21))

#### Test suite failed to run: SyntaxError: Unexpected token '<'

This happens because Jest cannot parse Svelte syntax right away, it needs to be
transformed by `svelte-jester` first, therefore, we must tell Jest to **NOT**
ignore `svelte-typewriter`, as by default, everything inside `node_modules` is
ignored and parsed as-is without any kind of pre-processing, this can be done
by setting the `transformIgnorePatterns` property on your Jest configuration,
example below:

```js
// jest.config.js
module.exports = {
    transformIgnorePatterns: ["/node_modules/(?!(svelte-typewriter)/)"]
}
```

> More details on [#73](https://github.com/henriquehbr/svelte-typewriter/issues/73)

## Contributing

1. Fork it
2. Create your feature branch: `git checkout -b fix/my-new-bug-fix`
3. Preview your changes by running the `dev` script on `package.json`
4. Commit your changes: `git commit -am 'fix: solve some issue'`
5. Push to the branch: `git push --set-upstream origin fix/my-new-bug-fix`
6. Submit a pull request
