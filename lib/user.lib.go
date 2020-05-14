package lib

import (
	"github.com/m-ssilva/api.go.gingonic/models"
	"github.com/m-ssilva/api.go.gingonic/repositories"
	"github.com/pkg/errors"
)

// UserLib is lib layer of all user endpoints
func UserLib(user models.User) error {
	err := repositories.CreateUser(user)
	if err != nil {
		return errors.New("EMAIL_ALREADY_USED")
	}
	return err
}
