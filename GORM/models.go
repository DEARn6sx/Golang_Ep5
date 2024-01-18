package gormEP5

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       int
}