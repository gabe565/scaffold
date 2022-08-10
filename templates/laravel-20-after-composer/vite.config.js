import { defineConfig } from 'vite';
import laravel from 'laravel-vite-plugin';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
	server: {
		host: true,
		strictPort: true,
		hmr: {
			host: 'hot.clevyr.run',
			clientPort: 443,
			protocol: 'wss',
		},
	},
	plugins: [
		laravel({
			input: 'resources/js/app.js',
			refresh: true,
		}),
		vue({
			template: {
				transformAssetUrls: {
					base: null,
					includeAbsolute: false,
				},
			},
		}),
	],
});
