package postgres

import (
	"fmt"
	"time"

	"github.com/gonrails/gonrails/utils/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(host, port, username, password, database string, gormConfig *gorm.Config) *gorm.DB {
	timezone := "Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=%s",
		host,
		port,
		username,
		database,
		password,
		timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	common.PanicError(err)
	return db
}

func ConfigureMySQL(db *gorm.DB, idles, connections int) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(idles)
	sqlDB.SetMaxOpenConns(connections)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)
}
