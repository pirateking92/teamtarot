import { getLongestTextElement } from '../helpers/getLongestTextElement'
import { writeEffect } from '../helpers/writeEffect'
import { onAnimationEnd } from '../helpers/onAnimationEnd'
import { animationSetup } from '../helpers/animationSetup'

// the name "default" cannot be used due to being a js keyword

/**
 * @typedef {object} Props
 * @property {number} [interval]
 * @property {boolean} [cursor]
 * @property {boolean} [keepCursorOnFinish]
 * @property {number} [delay]
 * @property {boolean} [showCursorOnDelay]
 * @property {boolean} [disabled]
 * @property {string} [element]
 * @property {number} [scrambleDuration]
 * @property {number} [scrambleSlowdown]
 * @property {number} [unwriteInterval]
 * @property {number} [wordInterval]
 */

/**
 * @typedef {{ update: () => void, destroy: () => void }} SvelteActionReturnType
 */

/**
 * @typedef {(node: HTMLElement, props: Props) => SvelteActionReturnType} TypewriterModeFn
 */

/**
 * @type {TypewriterModeFn}
 */
const concurrent = (node, props) => {
	const { options, elements } = animationSetup(node, props)

	const lastElementToFinish = getLongestTextElement(elements)
	onAnimationEnd(lastElementToFinish, () => options.dispatch('done'))

	for (const element of elements) {
		// "then" is required here to prevent blocking execution, thus keeping
		// the animation asynchronous
		writeEffect(element, options).then(() => {
			if (options.keepCursorOnFinish) {
				const isNotLongestElement = element.currentNode !== lastElementToFinish
				isNotLongestElement
					? element.currentNode.classList.replace('typing', 'finished-typing')
					: options.dispatch('done')
			} else {
				element.currentNode.classList.replace('typing', 'finished-typing')
			}

			const cursorHasTimeout = typeof options.keepCursorOnFinish === 'number'
			cursorHasTimeout &&
				setTimeout(() => {
					element.currentNode.classList.replace('typing', 'finished-typing')
				}, options.keepCursorOnFinish)
		})
	}

	return {
		update() {},
		destroy() {}
	}
}

export default concurrent
