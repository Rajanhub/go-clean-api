package main

import (
	//"example/go-api/controllers"
	"github.com/Rajanhub/goapi/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	_ = bootstrap.RootApp.Execute()
}
