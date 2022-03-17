package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickolasdaniel/golang-ginrest/controllers"
	"github.com/nickolasdaniel/golang-ginrest/models"
)

func main() {
	/* Use default gin router options */
	r := gin.Default()

	/* Connect to the database */
	models.ConnectDatabase()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello"})
	})
	/* Define the routes and its specific http request handler */
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
