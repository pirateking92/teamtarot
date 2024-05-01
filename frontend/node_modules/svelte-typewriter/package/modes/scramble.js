import { getLettersTimeout } from '../helpers/getLettersTimeout'
import { getRandomLetter } from '../helpers/getRandomLetter'
import { rng } from '../helpers/rng'
import { sleep } from '../helpers/sleep'
import { runOnEveryParentUntil } from '../helpers/runOnEveryParentUntil'
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
const scramble = async (node, props) => {
	const { options, elements } = animationSetup(node, props)

	const timeout = options.scrambleDuration

	await new Promise(resolve => {
		elements.forEach(async ({ currentNode, text }) => {
			let wordLetters = text.split('')
			const lettersTimeout = getLettersTimeout(wordLetters, timeout)
			const startingTime = Date.now()

			runOnEveryParentUntil(currentNode, options.parentElement, element => {
				element.classList.add('finished-typing')
			})

			while (Date.now() - startingTime < timeout) {
				const randomLetterIndex = rng(0, wordLetters.length)
				const randomLetterTimeout = lettersTimeout[randomLetterIndex]
				const isRandomLetterWhitespace = wordLetters[randomLetterIndex] === ' '
				const timeEllapsed = () => Date.now() - startingTime
				const didRandomLetterReachTimeout = () => timeEllapsed() >= randomLetterTimeout

				if (didRandomLetterReachTimeout() || isRandomLetterWhitespace) {
					const letterFinishedAnimation =
						wordLetters[randomLetterIndex] === text[randomLetterIndex]

					if (!letterFinishedAnimation)
						wordLetters[randomLetterIndex] = text[randomLetterIndex]
					else continue
				} else {
					wordLetters[randomLetterIndex] = getRandomLetter()
				}

				const scrambledText = wordLetters.join('')
				currentNode.innerHTML = scrambledText

				const finishedScrambling = scrambledText === text

				const letterInterval = options.scrambleSlowdown
					? Math.round(timeEllapsed() / 100)
					: 1

				await sleep(letterInterval)

				if (finishedScrambling) {
					resolve()
					break
				}
			}

			currentNode.innerHTML = text
		})
	})

	options.dispatch('done')

	return {
		update() {},
		destroy() {}
	}
}

export default scramble
