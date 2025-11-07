package bootstrap

import (
	"sync"

	"github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/client"
)

// Bootstrapper handles dependency initialization with lazy loading
type Bootstrapper struct {
	stripeSecretKey      string
	stripePublishableKey string
	stripeWebhookSecret  string

	stripeClient     *client.API
	stripeClientLock sync.Mutex
}

// NewBootstrapper creates a new bootstrapper instance
func NewBootstrapper(secretKey, publishableKey, webhookSecret string) *Bootstrapper {
	return &Bootstrapper{
		stripeSecretKey:      secretKey,
		stripePublishableKey: publishableKey,
		stripeWebhookSecret:  webhookSecret,
	}
}

// StripeClient lazily initializes and returns the Stripe client
func (b *Bootstrapper) StripeClient() *client.API {
	b.stripeClientLock.Lock()
	defer b.stripeClientLock.Unlock()

	if b.stripeClient != nil {
		return b.stripeClient
	}

	// Set the global Stripe API key
	stripe.Key = b.stripeSecretKey

	// Set app info for debugging
	stripe.SetAppInfo(&stripe.AppInfo{
		Name:    "stripe-samples/accept-a-payment/payment-element",
		Version: "0.0.2",
		URL:     "https://github.com/stripe-samples",
	})

	// Create and cache the Stripe client
	b.stripeClient = &client.API{}
	b.stripeClient.Init(b.stripeSecretKey, nil)

	return b.stripeClient
}

// StripePublishableKey returns the Stripe publishable key
func (b *Bootstrapper) StripePublishableKey() string {
	return b.stripePublishableKey
}

// StripeWebhookSecret returns the Stripe webhook secret
func (b *Bootstrapper) StripeWebhookSecret() string {
	return b.stripeWebhookSecret
}
