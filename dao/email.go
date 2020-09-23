package dao

import (
	"log"
	"time"
)

var DB = Init()

func createTable() error {
	DB.Exec("DROP email if exists ")
	sql := "create table email(id int auto_increment primary key, from_user varchar(25) not null, to_user varchar(25) not null, create_time datetime ON UPDATE CURRENT_TIMESTAMP, update_time datetime ON UPDATE CURRENT_TIMESTAMP);"
	_, err := DB.Exec(sql)
	if err == nil {
		log.Printf("create table email successsfully")
	}
	return err
}

func InsertEmail(fromUser, toUser string) error {
	sql := "insert into email(from_user, to_user, create_time, update_time) values(?,?,?,?);"
	_, err := DB.Exec(sql, fromUser, toUser, time.Now(), time.Now())
	if err != nil {
		log.Printf("insert fail, err:%v\n", err)
		return err
	}
	log.Printf("insert successfully")
	return nil
}
