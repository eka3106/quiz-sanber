package book

import "github.com/gin-gonic/gin"

func RouterBook(router *gin.Engine) {
	api := router.Group("/api")
	api.GET("/books", GetAllBooks)
	api.POST("/books", CreateBook)
	api.GET("/books/:id", GetBook)
	api.DELETE("/books/:id", DeleteBook)
}
