package server

import (
	"errors"
	"log"
	"net/http"
	"strings"

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

func (s Server) AuthenticateV2(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	// 'Bearer XXXXX'
	splits := strings.Split(authorizationHeader, " ")
	if len(splits) != 2 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	// splits[0] = 'Bearer'
	// splits[1] = 'XXXXX' < the idToken
	token, err := s.firebaseAuthClient.VerifyIDToken(c.Request.Context(), splits[1])
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	emailFound, found := token.Claims["email"]
	if !found {
		log.Println("impossible to find email claim")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	email, ok := emailFound.(string)
	if !ok {
		log.Println("impossible to check that concrete type under email claim is a string")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	isVerifiedFound, found := token.Claims["email_verified"]
	if !found {
		log.Println("impossible to find email_verified claim")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	emailVerified, ok := isVerifiedFound.(bool)
	if !ok {
		log.Println("impossible to check that concrete type under email_verified claim is a bool")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set(userKeyContext, user.User{
		ID:             token.UID,
		SigninProvider: token.Firebase.SignInProvider,
		Email:          email,
		EmailVerified:  emailVerified,
	})
	c.Next()
}

func (s Server) AuthenticateAdmin(c *gin.Context) {
	currentUser, err := s.currentUser(c)
	if err != nil {
		log.Println("impossible to retrieve user in context")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	var adminEmailFound bool
	for _, adminEmail := range s.adminEmails {
		if currentUser.Email == adminEmail {
			adminEmailFound = true
			break
		}
	}
	if !adminEmailFound {
		log.Println("impossible to find email in the admin list")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if !currentUser.EmailVerified {
		log.Println("email in admin list but not verified")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if currentUser.SigninProvider != "google.com" {
		log.Println("signin provider is not google.com")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
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
