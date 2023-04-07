package main

import (
	"challenge-3-simple-request-api-postgre2/database"
	"challenge-3-simple-request-api-postgre2/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			book := v1.Group("/book")
			{
				router.SetBookRoutes(book)
			}
		}
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
