import { apiFetch } from '$lib/server/api';
import type { Card } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const arcana = event.url.searchParams.get('arcana') ?? '';
	const suit = event.url.searchParams.get('suit') ?? '';

	const qs = new URLSearchParams();
	if (arcana) qs.set('arcana', arcana);
	if (suit) qs.set('suit', suit);
	const query = qs.toString();

	const res = await apiFetch(event, `/api/cards${query ? `?${query}` : ''}`);
	const { cards } = res.ok ? ((await res.json()) as { cards: Card[] }) : { cards: [] };
	return { cards, arcana, suit };
};
