package main

import (
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"

	"github.com/ArtemGretsov/golang-server-template/config"
)

func main() {
	serverConfig := config.Get()
	client, _ := sql.Open("postgres", serverConfig["POSTGRES_CONNECT_STRING"])

	for {
		err := client.DB().Ping()

		if err == nil {
			break
		}
	}
}
