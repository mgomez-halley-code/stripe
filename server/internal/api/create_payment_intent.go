package api

import (
	"net/http"

	"github.com/stripe/stripe-go/v83"
)

// HandleCreatePaymentIntent creates a new Stripe payment intent
func (h *Handler) HandleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	params := &stripe.PaymentIntentCreateParams{
		Amount:   stripe.Int64(1999),
		Currency: stripe.String("EUR"),
		AutomaticPaymentMethods: &stripe.PaymentIntentCreateAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Use the new client-based API
	pi, err := h.stripeClient.V1PaymentIntents.Create(r.Context(), params)
	if err != nil {
		// Try to safely cast a generic error to a stripe.Error
		if stripeErr, ok := err.(*stripe.Error); ok {
			h.logger.WithError(stripeErr).Error("Stripe error creating payment intent")
			h.writeJSONErrorMessage(w, stripeErr.Error(), 400)
		} else {
			h.logger.WithError(err).Error("Unknown error creating payment intent")
			h.writeJSONErrorMessage(w, "Unknown server error", 500)
		}
		return
	}

	h.logger.WithField("payment_intent_id", pi.ID).Info("Payment intent created successfully")
	h.writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}
