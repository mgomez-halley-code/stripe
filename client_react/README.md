# Stripe Payment Client

React + Vite frontend application for processing Stripe payments using the Payment Element.

## Setup

1. Install dependencies:
```bash
npm install
```

2. Start the development server:
```bash
npm start
```

The app will run on `http://localhost:5173`

## Testing

Use these Stripe test cards from the [official testing documentation](https://docs.stripe.com/testing):

### Successful Payment
- **Card Number**: `4242 4242 4242 4242`
- **Expiry**: Any future date (e.g., `12/34`)
- **CVC**: Any 3 digits (e.g., `123`)
- **ZIP**: Any 5 digits (e.g., `12345`)

### Test Specific Scenarios

**Require authentication (3D Secure)**
- Card: `4000 0025 0000 3155`

**Declined card**
- Card: `4000 0000 0000 9995`

**Insufficient funds**
- Card: `4000 0000 0000 9995`

**Card expired**
- Card: `4000 0000 0000 0069`

For more test cards and scenarios, visit: https://docs.stripe.com/testing
