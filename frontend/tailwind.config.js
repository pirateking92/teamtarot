/** @type {import('tailwindcss').Config} */
export default {
  content: ['src/routes/cards/+page.svelte'],
  theme: {
    fontFamily: {
      'serif': ['ui-serif', 'Georgia',],
      'medieval': [ 'MedievalSharp',],
    },
    extend: {
      screens: {
        'body': {'raw': 'body'},
      }
    },
  },
  plugins: [],
}

