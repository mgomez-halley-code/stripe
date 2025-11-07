package api

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/paymentintent"
)

// HandleCreatePaymentIntent creates a new Stripe payment intent
func (h *Handler) HandleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1999),
		Currency: stripe.String("EUR"),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		// Try to safely cast a generic error to a stripe.Error
		if stripeErr, ok := err.(*stripe.Error); ok {
			fmt.Printf("Stripe error occurred: %v\n", stripeErr.Error())
			writeJSONErrorMessage(w, stripeErr.Error(), 400)
		} else {
			fmt.Printf("Other error occurred: %v\n", err.Error())
			writeJSONErrorMessage(w, "Unknown server error", 500)
		}
		return
	}

	writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}
