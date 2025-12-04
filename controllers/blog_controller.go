package controllers

import (
	"go-blog/config"
	"go-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Blog Post
func CreateBlog(c *gin.Context) {
	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// Get All Blogs
func GetBlogs(c *gin.Context) {
	var blogs []models.Blog
	config.DB.Find(&blogs)
	c.JSON(http.StatusOK, blogs)
}

// Get Blog by ID
func GetBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	if err := config.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// Update Blog
func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	if err := config.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	var data models.Blog
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&blog).Updates(data)

	c.JSON(http.StatusOK, blog)
}

// Delete Blog
func DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Blog{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}
