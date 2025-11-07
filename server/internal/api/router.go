package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stripe-payment-server/internal/middleware"
)

// NewRouter creates and configures the HTTP router
func NewRouter(h *Handler, staticDir string) *mux.Router {
	r := mux.NewRouter()

	// Add middleware (order matters - logging first, then CORS)
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	// Define API routes
	r.HandleFunc("/config", h.HandleConfig).Methods("GET", "OPTIONS")
	r.HandleFunc("/create-payment-intent", h.HandleCreatePaymentIntent).Methods("POST", "OPTIONS")
	r.HandleFunc("/webhook", h.HandleWebhook).Methods("POST", "OPTIONS")

	// Static file server as fallback
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	return r
}
