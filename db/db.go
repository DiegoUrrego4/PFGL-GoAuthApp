package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}

	sqlFile, err := os.ReadFile("sql/users.sql")
	if err != nil {
		fmt.Println("There was an error reading users sql file", err)
		return nil, err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Print("There was an error executing the query\n", err)
		return nil, err
	}

	return db, nil
}
