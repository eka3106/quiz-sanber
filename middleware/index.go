package middleware

import (
	"quiz/config"
	"quiz/libs"
	"quiz/modules/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT() gin.HandlerFunc {
	return verifyToken
}

func verifyToken(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		c.Next()
		return
	}
	token := libs.ExtractToken(c)
	claims := &user.Claims{}
	// verify token
	secret := []byte(config.VarConfig.SecretJwt)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		c.Next()
		return
	}
	print(claims.Id)
	isExist := user.CheckToken(token, claims.Id)
	if !isExist {
		c.Next()
		println("error")
		return
	}
	println("No error")
	c.Set("claims", claims)
	c.Next()
}
