package book

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllBooks(c *gin.Context) {
	result, status, err := GetAll()
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
	}
	c.JSON(status, gin.H{"data": result})
}

func CreateBook(c *gin.Context) {
	if c.ContentType() == "multipart/form-data" {
		var book Books
		var err error
		_, isExist := c.Get("claims")
		if !isExist {
			c.JSON(401, gin.H{"error": "Not authorized"})
			return
		}
		book.Release_year, err = strconv.Atoi(c.PostForm("release_year"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid release year"})
			return
		}
		if book.Release_year < 1980 || book.Release_year > 2024 {
			c.JSON(400, gin.H{"error": "release year must be between 1980 and 2024"})
			return
		}

		book.Total_page, err = strconv.Atoi(c.PostForm("total_page"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid total page"})
			return
		}

		if book.Total_page > 100 {
			book.Thickness = "tebal"
		} else {
			book.Thickness = "tipis"
		}

		book.Price, err = strconv.Atoi(c.PostForm("price"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		book.Title = c.PostForm("title")
		book.Description = c.PostForm("description")
		book.Created_by = c.PostForm("created_by")
		book.Modified_by = c.PostForm("modified_by")

		book.Category_id, err = strconv.Atoi(c.PostForm("category_id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid category id"})
			return
		}
		file, err := c.FormFile("foto_buku")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		newFileName := uuid.New().String() + filepath.Ext(file.Filename)
		filePath := filepath.Join("uploads", newFileName)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		book.Image_url = filePath

		status, err := Create(book)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
		} else {
			c.JSON(status, gin.H{"message": "success update book"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Must be MultipartForm"})
		return
	}
}

func UpdateBook(c *gin.Context) {
	if c.ContentType() == "multipart/form-data" {
		var book Books
		var err error
		_, isExist := c.Get("claims")
		if !isExist {
			c.JSON(401, gin.H{"error": "Not authorized"})
			return
		}
		id := c.Param("id")
		book.Id, err = strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		book.Release_year, err = strconv.Atoi(c.PostForm("release_year"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid release year"})
			return
		}
		if book.Release_year < 1980 || book.Release_year > 2024 {
			c.JSON(400, gin.H{"error": "release year must be between 1980 and 2024"})
			return
		}

		book.Total_page, err = strconv.Atoi(c.PostForm("total_page"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid total page"})
			return
		}

		if book.Total_page > 100 {
			book.Thickness = "tebal"
		} else {
			book.Thickness = "tipis"
		}
		book.Modified_at = time.Now().Format(time.RFC3339)
		book.Price, err = strconv.Atoi(c.PostForm("price"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		book.Title = c.PostForm("title")
		book.Description = c.PostForm("description")
		book.Created_by = c.PostForm("created_by")
		book.Modified_by = c.PostForm("modified_by")

		book.Category_id, err = strconv.Atoi(c.PostForm("category_id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid category id"})
			return
		}
		fileLink := c.PostForm("foto_buku_link")
		if fileLink == "" {
			file, err := c.FormFile("foto_buku")
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			newFileName := uuid.New().String() + filepath.Ext(file.Filename)
			filePath := filepath.Join("uploads", newFileName)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			book.Image_url = filePath
		} else {
			book.Image_url = fileLink
		}
		status, err := Update(book)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
		} else {
			c.JSON(status, gin.H{"message": "success update book"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Must be MultipartForm"})
		return
	}
}

func GetBook(c *gin.Context) {
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

func DeleteBook(c *gin.Context) {
	_, isExist := c.Get("claims")
	if !isExist {
		c.JSON(401, gin.H{"error": "Not authorized"})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	status, err := Delete(idInt)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"message": "success delete book"})
}
