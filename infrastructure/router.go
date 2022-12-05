package infrastructure

import (
	"log"
	"net/http"

	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/utils"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(env *lib.Env) Router {

	if env.Environment != "local" && env.SentryDSN != "" {

		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         env.SentryDSN,
			Environment: `clean-backend-` + env.Environment,
		}); err != nil {
			//logger.Infof("Sentry initialization failed: %v\n", err)
			log.Printf("Sentry initialization failed: %v\n", err)
		}
	}
	appEnv := env.Environment

	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	httpRouter := gin.Default()

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		utils.SendSentryMsg(c, "Error")
		c.JSON(http.StatusOK, gin.H{"data": "clean architecture 📺 API Up and Running"})
	})
	// Attach sentry middleware

	return Router{httpRouter}
}
