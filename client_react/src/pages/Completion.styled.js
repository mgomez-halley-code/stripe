import styled from 'styled-components';

export const PageContainer = styled.div`
  width: 100%;
  text-align: center;
`;

export const PageTitle = styled.h1`
  font-size: 2em;
  margin-bottom: ${({ theme }) => theme.spacing.lg};
  color: ${({ theme }) => theme.colors.darkTerminal};
  font-weight: 700;
`;

export const HomeLink = styled.a`
  display: inline-block;
  margin: ${({ theme }) => theme.spacing.lg} 0;
  padding: ${({ theme }) => theme.spacing.sm} ${({ theme }) => theme.spacing.lg};
  background-color: ${({ theme }) => theme.colors.accent};
  color: ${({ theme }) => theme.colors.white};
  text-decoration: none;
  border-radius: ${({ theme }) => theme.radius};
  font-weight: 600;
  transition: all 0.2s ease;

  &:hover {
    filter: contrast(115%);
    text-decoration: none;
  }

  &:active {
    transform: translateY(0px) scale(0.98);
    filter: brightness(0.9);
  }
`;

export const StatusMessage = styled.div`
  font-family: ${({ theme }) => theme.fonts.mono};
  background-color: ${({ theme }) => theme.colors.terminalBg};
  color: ${({ theme }) => theme.colors.success};
  padding: ${({ theme }) => theme.spacing.lg};
  margin: ${({ theme }) => theme.spacing.lg} 0;
  border-radius: ${({ theme }) => theme.radius};
  font-size: ${({ theme }) => theme.fontSizes.terminal};
  text-align: left;
  display: ${({ $visible }) => ($visible ? 'block' : 'none')};

  a {
    color: ${({ theme }) => theme.colors.success};
    text-decoration: underline;

    &:hover {
      color: ${({ theme }) => theme.colors.white};
    }
  }
`;
