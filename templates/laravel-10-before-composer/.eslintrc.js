module.exports = {
	extends: [
		'plugin:vue/vue3-recommended',
		'google',
	],
	env: {
		node: true,
		browser: true,
		jest: true,
	},
	parser: 'vue-eslint-parser',
	parserOptions: {
		parser: '@babel/eslint-parser',
	},
	rules: {
		'object-curly-spacing': ['error', 'always'],
		'require-jsdoc': 'off',
		'indent': ['error', 'tab', { SwitchCase: 1 }],
		'max-len': 'off',
		'no-unused-vars': ['error', { varsIgnorePattern: '^_', argsIgnorePattern: '^_' }],
		'valid-jsdoc': 'off',
		'quotes': ['error', 'single'],
		'no-tabs': 'off',
		'vue/require-prop-types': 'off',
		'vue/multi-word-component-names': 'off',
		'vue/require-default-prop': 'off',
		'vue/no-v-html': 'off',
		'vue/html-indent': ['error', 'tab'],
	},
	root: true,
	ignorePatterns: [
		'resources/js/Jetstream/*',
		'resources/js/PageBuilder/*',
		'resources/js/Pages/Profile/Partials/TwoFactorAuthenticationForm.vue',
		'resources/js/Pages/Profile/Partials/ConnectedAccountsForm.vue',
		'resources/js/Pages/Default/Index.vue',
	],
};
