package main

import (
	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, product.Product{
			ID:               "42",
			Name:             "Test",
			Description:      "This is my product",
			PriceVATExcluded: money.New(100, "EUR"),
			VAT:              money.New(200, "EUR"),
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
