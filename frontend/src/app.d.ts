import type { User } from '$lib/types';

// See https://svelte.dev/docs/kit/types#app.d.ts
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null;
		}
		interface PageData {
			user?: User | null;
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
