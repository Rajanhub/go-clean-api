package controllers

import (
	"github.com/Rajanhub/goapi/api_errors"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/models"
	"github.com/Rajanhub/goapi/services"
	"github.com/Rajanhub/goapi/utils"
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
	posts, err := u.service.SetPaginationScope(utils.Paginate(c)).GetAllPost()
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

// GetOnePost gets one post
func (u *PostController) GetOnePost(c *gin.Context) {
	paramID := c.Param("id")

	postID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	post, err := u.service.GetOnePost(postID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})

}

// UpdatePost updates post
func (u *PostController) UpdatePost(c *gin.Context) {
	paramID := c.Param("id")

	postID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	post, err := u.service.GetOnePost(postID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// if err := lib.CustomBind(c.Request, &post); err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	if err := u.service.UpdatePost(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": post})
}

// DeletePost deletes post
func (u *PostController) DeletePost(c *gin.Context) {
	paramID := c.Param("id")

	postID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	if err := u.service.DeletePost(postID); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "post deleted"})
}
