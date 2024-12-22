package user

import (
	"quiz/libs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	if c.Request.MultipartForm != nil {
		c.JSON(400, gin.H{"error": "multipart form is not allowed"})
	} else {
		var user User
		c.BindJSON(&user)
		bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			user.Password = string(bytes)
			err = Create(user)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			} else {
				c.JSON(201, gin.H{"message": "success"})
			}
		}
	}
}

func Login(c *gin.Context) {
	if c.Request.MultipartForm != nil {
		c.JSON(400, gin.H{"error": "multipart form is not allowed"})
	} else {
		var user User
		c.BindJSON(&user)
		username := libs.CleanText(user.Username)
		user.Created_by = username
		user.Modified_by = username
		result, status, err := GetToken(username, user.Password)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
		} else {
			c.JSON(status, gin.H{"token": result.Token})
		}
	}
}

func Logout(c *gin.Context) {
	if c.Request.MultipartForm != nil {
		c.JSON(400, gin.H{"error": "multipart form is not allowed"})
	} else {
		claims, exist := c.Get("claims")
		if !exist {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}
		token := libs.ExtractToken(c)
		status, err := RemoveToken(token, claims.(*Claims).Id)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
		} else {
			c.JSON(status, gin.H{"message": "success"})
		}
	}
}
