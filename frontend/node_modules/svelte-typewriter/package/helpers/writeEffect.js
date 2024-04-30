import { typingInterval } from './typingInterval'
import { runOnEveryParentUntil } from '../helpers/runOnEveryParentUntil'

/** @type {import(types').TypewriterEffectFn} */
const writeEffect = async ({ currentNode, text }, options) => {
	runOnEveryParentUntil(currentNode, options.parentElement, element => {
		const classToAdd = currentNode === element ? 'typing' : 'finished-typing'
		element.classList.add(classToAdd)
	})

	const segmenter = new Intl.Segmenter('en', { granularity: 'grapheme' })
	const segmentedText = [...segmenter.segment(text)]

	let pendingClosingBracket = false

	for (const { segment: letter, index } of segmentedText) {
		if (letter === '>') {
			pendingClosingBracket = false
			continue
		} else if (letter === '<') {
			pendingClosingBracket = true
			continue
		} else if (pendingClosingBracket) {
			continue
		}

		currentNode.innerHTML = text.slice(0, index)
		await typingInterval(options.interval)
	}

	currentNode.innerHTML = text
}

export { writeEffect }
