import { error } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';
import type { Card } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const res = await apiFetch(event, `/api/cards/${event.params.nameShort}`);
	if (res.status === 404) error(404, 'No such card in the deck.');
	if (!res.ok) error(502, 'Could not consult the deck.');
	const { card } = (await res.json()) as { card: Card };
	return { card };
};
