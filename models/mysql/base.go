package mysql

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gonrails/gonrails/utils/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Open open a mysql database to use
func Open(host, port, username, password, name string) *gorm.DB {
	timezone := "'Asia/Shanghai'"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&time_zone=%s",
		username,
		password,
		host,
		port,
		name,
		url.QueryEscape(timezone),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	common.PanicError(err)
	return db
}

// ConfigureMySQL configure db's idles connections, max connections
func ConfigureMySQL(db *gorm.DB, idles, connections int) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(idles)
	sqlDB.SetMaxOpenConns(connections)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)
}
