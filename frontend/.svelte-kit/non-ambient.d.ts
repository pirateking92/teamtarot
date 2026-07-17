
// this file is generated — do not edit it


declare module "svelte/elements" {
	export interface HTMLAttributes<T> {
		'data-sveltekit-keepfocus'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-noscroll'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-preload-code'?:
			| true
			| ''
			| 'eager'
			| 'viewport'
			| 'hover'
			| 'tap'
			| 'off'
			| undefined
			| null;
		'data-sveltekit-preload-data'?: true | '' | 'hover' | 'tap' | 'off' | undefined | null;
		'data-sveltekit-reload'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-replacestate'?: true | '' | 'off' | undefined | null;
	}
}

export {};


declare module "$app/types" {
	type MatcherParam<M> = M extends (param : string) => param is (infer U extends string) ? U : string;

	export interface AppTypes {
		RouteId(): "/" | "/account" | "/api" | "/api/readings" | "/api/readings/[id]" | "/cards" | "/cards/[nameShort]" | "/daily" | "/history" | "/login" | "/logout" | "/reading" | "/reading/new" | "/reading/[id]" | "/register";
		RouteParams(): {
			"/api/readings/[id]": { id: string };
			"/cards/[nameShort]": { nameShort: string };
			"/reading/[id]": { id: string }
		};
		LayoutParams(): {
			"/": { id?: string | undefined; nameShort?: string | undefined };
			"/account": Record<string, never>;
			"/api": { id?: string | undefined };
			"/api/readings": { id?: string | undefined };
			"/api/readings/[id]": { id: string };
			"/cards": { nameShort?: string | undefined };
			"/cards/[nameShort]": { nameShort: string };
			"/daily": Record<string, never>;
			"/history": Record<string, never>;
			"/login": Record<string, never>;
			"/logout": Record<string, never>;
			"/reading": { id?: string | undefined };
			"/reading/new": Record<string, never>;
			"/reading/[id]": { id: string };
			"/register": Record<string, never>
		};
		Pathname(): "/" | "/account" | `/api/readings/${string}` & {} | "/cards" | `/cards/${string}` & {} | "/daily" | "/history" | "/login" | "/logout" | "/reading/new" | `/reading/${string}` & {} | "/register";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/bg/0-town.jpeg" | "/bg/1-alley.png" | "/bg/2-parlor.png" | "/cards/ar00.jpg" | "/cards/ar01.jpg" | "/cards/ar02.jpg" | "/cards/ar03.jpg" | "/cards/ar04.jpg" | "/cards/ar05.jpg" | "/cards/ar06.jpg" | "/cards/ar07.jpg" | "/cards/ar08.jpg" | "/cards/ar09.jpg" | "/cards/ar10.jpg" | "/cards/ar11.jpg" | "/cards/ar12.jpg" | "/cards/ar13.jpg" | "/cards/ar14.jpg" | "/cards/ar15.jpg" | "/cards/ar16.jpg" | "/cards/ar17.jpg" | "/cards/ar18.jpg" | "/cards/ar19.jpg" | "/cards/ar20.jpg" | "/cards/ar21.jpg" | "/cards/cu02.jpg" | "/cards/cu03.jpg" | "/cards/cu04.jpg" | "/cards/cu05.jpg" | "/cards/cu06.jpg" | "/cards/cu07.jpg" | "/cards/cu08.jpg" | "/cards/cu09.jpg" | "/cards/cu10.jpg" | "/cards/cuac.jpg" | "/cards/cuki.jpg" | "/cards/cukn.jpg" | "/cards/cupa.jpg" | "/cards/cuqu.jpg" | "/cards/pe02.jpg" | "/cards/pe03.jpg" | "/cards/pe04.jpg" | "/cards/pe05.jpg" | "/cards/pe06.jpg" | "/cards/pe07.jpg" | "/cards/pe08.jpg" | "/cards/pe09.jpg" | "/cards/pe10.jpg" | "/cards/peac.jpg" | "/cards/peki.jpg" | "/cards/pekn.jpg" | "/cards/pepa.jpg" | "/cards/pequ.jpg" | "/cards/sw02.jpg" | "/cards/sw03.jpg" | "/cards/sw04.jpg" | "/cards/sw05.jpg" | "/cards/sw06.jpg" | "/cards/sw07.jpg" | "/cards/sw08.jpg" | "/cards/sw09.jpg" | "/cards/sw10.jpg" | "/cards/swac.jpg" | "/cards/swki.jpg" | "/cards/swkn.jpg" | "/cards/swpa.jpg" | "/cards/swqu.jpg" | "/cards/tarot_back.png" | "/cards/tarot_deck.png" | "/cards/wa02.jpg" | "/cards/wa03.jpg" | "/cards/wa04.jpg" | "/cards/wa05.jpg" | "/cards/wa06.jpg" | "/cards/wa07.jpg" | "/cards/wa08.jpg" | "/cards/wa09.jpg" | "/cards/wa10.jpg" | "/cards/waac.jpg" | "/cards/waki.jpg" | "/cards/wakn.jpg" | "/cards/wapa.jpg" | "/cards/waqu.jpg" | "/favicon.png" | string & {};
	}
}