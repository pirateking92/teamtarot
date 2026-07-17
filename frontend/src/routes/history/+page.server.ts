import { redirect } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';
import type { Reading } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	if (!event.locals.user) redirect(303, '/login?redirectTo=/history');
	const res = await apiFetch(event, '/api/readings');
	const { readings } = res.ok ? ((await res.json()) as { readings: Reading[] }) : { readings: [] };
	return { readings };
};
