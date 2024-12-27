<script>
  import Main from '../components/Main.svelte';
  import Footer from '../components/Footer.svelte';

  let longUrl = '';
  let shortUrl = '';
  let customUrl = '';
  let errorMessage = '';
  let expiry = '';
  let validationError = '';
  let isCopied = false;
  let showAccordion = false;

  const shortenUrl = async () => {
    if (!longUrl) {
      validationError = 'Please enter a valid URL.';
      return;
    }
    validationError = '';

    try {
      const response = await fetch('/api/v1/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ long_url: longUrl, custom_url: customUrl, expiry }),
      });

      if (!response.ok) {
        const error = await response.text();
        if (response.status === 409) {
          errorMessage = 'The custom URL is already taken. Please try another.';
        } else {
          errorMessage = 'Error: Could not shorten the URL.';
        }
        throw new Error(errorMessage);
      }

      const data = await response.json();
      shortUrl = `${window.location.origin}/${data.short_url}`;
      showAccordion = false; // Collapse accordion
    } catch (error) {
      console.error(error);
    }
  };

  const copyToClipboard = async () => {
    if (shortUrl) {
      await navigator.clipboard.writeText(shortUrl);
      isCopied = true;
      setTimeout(() => (isCopied = false), 2000);
    }
  };
</script>

<div class="flex flex-col min-h-screen bg-gray-100">
  <main class="flex-grow flex items-center justify-center">
    <Main
      bind:longUrl
      bind:shortUrl
      bind:customUrl
      bind:errorMessage
      bind:expiry
      bind:validationError
      bind:isCopied
      bind:showAccordion
      {shortenUrl}
      {copyToClipboard}
    />
  </main>
  <Footer />
</div>
