package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func connect()  *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(dbHost)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_uts_pbp?parseTime=True&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return db
}