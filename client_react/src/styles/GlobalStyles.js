import { createGlobalStyle } from 'styled-components';

export const GlobalStyles = createGlobalStyle`
  @import url('https://fonts.googleapis.com/css?family=Raleway&display=swap');

  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    padding: 20px;
    font-family: ${({ theme }) => theme.fonts.primary};
    display: flex;
    justify-content: center;
    font-size: ${({ theme }) => theme.fontSizes.normal};
    color: ${({ theme }) => theme.colors.darkTerminal};
    background-color: ${({ theme }) => theme.colors.white};
  }

  main {
    width: 480px;
    max-width: 100%;
  }

  form > * {
    margin: ${({ theme }) => theme.spacing.xs} 0;
  }

  button {
    background: ${({ theme }) => theme.colors.accent};
    border-radius: ${({ theme }) => theme.radius};
    color: ${({ theme }) => theme.colors.white};
    border: 0;
    padding: ${({ theme }) => theme.spacing.sm} ${({ theme }) => theme.spacing.md};
    margin-top: ${({ theme }) => theme.spacing.md};
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s ease;
    display: block;
    font-size: 1em;

    &:hover:not(:disabled) {
      filter: contrast(115%);
    }

    &:active:not(:disabled) {
      transform: translateY(0px) scale(0.98);
      filter: brightness(0.9);
    }

    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }
  }

  input,
  select {
    display: block;
    font-size: ${({ theme }) => theme.fontSizes.medium};
    width: 100%;
    margin-bottom: ${({ theme }) => theme.spacing.xs};
    border: 1px solid ${({ theme }) => theme.colors.border};
    padding: 8px;
    border-radius: ${({ theme }) => theme.radius};
  }

  label {
    display: block;
  }

  a {
    color: ${({ theme }) => theme.colors.accent};
    font-weight: 900;
    text-decoration: none;

    &:hover {
      text-decoration: underline;
    }
  }

  small {
    font-size: ${({ theme }) => theme.fontSizes.small};
  }

  fieldset {
    border: 1px solid ${({ theme }) => theme.colors.border};
    padding: ${({ theme }) => theme.spacing.md};
    border-radius: ${({ theme }) => theme.radius};
  }
`;
