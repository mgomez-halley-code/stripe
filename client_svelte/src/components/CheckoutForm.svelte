<script>
  import { untrack } from 'svelte';
  import { confirmPayment, formatPaymentError } from '../api/payment.js';

  let { stripe, elements } = $props();

  let message = $state(null);
  let isLoading = $state(false);
  let paymentElement = $state();
  let linkAuthElement = $state();

  // Mount Stripe elements when component mounts
  $effect(() => {
    untrack(() => {
      if (!elements) return;

      // Create and mount the payment element
      paymentElement = elements.create('payment');
      paymentElement.mount('#payment-element');

      // Create and mount the link authentication element
      linkAuthElement = elements.create('linkAuthentication');
      linkAuthElement.mount('#link-authentication-element');
    });

    // Cleanup function - runs when component unmounts
    return () => {
      if (paymentElement) {
        paymentElement.unmount();
      }
      if (linkAuthElement) {
        linkAuthElement.unmount();
      }
    };
  });

  async function handleSubmit(e) {
    e.preventDefault();

    if (!stripe || !elements) {
      return;
    }

    isLoading = true;

    const returnUrl = `${window.location.origin}/completion`;
    const { error } = await confirmPayment(stripe, elements, returnUrl);

    if (error) {
      message = formatPaymentError(error);
    }

    isLoading = false;
  }
</script>

<form id="payment-form" onsubmit={handleSubmit}>
  <div id="link-authentication-element">
    <!-- Stripe Link Authentication Element will mount here -->
  </div>
  <div id="payment-element">
    <!-- Stripe Payment Element will mount here -->
  </div>
  <button disabled={isLoading || !stripe || !elements} id="submit">
    <span id="button-text">
      {#if isLoading}
        <div class="spinner" id="spinner"></div>
      {:else}
        Pay now
      {/if}
    </span>
  </button>
  {#if message}
    <div id="payment-message" class:error={true}>{message}</div>
  {/if}
</form>

<style>
  form {
    width: 100%;
    max-width: 500px;
    margin: 0 auto;
    text-align: left;
  }

  #payment-element {
    margin-bottom: 24px;
  }

  #link-authentication-element {
    margin-bottom: 24px;
  }

  button {
    width: 100%;
    background: #5469d4;
    font-family: Arial, sans-serif;
    color: #ffffff;
    border-radius: 4px;
    border: 0;
    padding: 12px 16px;
    font-size: 16px;
    font-weight: 600;
    cursor: pointer;
    display: block;
    transition: all 0.2s ease;
    box-shadow: 0px 4px 5.5px 0px rgba(0, 0, 0, 0.07);
  }

  button:hover:not(:disabled) {
    filter: contrast(115%);
  }

  button:disabled {
    opacity: 0.5;
    cursor: default;
  }

  #button-text {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .spinner {
    border: 2px solid #ffffff;
    border-radius: 50%;
    border-top: 2px solid transparent;
    width: 16px;
    height: 16px;
    animation: spinner 1s linear infinite;
  }

  @keyframes spinner {
    to {
      transform: rotate(360deg);
    }
  }

  #payment-message {
    color: rgb(105, 115, 134);
    font-size: 16px;
    line-height: 20px;
    padding-top: 12px;
    text-align: center;
  }

  #payment-message.error {
    color: #df1b41;
  }
</style>
