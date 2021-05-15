package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maximilienandile/backend-go-tuto/internal/email"

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
	if checkoutSession.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
		// the payment was a success
		// retrieve the checkout session from the DB
		sessionRetrieved, err := s.storage.GetCheckoutSession(checkoutSession.ID)
		if err != nil {
			log.Printf("impossible to find checkout session in DB: %s", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = s.storage.DeleteCart(sessionRetrieved.User.ID)
		if err != nil {
			log.Printf("impossible to delete the cart: %s", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		input := email.OrderConfirmationEmailInput{
			ProductName: "Gopher E-commerce",
			StoreLink:   s.frontendBaseUrl,
			LogoURL:     "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			OrderItems:  sessionRetrieved.Cart.Items,
		}
		// prepare the order confirmation email
		confirmationEmail, err := email.NewOrderConfirmationEmail(input)
		if err != nil {
			log.Printf("impossible to generate confirmation email: %s", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		log.Println("FROM:", s.emailFrom, "TO", checkoutSession.Customer.Email)
		err = s.emailSender.Send(email.SendInput{
			ToAddress:   checkoutSession.CustomerDetails.Email,
			FromAddress: s.emailFrom,
			HtmlBody:    confirmationEmail.BodyAsHTML,
			TextBody:    confirmationEmail.BodyAsText,
			Subject:     "Order Confirmed",
		})

		if err != nil {
			log.Printf("impossible to send email: %s", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	}

}
