package server

import (
	"fmt"
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
)

type Server struct {
	Engine *gin.Engine
	port   uint
}

type Config struct {
	Port uint
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	s := &Server{
		Engine: engine,
		port:   config.Port,
	}
	engine.GET("/categories", s.Categories)
	engine.GET("/products", s.Products)
	return s, nil
}

func (s *Server) Run() error {
	return s.Engine.Run(fmt.Sprintf(":%d", s.port))
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
	twoEuro := money.New(200, "EUR")
	fourEuros := money.New(400, "EUR")
	products := []product.Product{
		{
			ID:               "42",
			Name:             "Test",
			Image:            "https://www.practical-go-lessons.com/img/practical-go-lessons-book10.a8a05387.jpg",
			ShortDescription: "New",
			Description:      "This is my product",
			PriceVATExcluded: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			VAT: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			TotalPrice: product.Amount{
				Money:   fourEuros,
				Display: fourEuros.Display(),
			},
		},
		{
			ID:               "43",
			Name:             "Test",
			Description:      "This is my product",
			Image:            "https://www.practical-go-lessons.com/img/practical-go-lessons-book10.a8a05387.jpg",
			ShortDescription: "New",
			PriceVATExcluded: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			VAT: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			TotalPrice: product.Amount{
				Money:   fourEuros,
				Display: fourEuros.Display(),
			},
		},
		{
			ID:               "44",
			Name:             "Test",
			Image:            "https://www.practical-go-lessons.com/img/practical-go-lessons-book10.a8a05387.jpg",
			ShortDescription: "on Sale !",
			Description:      "This is my product",
			PriceVATExcluded: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			VAT: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			TotalPrice: product.Amount{
				Money:   fourEuros,
				Display: fourEuros.Display(),
			},
		},
	}
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.JSON(200, products)
}
