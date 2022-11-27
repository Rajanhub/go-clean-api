package routes

import (
	"log"

	"github.com/Rajanhub/goapi/controllers"
	"github.com/Rajanhub/goapi/lib"
)

// PostRoutes struct
type PostRoutes struct {
	handler        lib.RequestHandler
	postController controllers.PostController
}

// Setup post routes
func (s PostRoutes) Setup() {
	log.Println("Setting up routes")
	api := s.handler.Gin
	api.POST("/post", s.postController.SavePost)
	api.GET("/post", s.postController.GetPost)

}

// NewPostRoutes creates new post controller
func NewPostRoutes(
	handler lib.RequestHandler,
	postController controllers.PostController,

) PostRoutes {
	return PostRoutes{
		handler:        handler,
		postController: postController,
	}
}
