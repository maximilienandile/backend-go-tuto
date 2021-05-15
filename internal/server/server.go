package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maximilienandile/backend-go-tuto/internal/uniqueid"

	"github.com/maximilienandile/backend-go-tuto/internal/storage"

	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
)

type Server struct {
	Engine            *gin.Engine
	port              uint
	allowedOrigin     string
	storage           storage.Storage
	uniqueIDGenerator uniqueid.Generator
}

type Config struct {
	Port              uint
	AllowedOrigin     string
	Storage           storage.Storage
	UniqueIDGenerator uniqueid.Generator
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	s := &Server{
		Engine:            engine,
		port:              config.Port,
		allowedOrigin:     config.AllowedOrigin,
		storage:           config.Storage,
		uniqueIDGenerator: config.UniqueIDGenerator,
	}
	engine.Use(s.CORSMiddleware, s.MiddlewareServerModel, s.CheckRequest)
	// Create a new middleware
	// this middleware add a Header to the response
	// Header Name : X-Server-Model
	// Header value should be : Gin
	engine.GET("/categories", s.Categories)
	engine.GET("/products", s.Products)
	engine.POST("/admin/products", s.CreateProduct)
	engine.POST("/admin/categories", s.CreateCategories)
	engine.PUT("/admin/inventory", s.UpdateInventory)
	return s, nil
}

func (s *Server) Run() error {
	return s.Engine.Run(fmt.Sprintf(":%d", s.port))
}

func (s Server) CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", s.allowedOrigin)
}

func (s Server) MiddlewareServerModel(c *gin.Context) {
	c.Header("X-Server-Model", "Gin")
}

func (s Server) CheckRequest(c *gin.Context) {
	authValue := c.GetHeader("Authorization")
	if authValue != "ABC" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

}

func (s *Server) Categories(c *gin.Context) {
	categories, err := s.storage.Categories()
	if err != nil {
		log.Printf("impossible to get the products: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (s *Server) Products(c *gin.Context) {
	products, err := s.storage.Products()
	if err != nil {
		log.Printf("impossible to get the products: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, products)
}

func (s *Server) CreateProduct(c *gin.Context) {
	var productToAdd product.Product
	err := c.BindJSON(&productToAdd)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	productToAdd.ID = s.uniqueIDGenerator.Generate()
	err = s.storage.CreateProduct(productToAdd)
	if err != nil {
		//
		log.Printf("error occured while saving the product: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "impossible to persist product"})
		return
	}
	c.JSON(http.StatusOK, productToAdd)
}

func (s *Server) CreateCategories(c *gin.Context) {
	var categoryToSave category.Category
	err := c.BindJSON(&categoryToSave)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	categoryToSave.ID = s.uniqueIDGenerator.Generate()
	err = s.storage.CreateCategory(categoryToSave)
	if err != nil {
		log.Printf("error occured while saving the catgory: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "impossible to persist category"})
		return
	}
	c.JSON(http.StatusOK, categoryToSave)
}
