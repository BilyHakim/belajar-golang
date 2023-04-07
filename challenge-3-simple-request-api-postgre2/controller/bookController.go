package controller

import (
	"challenge-3-simple-request-api-postgre2/database"
	"challenge-3-simple-request-api-postgre2/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetAllBook(c *gin.Context) {
	var books []model.Book
	err := database.DB.Find(&books).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	var book model.Book
	err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.Create(&book).Error
	if err != nil {
		err = database.ErrorHandler(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	var book model.Book
	err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Save(&book).Error
	if err != nil {
		err = database.ErrorHandler(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book model.Book
	err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Delete(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
