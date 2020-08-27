package postgres

import (
	"fmt"
	"gonrails/utils/common"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(host, port, username, password, database string) *gorm.DB {
	timezone := "'Asia/Shanghai'"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=%s",
		host,
		port,
		username,
		database,
		password,
		url.QueryEscape(timezone),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	common.PanicError(err)
	return db
}

func ConfigureMySQL(db *gorm.DB, idles, connections int) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(idles)
	sqlDB.SetMaxOpenConns(connections)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)
}
