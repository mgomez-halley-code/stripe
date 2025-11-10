package main

import (
	"fmt"
	"net/http"

	"github.com/stripe-payment-server/internal/api"
	"github.com/stripe-payment-server/internal/bootstrap"
	"github.com/stripe-payment-server/internal/config"
)

func main() {
	// Setup logger
	logger := bootstrap.SetupLogger()

	// Load configuration
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		logger.Fatalf("Configuration validation failed: %v", err)
	}

	// Initialize bootstrapper with dependencies
	boot := bootstrap.NewBootstrapper(
		cfg.Stripe.SecretKey,
		cfg.Stripe.PublishableKey,
		cfg.Stripe.WebhookSecret,
	)

	// Initialize Stripe client (lazy loaded)
	stripeClient := boot.StripeClient()

	// Create API handler with stripe client and logger
	h := api.NewHandler(
		stripeClient,
		boot.StripePublishableKey(),
		boot.StripeWebhookSecret(),
		logger,
	)

	// Create and configure router with logger
	r := api.NewRouter(h, cfg.Server.StaticDir, logger)

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	logger.WithFields(map[string]interface{}{
		"address": addr,
		"router":  "gorilla/mux",
	}).Info("Starting server")

	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
