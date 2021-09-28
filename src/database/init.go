package database

import (
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
		DBInstance, err = _schemagen.Open("postgres", serverConfig["POSTGRES_CONNECT_STRING"])

		if err != nil {
			log.Fatal(err)
		}
	})

	return DBInstance
}
