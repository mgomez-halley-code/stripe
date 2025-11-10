package api

// Stripe webhook event types
const (
	EventCheckoutSessionCompleted = "checkout.session.completed"
	EventPaymentIntentSucceeded   = "payment_intent.succeeded"
	EventPaymentIntentFailed      = "payment_intent.payment_failed"
)
