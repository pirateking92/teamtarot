import { fail, redirect } from '@sveltejs/kit';
import { apiFetch, setSessionCookie } from '$lib/server/api';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = ({ locals }) => {
	if (locals.user) redirect(303, '/');
	return {};
};

export const actions: Actions = {
	default: async (event) => {
		const form = await event.request.formData();
		const email = String(form.get('email') ?? '').trim();
		const password = String(form.get('password') ?? '');
		const displayName = String(form.get('displayName') ?? '').trim();

		if (password.length < 8) {
			return fail(400, { error: 'Password must be at least 8 characters.', email, displayName });
		}

		const res = await apiFetch(event, '/api/auth/register', {
			method: 'POST',
			body: JSON.stringify({ email, password, displayName })
		});
		const data = await res.json().catch(() => ({}) as { token?: string; error?: string });
		if (!res.ok || !data.token) {
			return fail(res.status || 500, { error: data.error ?? 'Could not create account.', email, displayName });
		}

		setSessionCookie(event.cookies, data.token);
		redirect(303, '/');
	}
};
