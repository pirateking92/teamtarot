import { rng } from './rng'

export const getRandomLetter = () => {
	const possibleLetters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'.split(
		''
	)
	const letterIndexLimit = possibleLetters.length
	const randomLetterIndex = rng(0, letterIndexLimit)
	const randomLetter = possibleLetters[randomLetterIndex]
	return randomLetter
}
