package main

import (
	"quiz/databases"
	"quiz/middleware"
	"quiz/modules/book"
	"quiz/modules/category"
	"quiz/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	defer databases.DB.Close()

	router := gin.Default()

	//static file
	router.Static("/uploads", "./uploads")

	// route user

	router.Use(middleware.AuthJWT())

	user.RouterUser(router)

	// route category
	category.RouterCategory(router)

	// route book
	book.RouterBook(router)

	router.Run(":3000")
}
