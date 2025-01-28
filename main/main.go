package main

import (
	"fmt"
	"natural-language-app/config"
	"natural-language-app/database"
	"natural-language-app/processor"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	config.LoadDBCfg()

	// connect to DB
	if err := database.ConnectDB(); err != nil {
		fmt.Println(err)
	}

	// create tables
	db := database.GetDB()
	db.CreateTables()

	processor.Process()
}
