package repositories

import (
	"github.com/m-ssilva/api.go.gingonic/models"
	"github.com/m-ssilva/api.go.gingonic/server"
)

var dbConn = server.CreateDBConnection()

// CreateUser is a repository layer that makes access to database
func CreateUser(user models.User) error {
	result := dbConn.Create(&models.User{Name: user.Name, Email: user.Email, Password: user.Password})
	err := result.Error
	return err
}
