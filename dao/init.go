package dao

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"
// 用 database 开启一个链接
func Init() *sql.DB {
	log.Printf("init the db")

	//conStr := "root:root@tcp(139.9.187.243:3306)/test"
	conStr := "root:a13840419132@tcp(localhost:3306)/seckill"
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatalf("init db fail:%v\n", conStr)
		return nil
	}
	return db
}
