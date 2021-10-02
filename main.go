package main

import (
	"github.com/ArtemGretsov/golang-server-template/config"
	"github.com/ArtemGretsov/golang-server-template/internal/database"
	"github.com/ArtemGretsov/golang-server-template/internal/modules"
	"github.com/ArtemGretsov/golang-server-template/internal/server"
)

func main() {
	/* Init database */
	DB := database.DB()
	defer DB.Close()

	/* Int fiber app */
	serverConfig := config.Get()
	app := server.Init()

	/* Init fiber routing */
	modules.Router(app)
	err := app.Listen(serverConfig["APP_PORT"])

	if err != nil {
		panic(err)
	}
}
