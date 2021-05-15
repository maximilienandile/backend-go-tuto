package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) GetCartOfUser(c *gin.Context) {
	userFound, err := s.currentUser(c)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.JSON(http.StatusOK, userFound.ID)

}

func (s Server) UpdateCartOfUser(c *gin.Context) {

}
