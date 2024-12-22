package main

import (
	"quiz/databases"
	_ "quiz/docs"
	"quiz/middleware"
	"quiz/modules/book"
	"quiz/modules/category"
	"quiz/modules/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Quiz API Sanbercode
// @version 1.0
// @description This is a task quiz API Sanbercode
// @description To access the API, you need to register and login first
// @description To get image, you need image url and access it with /uploads/{image_url}
// @host quiz-sanber-production.up.railway.app
// @BasePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	defer databases.DB.Close()

	router := gin.Default()
	router.Use(cors.Default())
	//static file
	router.Static("/uploads", "./uploads")

	router.Use(middleware.AuthJWT())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// route user
	user.RouterUser(router)

	// route category
	category.RouterCategory(router)

	// route book
	book.RouterBook(router)

	router.Run(":3000")
}
