import { Routes, Route } from 'react-router-dom';
import { Payment, Completion } from '../pages';

export const AppRoutes = ({ stripePromise }) => {
  return (
    <Routes>
      <Route path="/" element={<Payment stripePromise={stripePromise} />} />
      <Route path="/completion" element={<Completion stripePromise={stripePromise} />} />
    </Routes>
  );
};
