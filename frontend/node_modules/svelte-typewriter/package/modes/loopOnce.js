import { animationSetup } from '../helpers/animationSetup'
import { writeEffect } from '../helpers/writeEffect'
import { typingInterval } from '../helpers/typingInterval'
import { unwriteEffect } from '../helpers/unwriteEffect'
import { runOnEveryParentUntil } from '../helpers/runOnEveryParentUntil'

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
const loopOnce = async (node, props) => {
	const { options, elements } = animationSetup(node, props)

	for (const element of elements) {
		const { currentNode, text } = element

		await writeEffect(element, options)
		const textWithUnescapedAmpersands = text.replaceAll('&', '&amp;')

		const fullyWritten = currentNode.innerHTML === textWithUnescapedAmpersands

		if (fullyWritten) {
			options.dispatch('done')
			await typingInterval(options.wordInterval)

			const isLastElement = elements.indexOf(element) === elements.length - 1

			if (!isLastElement) {
				await unwriteEffect(currentNode, options)
				runOnEveryParentUntil(currentNode, options.parentElement, element => {
					currentNode === element
						? element.classList.remove('typing')
						: element.classList.remove('finished-typing')
				})
			}

			runOnEveryParentUntil(currentNode, options.parentElement, element => {
				const cursorHasTimeout = typeof options.keepCursorOnFinish === 'number'
				cursorHasTimeout &&
					setTimeout(() => {
						element.classList.replace('typing', 'finished-typing')
					}, options.keepCursorOnFinish)
			})
		}
	}

	return {
		update() {},
		destroy() {}
	}
}

export default loopOnce
