package helper

import (
	"log"
	"time"
)

const (
	DateFormat             = "2006-01-02"
	DateTimeFormat         = "2006-01-02 15:04:05"
	DateTimeFormatWithZone = time.RFC3339
	DefaultPageSize        = 10
	DefaultPage            = 1
)


func TimeWithLocation(layout, name string, value string) time.Time {
	if "" == name {
		name = "Asia/Shanghai"
	}

	location, err := time.LoadLocation(name)

	if err != nil {
		log.Panic(err)
	}

	t, err := time.ParseInLocation(layout, value, location)

	if err != nil {
		log.Panic(err)
	}

	return t
}