<script>
    let longUrl = '';
    let shortUrl = '';
    let errorMessage = '';
  
    const shortenUrl = async () => {
      try {
        errorMessage = '';
        shortUrl = ''; // Clear any previous successful result
        const response = await fetch('/api/v1/shorten', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ long_url: longUrl }),
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
          alert('Shortened URL copied to clipboard!');
        }
      } catch (error) {
        alert('Failed to copy URL.');
        console.error(error);
      }
    };
  </script>
  
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
    <h1 class="text-4xl font-bold text-blue-600 mb-6">GoShort - URL Shortener</h1>
  
    <div class="w-full max-w-md space-y-4">
      <input
        type="url"
        bind:value="{longUrl}"
        placeholder="Enter your URL"
        class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:ring focus:ring-blue-300 focus:outline-none"
        required
      />
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
            class="bg-green-500 text-white font-semibold py-1 px-3 rounded-lg hover:bg-green-600 focus:ring focus:ring-green-300 focus:outline-none transition"
          >
            Copy
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
  