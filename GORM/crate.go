package gormEP5

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *Book) {
	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error Creating book: %v",result.Error)
	}
	fmt.Println("Create Book Successful")
}