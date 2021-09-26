package main

import (
	"github.com/ArtemGretsov/golang-server-template/config"
	"github.com/ArtemGretsov/golang-server-template/src/database"
	"github.com/ArtemGretsov/golang-server-template/src/modules"
	"github.com/ArtemGretsov/golang-server-template/src/server"
)

func main() {
	/* Init database */
	database.DB()

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
