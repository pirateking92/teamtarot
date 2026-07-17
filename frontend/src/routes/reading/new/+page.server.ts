import { fail, redirect } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';
import type { Spread } from '$lib/types';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const res = await apiFetch(event, '/api/spreads');
	if (!res.ok) {
		return { spreads: [] as Spread[], name: event.url.searchParams.get('name') ?? '' };
	}
	const { spreads } = (await res.json()) as { spreads: Spread[] };
	return { spreads, name: event.url.searchParams.get('name') ?? '' };
};

export const actions: Actions = {
	default: async (event) => {
		const form = await event.request.formData();
		const querentName = String(form.get('querentName') ?? '').trim();
		const story = String(form.get('story') ?? '').trim();
		const spreadType = String(form.get('spreadType') ?? 'past-present-future');

		const res = await apiFetch(event, '/api/readings', {
			method: 'POST',
			body: JSON.stringify({ querentName, story, spreadType })
		});
		if (!res.ok) {
			return fail(502, {
				error: 'Cassandra could not reach the cards. Please try again.',
				querentName,
				story,
				spreadType
			});
		}
		const { reading } = (await res.json()) as { reading: { id: string } };
		redirect(303, `/reading/${reading.id}`);
	}
};
