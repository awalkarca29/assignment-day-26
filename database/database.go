package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/assignment_day_26")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful!")
}
