import { fail, redirect } from '@sveltejs/kit';
import { apiFetch, setSessionCookie } from '$lib/server/api';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = ({ locals, url }) => {
	if (locals.user) redirect(303, '/');
	return { redirectTo: url.searchParams.get('redirectTo') ?? '/' };
};

export const actions: Actions = {
	default: async (event) => {
		const form = await event.request.formData();
		const email = String(form.get('email') ?? '').trim();
		const password = String(form.get('password') ?? '');
		const redirectTo = String(form.get('redirectTo') || '/');

		const res = await apiFetch(event, '/api/auth/login', {
			method: 'POST',
			body: JSON.stringify({ email, password })
		});
		const data = await res.json().catch(() => ({}) as { token?: string; error?: string });
		if (!res.ok || !data.token) {
			return fail(res.status || 500, { error: data.error ?? 'Could not sign in.', email });
		}

		setSessionCookie(event.cookies, data.token);
		redirect(303, redirectTo);
	}
};
