package router

import (
	"challenge-3-simple-request-api-postgre2/controller"
	"github.com/gin-gonic/gin"
)

func SetBookRoutes(router *gin.RouterGroup) {
	router.GET("/books", controller.GetAllBook)
	router.GET("/books/:id", controller.GetBookByID)
	router.POST("/books", controller.CreateBook)
	router.PUT("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)
}
