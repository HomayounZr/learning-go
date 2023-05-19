package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func main() {
	var err error
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("faile to connect to database")
	}

	db.AutoMigrate(&Book{})

	handler := newHandler(db)

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/books", handler.listBookHandler)

	r.POST("/books", handler.createBookHandler)

	r.DELETE("/books/:id", handler.deleteBookHandler)

	r.Run()
}

func (h *Handler) listBookHandler(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(auth, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var books []Book

	if result := h.db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &books)
}

func (h *Handler) createBookHandler(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(auth, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if result := h.db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &book)
}

func (h *Handler) deleteBookHandler(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(auth, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id := c.Param("id")

	if result := h.db.Delete(&Book{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func validateToken(token string) error {
	if token != "ACCESS_TOKEN" {
		return fmt.Errorf("provided token was invalid")
	}
	return nil
}
