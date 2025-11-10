package api

import (
	"io"
	"net/http"

	"github.com/stripe/stripe-go/v83/webhook"
)

// HandleWebhook processes Stripe webhook events
func (h *Handler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.WithError(err).Error("Failed to read webhook request body")
		return
	}

	event, err := webhook.ConstructEvent(b, r.Header.Get("Stripe-Signature"), h.webhookSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.WithError(err).Error("Failed to construct webhook event")
		return
	}

	h.logger.WithField("event_type", event.Type).Info("Webhook event received")

	switch event.Type {
	case EventCheckoutSessionCompleted:
		h.logger.Info("Checkout session completed")
	case EventPaymentIntentSucceeded:
		h.logger.Info("Payment intent succeeded")
	case EventPaymentIntentFailed:
		h.logger.Warn("Payment intent failed")
	}

	h.writeJSON(w, nil)
}
