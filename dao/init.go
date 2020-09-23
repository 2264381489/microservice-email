package dao

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func Init() *sql.DB {
	log.Printf("init the db")

	conStr := "root:root@tcp(139.9.187.243:3306)/test"
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatalf("init db fail:%v\n", conStr)
		return nil
	}
	return db
}
