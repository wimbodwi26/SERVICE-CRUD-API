package main

import (
	"backend-go/config"
	"backend-go/database"
	"backend-go/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r. Run(":" + config.GetEnv("APP_PORT", "8080"))
}