package main

import (
	"context"
	"log"
	"os"

	"github.com/ArtemGretsov/golang-server-template/src/database"
)

func main() {
	client := database.DB()
	defer client.Close()
	ctx := context.Background()
	// Dump migration changes to an SQL script.
	f, err := os.Create("migrate.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()
	if err := client.Schema.WriteTo(ctx, f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}