const purgecss = require('@fullhuman/postcss-purgecss')({
	content: [ `./public/**/*.html`, `./src/**/*.vue` ],
	defaultExtractor (content) {
		const contentWithoutStyleBlocks = content.replace(/<style[^]+?<\/style>/gi, '')
		return contentWithoutStyleBlocks.match(/[\w-/.:]+(?<!:)/g) || []
	},
	whitelist: [],
	whitelistPatterns: [ /-(leave|enter|appear)(|-(to|from|active))$/, /^(?!(|.*?:)cursor-move).+-move$/, /^router-link(|-exact)-active$/, /data-v-.*/, /mapboxgl.*/ ],
})

module.exports = {
	"plugins": [
		require('tailwindcss')('tailwind.config.js'),
		require('autoprefixer')(),
		...process.env.NODE_ENV === 'production'
			? [ purgecss ]
			: []
	]
}