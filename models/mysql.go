package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/one-hole/gonrails/config"
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

	db.LogMode(true)
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

func DB() *gorm.DB {
	return db
}

func init() {

	db = open(
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Name,
	)

	// db.AutoMigrate(&Game{})
	// db.AutoMigrate(&ScoreDetail{})
	// db.AutoMigrate(&BetOrder{})
	// db.AutoMigrate(&BetOrderAddition{})
	// db.AutoMigrate(&Tenant{})
	// db.AutoMigrate(&RevenueReport{})
	// db.AutoMigrate(&StatisticReports{})
}
