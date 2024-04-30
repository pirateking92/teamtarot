import { writeAndUnwriteText } from '../helpers/writeAndUnwriteText'
import { getRandomElement } from '../helpers/getRandomElement'
import { animationSetup } from '../helpers/animationSetup'
import { makeNestedStaticElementsVisible } from '../helpers/makeNestedStaticElementsVisible'

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
const loopRandom = async (node, props) => {
	const { options, elements } = animationSetup(node, props)

	while (true) {
		makeNestedStaticElementsVisible(node)
		const element = getRandomElement(elements)
		await writeAndUnwriteText(element, options)
	}

	return {
		update() {},
		destroy() {}
	}
}

export default loopRandom
