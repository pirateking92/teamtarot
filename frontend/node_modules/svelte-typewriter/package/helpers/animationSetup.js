import { createEventDispatcher } from 'svelte'

import { getElements } from './getElements'
import { makeNestedStaticElementsVisible } from './makeNestedStaticElementsVisible'

export const animationSetup = (node, props) => {
	const dispatch = createEventDispatcher()
	const options = { parentElement: node, dispatch, ...props }
	const elements = getElements(node, options)

	makeNestedStaticElementsVisible(node)

	return { options, elements }
}
