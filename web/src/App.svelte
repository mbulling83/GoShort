<script>
  let longUrl = '';
  let shortUrl = '';

  const shortenUrl = async () => {
    try {
      const response = await fetch('/api/v1/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ long_url: longUrl }),
      });

      if (!response.ok) throw new Error('Failed to shorten URL');

      const data = await response.json();
      shortUrl = data.short_url ? `http://localhost:8081/${data.short_url}` : '';
    } catch (error) {
      alert('Error: Could not shorten the URL.');
      console.error(error);
    }
  };
</script>


<style>
  body {
    font-family: Arial, sans-serif;
    text-align: center;
    padding: 20px;
  }

  input {
    padding: 10px;
    width: 300px;
    margin-right: 10px;
  }

  button {
    padding: 10px 20px;
    cursor: pointer;
  }

  .shortened {
    margin-top: 20px;
    font-size: 1.2rem;
  }
</style>

<h1>GoShort - URL Shortener</h1>

<div>
  <input
    type="url"
    bind:value="{longUrl}"
    placeholder="Enter your URL"
    required
  />
  <button on:click="{shortenUrl}">Shorten</button>
</div>

{#if shortUrl}
  <div class="shortened">
    <p>
      Your shortened URL:
      <a href="{shortUrl}" target="_blank">{shortUrl}</a>
    </p>
  </div>
{/if}
