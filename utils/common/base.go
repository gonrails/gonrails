package common

import (
	"database/sql"
	"log"
	"time"

	"encoding/json"
)

var (
	DateFormat             = "2006-01-02"
	DateTimeFormat         = "2006-01-02 15:04:05"
	DateTimeFormatWithZone = time.RFC3339
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func ToNullString(s string) NullString {

	return NullString{
		sql.NullString{
			String: s,
			Valid:  s != "",
		},
	}
}

type NullString struct {
	sql.NullString
}

func (v *NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.String = *s
	} else {
		v.Valid = false
	}

	return nil
}
