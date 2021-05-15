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

func (s Server) UpdateCartOfUser(c *gin.Context) {

}
