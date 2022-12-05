package console

import (
	"time"

	"github.com/Rajanhub/goapi/infrastructure"
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/middlewares"
	"github.com/Rajanhub/goapi/routes"
	"github.com/Rajanhub/goapi/seeds"
	"github.com/spf13/cobra"
)

type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		env *lib.Env,
		router infrastructure.Router,
		routes routes.Routes,
		logger lib.Logger,
		seeds seeds.Seeds,

	) {
		logger.Info(`+-----------------------+`)
		logger.Info(`| GO CLEAN ARCHITECTURE |`)
		logger.Info(`+-----------------------+`)

		// Using time zone as specified in env file
		loc, _ := time.LoadLocation(env.TimeZone)
		time.Local = loc

		middleware.Setup()
		routes.Setup()
		seeds.Setup()
		logger.Info("Running server")
		if env.ServerPort == "" {
			if err := router.Run(); err != nil {
				logger.Fatal(err)
				return
			}
		} else {
			if err := router.Run(":" + env.ServerPort); err != nil {
				logger.Fatal(err)
				return
			}
		}
	}
}

func NewServeCommand() lib.Command {
	return &ServeCommand{}
}
