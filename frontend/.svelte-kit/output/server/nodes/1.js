

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.Dean8tJb.js","_app/immutable/chunks/scheduler.C52glMjA.js","_app/immutable/chunks/index.M9Shu6X3.js","_app/immutable/chunks/entry.B20SFMH4.js"];
export const stylesheets = [];
export const fonts = [];
