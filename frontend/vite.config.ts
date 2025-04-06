import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	server: {
        allowedHosts: ['5f3f-152-58-188-234.ngrok-free.app', 'localhost:8080'] // Add your host here
    }
});
