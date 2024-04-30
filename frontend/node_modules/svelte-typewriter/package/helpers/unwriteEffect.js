import { typingInterval } from './typingInterval'

/** @type {import(types').UnwriteEffect} */
const unwriteEffect = async (currentNode, options) => {
	const text = currentNode.innerHTML.replaceAll('&amp;', '&')
	for (let index = text.length - 1; index >= 0; index--) {
		const letter = text[index]

		letter === '>' && (index = text.lastIndexOf('<', index))

		const letterSizeInBytes = new Blob([letter]).size
		const isSpecialCharacter = letterSizeInBytes > 1

		if (isSpecialCharacter) continue

		currentNode.innerHTML = text.slice(0, index)
		await typingInterval(options.unwriteInterval ? options.unwriteInterval : options.interval)
	}

	currentNode.innerHTML = text
}

export { unwriteEffect }
