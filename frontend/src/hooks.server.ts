import type { Handle } from '@sveltejs/kit';
import { apiFetch } from '$lib/server/api';

// Populate event.locals.user from the session cookie on every request by
// asking the backend who the token belongs to.
export const handle: Handle = async ({ event, resolve }) => {
	event.locals.user = null;

	const token = event.cookies.get('session');
	if (token) {
		try {
			const res = await apiFetch(event, '/api/auth/me');
			if (res.ok) {
				const { user } = await res.json();
				event.locals.user = user;
			}
		} catch {
			// Backend unreachable — treat as logged out rather than erroring the page.
		}
	}

	return resolve(event);
};
