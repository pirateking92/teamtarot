

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/layout.svelte.js')).default;
export const imports = ["_app/immutable/nodes/0.BN_7um98.js","_app/immutable/chunks/scheduler.CgzrBarl.js","_app/immutable/chunks/index.B7BUR2Ah.js"];
export const stylesheets = [];
export const fonts = [];
