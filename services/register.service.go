package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-ssilva/api.go.gingonic/lib"
	models "github.com/m-ssilva/api.go.gingonic/models"
)

// RegisterService is service layer of register endpoint
func RegisterService(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := lib.UserLib(user)

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "User created"})
	}
}
