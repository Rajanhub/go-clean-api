package routes

import (
	"github.com/Rajanhub/goapi/controllers"
	"github.com/Rajanhub/goapi/infrastructure"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/middlewares"
)

// PostRoutes struct
type PostRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	postController controllers.PostController
	authMiddleware middlewares.FirebaseAuthMiddleware
}

// NewPostRoutes creates new post controller
func NewPostRoutes(
	logger lib.Logger,
	router infrastructure.Router,
	postController controllers.PostController,
	authMiddleware middlewares.FirebaseAuthMiddleware,

) PostRoutes {
	return PostRoutes{
		logger:         logger,
		handler:        router,
		postController: postController,
		authMiddleware: authMiddleware,
	}
}

// Setup post routes
func (s PostRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Group("/api") //.Use(s.authMiddleware.HandleAuthWithRole(constants.RoleIsAdmin))
	api.POST("/post", s.postController.SavePost)
	api.GET("/post", s.postController.GetPost)
	api.GET("/post/:id", s.postController.GetOnePost)
	api.PUT("/post/:id", s.postController.UpdatePost)
	api.DELETE("/post/:id", s.postController.DeletePost)

}
