package main

import (
	"database/sql"
	"github.com/kaungmyathan22/golang-metro-train/database"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// Create tables
	database.Initialize(db)
}
