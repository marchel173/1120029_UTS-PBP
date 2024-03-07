package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect()  *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(dbHost)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=True&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return db
}

func connectGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=True&loc=Asia%2FJakarta"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}