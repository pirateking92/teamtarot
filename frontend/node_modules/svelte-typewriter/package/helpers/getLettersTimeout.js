import { rng } from './rng'

// returns a array with a timeout (in ms) for each letter of the word
export const getLettersTimeout = (textLetters, timeout) => {
	const minimumTimeoutPossible = timeout / 3
	// TODO: find a better way to deal with this instead of explicitly reducing the maximum timeout
	// otherwise, at the end of the animation, one or two characters remain scrambled
	const lettersTimeout = textLetters.map(() => rng(minimumTimeoutPossible, timeout - 100))
	return lettersTimeout
}
