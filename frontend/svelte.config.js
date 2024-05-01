import adapter from '@sveltejs/adapter-auto';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import sveltePreprocess from 'svelte-preprocess';
/** @type {import('@sveltejs/kit').Config} */
const config = {
    kit: {
        adapter: adapter()
    },
    vite: {
        plugins: [
            svelte({
                preprocess: sveltePreprocess({
                    postcss: true
                })
            })
        ]
    }
};
export default config;