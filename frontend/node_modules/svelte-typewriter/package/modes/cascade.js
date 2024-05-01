import { writeEffect } from '../helpers/writeEffect'
import { cleanChildText } from '../helpers/cleanChildText'
import { animationSetup } from '../helpers/animationSetup'

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
const cascade = async (node, props) => {
	const { options, elements } = animationSetup(node, props)

	cleanChildText(elements)

	for (const element of elements) {
		await writeEffect(element, options)

		if (options.keepCursorOnFinish) {
			const isLastElement = elements.indexOf(element) === elements.length - 1
			!isLastElement && element.currentNode.classList.replace('typing', 'finished-typing')
		} else {
			element.currentNode.classList.replace('typing', 'finished-typing')
		}

		const cursorHasTimeout = typeof options.keepCursorOnFinish === 'number'
		cursorHasTimeout &&
			setTimeout(() => {
				element.currentNode.classList.replace('typing', 'finished-typing')
			}, options.keepCursorOnFinish)
	}

	options.dispatch('done')

	return {
		update() {},
		destroy() {}
	}
}

export default cascade
