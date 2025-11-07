package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig
	Stripe StripeConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port      string
	Host      string
	StaticDir string
}

// StripeConfig holds Stripe API configuration
type StripeConfig struct {
	SecretKey      string
	PublishableKey string
	WebhookSecret  string
}

// Load reads configuration from environment variables
func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Server: ServerConfig{
			Port:      getEnv("PORT", "4242"),
			Host:      getEnv("HOST", "0.0.0.0"),
			StaticDir: getEnv("STATIC_DIR", "../client/dist"),
		},
		Stripe: StripeConfig{
			SecretKey:      getEnv("STRIPE_SECRET_KEY", ""),
			PublishableKey: getEnv("STRIPE_PUBLISHABLE_KEY", ""),
			WebhookSecret:  getEnv("STRIPE_WEBHOOK_SECRET", ""),
		},
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Validate checks if required configuration values are present
func (c *Config) Validate() error {
	if c.Stripe.SecretKey == "" {
		log.Fatal("STRIPE_SECRET_KEY is required")
	}
	if c.Stripe.PublishableKey == "" {
		log.Fatal("STRIPE_PUBLISHABLE_KEY is required")
	}
	return nil
}
