package routes

import (
	"github.com/Rajanhub/goapi/controllers"
	"github.com/Rajanhub/goapi/infrastructure"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/middlewares"
)

// UserRoutes struct
type UserRoutes struct {
	logger               lib.Logger
	handler              infrastructure.Router
	userController       controllers.UserController
	authMiddleware       middlewares.FirebaseAuthMiddleware
	PaginationMiddleware middlewares.PaginationMiddleware
	uploadMiddleware     middlewares.UploadMiddleware
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger lib.Logger,
	router infrastructure.Router,
	userController controllers.UserController,
	authMiddleware middlewares.FirebaseAuthMiddleware,
	PaginationMiddleware middlewares.PaginationMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
) UserRoutes {
	return UserRoutes{
		logger:               logger,
		handler:              router,
		userController:       userController,
		authMiddleware:       authMiddleware,
		PaginationMiddleware: PaginationMiddleware,
		uploadMiddleware:     uploadMiddleware,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Group("/api") //.Use(s.authMiddleware.HandleAuthWithRole(constants.RoleIsAdmin))
	api.POST("/user", s.userController.SaveUser)
	api.GET("/user", s.PaginationMiddleware.Handle(), s.userController.GetUser)
	api.GET("/user/:id", s.userController.GetOneUser)
	api.PUT("/user/:id", s.userController.UpdateUser)
	api.DELETE("/user/:id", s.userController.DeleteUser)
	api.PUT("/user/upload/:id",
		s.uploadMiddleware.Push(s.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
		s.userController.UploadProfilePic)
}
