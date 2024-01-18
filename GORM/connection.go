package gormEP5

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"      // or the Docker service name if running in another container
	port     = 5432             // default PostgreSQL port
	user     = "mydearuser"     // as defined in docker-compose.yml
	password = "mydearpassword" // as defined in docker-compose.yml
	dbname   = "mydeardatabase" // as defined in docker-compose.yml
)

var Db *gorm.DB // Declare db as a package-level variable

func ConnectDB() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})

	if err != nil {
		//panic : break process!!
		panic("failed to connect to database")
	}
	Db = dbs
	// Migrate the schema
	Db.AutoMigrate(&Book{}) //AutoMigrate : Compare DB and Stuck //unsupport modify colums
	fmt.Println("Database migration completed!")
}
