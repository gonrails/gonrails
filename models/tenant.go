package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tenant struct {
	gorm.Model
}

func (t Tenant) Count() {
	var count int
	db.Table(Tenant{}.TableName()).Count(&count)

	fmt.Println(count)
}

func (Tenant) TableName() string {
	return "tenants"
}
