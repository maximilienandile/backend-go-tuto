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
	c.Status(http.StatusOK)
}
