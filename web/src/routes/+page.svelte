<script>
    let longUrl = '';
    let shortUrl = '';
    let errorMessage = '';
    let expiry = ''; // Local datetime format
    let validationError = ''; // URL validation error
    let isCopied = false; // State for the copy button animation
    let showExpiry = false; // State to toggle the accordion
  
    const isValidUrl = (url) => {
      try {
        new URL(url);
        return true;
      } catch {
        return false;
      }
    };
  
    const shortenUrl = async () => {
      // Validate the URL
      if (!isValidUrl(longUrl)) {
        validationError = 'Please enter a valid URL.';
        return;
      }
      validationError = ''; // Clear any previous validation error
  
      try {
        errorMessage = '';
        shortUrl = ''; // Clear any previous successful result
  
        // Convert expiry to RFC3339 format if provided
        const formattedExpiry = expiry ? new Date(expiry).toISOString() : null;
  
        const response = await fetch('/api/v1/shorten', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            long_url: longUrl,
            expiry: formattedExpiry,
          }),
        });
  
        if (!response.ok) throw new Error('Failed to shorten URL');
  
        const data = await response.json();
        shortUrl = data.short_url ? `${window.location.origin}/${data.short_url}` : '';
      } catch (error) {
        errorMessage = 'Error: Could not shorten the URL.';
        console.error(error);
      }
    };
  
    const copyToClipboard = async () => {
      try {
        if (shortUrl) {
          await navigator.clipboard.writeText(shortUrl);
          isCopied = true; // Set state to show "Copied!" message
          setTimeout(() => (isCopied = false), 2000); // Reset state after 2 seconds
        }
      } catch (error) {
        console.error('Failed to copy URL.', error);
      }
    };
  </script>
  
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
    <h1 class="text-4xl font-bold text-blue-600 mb-6">GoShort - URL Shortener</h1>
  
    <div class="w-full max-w-md space-y-4">
      <div>
        <label for="long-url" class="block text-gray-700 font-medium mb-2">Enter URL</label>
        <input
          id="long-url"
          type="url"
          bind:value="{longUrl}"
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
          on:click="{() => (showExpiry = !showExpiry)}"
          class="w-full bg-gray-200 text-gray-700 font-semibold py-2 px-4 rounded-lg hover:bg-gray-300 focus:ring focus:ring-gray-300 focus:outline-none transition flex items-center justify-between"
        >
          <span>Set Expiry (Optional)</span>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 transform transition-transform"
            viewBox="0 0 20 20"
            fill="currentColor"
            class:rotate-180="{showExpiry}"
          >
            <path fill-rule="evenodd" d="M10 3a1 1 0 01.707.293l6 6a1 1 0 11-1.414 1.414L10 5.414 4.707 10.707A1 1 0 013.293 9.293l6-6A1 1 0 0110 3z" clip-rule="evenodd" />
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
              bind:value="{expiry}"
              class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:ring focus:ring-blue-300 focus:outline-none"
            />
          </div>
        {/if}
      </div>
  
      <button
        on:click="{shortenUrl}"
        class="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg hover:bg-blue-700 focus:ring focus:ring-blue-300 focus:outline-none transition"
      >
        Shorten
      </button>
    </div>
  
    {#if shortUrl}
      <div class="mt-6 text-center bg-white p-4 rounded-lg shadow-lg">
        <p class="text-lg font-medium">Your shortened URL:</p>
        <div class="flex items-center justify-center space-x-4">
          <a
            href="{shortUrl}"
            target="_blank"
            class="text-blue-600 underline break-all"
          >
            {shortUrl}
          </a>
          <button
            on:click="{copyToClipboard}"
            class="bg-green-500 text-white font-semibold py-1 px-3 rounded-lg hover:bg-green-600 focus:ring focus:ring-green-300 focus:outline-none transition flex items-center space-x-2"
          >
            {#if isCopied}
              <span class="animate-pulse">Copied!</span>
            {:else}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path d="M8 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V4a2 2 0 00-2-2H8zm-3 3a2 2 0 012-2h5a2 2 0 012 2v11a2 2 0 01-2 2H7a2 2 0 01-2-2V5z" />
              </svg>
              <span>Copy</span>
            {/if}
          </button>
        </div>
      </div>
    {/if}
  
    {#if errorMessage}
      <div class="mt-4 text-red-500 font-medium">{errorMessage}</div>
    {/if}
  </div>
  
  <style>
    a {
      word-break: break-all;
    }
  </style>
  