package models

import (
	"kalista/utils/common"
	"time"
)

type Tenant struct {
	ID       uint              `gorm:"primary_key" json:"id"`
	PubKey   common.NullString `gorm:"column:public_key" sql:"type:text" json:"pub_key"`
	ApiPath  common.NullString `gorm:"column:api_path" json:"api_path"`
	ApiKey   string            `gorm:"column:api_key;unique_index" sql:"not null" json:"api_key"`
	ExpireAt time.Time         `gorm:"column:expire_time" json:"expire_at"`
	Name     common.NullString `json:"name"`
}

// ---------------- CallBacks ----------------

func (t *Tenant) BeforeCreate() (err error) {
	t.ApiKey = t.generateApiKey()
	return
}

// Create
func NewTenant() {
	t := &Tenant{
		Name:     common.ToNullString("string"),
		ExpireAt: time.Now(),
	}

	db.Create(t)
}

// 这里返回 Tenant 列表
func Tenants(offset, limit int) ([]*Tenant, error) {
	tenants := make([]*Tenant, 0)

	if err := db.Model(&Tenant{}).Offset(offset).Limit(limit).Order("id asc").Find(&tenants).Error; err != nil {
		return tenants, err
	}

	return tenants, nil
}

func (Tenant) TableName() string {
	return "tenants"
}

func (t Tenant) generateApiKey() string {
	apiKey := common.RandStr(64)
	for {

		if db.Where(&Tenant{ApiKey: apiKey}).First(&t).RecordNotFound() {
			break
		} else {
			apiKey = common.RandStr(64)
		}
	}
	return apiKey
}
