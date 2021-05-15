package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrUpdateCartInput struct {
	ProductID string `json:"productId"`
	Delta     int    `json:"delta"`
}

func (s *Server) CreateOrUpdateCart(c *gin.Context) {
	currentUser, err := s.currentUser(c)
	if err != nil {
		log.Printf("error occured while retrieving current user: %s \n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var input CreateOrUpdateCartInput
	err = c.BindJSON(&input)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	newCart, err := s.storage.CreateOrUpdateCart(currentUser.ID, input.ProductID, input.Delta)
	if err != nil {
		log.Printf("error occured while updating cart in db: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server Error"})
		return
	}
	c.JSON(http.StatusOK, newCart)
}

func (s *Server) GetCart(c *gin.Context) {
	currentUser, err := s.currentUser(c)
	if err != nil {
		log.Printf("error occured while retrieving current user: %s \n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	cart, err := s.storage.GetCart(currentUser.ID)
	if err != nil {
		log.Printf("error occured while updating cart in db: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server Error"})
		return
	}
	c.JSON(http.StatusOK, cart)
}
