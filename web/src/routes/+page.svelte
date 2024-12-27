<script>
    import Main from '../components/Main.svelte';
    import Footer from '../components/Footer.svelte';
  
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
      if (!isValidUrl(longUrl)) {
        validationError = 'Please enter a valid URL.';
        return;
      }
      validationError = ''; // Clear validation errors
  
      try {
        errorMessage = '';
        shortUrl = ''; // Clear previous result
  
        const formattedExpiry = expiry ? new Date(expiry).toISOString() : null;
  
        const response = await fetch('/api/v1/shorten', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ long_url: longUrl, expiry: formattedExpiry }),
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
          isCopied = true;
          setTimeout(() => (isCopied = false), 2000);
        }
      } catch (error) {
        console.error('Failed to copy URL.', error);
      }
    };
  </script>
  
  <div class="flex flex-col min-h-screen bg-gray-100">
    <main class="flex-grow flex items-center justify-center">
      <Main
        bind:longUrl
        bind:shortUrl
        bind:errorMessage
        bind:expiry
        bind:validationError
        bind:isCopied
        bind:showExpiry
        {shortenUrl}
        {copyToClipboard}
      />
    </main>
    <Footer />
  </div>
  