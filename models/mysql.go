package models

import (
	"fmt"
	"kalista/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func open(host, port, username, password, name string) *gorm.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)

	db, err := gorm.Open("mysql", dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	configureMySQL(db)
	return db
}

func configureMySQL(db *gorm.DB) {
	db.DB().SetMaxIdleConns(config.MySQL.Idles)
	db.DB().SetMaxOpenConns(config.MySQL.Connections)
	db.DB().SetConnMaxLifetime(time.Minute * 1)
}

func Close() {
	db.Close()
}

func init() {

	db = open(
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Name,
	)

	db.AutoMigrate(&ScoreDetail{})
	db.AutoMigrate(&BetOrder{})
}
