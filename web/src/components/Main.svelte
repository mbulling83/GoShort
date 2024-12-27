<script>
	export let longUrl = '';
	export let shortUrl = '';
	export let errorMessage = '';
	export let expiry = '';
	export let validationError = '';
	export let isCopied = false;
	export let showExpiry = false;

	export let shortenUrl;
	export let copyToClipboard;
</script>

<div class="flex flex-col items-center justify-center bg-gray-100 p-4">
	<h1 class="text-4xl font-bold text-blue-600 mb-6">GoShort - URL Shortener</h1>

	<div class="w-full max-w-md space-y-4">
		<div>
			<input
				id="long-url"
				type="url"
				bind:value={longUrl}
				placeholder="Enter your URL"
				class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:ring focus:ring-blue-300 focus:outline-none"
				required
			/>
			{#if validationError}
				<p class="text-red-500 text-sm mt-1">{validationError}</p>
			{/if}
		</div>

		<div>
			<button
				type="button"
				on:click={() => (showExpiry = !showExpiry)}
				class="w-full bg-gray-200 text-gray-700 font-semibold py-2 px-4 rounded-lg hover:bg-gray-300 focus:ring focus:ring-gray-300 focus:outline-none transition flex items-center justify-between"
			>
				<span>Set Expiry (Optional)</span>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5 transform transition-transform"
					viewBox="0 0 20 20"
					fill="currentColor"
					class:rotate-180={showExpiry}
				>
					<path
						fill-rule="evenodd"
						d="M10 3a1 1 0 01.707.293l6 6a1 1 0 11-1.414 1.414L10 5.414 4.707 10.707A1 1 0 013.293 9.293l6-6A1 1 0 0110 3z"
						clip-rule="evenodd"
					/>
				</svg>
			</button>
			{#if showExpiry}
				<div class="mt-2">
					<label for="expiry" class="block text-gray-700 font-medium mb-2">
						Expiry Date & Time
					</label>
					<input
						id="expiry"
						type="datetime-local"
						bind:value={expiry}
						class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:ring focus:ring-blue-300 focus:outline-none"
					/>
				</div>
			{/if}
		</div>

		<button
			on:click={shortenUrl}
			class="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg hover:bg-blue-700 focus:ring focus:ring-blue-300 focus:outline-none transition"
		>
			Shorten
		</button>
	</div>

	{#if shortUrl}
		<div class="mt-6 text-center bg-white p-6 rounded-lg shadow-lg">
			<p class="text-lg font-medium mb-4">Your shortened URL is ready! 🚀</p>
			<div class="flex flex-col items-center space-y-4">
				<div
					class="bg-gray-100 text-gray-800 px-4 py-2 rounded-md border border-gray-300 font-mono text-sm break-all w-full max-w-lg"
				>
					<a href={shortUrl} target="_blank" class="hover:text-blue-500">
						{shortUrl}
					</a>
				</div>
				<button
					on:click={copyToClipboard}
					class="bg-green-500 text-white font-semibold py-2 px-4 rounded-lg hover:bg-green-600 focus:ring focus:ring-green-300 focus:outline-none transition flex items-center space-x-2"
				>
					{#if isCopied}
						<span class="animate-pulse">Copied! ✅</span>
					{:else}
						<span>📋 Copy URL</span>
					{/if}
				</button>
			</div>
		</div>
	{/if}

	{#if errorMessage}
		<div class="mt-4 text-red-500 font-medium">{errorMessage}</div>
	{/if}
</div>
