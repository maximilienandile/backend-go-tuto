package server

import (
	"fmt"
	"net/http"

	"github.com/maximilienandile/backend-go-tuto/internal/user"

	"github.com/gin-gonic/gin"
)

const userKey = "user"

func (s Server) Authenticate(c *gin.Context) {
	userRetrieved, err := s.checkAuthenticationToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set(userKey, userRetrieved)
	c.Next()
}

func (s Server) checkAuthenticationToken(token string) (user.User, error) {
	// TODO : check here the authentication token
	// NOT secure
	return user.User{ID: token}, nil
}

func (s Server) currentUser(c *gin.Context) (user.User, error) {
	userContext, exists := c.Get(userKey)
	if !exists {
		return user.User{}, fmt.Errorf("cannot retrieve user in context of req")
	}
	userRetrieved, ok := userContext.(user.User)
	if !ok {
		return user.User{}, fmt.Errorf("cannot convert user in context into a valid user.User element")
	}
	return userRetrieved, nil
}
