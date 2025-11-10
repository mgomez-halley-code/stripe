<script>
  import { untrack } from 'svelte';
  import { Router, Route } from 'svelte-routing';
  import { loadStripe } from '@stripe/stripe-js';
  import { fetchPublishableKey } from './api/stripe.js';
  import Payment from './pages/Payment.svelte';
  import Completion from './pages/Completion.svelte';

  let stripePromise = $state(null);
  let isReady = $state(false);

  // Initialize Stripe on mount
  $effect(() => {
    untrack(async () => {
      const publishableKey = await fetchPublishableKey();
      stripePromise = loadStripe(publishableKey);
      isReady = true;
    });
  });
</script>

<main>
  {#if isReady}
    <Router>
      <Route path="/">
        <Payment {stripePromise} />
      </Route>
      <Route path="/completion">
        <Completion {stripePromise} />
      </Route>
    </Router>
  {:else}
    <div class="loading">
      <p>Loading Stripe...</p>
    </div>
  {/if}
</main>

<style>
  :global(body) {
    padding: 20px;
    font-family: 'Raleway', sans-serif;
    display: flex;
    justify-content: center;
    font-size: 1.2em;
    color: #0A2540;
    background-color: #F6F9FC;
  }

  main {
    width: 480px;
  }

  .loading {
    text-align: center;
    padding: 2rem;
  }
</style>
