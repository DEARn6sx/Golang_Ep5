package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"      // or the Docker service name if running in another container
	port     = 5432             // default PostgreSQL port
	user     = "mydearuser"     // as defined in docker-compose.yml
	password = "mydearpassword" // as defined in docker-compose.yml
	dbname   = "mydeardatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic : break process!!
		panic("failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&Book{}) //AutoMigrate : Compare DB and Stuck //unsupport modify colums
	fmt.Println("Database migration completed!")
}
