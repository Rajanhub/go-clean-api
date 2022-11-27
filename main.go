package main

import (
	//"example/go-api/controllers"
	"github.com/Rajanhub/goapi/bootstrap"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		bootstrap.CommonModules,
	).Run()

}
