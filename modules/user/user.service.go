package user

import (
	"quiz/libs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ResponseUser struct {
	Token string `json:"token"`
}

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser is a function to create a user godoc
// @Summary Create a user
// @Description Create a user in database
// @Tags Users
// @Accept json
// @Produce json
// @Security None
// @Param  Register body RequestUser true "Register"
// @Success 201 {string} message: "success"
// @Router /users/register [post]
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

// Login is a function to login godoc
// @Summary Login
// @Description Login in database
// @Tags Users
// @Accept json
// @Produce json
// @Security None
// @Param  Login body RequestUser true "Login"
// @Success 200 {string} token
// @Router /users/login [post]
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

// Logout is a function to logout godoc
// @Summary Logout
// @Description Logout in database
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {string} message: "success"
// @Router /users/logout [post]
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
