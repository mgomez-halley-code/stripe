/**
 * Confirms a Stripe payment with the provided elements
 * @param {Object} stripe - Stripe instance
 * @param {Object} elements - Stripe Elements instance
 * @param {string} returnUrl - URL to redirect to after payment
 * @returns {Promise<{error: Object|null}>}
 */
export async function confirmPayment(stripe, elements, returnUrl) {
  if (!stripe || !elements) {
    return {
      error: {
        message: 'Payment system not initialized'
      }
    };
  }

  const { error } = await stripe.confirmPayment({
    elements,
    confirmParams: {
      return_url: returnUrl,
    },
  });

  return { error };
}

/**
 * Formats payment error message for display
 * @param {Object} error - Stripe error object
 * @returns {string} User-friendly error message
 */
export function formatPaymentError(error) {
  if (!error) return '';

  if (error.type === 'card_error' || error.type === 'validation_error') {
    return error.message;
  }

  return 'An unexpected error occurred.';
}
