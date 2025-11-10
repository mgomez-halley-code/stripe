/**
 * Stripe API service
 */

/**
 * Fetch Stripe publishable key from the server
 * @returns {Promise<string>} The publishable key
 */
export async function fetchPublishableKey() {
  const response = await fetch('/api/config');
  const { publishableKey } = await response.json();
  return publishableKey;
}

/**
 * Create a payment intent
 * @returns {Promise<string>} The client secret
 */
export async function createPaymentIntent() {
  const response = await fetch('/api/create-payment-intent', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });
  const { clientSecret } = await response.json();
  return clientSecret;
}
