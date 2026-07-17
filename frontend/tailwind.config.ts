import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				display: ['MedievalSharp', 'cursive'],
				body: ['MedievalSharp', 'Arial', 'Helvetica', 'sans-serif']
			},
			colors: {
				// The Cassandra palette, extracted from the original bespoke CSS.
				ink: 'rgb(197, 176, 176)', // muted rose-grey text
				blood: {
					DEFAULT: '#871616',
					dark: '#200505',
					deep: '#130500'
				}
			},
			boxShadow: {
				ember: '0 0 1em 4px rgba(207, 49, 5, 0.5)',
				amber: '0 0 1em 4px rgba(197, 126, 34, 0.5)'
			},
			keyframes: {
				glow: {
					'0%, 100%': { boxShadow: '0 0 1em 2px rgba(207, 49, 5, 0.35)' },
					'50%': { boxShadow: '0 0 1.4em 6px rgba(207, 49, 5, 0.6)' }
				}
			},
			animation: {
				glow: 'glow 3s ease-in-out infinite'
			}
		}
	},
	plugins: []
} satisfies Config;
