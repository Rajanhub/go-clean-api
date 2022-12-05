package bootstrap

import (
	"github.com/Rajanhub/goapi/controllers"
	"github.com/Rajanhub/goapi/infrastructure"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/middlewares"
	"github.com/Rajanhub/goapi/repository"
	"github.com/Rajanhub/goapi/routes"
	"github.com/Rajanhub/goapi/seeds"
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
	seeds.Module,
	middlewares.Module,
)
