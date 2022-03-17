package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickolasdaniel/test-golang-versioning/models"
)

// GET /books
// Get all books

/* Create struct for valid input */
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"requried"`
}

/* Same struct as CreateBookInput except the binding is not required anymore */
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

/* Function to find all the existing books */
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

/* Function to create a particular book */
func CreateBook(c *gin.Context) {
	//Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

/* Function to find a particular book */
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

/* Function to update a particular book */
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

/* Function to delete a particular book */
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
