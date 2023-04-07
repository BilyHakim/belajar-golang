package model

import "github.com/jinzhu/gorm"

// Model Buku
type Book struct {
	gorm.Model
	Title       string `gorm:"not null;unique"`
	Author      string `gorm:"not null"`
	Description string `gorm:"not null"`
}
