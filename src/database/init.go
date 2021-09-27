package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"

	"github.com/ArtemGretsov/golang-server-template/config"
	"github.com/ArtemGretsov/golang-server-template/src/database/_schemagen"
)

var DBInstance *_schemagen.Client
var once sync.Once

func DB() *_schemagen.Client {
	once.Do(func() {
		var err error
		serverConfig := config.Get()
		dbConnectString := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			serverConfig["DB_USER"],
			serverConfig["DB_PASS"],
			serverConfig["DB_HOST"],
			serverConfig["DB_PORT"],
			serverConfig["DB_NAME"],
		)

		DBInstance, err = _schemagen.Open("postgres", dbConnectString)

		if err != nil {
			log.Fatal(err)
		}
	})

	return DBInstance
}
