package controllers

import (
	"log"

	"github.com/Rajanhub/goapi/api_errors"
	"github.com/Rajanhub/goapi/constants"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/models"
	"github.com/Rajanhub/goapi/services"
	"github.com/Rajanhub/goapi/utils"
	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service *services.UserService
}

// NewUserController creates new user controller
func NewUserController(userService *services.UserService) UserController {
	return UserController{
		service: userService,
	}
}

// GetUser gets the user
func (u *UserController) GetUser(c *gin.Context) {
	users, err := u.service.SetPaginationScope(utils.Paginate(c)).GetAllUser()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

// SaveUser saves the user
func (u *UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := u.service.Create(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// GetOneUser gets one user
func (u *UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	user, err := u.service.GetOneUser(userID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// UpdateUser updates user
func (u *UserController) UpdateUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	user, err := u.service.GetOneUser(userID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(c.Request)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// if err := lib.CustomBind(c.Request, &user); err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	if err := u.service.UpdateUser(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u *UserController) UploadProfilePic(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	user, err := u.service.GetOneUser(userID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	metadata, _ := c.MustGet(constants.File).(lib.UploadedFiles)
	user.ProfilePic = lib.SignedURL(metadata.GetFile("file").URL)

	if err := u.service.UpdateUser(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes user
func (u *UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": api_errors.ErrInvalidUUID.Error(),
		})
		return
	}

	if err := u.service.DeleteUser(userID); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
