package bootstrap

import (
	"context"
	"fmt"

	"github.com/Rajanhub/goapi/controllers"
	"github.com/Rajanhub/goapi/infrastructure"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/repository"
	"github.com/Rajanhub/goapi/routes"
	"github.com/Rajanhub/goapi/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	infrastructure.Module,
	lib.Module,
	repository.Module,
	controllers.Module,
	routes.Module,
	services.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h lib.RequestHandler, postRoute routes.PostRoutes) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application in :5000")
				postRoute.Setup()

				go h.Gin.Run(":5000")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
