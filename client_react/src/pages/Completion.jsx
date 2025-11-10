import { usePaymentStatus } from '../customHooks';
import { PageContainer, PageTitle, HomeLink, StatusMessage } from './Completion.styled';

function Completion(props) {
  const { stripePromise } = props;
  const messageBody = usePaymentStatus(stripePromise);

  return (
    <PageContainer>
      <PageTitle>Thank you!</PageTitle>
      <HomeLink href="/">Return to Home</HomeLink>
      <StatusMessage role="alert" $visible={!!messageBody}>
        {messageBody}
      </StatusMessage>
    </PageContainer>
  );
}

export default Completion;
