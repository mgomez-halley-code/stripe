package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v83"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	stripeClient   *stripe.Client
	publishableKey string
	webhookSecret  string
	logger         *logrus.Logger
}

// NewHandler creates a new handler instance
func NewHandler(stripeClient *stripe.Client, publishableKey, webhookSecret string, logger *logrus.Logger) *Handler {
	return &Handler{
		stripeClient:   stripeClient,
		publishableKey: publishableKey,
		webhookSecret:  webhookSecret,
		logger:         logger,
	}
}

// ErrorResponseMessage represents the structure of the error object sent in failed responses
type ErrorResponseMessage struct {
	Message string `json:"message"`
}

// ErrorResponse represents the structure of the error object sent in failed responses
type ErrorResponse struct {
	Error *ErrorResponseMessage `json:"error"`
}

// HandleConfig returns the Stripe publishable key
func (h *Handler) HandleConfig(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, struct {
		PublishableKey string `json:"publishableKey"`
	}{
		PublishableKey: h.publishableKey,
	})
}

// writeJSON writes a JSON response
func (h *Handler) writeJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.WithError(err).Error("Failed to encode JSON response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		h.logger.WithError(err).Error("Failed to write JSON response")
	}
}

// writeJSONError writes a JSON error response
func (h *Handler) writeJSONError(w http.ResponseWriter, v interface{}, code int) {
	w.WriteHeader(code)
	h.writeJSON(w, v)
}

// writeJSONErrorMessage writes a JSON error message response
func (h *Handler) writeJSONErrorMessage(w http.ResponseWriter, message string, code int) {
	resp := &ErrorResponse{
		Error: &ErrorResponseMessage{
			Message: message,
		},
	}
	h.writeJSONError(w, resp, code)
}
