import { error } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';
import type { DailyCard } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const res = await apiFetch(event, '/api/daily');
	if (!res.ok) error(502, 'Cassandra could not turn today’s card.');
	const { daily } = (await res.json()) as { daily: DailyCard };
	return { daily };
};
