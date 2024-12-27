import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
		  // Configure the static site generation
		  pages: 'build',       // Output directory for prerendered pages
		  assets: 'build',      // Output directory for static assets
		  fallback: 'index.html', // Enables SPA fallback
		  precompress: false,   // Disables precompression (gzip/brotli)
		  strict: true          // Ensures all routes are accounted for
		}),
	  }
};

export default config;
