package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"

	"github.com/ArtemGretsov/golang-server-template/config"
)

var DBInstance *sqlx.DB
var once sync.Once

func DB() *sqlx.DB {
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
		DBInstance, err = sqlx.Connect("postgres", dbConnectString)

		if err != nil {
			panic(err)
		}
	})

	return DBInstance
}
