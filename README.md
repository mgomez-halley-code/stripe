# Stripe Payment Integration Demo

A full-stack payment processing application demonstrating Stripe integration with multiple frontend frameworks and a Go backend.

> **Note:** This is a demo project for educational and demonstration purposes. Unit tests are not included in this implementation.

## Project Overview

This repository contains a complete payment processing system with:
- **Backend**: Go server handling Stripe Payment Intents and webhooks
- **Frontend Options**:
  - React with styled-components
  - Svelte with svelte-routing

## Architecture

```
Stripe/
├── server/              # Go backend server
├── client_react/        # React + styled-components frontend
└── client_svelte/       # Svelte + svelte-routing frontend
```

## Quick Start

### 1. Setup the Server

Navigate to the server directory and follow the setup instructions:

**[📖 Server Documentation](server/README.md)**

```bash
cd server
# Create .env file with your Stripe keys
go run cmd/server/main.go
```

The server runs on `http://localhost:4242`

### 2. Choose a Frontend

#### Option A: React Client

**[📖 React Client Documentation](client_react/README.md)**

```bash
cd client_react
npm install
npm start
```

**Features:**
- React 17 with hooks
- Styled-components for CSS-in-JS
- React Router for navigation
- Custom hooks for Stripe integration
- Organized folder structure (customHooks, components, pages, routes, styles)

#### Option B: Svelte Client

**[📖 Svelte Client Documentation](client_svelte/README.md)**

```bash
cd client_svelte
npm install
npm run dev
```

**Features:**
- Svelte 5 with runes ($state, $props, $effect)
- Svelte-routing for navigation
- Component-scoped styles
- Clean API service layer

## Tech Stack

### Backend
- **Language**: Go 1.23+
- **Framework**: Standard library with gorilla/mux
- **Payment**: Stripe Go SDK
- **Architecture**: Clean architecture with internal packages

### Frontend - React
- **Framework**: React 17
- **Styling**: Styled-components
- **Routing**: React Router v6
- **Build Tool**: Vite
- **Payment**: Stripe.js & React Stripe.js

### Frontend - Svelte
- **Framework**: Svelte 5
- **Routing**: svelte-routing
- **Build Tool**: Vite
- **Payment**: Stripe.js

## Features

✅ Payment Intent creation
✅ Stripe Elements integration
✅ Payment confirmation flow
✅ Webhook handling
✅ Error handling
✅ Loading states
✅ Responsive design
✅ Client-side routing

## API Endpoints

The Go server exposes the following endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/config` | Returns Stripe publishable key |
| POST | `/api/create-payment-intent` | Creates a payment intent for €19.99 |
| POST | `/api/webhook` | Handles Stripe webhook events |

## Payment Flow

1. **Client loads** → Fetches Stripe publishable key from server
2. **Payment page** → Creates payment intent on server
3. **User enters card** → Stripe Elements handles card input
4. **Submit payment** → Client confirms payment with Stripe
5. **Redirect** → User redirected to completion page
6. **Webhook** → Server receives payment confirmation from Stripe

## Testing

### Test Cards

Use these test cards from [Stripe Testing Docs](https://docs.stripe.com/testing):

| Card Number | Type | Result |
|-------------|------|--------|
| `4242 4242 4242 4242` | Visa | Success |
| `4000 0025 0000 3155` | Visa | Requires 3D Secure |
| `4000 0000 0000 9995` | Visa | Declined |

**Card Details:**
- Expiry: Any future date (e.g., 12/34)
- CVC: Any 3 digits (e.g., 123)
- ZIP: Any 5 digits (e.g., 12345)

### Webhook Testing

Install the [Stripe CLI](https://stripe.com/docs/stripe-cli) and forward webhooks to your local server:

```bash
stripe listen --forward-to localhost:3000/api/webhook
```

Copy the webhook signing secret to your server's `.env` file.

## Environment Variables

### Server (.env)
```bash
STRIPE_SECRET_KEY=sk_test_...
STRIPE_PUBLISHABLE_KEY=pk_test_...
STRIPE_WEBHOOK_SECRET=whsec_...
PORT=4242
HOST=0.0.0.0
```

## Project Structure

### Server Structure
```
server/
├── cmd/server/          # Application entry point
├── internal/
│   ├── api/            # API handlers and routing
│   ├── bootstrap/      # Dependency initialization
│   ├── config/         # Configuration management
│   └── middleware/     # HTTP middleware
└── go.mod
```

### React Client Structure
```
client_react/src/
├── customHooks/        # Custom React hooks
├── components/         # Reusable components
├── pages/             # Page components
├── routes/            # Route configuration
├── styles/            # Styled-components theme & global styles
├── App.jsx
└── index.jsx
```

### Svelte Client Structure
```
client_svelte/src/
├── api/               # API service layer
├── components/        # Svelte components
├── pages/            # Page components
├── App.svelte
└── main.js
```

## Development

### Prerequisites
- Go 1.23 or higher
- Node.js 18 or higher
- Stripe account (test mode)

### Running All Services

**Terminal 1 - Server:**
```bash
cd server
go run cmd/server/main.go
```

**Terminal 2 - React Client:**
```bash
cd client_react
npm start
```

**Or Terminal 2 - Svelte Client:**
```bash
cd client_svelte
npm run dev
```

## Documentation Links

- [Server README](server/README.md) - Go backend documentation
- [React Client README](client_react/README.md) - React frontend documentation
- [Svelte Client README](client_svelte/README.md) - Svelte frontend documentation
- [Stripe Documentation](https://stripe.com/docs)
- [Stripe Testing](https://docs.stripe.com/testing)

## Security Notes

⚠️ **Important:**
- Never commit your `.env` file or expose secret keys
- Use test mode keys for development
- Validate webhooks using the webhook secret
- Always process payments server-side
- Never trust client-side payment amounts

## License

This is a demo project for educational purposes.
