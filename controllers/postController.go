package controllers

import (
	"github.com/Rajanhub/goapi/models"
	"github.com/Rajanhub/goapi/services"
	"github.com/gin-gonic/gin"
)

// PostController data type
type PostController struct {
	service *services.PostService
}

// NewPostController creates new post controller
func NewPostController(postService *services.PostService) PostController {
	return PostController{
		service: postService,
	}
}

// GetPost gets the post
func (u *PostController) GetPost(c *gin.Context) {
	posts, err := u.service.GetAllPost()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// SavePost saves the post
func (u *PostController) SavePost(c *gin.Context) {
	post := models.Post{}
	if err := c.Bind(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := u.service.Create(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": "post created"})
}
