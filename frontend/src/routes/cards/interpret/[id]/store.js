// store.js
import { writable } from 'svelte/store';

// Create a writable store
export const getFinalInterpretation = () => {
    async function load() {
        // Fetch interpretation based on id
        const res = await fetch(`http://localhost:8082/cards/interpret/05ded632-c07d-42d5-b651-c098139b7921`);
        const data = await res.text();

        // Set the interpretation in the store
        console.log(data);
    }
    load()
}

