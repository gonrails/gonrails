package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/one-hole/gonrails/config"
	"github.com/one-hole/gonrails/utils/common"
)

func Open(host, port, username, password, name string) *gorm.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)

	db, err := gorm.Open("mysql", dbConfig)
	common.PanicError(err)

	configureMySQL(db)

	return db
}

func configureMySQL(db *gorm.DB) {
	db.DB().SetMaxIdleConns(config.MySQL.Idles)
	db.DB().SetMaxOpenConns(config.MySQL.Connections)
	db.DB().SetConnMaxLifetime(time.Minute * 1)
}
