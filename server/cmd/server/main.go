package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stripe-payment-server/internal/api"
	"github.com/stripe-payment-server/internal/bootstrap"
	"github.com/stripe-payment-server/internal/config"
)

func main() {
	// Load configuration
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatal("Configuration validation failed:", err)
	}

	// Initialize bootstrapper with dependencies
	boot := bootstrap.NewBootstrapper(
		cfg.Stripe.SecretKey,
		cfg.Stripe.PublishableKey,
		cfg.Stripe.WebhookSecret,
	)

	// Initialize Stripe client (lazy loaded)
	stripeClient := boot.StripeClient()

	// Create API handler with stripe client
	h := api.NewHandler(
		stripeClient,
		boot.StripePublishableKey(),
		boot.StripeWebhookSecret(),
	)

	// Create and configure router
	r := api.NewRouter(h, cfg.Server.StaticDir)

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server running at %s with Gorilla Mux", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
