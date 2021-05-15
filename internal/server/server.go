package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
)

type Server struct {
}

type Config struct {
}

func New(config Config) (*Server, error) {
	return &Server{}, nil
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
