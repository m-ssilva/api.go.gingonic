package models

import "github.com/jinzhu/gorm"

// User model struct
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"type:varchar(100)"`
}
