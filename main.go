package main

import gormEP5 "DEAREP5/GORM"


func main() {
	gormEP5.ConnectDB()

	
newBook := &gormEP5.Book{
	Name: "TEST",
	Author: "Test",
	Description: "TEST",
	Price: 1,
}
	gormEP5.CreateBook(gormEP5.Db,newBook)
}
