package app

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code     string
	Price    *int    `gorm:"not null"`
	Company  *string `gorm:"not null"`
	Customer string
}
