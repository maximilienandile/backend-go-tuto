package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateInventoryInput struct {
	ProductID string `json:"productId"`
	Delta     int    `json:"delta"`
}

func (s *Server) UpdateInventory(c *gin.Context) {
	var updateInput UpdateInventoryInput
	err := c.BindJSON(&updateInput)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	err = s.storage.UpdateInventory(updateInput.ProductID, updateInput.Delta)
	if err != nil {
		//
		log.Printf("error occured while saving the product: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "impossible to persist product"})
		return
	}
	c.Status(http.StatusNoContent)

}
