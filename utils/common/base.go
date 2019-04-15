package common

import (
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
