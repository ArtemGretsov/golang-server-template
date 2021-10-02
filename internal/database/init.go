package database

import (
	_ "github.com/lib/pq"
	"sync"

	"github.com/ArtemGretsov/golang-server-template/config"
	"github.com/ArtemGretsov/golang-server-template/internal/database/schemagen"
)

var DBInstance *schemagen.Client
var once sync.Once

func DB() *schemagen.Client {
	once.Do(func() {
		var err error
		serverConfig := config.Get()

		DBInstance, err = schemagen.Open("postgres", serverConfig["POSTGRES_CONNECT_STRING"])

		if err != nil {
			panic(err)
		}
	})

	return DBInstance
}
