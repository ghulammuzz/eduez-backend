package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// _ "github.com/go-sql-driver/mysql"

var DB *sql.DB

func InitDB() {
	var err error
	// DB, err = sql.Open("mysql", "root:secret@tcp(localhost:3306)/eduze")
	DB, err = sql.Open("sqlite3", "eduze.db")
	if err != nil {
		log.Fatal(err)
	}

	// DB.SetMaxIdleConns(10)
	// DB.SetMaxOpenConns(10)

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect DB")
}

func CloseDB() {
	DB.Close()
}
