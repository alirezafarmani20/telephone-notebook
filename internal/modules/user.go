package modules

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// create a module for users
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	PhoneNumber string  `gorm:"unique"`
}