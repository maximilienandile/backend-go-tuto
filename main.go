package main

import (
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.GET("/categories", func(c *gin.Context) {
		categories := []category.Category{
			{
				ID:          "42",
				Name:        "Plushies",
				Description: "kdsjdjsidjisdj",
			},
			{
				ID:          "43",
				Name:        "T-Shirts",
				Description: "kdsjdjsidjisdj",
			},
		}
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		c.JSON(http.StatusOK, categories)
	})
	r.GET("/products", func(c *gin.Context) {
		products := []product.Product{
			{
				ID:               "42",
				Name:             "Test",
				Description:      "This is my product",
				PriceVATExcluded: money.New(100, "EUR"),
				VAT:              money.New(200, "EUR"),
			},
			{
				ID:               "43",
				Name:             "Test",
				Description:      "This is my product",
				PriceVATExcluded: money.New(100, "EUR"),
				VAT:              money.New(250, "EUR"),
			},
			{
				ID:               "44",
				Name:             "Test",
				Description:      "This is my product",
				PriceVATExcluded: money.New(189, "EUR"),
				VAT:              money.New(200, "EUR"),
			},
		}
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		c.JSON(200, products)
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
