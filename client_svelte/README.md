# Stripe Payment Client (Svelte + Vite)

A Svelte-based payment client for accepting payments via Stripe.

## Setup

1. Install dependencies:
```bash
npm install
```

2. Make sure your Go server is running on port 4242:
```bash
cd ../server
go run cmd/server/main.go
```

3. Run the development client:
```bash
npm run dev
```

The app will be available at `http://localhost:3000`

## Build for Production

```bash
npm run build
```

The built files will be in the `dist/` directory.

## Features

- Payment form with Stripe Elements
- Client-side routing (no dependencies)
- Component-specific styles organized in `/src/styles/`
- API service layer for backend communication
- Payment completion page with status

## Test Cards

Use test cards from https://docs.stripe.com/testing:

- **Success**: `4242 4242 4242 4242`
- **3D Secure**: `4000 0025 0000 3155`
- **Declined**: `4000 0000 0000 9995`
