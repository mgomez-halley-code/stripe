# Stripe Payment Server

Go backend server that handles Stripe payment intents and webhooks.

## Setup

1. Create a `.env` file in the server directory:
```bash
STRIPE_SECRET_KEY=sk_test_your_secret_key_here
STRIPE_PUBLISHABLE_KEY=pk_test_your_publishable_key_here
STRIPE_WEBHOOK_SECRET=whsec_your_webhook_secret_here
STATIC_DIR=../client/dist
```

Get your API keys from the [Stripe Dashboard](https://dashboard.stripe.com/test/apikeys)

2. Install Go dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The server will run on `http://localhost:4242`

## API Endpoints

- `GET /config` - Returns the publishable key
- `POST /create-payment-intent` - Creates a payment intent for €19.99
- `POST /webhook` - Handles Stripe webhook events

## Testing with Stripe CLI

Install the [Stripe CLI](https://stripe.com/docs/stripe-cli) and forward webhooks:

```bash
stripe listen --forward-to localhost:4242/webhook
```

This will give you a webhook secret to use in your `.env` file.

## Test Cards

Use test cards from https://docs.stripe.com/testing:

- **Success**: `4242 4242 4242 4242`
- **3D Secure**: `4000 0025 0000 3155`
- **Declined**: `4000 0000 0000 9995`
