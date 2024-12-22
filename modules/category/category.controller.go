package category

import "github.com/gin-gonic/gin"

func RouterCategory(router *gin.Engine) {
	api := router.Group("/api")
	api.GET("/categories", GetAllCategories)
	api.GET("/categories/:id", GetCategory)
	api.GET("/categories/:id/books", GetBooksByCategory)
	api.POST("/categories", CreateCategory)
	api.PUT("/categories/:id", UpdateCategory)
	api.DELETE("/categories/:id", DeleteCategory)
}
