package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/maximilienandile/backend-go-tuto/internal/checkout"

	"github.com/stripe/stripe-go/v72/checkout/session"

	"github.com/stripe/stripe-go/v72"

	"github.com/gin-gonic/gin"
)

type CheckoutInput struct {
	Provider            string `json:"provider"`
	ShippingCountryCode string `json:"shippingCountryCode"`
	Currency            string `json:"currency"`
}

type CheckoutOutput struct {
	SessionId string `json:"sessionId"`
}

func (s Server) Checkout(c *gin.Context) {
	var input CheckoutInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	// retrieve current user
	currentUser, err := s.currentUser(c)
	if err != nil {
		log.Printf("impossible to retrieve current user: %s", err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	// retrieve cart of the current user
	cartRetrieved, err := s.storage.GetCart(currentUser.ID)
	if err != nil {
		log.Printf("impossible to retrieve cart of current user: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	lines := make([]*stripe.CheckoutSessionLineItemParams, 0)
	for _, item := range cartRetrieved.Items {
		var line stripe.CheckoutSessionLineItemParams
		line.PriceData = &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String(cartRetrieved.CurrencyCode),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Description: stripe.String(item.ShortDescription),
				Name:        stripe.String(item.Title),
			},
			UnitAmount: &item.UnitPriceVATInc.Amount,
		}
		qty := int64(item.Quantity)
		line.Quantity = &qty
		lines = append(lines, &line)
	}

	params := stripe.CheckoutSessionParams{}
	// fill the parameters of the checkout session
	// that is to say which item are we going to sell, for which price, ...
	params.SuccessURL = stripe.String(fmt.Sprintf("%s/#/success", s.frontendBaseUrl))
	params.CancelURL = stripe.String(fmt.Sprintf("%s/#/cart", s.frontendBaseUrl))
	params.Mode = stripe.String("payment")
	params.LineItems = lines
	params.ShippingAddressCollection = &stripe.CheckoutSessionShippingAddressCollectionParams{
		AllowedCountries: []*string{
			stripe.String("FR"), stripe.String("US"), stripe.String("DE"),
		},
	}

	stripe.Key = s.stripeSecretKey
	checkoutSession, err := session.New(&params)
	if err != nil {
		log.Printf("error while building the checkout session: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// create a new checkout session to be stored in the database
	sessionToStore := checkout.Session{
		ID:        checkoutSession.ID,
		CreatedAt: time.Now().Format(time.RFC3339),
		Provider:  "STRIPE",
		Cart:      cartRetrieved,
		User:      currentUser,
	}
	err = s.storage.CreateCheckoutSession(sessionToStore)
	if err != nil {
		log.Printf("error while saving the checkout session to DB: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, CheckoutOutput{SessionId: checkoutSession.ID})
}
