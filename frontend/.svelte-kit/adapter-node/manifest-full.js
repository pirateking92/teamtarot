export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["assets/.DS_Store","assets/0-town.jpeg","assets/1-alley.png","assets/2-parlor.png","assets/ar00.jpg","assets/ar01.jpg","assets/ar02.jpg","assets/ar03.jpg","assets/ar04.jpg","assets/ar05.jpg","assets/ar06.jpg","assets/ar07.jpg","assets/ar08.jpg","assets/ar09.jpg","assets/ar10.jpg","assets/ar11.jpg","assets/ar12.jpg","assets/ar13.jpg","assets/ar14.jpg","assets/ar15.jpg","assets/ar16.jpg","assets/ar17.jpg","assets/ar18.jpg","assets/ar19.jpg","assets/ar20.jpg","assets/ar21.jpg","assets/cu02.jpg","assets/cu03.jpg","assets/cu04.jpg","assets/cu05.jpg","assets/cu06.jpg","assets/cu07.jpg","assets/cu08.jpg","assets/cu09.jpg","assets/cu10.jpg","assets/cuac.jpg","assets/cuki.jpg","assets/cukn.jpg","assets/cupa.jpg","assets/cuqu.jpg","assets/pe02.jpg","assets/pe03.jpg","assets/pe04.jpg","assets/pe05.jpg","assets/pe06.jpg","assets/pe07.jpg","assets/pe08.jpg","assets/pe09.jpg","assets/pe10.jpg","assets/peac.jpg","assets/peki.jpg","assets/pekn.jpg","assets/pepa.jpg","assets/pequ.jpg","assets/sw02.jpg","assets/sw03.jpg","assets/sw04.jpg","assets/sw05.jpg","assets/sw06.jpg","assets/sw07.jpg","assets/sw08.jpg","assets/sw09.jpg","assets/sw10.jpg","assets/swac.jpg","assets/swki.jpg","assets/swkn.jpg","assets/swpa.jpg","assets/swqu.jpg","assets/tarot_back.png","assets/tarot_deck.png","assets/wa02.jpg","assets/wa03.jpg","assets/wa04.jpg","assets/wa05.jpg","assets/wa06.jpg","assets/wa07.jpg","assets/wa08.jpg","assets/wa09.jpg","assets/wa10.jpg","assets/waac.jpg","assets/waki.jpg","assets/wakn.jpg","assets/wapa.jpg","assets/waqu.jpg","favicon.png"]),
	mimeTypes: {".jpeg":"image/jpeg",".png":"image/png",".jpg":"image/jpeg"},
	_: {
		client: {"start":"_app/immutable/entry/start.H3F_7lVo.js","app":"_app/immutable/entry/app.DRvXX_yX.js","imports":["_app/immutable/entry/start.H3F_7lVo.js","_app/immutable/chunks/entry.WE1SeAUf.js","_app/immutable/chunks/scheduler.CgzrBarl.js","_app/immutable/entry/app.DRvXX_yX.js","_app/immutable/chunks/preload-helper.D6kgxu3v.js","_app/immutable/chunks/scheduler.CgzrBarl.js","_app/immutable/chunks/index.B7BUR2Ah.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js')),
			__memo(() => import('./nodes/4.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/cards",
				pattern: /^\/cards\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/cassandra",
				pattern: /^\/cassandra\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
