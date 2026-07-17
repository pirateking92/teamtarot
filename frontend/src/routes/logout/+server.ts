import { redirect } from '@sveltejs/kit';
import { apiFetch, clearSessionCookie } from '$lib/server/api';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async (event) => {
	try {
		await apiFetch(event, '/api/auth/logout', { method: 'POST' });
	} catch {
		// best effort — clear locally regardless
	}
	clearSessionCookie(event.cookies);
	redirect(303, '/');
};
