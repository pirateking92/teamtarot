import { json } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';
import type { RequestHandler } from './$types';

// Same-origin proxy so the browser can poll a reading's interpretation status
// without talking to the Go backend directly.
export const GET: RequestHandler = async (event) => {
	const res = await apiFetch(event, `/api/readings/${event.params.id}`);
	const data = await res.json();
	return json(data, { status: res.status });
};
