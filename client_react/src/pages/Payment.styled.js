import styled from 'styled-components';

export const PageContainer = styled.div`
  width: 100%;
`;

export const PageTitle = styled.h1`
  font-size: 2em;
  margin-bottom: ${({ theme }) => theme.spacing.lg};
  color: ${({ theme }) => theme.colors.darkTerminal};
  font-weight: 700;
`;

export const LoadingMessage = styled.div`
  text-align: center;
  padding: ${({ theme }) => theme.spacing.xl};
  color: ${({ theme }) => theme.colors.darkTerminal};
  font-size: 1em;
`;
