package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/backend-go-tuto/internal/user"
)

var usernamePassword = map[string]string{
	"maximilien": "1234",
	"john":       "456",
}

const userKeyContext = "user"

func (s Server) Authenticate(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	passwordStored, found := usernamePassword[username]
	if !found {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if password != passwordStored {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set(userKeyContext, user.User{ID: username})
	c.Next()
}

func (s Server) currentUser(c *gin.Context) (user.User, error) {
	userContext, exists := c.Get(userKeyContext)
	if !exists {
		return user.User{}, errors.New("no user was retrieved in the context")
	}
	userRetrieved, ok := userContext.(user.User)
	if !ok {
		return user.User{}, errors.New("impossible to convert the user in context into a user.User")
	}
	return userRetrieved, nil
}
