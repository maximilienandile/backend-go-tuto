package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/maximilienandile/backend-go-tuto/internal/storage"

	"github.com/gin-gonic/gin"
)

func (s Server) GetCartOfUser(c *gin.Context) {
	userFound, err := s.currentUser(c)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	cartRetrieved, err := s.storage.GetCart(userFound.ID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Printf("impossible to retrieve the cart: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, cartRetrieved)

}

type UpdateCartOfUserInput struct {
	ProductID string `json:"productId"`
	Delta     int    `json:"delta"`
}

func (s Server) UpdateCartOfUser(c *gin.Context) {
	var input UpdateCartOfUserInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	currentUser, err := s.currentUser(c)
	if err != nil {
		log.Printf("impossible to retrieve current user: %s", err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	cartUpdated, err := s.storage.CreateOrUpdateCart(currentUser.ID, input.ProductID, input.Delta)
	if err != nil {
		log.Printf("impossible to create or update cart: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, cartUpdated)
}
