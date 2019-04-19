package models

import (
	"database/sql"
	"fmt"
	"kalista/utils/common"
	"log"
	"time"
)

type Tenant struct {
	ID       uint           `gorm:"primary_key"`
	PubKey   sql.NullString `gorm:"column:public_key" sql:"type:text" `
	ApiPath  sql.NullString `gorm:"column:api_path"`
	ApiKey   sql.NullString `gorm:"column:api_key;unique_index" sql:"not null"`
	ExpireAt time.Time      `gorm:"column:expire_time"`
	Name     sql.NullString
}

func (t Tenant) Count() {
	var count int
	db.Table(Tenant{}.TableName()).Count(&count)

	fmt.Println(count)
}

// ---------------- CallBacks ----------------

func (t *Tenant) BeforeCreate() (err error) {

	log.Println("Callback")
	t.ApiKey = t.generateApiKey()

	return
}

func NewTenant() {
	log.Println("Creating new tenant")
	t := &Tenant{
		Name:     common.ToNullString("string"),
		ExpireAt: time.Now(),
	}

	db.Create(t)
}

func (Tenant) TableName() string {
	return "tenants"
}

func (t Tenant) generateApiKey() sql.NullString {
	apiKey := common.ToNullString(common.RandStr(64))
	for {

		if db.Where(&Tenant{ApiKey: apiKey}).First(&t).RecordNotFound() {
			break
		} else {
			apiKey = common.ToNullString(common.RandStr(64))
		}
	}
	return apiKey
}
