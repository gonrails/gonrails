package common

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var (
	DateFormat             = "2006-01-02"
	DateTimeFormat         = "2006-01-02 15:04:05"
	DateTimeFormatWithZone = time.RFC3339
	p                      = fmt.Println
	l                      = log.Println
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}
