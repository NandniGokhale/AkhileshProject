package main

import (
	"fmt"
	"github.com/akhilesh-ingle-ge/pkg/controller"
	"github.com/akhilesh-ingle-ge/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDB()
	router := gin.Default()

	router.GET("/books", controller.GetAllBooks)
	router.GET("/books/test", controller.GetAllBooks)
	router.GET("/books/test/new", controller.GetAllBooks)
	router.GET("/books/:id", controller.GetBookById)
	router.POST("/books", controller.CreateBooks)
	router.PATCH("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)
	fmt.Println("Golden Eagle It")

	// log.Fatal(http.ListenAndServe(":8080", router))
	router.Run(":8080")
}
