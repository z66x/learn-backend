<script lang="ts">
  let url = '';
  let shortUrl = '';
  let error = '';
  let loading = false;

  async function shorten() {
    if (!url) return;
    loading = true;
    error = '';
    shortUrl = '';

    try {
      const res = await fetch('http://localhost:8080/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url })
      });

      if (!res.ok) throw new Error('Failed to shorten URL');

      const data = await res.json();
      shortUrl = data.short_url;
    } catch (e) {
      error = 'Something went wrong. Please try again.';
    } finally {
      loading = false;
    }
  }
</script>

<main>
  <div class="container">
    <h1>URL Shortener</h1>
    <p class="subtitle">Paste a long URL and get a short one instantly</p>

    <div class="input-group">
      <input
        type="url"
        bind:value={url}
        placeholder="https://example.com/very/long/url"
        on:keydown={(e) => e.key === 'Enter' && shorten()}
      />
      <button on:click={shorten} disabled={loading}>
        {loading ? 'Shortening...' : 'Shorten'}
      </button>
    </div>

    {#if shortUrl}
      <div class="result">
        <p>Your short URL:</p>
        <a href={shortUrl} target="_blank">{shortUrl}</a>
        <button class="copy" on:click={() => navigator.clipboard.writeText(shortUrl)}>
          Copy
        </button>
      </div>
    {/if}

    {#if error}
      <p class="error">{error}</p>
    {/if}
  </div>
</main>

<style>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  main {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #0f0f0f;
    font-family: 'Inter', system-ui, sans-serif;
  }

  .container {
    width: 100%;
    max-width: 580px;
    padding: 2rem;
  }

  h1 {
    font-size: 2rem;
    font-weight: 700;
    color: #fff;
    margin-bottom: 0.5rem;
  }

  .subtitle {
    color: #666;
    margin-bottom: 2rem;
    font-size: 0.95rem;
  }

  .input-group {
    display: flex;
    gap: 0.5rem;
  }

  input {
    flex: 1;
    padding: 0.75rem 1rem;
    border-radius: 8px;
    border: 1px solid #222;
    background: #1a1a1a;
    color: #fff;
    font-size: 0.95rem;
    outline: none;
    transition: border-color 0.2s;
  }

  input:focus {
    border-color: #555;
  }

  button {
    padding: 0.75rem 1.25rem;
    border-radius: 8px;
    border: none;
    background: #fff;
    color: #000;
    font-weight: 600;
    font-size: 0.9rem;
    cursor: pointer;
    transition: opacity 0.2s;
  }

  button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  button:hover:not(:disabled) {
    opacity: 0.85;
  }

  .result {
    margin-top: 1.5rem;
    padding: 1rem;
    border-radius: 8px;
    background: #1a1a1a;
    border: 1px solid #222;
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }

  .result p {
    color: #666;
    font-size: 0.85rem;
    white-space: nowrap;
  }

  .result a {
    color: #fff;
    text-decoration: none;
    font-size: 0.95rem;
    flex: 1;
  }

  .result a:hover {
    text-decoration: underline;
  }

  .copy {
    padding: 0.4rem 0.75rem;
    font-size: 0.8rem;
    background: #2a2a2a;
    color: #fff;
    border: 1px solid #333;
  }

  .error {
    margin-top: 1rem;
    color: #ff4444;
    font-size: 0.9rem;
  }
</style>