package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckoutInput struct {
	Provider            string `json:"provider"`
	ShippingCountryCode string `json:"shippingCountryCode"`
	Currency            string `json:"currency"`
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
	log.Println(cartRetrieved)
	c.Status(http.StatusOK)
}
