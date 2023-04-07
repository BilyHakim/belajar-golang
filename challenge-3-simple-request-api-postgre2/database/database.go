package database

import (
	"challenge-3-simple-request-api-postgre2/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

var DB *gorm.DB

// Konek
func Connect() error {
	db, err := gorm.Open("postgres", "user=postgres password=latihan dbname=crud_postgre2 sslmode=disable")

	if err != nil {
		return err
	}

	err = db.DB().Ping()
	if err != nil {
		return err
	}
	// Set maximum number of idle connections in the connection pool
	db.DB().SetMaxIdleConns(10)

	// Set maximum number of open connections in the connection pool
	db.DB().SetMaxOpenConns(100)

	// Set connection life time
	db.DB().SetConnMaxLifetime(300)

	// AutoMigrate Models
	db.AutoMigrate(&model.Book{})

	DB = db
	return nil
}

// Error Handling
func ErrorHandler(err error) error {
	if err, ok := err.(*pq.Error); ok {
		switch err.Code.Name() {
		case "unique_violation":
			return fmt.Errorf("Already Exists")
		default:
			return fmt.Errorf("Database Error")
		}
	}
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
