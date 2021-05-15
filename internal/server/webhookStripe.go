package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"

	"github.com/stripe/stripe-go/v72/webhook"

	"github.com/gin-gonic/gin"
)

const stripeSignatureHeaderName = "Stripe-Signature"
const stripeWebhookCheckoutSessionCompleted = "checkout.session.completed"

func (s Server) HandleStripeWebhook(c *gin.Context) {
	rawWebhookData, err := c.GetRawData()
	if err != nil {
		log.Printf("impossible to retrieve raw data: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	stripeSignature := c.GetHeader(stripeSignatureHeaderName)
	if stripeSignature == "" {
		log.Printf("no signature in the request: %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	event, err := webhook.ConstructEvent(rawWebhookData, stripeSignature, s.stripeWebhookSigningSecretKey)
	if err != nil {
		log.Printf("impossible to construct event: %s", err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if event.Type != stripeWebhookCheckoutSessionCompleted {
		log.Printf("webhook type is equal to %s we only handle this type %s", event.Type, stripeWebhookCheckoutSessionCompleted)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var checkoutSession stripe.CheckoutSession
	err = json.Unmarshal(event.Data.Raw, &checkoutSession)
	if err != nil {
		log.Printf("impossible to unmarshall raw data into a stripe.CheckoutSession: %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

}
