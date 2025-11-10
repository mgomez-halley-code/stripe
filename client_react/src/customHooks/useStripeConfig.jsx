import { useState, useEffect } from 'react';
import { loadStripe } from '@stripe/stripe-js';

export const useStripeConfig = () => {
  const [stripePromise, setStripePromise] = useState(null);

  useEffect(() => {
    fetch("/api/config").then(async (r) => {
      const { publishableKey } = await r.json();
      setStripePromise(loadStripe(publishableKey));
    });
  }, []);

  return stripePromise;
};
