package models

import (
	. "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"grabathon/config"
	"log"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */
var db *DB

func init() {
	var err error
	db, err = Open("mysql", config.DBConnectionString)
	if err != nil {
		log.Fatalln(err)
	}
}