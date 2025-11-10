<script>
  import { untrack } from 'svelte';
  import { Link } from 'svelte-routing';
  import { getPaymentIntent } from '../api/paymentIntent.js';

  let { stripePromise } = $props();

  let paymentIntentId = $state('');
  let paymentStatus = $state('');
  let message = $state('');
  let loading = $state(true);

  // Load payment intent on mount
  async function loadPaymentIntent() {
    try {
      const { error, paymentIntent } = await getPaymentIntent(stripePromise);

      if (error) {
        message = `> ${error.message}`;
      } else if (paymentIntent) {
        paymentIntentId = paymentIntent.id;
        paymentStatus = paymentIntent.status;
        message = `> Payment ${paymentStatus}: ${paymentIntentId}`;
      }
    } catch (err) {
      message = `> Unexpected error: ${err.message || String(err)}`;
    } finally {
      loading = false;
    }
  }

  // Run once on mount
  $effect(() => {
    untrack(() => loadPaymentIntent());
  });
</script>

<div class="completion-container">
  <h1>Thank you!</h1>
  <Link to="/" class="home-link">Return to Home</Link>

  {#if loading}
    <p>Loading payment details...</p>
  {:else}
    <div id="messages" class="visible" role="alert">
      <p>{message}</p>

      {#if paymentIntentId}
        <p>
          View in dashboard:
          <a
            href="https://dashboard.stripe.com/test/payments/{paymentIntentId}"
            target="_blank"
            rel="noreferrer"
          >
            {paymentIntentId}
          </a>
        </p>
      {/if}
    </div>
  {/if}
</div>

<style>
  .completion-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 2rem;
    margin: 2rem auto;
    max-width: 600px;
  }

  #messages {
    display: none;
    background-color: #0a253c;
    color: #00d4ff;
    padding: 20px;
    border-radius: 6px;
    font-family: monospace;
    width: 100%;
    text-align: center;
  }

  #messages.visible {
    display: block;
  }

  .home-link {
    display: inline-block;
    padding: 10px 20px;
    background-color: #5469d4;
    color: white;
    text-decoration: none;
    border-radius: 6px;
    transition: all 0.2s ease;
  }

  .home-link:hover {
    filter: contrast(115%);
  }
</style>
