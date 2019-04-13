package models

import (
	"github.com/jinzhu/gorm"
)

type Tenant struct {
	gorm.Model
}

func (Tenant) TableName() string {
	return "tenants"
}
