package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateInventoryInput struct {
	ProductId string `json:"productId"`
	Delta     int    `json:"delta"`
}

func (s *Server) UpdateInventory(c *gin.Context) {
	var input UpdateInventoryInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	err = s.storage.UpdateInventory(input.ProductId, input.Delta)
	if err != nil {
		log.Printf("impossible to update inventory: %s", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "impossible to update inventory"})
		return
	}
	c.Status(http.StatusNoContent)
}
