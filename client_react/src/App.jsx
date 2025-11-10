import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider } from 'styled-components';
import { useStripeConfig } from './customHooks';
import { AppRoutes } from './routes';
import { GlobalStyles } from './styles/GlobalStyles';
import { theme } from './styles/theme';

function App() {
  const stripePromise = useStripeConfig();

  return (
    <ThemeProvider theme={theme}>
      <GlobalStyles />
      <main>
        <BrowserRouter>
          <AppRoutes stripePromise={stripePromise} />
        </BrowserRouter>
      </main>
    </ThemeProvider>
  );
}

export default App;