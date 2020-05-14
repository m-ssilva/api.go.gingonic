package server

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/m-ssilva/api.go.gingonic/models"

	// Mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateDBConnection returns a database connection
func CreateDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/apiGoLang")
	if err != nil {
		log.Printf("Connection to database failed, error: %v", err)
	}
	db.AutoMigrate(&models.User{})
	return db
}
