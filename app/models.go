package app

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code       string
	Price      *int `gorm:"not null"`
	Customer   Customer
	CustomerId int
}

type Customer struct {
	gorm.Model
	Name        string
	PhoneNumber string
}
