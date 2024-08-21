

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/layout.svelte.js')).default;
export const imports = ["_app/immutable/nodes/0.BDkfXnre.js","_app/immutable/chunks/scheduler.C52glMjA.js","_app/immutable/chunks/index.M9Shu6X3.js"];
export const stylesheets = [];
export const fonts = [];
