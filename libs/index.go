package libs

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func CleanText(text string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return regex.ReplaceAllString(text, "")
}

func ExtractToken(c *gin.Context) string {
	bearer := c.GetHeader("Authorization")
	token := strings.Replace(bearer, "Bearer ", "", 1)
	return token

}
