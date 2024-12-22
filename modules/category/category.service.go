package category

import (
	"quiz/modules/book"
	"quiz/modules/user"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCategories struct {
	Data  []Category `json:"data"`
	Error error      `json:"error,omitempty"`
}
type ResponseOneCategories struct {
	Data  Category `json:"data"`
	Error error    `json:"error,omitempty"`
}

type ResponseBooksByCategory struct {
	Data  []book.Books `json:"data"`
	Error error        `json:"error,omitempty"`
}

// GetAllCategories is a function to get all categories godoc
// @Summary Get all categories
// @Description Get all categories in database
// @Tags Categories
// @Accept json
// @Produce json
// @Security None
// @Success 200 {object} ResponseCategories
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	result, status, err := GetAll()
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
	}
	c.JSON(status, gin.H{"data": result})
}

// CreateCategory is a function to create a category godoc
// @Summary Create a category
// @Description Create a category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Security Bearer
// @Param name body string true "Name"
// @Success 201 {string} message: "success create category"
// @Router /categories [post]
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

// UpdateCategory is a function to update a category godoc
// @Summary Update a category
// @Description Update a category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Success 200 {string} message: "success update category"
// @Router /categories/{id} [put]
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

// GetCategory is a function to get a category godoc
// @Summary Get a category
// @Description Get a category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Security None
// @Success 200 {object} ResponseOneCategories
// @Router /categories/{id} [get]
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

// DeleteCategory is a function to delete a category godoc
// @Summary Delete a category
// @Description Delete a category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Security Bearer
// @Success 200 {string} message: "success delete category"
// @Router /categories/{id} [delete]
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

// GetBooksByCategory is a function to get books by category godoc
// @Summary Get books by category
// @Description Get books by category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Security None
// @Success 200 {object} ResponseBooksByCategory
// @Router /categories/{id}/books [get]
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
