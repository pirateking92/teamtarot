

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.DYgd1Z6E.js","_app/immutable/chunks/scheduler.CgzrBarl.js","_app/immutable/chunks/index.B7BUR2Ah.js","_app/immutable/chunks/entry.WE1SeAUf.js"];
export const stylesheets = [];
export const fonts = [];
