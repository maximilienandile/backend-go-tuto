package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/maximilienandile/backend-go-tuto/internal/email"

	"firebase.google.com/go/v4/auth"

	"github.com/maximilienandile/backend-go-tuto/internal/extMoney"

	"github.com/maximilienandile/backend-go-tuto/internal/uniqueid"

	"github.com/maximilienandile/backend-go-tuto/internal/storage"

	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
)

type Server struct {
	Engine                        *gin.Engine
	port                          uint
	allowedOrigin                 string
	storage                       storage.Storage
	uniqueIDGenerator             uniqueid.Generator
	firebaseAuthClient            *auth.Client
	stripeSecretKey               string
	stripeWebhookSigningSecretKey string
	frontendBaseUrl               string
	emailSender                   email.Sender
}

type Config struct {
	Port                          uint
	AllowedOrigin                 string
	Storage                       storage.Storage
	UniqueIDGenerator             uniqueid.Generator
	FirebaseAuthClient            *auth.Client
	StripeSecretKey               string
	StripeWebhookSigningSecretKey string
	FrontendBaseUrl               string
	EmailSender                   email.Sender
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	s := &Server{
		Engine:                        engine,
		port:                          config.Port,
		allowedOrigin:                 config.AllowedOrigin,
		storage:                       config.Storage,
		uniqueIDGenerator:             config.UniqueIDGenerator,
		firebaseAuthClient:            config.FirebaseAuthClient,
		stripeSecretKey:               config.StripeSecretKey,
		stripeWebhookSigningSecretKey: config.StripeWebhookSigningSecretKey,
		frontendBaseUrl:               config.FrontendBaseUrl,
		emailSender:                   config.EmailSender,
	}
	engine.Use(s.CORSMiddleware, s.MiddlewareServerModel)
	// Create a new middleware
	// this middleware add a Header to the response
	// Header Name : X-Server-Model
	// Header value should be : Gin
	engine.GET("/categories", s.Categories)
	engine.GET("/products", s.Products)
	engine.GET("/product/:id", s.GetProductByID)
	engine.POST("/admin/products", s.CreateProduct)
	engine.PUT("/admin/product/:productId", s.UpdateProduct)
	engine.POST("/admin/categories", s.CreateCategories)
	engine.PUT("/admin/inventory", s.UpdateInventory)
	engine.GET("/me/cart", s.AuthenticateV2, s.GetCartOfUser)
	engine.PUT("/me/cart", s.AuthenticateV2, s.UpdateCartOfUser)
	engine.POST("/checkout", s.AuthenticateV2, s.Checkout)
	engine.POST("/webhooks/stripe", s.HandleStripeWebhook)
	engine.GET("/testEmail", s.SendTestEmail)
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

type UpdateProductInput struct {
	Name             string            `json:"name"`
	Image            string            `json:"image"`
	ShortDescription string            `json:"shortDescription"`
	Description      string            `json:"description"`
	PriceVATExcluded extMoney.ExtMoney `json:"priceVatExcluded"`
	VAT              extMoney.ExtMoney `json:"vat"`
	TotalPrice       extMoney.ExtMoney `json:"totalPrice"`
}

func (s *Server) UpdateProduct(c *gin.Context) {
	var input UpdateProductInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}
	productID := c.Param("productId")
	if productID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "productId is mandatory"})
		return
	}
	err = s.storage.UpdateProduct(storage.UpdateProductInput{
		ProductID:        productID,
		Name:             input.Name,
		Image:            input.Image,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		PriceVATExcluded: input.PriceVATExcluded,
		VAT:              input.VAT,
		TotalPrice:       input.TotalPrice,
	})
	if err != nil {
		log.Printf("error occured while updating the product: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "impossible to update product"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) GetProductByID(c *gin.Context) {
	// first we need to the id
	productID := c.Param("id")
	if productID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "productId is mandatory"})
		return
	}
	productFound, err := s.storage.Product(productID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "product not retrieved"})
			return
		}
		log.Printf("impossible to retrieve product: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, productFound)
}

func (s Server) SendTestEmail(c *gin.Context) {
	err := s.emailSender.Send(email.SendInput{
		ToAddress:   "maximilien.andile.demo@gmail.com",
		FromAddress: "maximilien.andile.demo@gmail.com",
		HtmlBody:    "TEST",
		TextBody:    "TEST",
		Subject:     "This is a test from the backend",
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
