package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/ArtemGretsov/golang-server-template/config"
	"github.com/ArtemGretsov/golang-server-template/src/database"
)

const currentDir = "./src/database/migrations/"
const currentFileName = "migration.go"

func main() {
	topLevelArg := os.Args[1]

	switch topLevelArg {
	case "generate":
		GenerateMigration()
	case "up":
		Up()
	case "down":
		Down()
	}
}

/* Generate migration */

func GenerateMigration() {
	client := database.DB()
	defer client.Close()

	timestamp := time.Now().UnixNano()
	migrationName := os.Args[2]
	upMigrationName := fmt.Sprintf("%d_%s.up.sql", timestamp, migrationName)
	upMigrationPath := path.Join(currentDir, upMigrationName)
	downMigrationName := fmt.Sprintf("%d_%s.down.sql", timestamp, migrationName)
	downMigrationPath := path.Join(currentDir, downMigrationName)

	creatingFile, err := os.Create(upMigrationPath)
	checkErr(err)

	err = client.Schema.WriteTo(context.Background(), creatingFile)
	checkErr(err)
	creatingFile.Close()

	openingFile, err := os.Open(upMigrationPath)
	checkErr(err)

	migrationContent, err := ioutil.ReadAll(openingFile)
	checkErr(err)
	openingFile.Close()

	rawMigrationContent := regexp.
		MustCompile(`BEGIN;\n(.+)\nCOMMIT;\n`).
		ReplaceAll(migrationContent, []byte("$1"))

	err = os.WriteFile(upMigrationPath, rawMigrationContent, 0644)
	checkErr(err)

	err = os.WriteFile(downMigrationPath, []byte(""), 0644)
	checkErr(err)

	fmt.Println()
	color.Green("Migration have been successfully generated.")
	fmt.Println("Please write migration * .down.sql manually!")
	fmt.Println()
}
/* ------------------ */


/* Migration Up */

func Up() {
	fmt.Println()
	fmt.Println("Start migration up...")

	serverConfig := config.Get()
	db, err := sql.Open("postgres", serverConfig["POSTGRES_CONNECT_STRING"])
	checkErr(err)
	defer db.Close()

	_, err = db.Query(`
			create table if not exists migrations(
    		id  bigint generated by default as identity constraint migrations_pkey primary key,
    		date_create timestamp default NOW(),
   			name    varchar not null
			)
  `)
	checkErr(err)

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	checkErr(err)

	_, err = tx.ExecContext(ctx, "lock table migrations in exclusive mode")
	checkErr(err)

	files := getMigrationFileNames()
	migrationNameHashTable := map[string]bool{}
	rows, err := tx.QueryContext(ctx, "select name from migrations")
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		checkErr(err)
		migrationNameHashTable[name] = true
	}

	var filesForExec []string

	for _, fileName := range files {
		isUp := strings.Contains(fileName, "up.sql")

		if !isUp {
			continue
		}

		migrationName := getMigrationName(fileName)

		if !migrationNameHashTable[migrationName] {
			filesForExec = append(filesForExec, fileName)
		}
	}

	fileForExecLength := len(filesForExec)

	if fileForExecLength == 0 {
		fmt.Println("No migrations found to update.")
	}

	if fileForExecLength != 0 {
		fmt.Printf("Found %d migrations to run.",  fileForExecLength)
		fmt.Println()
		fmt.Println()
	}

	for _, migrationName := range filesForExec {
		migrationPath := path.Join(currentDir, migrationName)
		migrationFile, err := os.Open(migrationPath)
		checkErr(err)

		migrationContent, err := ioutil.ReadAll(migrationFile)
		checkErr(err)

		_, err = tx.ExecContext(ctx, string(migrationContent))
		if err != nil {
			color.Red(" × %s - error!", migrationName)
			fmt.Println()
			checkErr(err)
		}

		_, err = tx.Exec("insert into migrations(name) values($1)", getMigrationName(migrationName))
		color.Green(" ✓ %s - successful!", migrationName)
	}

	err = tx.Commit()
	fmt.Println()
	checkErr(err)
}

func getMigrationName(name string) string {
	return regexp.
		MustCompile(`(.+)\.up\.sql`).
		ReplaceAllString(name, "$1")
}

func getMigrationFileNames() []string {
	var files []string
	allFiles, err := ioutil.ReadDir(currentDir)
	checkErr(err)

	for _, f := range allFiles {
		if currentFileName != f.Name() {
			files = append(files, f.Name())
		}
	}

	sort.Strings(files)
	return files
}
/* ------------------ */


/* Down migration */

func Down()  {
	fmt.Println()
	fmt.Println("Start migration down...")
	fmt.Println()

	serverConfig := config.Get()
	db, err := sql.Open("postgres", serverConfig["POSTGRES_CONNECT_STRING"])
	checkErr(err)
	defer db.Close()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	checkErr(err)

	_, err = tx.ExecContext(ctx, "lock table migrations in exclusive mode")
	checkErr(err)

	rows, err := tx.QueryContext(ctx, "select name from migrations order by date_create desc, name desc limit 1")
	defer rows.Close()
	checkErr(err)

	var migrationName string

	for rows.Next() {
		err = rows.Scan(&migrationName)
		checkErr(err)
	}

	if migrationName == "" {
		fmt.Println("No migrations found to down.")

		err = tx.Commit()
		checkErr(err)
		return
	}

	downMigrationName := migrationName + ".down.sql"
	downMigrationPath := path.Join(currentDir, downMigrationName)

	downMigrationFile, err := os.Open(downMigrationPath)
	checkErr(err)

	downMigrationContent, err := ioutil.ReadAll(downMigrationFile)
	checkErr(err)

	_, err = tx.ExecContext(ctx, string(downMigrationContent))

	if err != nil {
		color.Red(" × %s - error!", downMigrationName)
	}
	checkErr(err)

	_, err = tx.ExecContext(ctx, "delete from migrations where name=$1", migrationName)

	err = tx.Commit()
	checkErr(err)

	color.Green(" ✓ %s - successful!", downMigrationName)
	fmt.Println()
}
/* ------------------ */

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
