import { writeEffect } from '../helpers/writeEffect'
import { unwriteEffect } from '../helpers/unwriteEffect'
import { runOnEveryParentUntil } from '../helpers/runOnEveryParentUntil'
import { typingInterval } from './typingInterval'

const writeAndUnwriteText = async ({ currentNode, text }, options) => {
	await writeEffect({ currentNode, text }, options)
	const textWithUnescapedAmpersands = text.replaceAll('&', '&amp;')

	const fullyWritten = currentNode.innerHTML === textWithUnescapedAmpersands

	if (fullyWritten) {
		options.dispatch('done', 'write')
		await typingInterval(options.wordInterval)
		await unwriteEffect(currentNode, options)
		options.dispatch('done', 'unwrite')
	}

	runOnEveryParentUntil(currentNode, options.parentElement, element => {
		currentNode === element
			? element.classList.remove('typing')
			: element.classList.remove('finished-typing')
	})
}

export { writeAndUnwriteText }
