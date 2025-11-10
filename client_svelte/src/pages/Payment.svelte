<script>
  import { untrack } from 'svelte';
  import CheckoutForm from '../components/CheckoutForm.svelte';
  import { createPaymentIntent } from '../api/stripe.js';

  let { stripePromise } = $props();

  let clientSecret = $state('');
  let stripe = $state(null);
  let elements = $state(null);
  let ready = $state(false);

  // Initialize payment on mount
  async function initializePayment() {
    // Create PaymentIntent as soon as the page loads
    clientSecret = await createPaymentIntent();

    // Wait for Stripe to load
    if (stripePromise) {
      stripe = await stripePromise;

      if (stripe && clientSecret) {
        elements = stripe.elements({ clientSecret });
        ready = true;
      }
    }
  }

  // Run once on mount
  $effect(() => {
    untrack(() => initializePayment());
  });
</script>

<div>
  <h1>Payment</h1>
  {#if ready}
    <CheckoutForm {stripe} {elements} />
  {:else}
    <p>Loading payment form...</p>
  {/if}
</div>
