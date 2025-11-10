/**
 * Utility functions for handling Stripe Payment Intents
 */

/**
 * Retrieve payment intent from URL query parameters
 * @param {Promise} stripePromise - Promise that resolves to Stripe instance
 * @returns {Promise<{error?: object, paymentIntent?: object}>}
 */
export async function getPaymentIntent(stripePromise) {
  if (!stripePromise) {
    return { error: { message: 'Stripe not initialized' } };
  }

  const stripe = await stripePromise;

  const url = new URL(window.location);
  const clientSecret = url.searchParams.get('payment_intent_client_secret');

  if (!clientSecret) {
    return { error: { message: 'No payment intent found' } };
  }

  return stripe.retrievePaymentIntent(clientSecret);
}
