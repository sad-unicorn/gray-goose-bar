package main

import (
	"fmt"
	"github.com/sad-unicorn/gray-goose-bar/cwapi"
	"github.com/sad-unicorn/gray-goose-bar/db"
	"github.com/sad-unicorn/gray-goose-bar/router"
	"log"
	"os"
)

func main() {
	initDatabase()
	initApi()
	router.StartBot(requireEnv("BOT_TOKEN"))
}

func initApi() {
	login := os.Getenv("API_LOGIN")
	password := os.Getenv( "API_PASSWORD")

	if login != "" && password != "" {
		log.Println("Connecting to API as ", login)
		cwapi.InitAPI(login, password)
	}
}

func initDatabase() {
	dbUser := requireEnv("DATABASE_USER")
	dbPass := requireEnv("DATABASE_PASS")
	dbName := requireEnv("DATABASE_NAME")
	dbHost := requireEnv("DATABASE_HOST")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", dbUser, dbPass, dbHost, dbName)
	db.InitDB(dbUrl)

	res, err := db.QueryForInt("SELECT 1")
	if err != nil {
		panic("Database is failed to `select 1`: " + err.Error())
	}

	fmt.Println("Db query: ", res)

	db.InitMigrations()
}

func requireEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("No " + key + " environment variable is set")
	}
	return val
}
