import styled from 'styled-components';

export const StyledForm = styled.form`
  border: ${({ theme }) => theme.colors.lightGrey} solid 1px;
  border-radius: ${({ theme }) => theme.radius};
  padding: ${({ theme }) => theme.spacing.lg};
  margin: ${({ theme }) => theme.spacing.lg} 0;
  box-shadow: ${({ theme }) => theme.shadows.card};
`;

export const SubmitButton = styled.button`
  width: 100%;
  margin-top: ${({ theme }) => theme.spacing.md};
`;

export const ButtonText = styled.span`
  display: flex;
  align-items: center;
  justify-content: center;
`;

export const Spinner = styled.div`
  border: 2px solid ${({ theme }) => theme.colors.white};
  border-top: 2px solid transparent;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  animation: spin 0.6s linear infinite;

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
`;

export const ErrorMessage = styled.div`
  color: #df1b41;
  font-size: 0.9em;
  margin-top: ${({ theme }) => theme.spacing.md};
  padding: ${({ theme }) => theme.spacing.sm};
  background-color: #fef0f0;
  border-radius: ${({ theme }) => theme.radius};
  border-left: 3px solid #df1b41;
`;
