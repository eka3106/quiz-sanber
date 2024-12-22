package user

import "github.com/gin-gonic/gin"

func RouterUser(router *gin.Engine) {
	api := router.Group("/api/users")
	api.POST("/register", CreateUser)
	api.POST("/login", Login)
	api.POST("/logout", Logout)
}
