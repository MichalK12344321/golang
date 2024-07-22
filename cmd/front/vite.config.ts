/// <reference types="vitest" />
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';

export default defineConfig({
	plugins: [react()],
	resolve: {
		alias: {
			'@': '/src',
		},
	},
	test: {
		globals: true,
		environment: 'happy-dom',
		setupFiles: './tests/setup.ts',
	},
});
