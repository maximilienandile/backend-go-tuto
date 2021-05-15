package server

import (
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
)

type Server struct {
	engine *gin.Engine
}

type Config struct {
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	s := &Server{
		engine: engine,
	}
	engine.GET("/categories", s.Categories)
	engine.GET("/products", s.Products)
	return s, nil
}

func (s *Server) Run() error {
	return s.engine.Run(":9090")
}

func (s *Server) Categories(c *gin.Context) {
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
}

func (s *Server) Products(c *gin.Context) {
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
}
