// Server-side helper for talking to the Go backend (the BFF boundary).
// The browser never calls the Go API directly; SvelteKit server code does,
// forwarding the session token so the backend can identify the user.
import { env } from '$env/dynamic/private';
import { dev } from '$app/environment';
import type { Cookies, RequestEvent } from '@sveltejs/kit';

const BACKEND_URL = env.BACKEND_URL ?? 'http://localhost:8082';

export const SESSION_COOKIE = 'session';
const SESSION_MAX_AGE = 60 * 60 * 24 * 30; // 30 days

/** Store the backend session token in a first-party httpOnly cookie. */
export function setSessionCookie(cookies: Cookies, token: string): void {
	cookies.set(SESSION_COOKIE, token, {
		path: '/',
		httpOnly: true,
		sameSite: 'lax',
		secure: !dev,
		maxAge: SESSION_MAX_AGE
	});
}

/** Remove the session cookie (sign out). */
export function clearSessionCookie(cookies: Cookies): void {
	cookies.delete(SESSION_COOKIE, { path: '/' });
}

/** Absolute URL for a backend API path (e.g. "/api/cards"). */
export function backendUrl(path: string): string {
	return `${BACKEND_URL}${path}`;
}

/**
 * Fetch a backend API path from server code, forwarding the current session
 * token as a bearer token. Uses the event's fetch so SvelteKit can track it.
 */
export function apiFetch(
	event: Pick<RequestEvent, 'fetch' | 'cookies'>,
	path: string,
	init: RequestInit = {}
): Promise<Response> {
	const headers = new Headers(init.headers);
	const token = event.cookies.get(SESSION_COOKIE);
	if (token) {
		headers.set('Authorization', `Bearer ${token}`);
	}
	if (init.body && !headers.has('Content-Type')) {
		headers.set('Content-Type', 'application/json');
	}
	return event.fetch(backendUrl(path), { ...init, headers });
}
