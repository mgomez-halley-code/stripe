import { Elements } from '@stripe/react-stripe-js';
import { CheckoutForm } from '../components';
import { usePaymentIntent } from '../customHooks';
import { PageContainer, PageTitle, LoadingMessage } from './Payment.styled';

function Payment(props) {
  const { stripePromise } = props;
  const clientSecret = usePaymentIntent();

  return (
    <PageContainer>
      <PageTitle>Payment</PageTitle>
      {clientSecret && stripePromise ? (
        <Elements stripe={stripePromise} options={{ clientSecret, }}>
          <CheckoutForm />
        </Elements>
      ) : (
        <LoadingMessage>Loading payment form...</LoadingMessage>
      )}
    </PageContainer>
  );
}

export default Payment;
