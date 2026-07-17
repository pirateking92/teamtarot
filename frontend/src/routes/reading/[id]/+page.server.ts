import { error } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';
import type { Reading } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const res = await apiFetch(event, `/api/readings/${event.params.id}`);
	if (res.status === 404) {
		error(404, 'That reading has faded from Cassandra’s memory.');
	}
	if (!res.ok) {
		error(502, 'Cassandra could not recall that reading.');
	}
	const { reading } = (await res.json()) as { reading: Reading };
	return { reading };
};
