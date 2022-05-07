module.exports = {
	testRegex: '/resources/js/.*.spec.js$',
	moduleFileExtensions: [
		'js',
		'json',
		'vue',
	],
	testEnvironment: 'jsdom',
	transform: {
		'^.+\\.vue$': 'vue3-jest',
		'^.+\\.js$': 'babel-jest',
	},
	moduleNameMapper: {
		'^@/(.*)$': '<rootDir>/resources/js/$1',
	},
	roots: [
		'<rootDir>/resources/js/',
	],
};
