import { defineConfig } from 'vite';
import laravel from 'laravel-vite-plugin';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
	resolve: {
		alias: {
			'@': 'resources/js',
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
