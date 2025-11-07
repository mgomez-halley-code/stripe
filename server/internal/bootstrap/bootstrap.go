package bootstrap

import (
	"sync"

	"github.com/stripe/stripe-go/v83"
)

// Bootstrapper handles dependency initialization with lazy loading
type Bootstrapper struct {
	stripeSecretKey      string
	stripePublishableKey string
	stripeWebhookSecret  string

	stripeClient     *stripe.Client
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
func (b *Bootstrapper) StripeClient() *stripe.Client {
	b.stripeClientLock.Lock()
	defer b.stripeClientLock.Unlock()

	if b.stripeClient != nil {
		return b.stripeClient
	}

	// Set app info for debugging (global setting)
	stripe.SetAppInfo(&stripe.AppInfo{
		Name:    "mary-gomez-stripe",
		Version: "1.0.0",
		URL:     "https://github.com/mgomez-halley-code/stripe",
	})

	// Create the new stripe.Client
	b.stripeClient = stripe.NewClient(b.stripeSecretKey)

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
