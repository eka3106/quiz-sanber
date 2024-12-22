package category

import (
	"quiz/modules/user"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	result, status, err := GetAll()
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
	}
	c.JSON(status, gin.H{"data": result})
}

func CreateCategory(c *gin.Context) {
	if c.ContentType() == "multipart/form-data" {
		c.JSON(400, gin.H{"error": "multipart form is not allowed"})
		return
	}
	_, isExist := c.Get("claims")
	if !isExist {
		c.JSON(401, gin.H{"error": "Not authorized"})
		return
	}
	var category Category
	c.BindJSON(&category)
	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(400, gin.H{"error": "Not authorized"})
		return
	}
	username := claims.(*user.Claims).Username
	category.Created_by = username
	category.Modified_by = username

	status, err := Create(category)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
	} else {
		c.JSON(status, gin.H{"message": "success create category"})
	}
}

func UpdateCategory(c *gin.Context) {
	if c.ContentType() == "multipart/form-data" {
		c.JSON(400, gin.H{"error": "multipart form is not allowed"})
		return
	}
	_, isExist := c.Get("claims")
	if !isExist {
		c.JSON(401, gin.H{"error": "Not authorized"})
		return
	}
	var category Category
	var err error
	id := c.Param("id")
	category.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	c.BindJSON(&category)
	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(400, gin.H{"error": "Not authorized"})
		return
	}
	username := claims.(*user.Claims).Username
	category.Modified_by = username
	category.Modified_at = time.Now().Format(time.RFC3339)

	status, err := Update(category)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
	} else {
		c.JSON(status, gin.H{"message": "success update category"})
	}
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	result, status, err := GetOne(idInt)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"data": result})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	_, isExist := c.Get("claims")
	if !isExist {
		c.JSON(401, gin.H{"error": "Not authorized"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	status, err := Delete(idInt)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
	} else {
		c.JSON(status, gin.H{"message": "success delete category"})
	}
}

func GetBooksByCategory(c *gin.Context) {
	if c.Request.MultipartForm != nil {
		c.JSON(400, gin.H{"error": "multipart form is not allowed"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	result, status, err := GetBooks(idInt)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"data": result})
}
